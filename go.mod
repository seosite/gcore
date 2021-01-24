module github.com/seosite/gcore

go 1.15

require (
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fsnotify/fsnotify v1.4.7
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ole/go-ole v1.2.5 // indirect
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.2
	github.com/jordan-wright/email v4.0.1-0.20210109023952-943e75fe5223+incompatible
	github.com/json-iterator/go v1.1.10
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible // indirect
	github.com/lestrrat-go/strftime v1.0.4 // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/prometheus/client_golang v1.9.0
	github.com/qiniu/api.v7/v7 v7.8.1
	github.com/shirou/gopsutil v3.20.12+incompatible
	github.com/spaolacci/murmur3 v1.1.0
	github.com/spf13/cast v1.3.1
	github.com/spf13/viper v1.7.1
	github.com/tencentyun/cos-go-sdk-v5 v0.7.18
	github.com/unrolled/secure v1.0.8
	go.uber.org/zap v1.16.0
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.11
	gorm.io/plugin/prometheus v0.0.0-20210112035011-ae3013937adc
)

//replace github.com/seosite/gcore => github.com/seosite/gcore v0.0.2 // 绝对路径 或 相对路径 都可以
