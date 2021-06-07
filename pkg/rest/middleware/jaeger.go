package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/seosite/gcore/pkg/app"
	"github.com/seosite/gcore/pkg/core/jaeger_trace"
)

func WithJaeger() gin.HandlerFunc {

	return func(c *gin.Context) {
		fmt.Println("app.Config.ThirdService.JaegerTrace.IsOpen....", app.Config.ThirdService.JaegerTrace.IsOpen)
		if app.Config.ThirdService.JaegerTrace.IsOpen == 1 {

			var parentSpan opentracing.Span

			tracer, closer := jaeger_trace.NewJaegerTracer(app.Config.Server.Name, app.Config.ThirdService.JaegerTrace.HostPort)
			defer closer.Close()

			spCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
			if err != nil {
				parentSpan = tracer.StartSpan(c.Request.URL.Path)
				defer parentSpan.Finish()
			} else {
				parentSpan = opentracing.StartSpan(
					c.Request.URL.Path,
					opentracing.ChildOf(spCtx),
					opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
					ext.SpanKindRPCServer,
				)
				defer parentSpan.Finish()
			}
			c.Set("Tracer", tracer)
			c.Set("ParentSpanContext", parentSpan.Context())
		}
		c.Next()
	}
}
