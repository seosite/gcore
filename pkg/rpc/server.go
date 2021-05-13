package rpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"runtime/debug"
	"strings"
	"sync/atomic"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/seosite/gcore/pkg/app"
	"github.com/seosite/gcore/pkg/core/ecode"
	"github.com/seosite/gcore/pkg/core/threading"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	alerting         uint32
	lastTryAlertTime = time.Now().AddDate(-1, -1, -1)
)

// Server rest server
type Server struct {
	Opt         ServerOpt
	RPCAddress  string
	RPCEngine   *grpc.Server
	RPCPromReg  *prometheus.Registry
	RPCMetrics  *grpc_prometheus.ServerMetrics
	HTTPContext context.Context
	HTTPMux     *runtime.ServeMux
	HTTPOpts    []grpc.DialOption
}

// ServerOpt server options
type ServerOpt struct {
	ConfigFile    string
	MigrateTables map[string][]interface{}
}

// Default integrate with default config/logger/db/redis/middleware/routes
func Default(serverOpt ServerOpt) *Server {
	s := &Server{
		Opt: serverOpt,
	}
	// init conf
	app.InitConfig(s.Opt.ConfigFile)
	// init logger
	app.InitLogger()
	// init third services
	app.InitThird()
	// init db
	app.InitDb()
	app.Migrate(s.Opt.MigrateTables)
	// init redis
	app.InitRedis()
	// init cos
	app.InitCos()

	// grpc prometheus
	s.RPCMetrics = grpc_prometheus.NewServerMetrics()
	s.RPCPromReg = prometheus.NewRegistry()
	s.RPCPromReg.MustRegister(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}))
	s.RPCPromReg.MustRegister(prometheus.NewGoCollector())
	s.RPCPromReg.MustRegister(s.RPCMetrics)
	// grpc server
	s.RPCAddress = fmt.Sprintf("localhost:%d", app.Config.Server.RPCPort)
	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(func(p interface{}) (err error) {
			serveEnd(p)
			return status.Errorf(codes.Unknown, "panic triggered: %v", p)
		}),
	}
	s.RPCEngine = grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(opts...),
			s.RPCMetrics.UnaryServerInterceptor(),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_recovery.StreamServerInterceptor(opts...),
			s.RPCMetrics.StreamServerInterceptor(),
		),
	)
	// http server
	s.HTTPContext = context.Background()
	s.HTTPMux = runtime.NewServeMux()
	s.HTTPOpts = []grpc.DialOption{grpc.WithInsecure()}

	return s
}

func (s *Server) preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	// fmt.Printf("preflight request for %s", r.URL.Path)
	return
}

func (s *Server) allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				s.preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

// Run run server
func (s *Server) Run() {
	// http serve
	threading.GoSafe(func() {
		defer serveEnd(recover())
		// init metrics
		s.RPCMetrics.InitializeMetrics(s.RPCEngine)

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "ok")
		})
		http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "pong")
		})
		http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
			promhttp.HandlerFor(s.RPCPromReg, promhttp.HandlerOpts{}).ServeHTTP(w, r)
		})
		fmt.Println("Running http server on :", app.Config.Server.Port)
		http.ListenAndServe(":"+cast.ToString(app.Config.Server.Port), nil)
	})

	// rpc gateway serve
	threading.GoSafe(func() {
		defer serveEnd(recover())
		fmt.Println("Running rpc gateway server on :", app.Config.Server.RPCGatewayPort)
		http.ListenAndServe(":"+cast.ToString(app.Config.Server.RPCGatewayPort), s.allowCORS(s.HTTPMux))
	})

	// rpc serve
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", app.Config.Server.RPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()
	fmt.Println("Running rpc server on :", app.Config.Server.RPCPort)
	if err := s.RPCEngine.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func serveEnd(err interface{}) {
	if err != nil {
		errDetail := string(debug.Stack())
		// log panic
		app.Logger.DPanic(
			"recover from panic",
			// zap.String("client_ip", ctx.ClientIP()),
			// zap.String("method", ctx.Request.Method),
			// zap.String("uri", ctx.Request.RequestURI),
			// zap.Int("ecode", ecode.EcodeRESTPanic),
			zap.Any("emsg", err),
			zap.String("edetail", errDetail),
		)

		// send alert message to work wechat, with 1 minute interval
		alertUsers := app.Config.Server.AlertUsers
		if canSendAlert(alertUsers) &&
			atomic.CompareAndSwapUint32(&alerting, 0, 1) &&
			canSendAlert(alertUsers) { // double check
			defer atomic.StoreUint32(&alerting, 0)

			lastTryAlertTime = time.Now()
			content := "App: %s\nEnv: %s\nTime: %s\nError: %s"
			content = fmt.Sprintf(content, app.Config.Server.Name, app.Config.Server.Env, lastTryAlertTime.Format("2006-01-02 15:04:05"), errDetail)
			alertErr := app.Sso.SendWorkWechatMsg(alertUsers, fmt.Sprintf("%v", content))
			if alertErr != nil {
				app.Logger.Error(
					"alert work wechat failed",
					// zap.String("client_ip", ctx.ClientIP()),
					// zap.String("method", ctx.Request.Method),
					// zap.String("uri", ctx.Request.RequestURI),
					zap.Any("alertUsers", alertUsers),
					zap.Int("ecode", ecode.EcodeRESTAlertWechatErr),
					zap.Any("emsg", alertErr),
				)
			}
		}
	}
}

// canSendAlert can try alert with 1 minute interval
func canSendAlert(alertUsers []string) bool {
	return len(alertUsers) > 0 && time.Since(lastTryAlertTime).Minutes() >= 1
}
