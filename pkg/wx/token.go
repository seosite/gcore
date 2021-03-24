package wx

import (
	"context"
	"errors"
	"time"

	"github.com/seosite/gcore/pkg/app"
	"github.com/seosite/gcore/pkg/core/rpcx"
	wxaccesstoken "github.com/seosite/gcore/pkg/core/third/proto/growth/wx_accesstoken"
)

var tokenClient wxaccesstoken.AccessTokenServiceClient

func getTokenCacheKey(key string) string {
	return app.Config.Server.Name + ":wxtoken:" + key
}

func getTokenRPCClient(address string) (wxaccesstoken.AccessTokenServiceClient, error) {
	if tokenClient == nil {
		grpcx, err := rpcx.NewGRPC(address)
		if err != nil {
			return nil, err
		}
		tokenClient = wxaccesstoken.NewAccessTokenServiceClient(grpcx.Conn)
	}
	return tokenClient, nil
}

func newAccessTokenHandlerByAppCode(appCode string) (*AccessTokenHandler, error) {
	address := app.Config.ThirdService.WxCenter.Address
	c, err := getTokenRPCClient(address)
	if err != nil {
		return nil, err
	}
	return &AccessTokenHandler{
		appCode:     appCode,
		tokenClient: c,
	}, nil
}

func newAccessTokenHandlerByAppIDSecret(appID, appSecret string) (*AccessTokenHandler, error) {
	address := app.Config.ThirdService.WxCenter.Address
	c, err := getTokenRPCClient(address)
	if err != nil {
		return nil, err
	}
	return &AccessTokenHandler{
		appID:       appID,
		appSecret:   appSecret,
		tokenClient: c,
	}, nil
}

// AccessTokenHandler .
type AccessTokenHandler struct {
	appCode     string
	appID       string
	appSecret   string
	tokenClient wxaccesstoken.AccessTokenServiceClient
}

// GetAccessToken 获取应用access token，本地缓存5分钟
func (h *AccessTokenHandler) GetAccessToken() (accessToken string, err error) {
	redisClient := app.DefaultRedis()
	cacheKey := getTokenCacheKey("code:" + h.appCode)
	cacheValue := redisClient.Get(cacheKey).Val()

	if len(cacheValue) == 0 {
		// 获取rpc数据
		if len(h.appCode) > 0 {
			// 通过appcode获取
			accessToken, err = h.getAccessTokenByAppCode()
		} else {
			// 通过app id & secret获取
			accessToken, err = h.getAccessTokenByAppIDSecret()
		}
		if err != nil {
			return "", err
		}
		// 更新缓存
		redisClient.Set(cacheKey, accessToken, time.Minute*5)
	} else {
		// 获取缓存数据
		accessToken = cacheValue
	}

	return accessToken, nil
}

func (h *AccessTokenHandler) getAccessTokenByAppCode() (accessToken string, err error) {
	r, err := h.tokenClient.GetByAppCode(context.Background(), &wxaccesstoken.GetByAppCodeRequest{
		AppCode: h.appCode,
	})
	if err != nil {
		app.Logger.Error(err.Error())
		return
	}
	if r.GetCode() != 0 {
		err = errors.New(r.Msg)
		app.Logger.Error(err.Error())
		return
	}
	accessToken = r.GetData()
	return
}

func (h *AccessTokenHandler) getAccessTokenByAppIDSecret() (accessToken string, err error) {
	r, err := h.tokenClient.GetByAppIDSecret(context.Background(), &wxaccesstoken.GetByAppIDSecretRequest{
		AppId:     h.appID,
		AppSecret: h.appSecret,
	})
	if err != nil {
		app.Logger.Error(err.Error())
		return
	}
	if r.GetCode() != 0 {
		err = errors.New(r.Msg)
		app.Logger.Error(err.Error())
		return
	}
	accessToken = r.GetData()
	return
}
