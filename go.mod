module github.com/seosite/gcore

go 1.15

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.0 // indirect
	github.com/fatih/structs v1.1.0
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/golang/protobuf v1.4.3
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.1-0.20190118093823-f849b5445de4
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.3.0
	github.com/hashicorp/go-retryablehttp v0.6.8
	github.com/json-iterator/go v1.1.10
	github.com/mitchellh/mapstructure v1.1.2
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/opentracing/opentracing-go v1.1.0
	github.com/prometheus/client_golang v1.9.0
	github.com/segmentio/kafka-go v0.4.10
	github.com/spaolacci/murmur3 v1.1.0
	github.com/spf13/cast v1.3.1
	github.com/spf13/viper v1.7.1
	github.com/tencentyun/cos-go-sdk-v5 v0.7.18
	github.com/uber/jaeger-client-go v2.29.0+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.uber.org/zap v1.16.0
	google.golang.org/grpc v1.36.0
	google.golang.org/protobuf v1.25.1-0.20201208041424-160c7477e0e8
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.11
	gorm.io/plugin/prometheus v0.0.0-20210112035011-ae3013937adc
)

//replace github.com/seosite/gcore => github.com/seosite/gcore v0.0.8 // 绝对路径 或 相对路径 都可以
