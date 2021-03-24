package wx

import (
	wechat "github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
)

var accounts = map[string]*officialaccount.OfficialAccount{}

// GetOfficialAccountByAppCode get wechat official account by app code
func GetOfficialAccountByAppCode(appCode string) (*officialaccount.OfficialAccount, error) {
	if account, ok := accounts[appCode]; ok {
		return account, nil
	}

	appInfo, err := GetAppInfo(appCode)
	if err != nil {
		return nil, err
	}
	tokenHandler, err := newAccessTokenHandlerByAppCode(appCode)
	if err != nil {
		return nil, err
	}
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &offConfig.Config{
		AppID:     appInfo.WxAppID,
		AppSecret: appInfo.WxAppSecret,
		Token:     appInfo.WxToken,
		Cache:     memory,
	}
	officialAccount := wc.GetOfficialAccount(cfg)
	officialAccount.SetAccessTokenHandle(tokenHandler)
	accounts[appCode] = officialAccount
	return accounts[appCode], nil
}

// GetOfficialAccountAccessToken 获取accesstoken
func GetOfficialAccountAccessToken(appCode string) (string, error) {
	oa, err := GetOfficialAccountByAppCode(appCode)
	if err != nil {
		return "", err
	}
	return oa.GetAccessToken()
}
