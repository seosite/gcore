package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/seosite/gcore/pkg/app"
	"github.com/seosite/gcore/pkg/core/ecode"
	"github.com/seosite/gcore/pkg/core/third"
	"go.uber.org/zap"
)

var (
	alerting         uint32
	lastTryAlertTime = time.Now().AddDate(-1, -1, -1)
)

// BaseRecover recover with both log and wechat alert
func BaseRecover(log *zap.Logger, sso *third.Sso, alertUsers []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				errDetail := string(debug.Stack())
				// log panic
				log.DPanic(
					"recover from panic",
					zap.String("client_ip", ctx.ClientIP()),
					zap.String("method", ctx.Request.Method),
					zap.String("uri", ctx.Request.RequestURI),
					zap.Int("ecode", ecode.EcodeRESTPanic),
					zap.Any("emsg", err),
					zap.String("edetail", errDetail),
				)

				// send alert message to work wechat, with 1 minute interval
				if canSendAlert(alertUsers) &&
					atomic.CompareAndSwapUint32(&alerting, 0, 1) &&
					canSendAlert(alertUsers) { // double check
					defer atomic.StoreUint32(&alerting, 0)

					lastTryAlertTime = time.Now()
					content := "App: %s\nEnv: %s\nTime: %s\nError: %s"
					content = fmt.Sprintf(content, app.Config.Server.Name, app.Config.Server.Env, lastTryAlertTime.Format("2006-01-02 15:04:05"), errDetail)
					alertErr := sso.SendWorkWechatMsg(alertUsers, fmt.Sprintf("%v", content))
					if alertErr != nil {
						log.Error(
							"alert work wechat failed",
							zap.String("client_ip", ctx.ClientIP()),
							zap.String("method", ctx.Request.Method),
							zap.String("uri", ctx.Request.RequestURI),
							zap.Any("alertUsers", alertUsers),
							zap.Int("ecode", ecode.EcodeRESTAlertWechatErr),
							zap.Any("emsg", alertErr),
						)
					}
				}

				// abort
				ctx.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		ctx.Next()
	}
}

// canSendAlert can try alert with 1 minute interval
func canSendAlert(alertUsers []string) bool {
	return len(alertUsers) > 0 && time.Since(lastTryAlertTime).Minutes() >= 1
}
