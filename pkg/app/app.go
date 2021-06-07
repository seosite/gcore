package app

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/seosite/gcore/pkg/core/collection"
	"github.com/seosite/gcore/pkg/core/env"
	"github.com/seosite/gcore/pkg/core/tencentyun"
	"github.com/seosite/gcore/pkg/core/third"
	"github.com/seosite/gcore/pkg/rest/jisu"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	// Config app config
	Config Conf
	// Logger app logger
	Logger *zap.Logger
	// Db app db pool
	Db map[string]*gorm.DB
	// Redis app redis clients
	Redis map[string]*redis.Client
	// Cache app memory cache
	Cache *collection.Cache
	// Cos app cos
	Cos *tencentyun.CosStorage
	// Viper app viper
	Viper *viper.Viper

	// Sso sso service
	Sso *third.Sso
	// Analytics bigdata analytics service
	Analytics *third.Analytics

	// AppLogFile default app log file
	AppLogFile = "./storage/logs/app.log"

	JisuAPI *jisu.JisuAPI
)

// ------ config ------

// GetConfig get config by name with dotable
func GetConfig(key string) interface{} {
	return Viper.Get(key)
}

// ------ env ------

// IsLocal check env is local or not
func IsLocal() bool {
	return Config.Server.Env == env.TypeLocal
}

// IsQa check env is qa or not
func IsQa() bool {
	return Config.Server.Env == env.TypeQa
}

// IsProd check env is prod or not
func IsProd() bool {
	return Config.Server.Env == env.TypeProd
}

// ------ database ------

// Migrate migrate database tables, MUST BE SUCCESS
func Migrate(dbTables map[string][]interface{}) {
	for name, tables := range dbTables {
		if err := UseDb(name).AutoMigrate(tables...); err != nil {
			panic(fmt.Errorf("DB[%s] migrate error: %s", name, err))
		}
	}
}

// UseDb use db with name that MUST EXISTS
func UseDb(name string) *gorm.DB {
	if db, found := Db[name]; found {
		return db
	}
	panic(fmt.Errorf("DB[%s] not found", name))
}

// DefaultDb use default db that MUST EXISTS
func DefaultDb() *gorm.DB {
	return UseDb("default")
}

// SlaveDb use slave db that MUST EXISTS
func SlaveDb() *gorm.DB {
	return UseDb("slave")
}

// ------ redis ------

// UseRedis use redis with name that MUST EXISTS
func UseRedis(name string) *redis.Client {
	if redis, found := Redis[name]; found {
		return redis
	}
	panic(fmt.Errorf("Redis[%s] not found", name))
}

// DefaultRedis use default redis that MUST EXISTS
func DefaultRedis() *redis.Client {
	return UseRedis("default")
}

// InitRedisByDB select redis db
func DefaultRedisByDB(db int, redisName string) *redis.Client {
	configs := Config.Redis
	redisNum := len(configs)
	if redisNum == 0 {
		return nil
	}

	if config, ok := configs[redisName]; ok {
		client := redis.NewClient(&redis.Options{
			Addr:     config.Addr,
			Password: config.Password,
			DB:       db,
		})
		_, err := client.Ping().Result()
		if err != nil {
			panic(fmt.Errorf("Ping redis[%s] error: %s", redisName, err))
		}

		Redis[redisName] = client
		return client
	}

	return nil
}
