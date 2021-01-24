package app

import (
	"github.com/seosite/gcore/pkg/core/env"
	"github.com/seosite/gcore/pkg/core/logx"
)

// Conf app config
type Conf struct {
	Server       ServerConf           `mapstructure:"server" json:"server" yaml:"server"`
	Log          LogConf              `mapstructure:"log" json:"log" yaml:"log"`
	Mysql        map[string]MysqlConf `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis        map[string]RedisConf `mapstructure:"redis" json:"redis" yaml:"redis"`
	Cache        CacheConf            `mapstructure:"cache" json:"cache" yaml:"cache"`
	Cos          CosConf              `mapstructure:"cos" json:"cos" yaml:"cos"`
	JWT          JWTConf              `mapstructure:"jwt_conf" json:"jwt_conf" yaml:"jwt_conf"`
	ThirdService ThirdServiceConf     `mapstructure:"thirdService" json:"thirdService" yaml:"thirdService"`
}

// ServerConf server config
type ServerConf struct {
	Name         string   `mapstructure:"name" json:"name" yaml:"name"`
	Port         int      `mapstructure:"port" json:"port" yaml:"port"`
	Env          env.Type `mapstructure:"env" json:"env" yaml:"env"`
	AllowOrigins []string `mapstructure:"allowOrigins" json:"allowOrigins" yaml:"allowOrigins"`
	AlertUsers   []string `mapstructure:"alertUsers" json:"alertUsers" yaml:"alertUsers"`
}

// LogConf log config
type LogConf struct {
	Mod logx.Mod `mapstructure:"mod" json:"mod" yaml:"mod"`
}

// MysqlConf mysql config
type MysqlConf struct {
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	Dbname       string `mapstructure:"dbName" json:"dbName" yaml:"dbName"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	MaxIdleConns int    `mapstructure:"maxIdleConns" json:"maxIdleConns" yaml:"maxIdleConns"`
	MaxOpenConns int    `mapstructure:"maxOpenConns" json:"maxOpenConns" yaml:"maxOpenConns"`
}

// RedisConf redis config
type RedisConf struct {
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
}

// CacheConf cache config
type CacheConf struct {
	Size   int `mapstructure:"size" json:"size" yaml:"size"`
	Expire int `mapstructure:"expire" json:"expire" yaml:"expire"`
}

// CosConf cos config
type CosConf struct {
	SecretID  string `mapstructure:"secretId" json:"secretId" yaml:"secretId"`
	SecretKey string `mapstructure:"secretKey" json:"secretKey" yaml:"secretKey"`
	Region    string `mapstructure:"region" json:"region" yaml:"region"`
	Bucket    string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	RootPath  string `mapstructure:"rootPath" json:"rootPath" yaml:"rootPath"`
	Timeout   int    `mapstructure:"timeout" json:"timeout" yaml:"timeout"`
}

type JWTConf struct {
	IsOpen       int64  `mapstructure:"isOpen" json:"isOpen" yaml:"isOpen"`
	JwtID        string `mapstructure:"jwtID" json:"jwtID" yaml:"jwtID"`
	JwtTokenName string `mapstructure:"jwtTokenName" json:"jwtTokenName" yaml:"jwtTokenName"`
	SigningKey   string `mapstructure:"signingKey" json:"signingKey" yaml:"signingKey"`
	ExpiresTime  int64  `mapstructure:"expires-time" json:"expiresTime" yaml:"expires-time"`
	BufferTime   int64  `mapstructure:"buffer-time" json:"bufferTime" yaml:"buffer-time"`
}

// ThirdServiceConf third service domain config
type ThirdServiceConf struct {
	Sso       ThirdSsoConf       `mapstructure:"sso" json:"sso" yaml:"sso"`
	Analytics ThirdAnalyticsConf `mapstructure:"analytics" json:"analytics" yaml:"analytics"`
}

// ThirdSsoConf sso config
type ThirdSsoConf struct {
	Domain string `mapstructure:"domain" json:"domain" yaml:"domain"`
}

// ThirdAnalyticsConf analytics config
type ThirdAnalyticsConf struct {
	Domain   string `mapstructure:"domain" json:"domain" yaml:"domain"`
	Version  string `mapstructure:"version" json:"version" yaml:"version"`
	AppID    int    `mapstructure:"appid" json:"appid" yaml:"appid"`
	Platform int    `mapstructure:"platform" json:"platform" yaml:"platform"`
}
