package rest

import (
	"strconv"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"

	"github.com/seosite/gcore/pkg/app"
	"github.com/seosite/gcore/pkg/rest/middleware"
	"github.com/seosite/gcore/pkg/rest/router"
)

// Server rest server
type Server struct {
	Opt    ServerOpt
	Engine *gin.Engine
}

// ServerOpt server options
type ServerOpt struct {
	ConfigFile    string
	Middlewares   []gin.HandlerFunc
	Router        func(r *gin.Engine)
	MigrateTables map[string][]interface{}
}

// Default integrate with default config/logger/db/redis/middleware/routes
func Default(serverOpt ServerOpt) *Server {
	s := &Server{
		Opt: serverOpt,
	}
	// init conf
	app.InitConfig(s.Opt.ConfigFile)
	// init env stage
	var middlewareLogger gin.HandlerFunc
	if app.IsProd() {
		gin.SetMode(gin.ReleaseMode)
		middlewareLogger = middleware.BaseLogger
	} else {
		gin.ForceConsoleColor()
		// middlewareLogger = middleware.DebugLogger
		middlewareLogger = middleware.BaseLogger
	}
	// init engine
	s.Engine = gin.New()
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
	// init middlewares
	s.Engine.Use(middlewareLogger)
	s.Engine.Use(middleware.BaseRecover(app.Logger, app.Sso, app.Config.Server.AlertUsers))
	s.Engine.Use(cors.Default())
	s.Engine.Use(s.Opt.Middlewares...)
	// init routes
	router.WithKeepalive(s.Engine)
	router.WithMetrics(s.Engine)
	s.Opt.Router(s.Engine)

	return s
}

// Run run server
func (s *Server) Run() {
	s.Engine.Run(":" + strconv.Itoa(app.Config.Server.Port))
}
