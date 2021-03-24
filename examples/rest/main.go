package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seosite/gcore/pkg/rest"
)

func main() {
	opt := rest.ServerOpt{
		ConfigFile:    initConfig(),
		Router:        initRoutes,
		MigrateTables: initTables(),
	}

	server := rest.Default(opt)

	initOthers()

	server.Run()
}

func initConfig() string {
	return "./dev.yaml"
}

func initRoutes(r *gin.Engine) {
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "test")
	})
}

func initTables() map[string][]interface{} {
	dbMigrate := make(map[string][]interface{})
	// dbMigrate["default"] = []interface{}{
	// 	model.User{},
	// }

	return dbMigrate
}

func initOthers() {
	// do others
}
