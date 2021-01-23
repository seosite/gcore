package app

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/seosite/gcore/pkg/core/collection"
	"github.com/seosite/gcore/pkg/core/logx"
	"github.com/seosite/gcore/pkg/core/tencentyun"
	"github.com/seosite/gcore/pkg/core/third"
	"github.com/seosite/gcore/pkg/core/zapx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	driver_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/prometheus"
)

// InitConfig init app config
func InitConfig(configFile string) {
	v := viper.New()

	v.SetConfigFile(configFile)

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Read config file error: %s", err))
	}

	if err := v.Unmarshal(&Config); err != nil {
		panic(fmt.Errorf("Parse config file error: %s", err))
	}

	fmt.Println("Load config file success.")

	Viper = v
}

// InitLogger init app logger
func InitLogger() {
	debug := !IsProd()
	level := -1
	if !debug {
		level = 0
	}

	var logger *zap.Logger
	var err error
	switch Config.Log.Mod {
	case logx.ModFileSizeRotate:
		cfg := &zapx.RotateLogConfig{
			Level:      level,
			Filename:   AppLogFile,
			MaxSize:    100, // MB
			MaxBackups: 10,
			MaxAge:     15, // day
		}
		logger, err = zapx.NewRotateZap(cfg)
	case logx.ModK8s:
		cfg := &zapx.K8sLogConfig{
			Level:      level,
			Dev:        debug,
			AccessFile: "/opt/app-logs/" + Config.Server.Name + "-access.log",
			ErrorFile:  "/opt/app-logs/" + Config.Server.Name + "-error.log",
		}
		logger, err = zapx.NewK8sZap(cfg)
	default:
		cfg := &zapx.LogConfig{
			Level:       level,
			OutputPaths: []string{"stdout"},
			Dev:         debug,
		}
		logger, err = zapx.NewZap(cfg)
	}
	if err != nil {
		panic(fmt.Errorf("Init logger error: %s", err))
	}
	zapx.Async()

	Logger = logger
}

// InitDb init app db
func InitDb() {
	configs := Config.Mysql
	dbNum := len(configs)
	if dbNum == 0 {
		return
	}

	Db = make(map[string]*gorm.DB, dbNum)
	for name, config := range configs {
		loglevel := logger.Info
		if IsProd() {
			loglevel = logger.Silent
		}
		dsn := config.Username + ":" + config.Password + "@(" + config.Path + ")/" + config.Dbname + "?" + config.Config
		db, err := gorm.Open(driver_mysql.Open(dsn), &gorm.Config{
			Logger:                                   logger.Default.LogMode(loglevel), // @TODO change to app logger
			SkipDefaultTransaction:                   true,
			PrepareStmt:                              true,
			DisableForeignKeyConstraintWhenMigrating: true,
		})
		if err != nil {
			panic(fmt.Errorf("Connect mysql[%s] error: %s", name, err))
		}

		// export metrics
		db.Use(prometheus.New(prometheus.Config{
			DBName:          "default",                  // `DBName` as metrics label
			RefreshInterval: 15,                         // refresh metrics interval (default 15 seconds)
			StartServer:     false,                      // start http server to expose metrics
			HTTPServerPort:  uint32(Config.Server.Port), // configure http server port, default port 8080 (if you have configured multiple instances, only the first `HTTPServerPort` will be used to start server)
			MetricsCollector: []prometheus.MetricsCollector{
				&prometheus.MySQL{
					Prefix:        "gorm_status_",
					Interval:      100,
					VariableNames: []string{"Threads_running"},
				},
			},
		}))

		sqlDb, err := db.DB()
		if err != nil {
			panic(fmt.Errorf("Ping mysql[%s] error: %s", name, err))
		}

		sqlDb.SetMaxIdleConns(config.MaxIdleConns)
		sqlDb.SetMaxOpenConns(config.MaxOpenConns)
		sqlDb.SetConnMaxLifetime(time.Hour)

		Db[name] = db
	}
}

// InitRedis init app redis client
func InitRedis() {
	configs := Config.Redis
	redisNum := len(configs)
	if redisNum == 0 {
		return
	}

	Redis = make(map[string]*redis.Client, redisNum)
	for name, config := range configs {
		client := redis.NewClient(&redis.Options{
			Addr:     config.Addr,
			Password: config.Password,
			DB:       config.DB,
		})
		_, err := client.Ping().Result()
		if err != nil {
			panic(fmt.Errorf("Ping redis[%s] error: %s", name, err))
		}

		Redis[name] = client
	}
}

// InitCache init app cache
func InitCache() {
	config := Config.Cache
	if config.Size == 0 || config.Expire == 0 {
		return
	}

	expire := time.Duration(config.Expire) * time.Second
	cache, err := collection.NewCache(expire, collection.WithLimit(config.Size))
	if err != nil {
		panic(fmt.Errorf("Init cache error: %s", err))
	}

	Cache = cache
}

// InitCos init app default cos
func InitCos() {
	config := Config.Cos
	if config.SecretID == "" {
		return
	}

	timeout := time.Duration(config.Timeout) * time.Second
	cosStorage := tencentyun.NewCosStorage(config.SecretID, config.SecretKey, config.Region, config.Bucket, config.RootPath, timeout)

	Cos = cosStorage
}

// InitMiddleware init router middleware
// func InitMiddleware(r *gin.Engine) {
// 	middleware.Init(r)
// }

// // InitRouter init router
// func InitRouter(r *gin.Engine) {
// 	router.Init(r)
// }

// InitThird init third services
func InitThird() {
	// sso
	Sso = &third.Sso{Domain: Config.ThirdService.Sso.Domain}

	// analytics
	analyticsEnv := third.AnalyticsEnvTest
	if IsProd() {
		analyticsEnv = third.AnalyticsEnvDefault
	}
	Analytics = &third.Analytics{
		Logger:   Logger,
		Domain:   Config.ThirdService.Analytics.Domain,
		Env:      analyticsEnv,
		Version:  Config.ThirdService.Analytics.Version,
		AppID:    Config.ThirdService.Analytics.AppID,
		Platform: Config.ThirdService.Analytics.Platform,
	}

}
