package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// WithKeepalive keepalive router
var WithKeepalive = func(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}
