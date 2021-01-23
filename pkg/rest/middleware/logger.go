package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/seosite/gcore/pkg/core/jsonx"
)

const (
	timeFormatter = "2006-01-02 15:04:05.000000"
)

// BaseLogFormatter basic request log formatter for gin
var BaseLogFormatter = func(param gin.LogFormatterParams) string {
	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		param.ClientIP,
		// param.TimeStamp.Format(time.RFC1123),
		param.TimeStamp.Format(timeFormatter),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}

// DebugLogFormatter debug json formatter with request details
var DebugLogFormatter = func(param gin.LogFormatterParams) string {
	keys := param.Keys
	val, _ := jsonx.Marshal(keys)
	jsonFormat := `{"level":"debug","ts":"%s","method":"%s","uri":"%s","status":%3d,"latency":"%v","client_ip":"%s","error_msg":"%s","keys":%s,"size":%d}` + "\n"
	return fmt.Sprintf(jsonFormat,
		param.TimeStamp.Format(timeFormatter),
		param.Method,
		param.Request.RequestURI,
		param.StatusCode,
		param.Latency,
		param.ClientIP,
		param.ErrorMessage,
		string(val),
		param.BodySize,
	)
}

// BaseLogger base logger middleware
var BaseLogger = gin.LoggerWithFormatter(BaseLogFormatter)

// DebugLogger debug json logger
var DebugLogger = gin.LoggerWithFormatter(DebugLogFormatter)
