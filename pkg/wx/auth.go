package wx

import (
	"net/url"

	"github.com/seosite/gcore/pkg/app"
	"github.com/seosite/gcore/pkg/core/jsonx"
	"github.com/seosite/gcore/pkg/core/netx"
)

type getAuthAccessTokenResp struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenID       string `json:"openid"`
	Scope        string `json:"scope"`
}

// GetAuthOpenID 获取auth授权登录返回的openid
func GetAuthOpenID(appID, appSecret, code string) (string, error) {
	var (
		getAuthAccessTokenURL = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + appID + "&secret=" + appSecret + "&code=" + code + "&grant_type=authorization_code"
	)

	client := netx.NewRetryClient()
	resp, err := client.Get(getAuthAccessTokenURL)
	if err != nil {
		app.Logger.Error(err.Error())
		return "", err
	}
	var getAuthAccessTokenResp getAuthAccessTokenResp
	err = jsonx.Unmarshal([]byte(resp), &getAuthAccessTokenResp)
	if err != nil {
		app.Logger.Error(err.Error())
		return "", err
	}

	return getAuthAccessTokenResp.OpenID, nil
}

// GetH5AuthURL 前端进行微信认证，认证成功后跳转到目标地址
func GetH5AuthURL(appID, redirectURI, destURL string) string {
	return "https://open.weixin.qq.com/connect/oauth2/authorize?appid=" + appID +
		"&redirect_uri=" + redirectURI +
		"&response_type=code&scope=snsapi_base&state=" + url.QueryEscape(destURL)
}
