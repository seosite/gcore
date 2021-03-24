package wx

import (
	"context"
	"errors"
	"time"

	"github.com/seosite/gcore/pkg/app"
	"github.com/seosite/gcore/pkg/core/jsonx"
	"github.com/seosite/gcore/pkg/core/rpcx"
	wxapp "github.com/seosite/gcore/pkg/core/third/proto/growth/wx_app"
	"github.com/spf13/cast"
)

var appClient wxapp.AppServiceClient

func getAppCacheKey(key string) string {
	return app.Config.Server.Name + ":wxapp:" + key
}

func getAppRPCClient(address string) (wxapp.AppServiceClient, error) {
	if appClient == nil {
		grpcx, err := rpcx.NewGRPC(address)
		if err != nil {
			return nil, err
		}
		appClient = wxapp.NewAppServiceClient(grpcx.Conn)
	}
	return appClient, nil
}

// GetAppInfo 获取应用信息，本地缓存10分钟
func GetAppInfo(appCode string) (*wxapp.App, error) {
	var appInfo *wxapp.App

	if appCode == "" {
		return nil, errors.New("Invalid app code")
	}

	redisClient := app.DefaultRedis()
	cacheKey := getAppCacheKey("code:" + appCode)
	cacheValue := redisClient.Get(cacheKey).Val()

	if len(cacheValue) == 0 {
		// 获取rpc数据
		address := app.Config.ThirdService.WxCenter.Address
		appClient, err := getAppRPCClient(address)
		if err != nil {
			return nil, err
		}
		r, err := appClient.GetOne(context.Background(), &wxapp.App{
			Code: appCode,
		})
		if err != nil {
			return nil, err
		}
		if r.GetCode() != 0 {
			return nil, errors.New(r.GetMsg())
		}
		appInfo = r.GetData()
		// 更新缓存
		redisClient.Set(cacheKey, jsonx.MarshalToString(appInfo), time.Minute*10)
	} else {
		// 获取缓存数据
		if err := jsonx.Unmarshal([]byte(cacheValue), &appInfo); err != nil {
			return nil, err
		}
	}

	return appInfo, nil
}

// GetCodeNameList 获取应用名称code列表，本地缓存10分钟
func GetCodeNameList(platform int64) (map[string]string, error) {
	var appList map[string]string

	redisClient := app.DefaultRedis()
	cacheKey := getAppCacheKey("platform:" + cast.ToString(platform))
	cacheValue := redisClient.Get(cacheKey).Val()

	if len(cacheValue) == 0 {
		// 获取rpc数据
		address := app.Config.ThirdService.WxCenter.Address
		appClient, err := getAppRPCClient(address)
		if err != nil {
			return nil, err
		}
		r, err := appClient.GetCodeNameList(context.Background(), &wxapp.GetCodeNameListRequest{
			Platform: int32(platform),
		})
		if err != nil {
			return nil, err
		}
		if r.GetCode() != 0 {
			return nil, errors.New(r.GetMsg())
		}
		appList = r.GetData()
		// 更新缓存
		redisClient.Set(cacheKey, jsonx.MarshalToString(appList), time.Minute*10)
	} else {
		// 获取缓存数据
		if err := jsonx.Unmarshal([]byte(cacheValue), &appList); err != nil {
			return nil, err
		}
	}

	return appList, nil
}
