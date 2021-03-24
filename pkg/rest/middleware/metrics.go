package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/seosite/gcore/pkg/core/metricsx"
)

const (
	handler = "gin_http"
)

var (
	httpRequestTotal      *metricsx.MetricsCounter
	httpResponseTimeUs    *metricsx.MetricsHistogram
	httpResponseSizeBytes *metricsx.MetricsHistogram
	httpRequestSizeBytes  *metricsx.MetricsHistogram
)

func Metrics() gin.HandlerFunc {
	metricsInit()
	return func(c *gin.Context) {
		requestCon := c.Request
		responseCon := c.Writer
		if requestCon.URL.Path == "/metrics" {
			c.Next()
			return
		}
		var headlen int
		for k, v := range requestCon.Header {
			headlen += len(k)
			for _, vv := range v {
				headlen += len(vv)
			}
		}
		requestSize := float64(c.Request.ContentLength) + float64(headlen)
		timeNow := time.Now()
		c.Next()
		label := []string{handler, strconv.Itoa(responseCon.Status()), requestCon.URL.Path, requestCon.Method}
		labelb := []string{handler, requestCon.URL.Path, requestCon.Method}
		httpRequestTotal.MetricsAdd(label, 1) // 往metric里面添加数据
		httpRequestSizeBytes.MetricsAdd(labelb, requestSize)
		httpResponseSizeBytes.MetricsAdd(labelb, float64(responseCon.Size()))
		endTime := time.Since(timeNow)
		t := float64(endTime) / 1000
		httpResponseTimeUs.MetricsAdd(labelb, t)

		return
	}
}

func metricsInit() {
	s := []string{"handler", "code", "path", "method"}
	httpRequestTotal = &metricsx.MetricsCounter{ // 构建Counter类型的metrics初始化数据
		Metrics: metricsx.Metrics{
			Name:  "http_request_total",
			Help:  "How many HTTP requests processed,partitioned by status code and HTTP method.",
			Label: s,
		},
	}
	httpRequestTotal.NewCounterVec() // 初始化一个Counter类型的metric

	s = []string{"handler", "path", "method"}
	httpResponseTimeUs = &metricsx.MetricsHistogram{
		Metrics: metricsx.Metrics{
			Name:  "http_response_time_us",
			Help:  "How much time costs(us)",
			Label: s,
		},
		Buckets: prometheus.LinearBuckets(10000, 30000, 10),
	}
	httpResponseTimeUs.NewHistogramVec()

	httpResponseSizeBytes = &metricsx.MetricsHistogram{
		Metrics: metricsx.Metrics{
			Name:  "http_response_size_bytes",
			Help:  "The size of http response",
			Label: s,
		},
		Buckets: prometheus.LinearBuckets(256, 256, 10),
	}
	httpResponseSizeBytes.NewHistogramVec()

	httpRequestSizeBytes = &metricsx.MetricsHistogram{
		Metrics: metricsx.Metrics{
			Name:  "http_request_size_bytes",
			Help:  "The size of http request",
			Label: s,
		},
		Buckets: prometheus.LinearBuckets(256, 256, 10),
	}
	httpRequestSizeBytes.NewHistogramVec()
}
