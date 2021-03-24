package wx

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/seosite/gcore/pkg/core/jsonx"
)

const (
	// QrcodeTypeExpire 临时二维码
	QrcodeTypeExpire = 1
	// QrcodeTypePermanent 永久二维码
	QrcodeTypePermanent = 2
)

// qrTicketResp 微信二维码ticket返回
type qrTicketResp struct {
	ExpireSeconds int64  `json:"expire_seconds"`
	Ticket        string `json:"ticket"`
	URL           string `json:"url"`
}

// NewQrcodeSceneStr 新建自定义字符串二维码
func NewQrcodeSceneStr(appCode, sceneStr string, qrType int64) (string, error) {
	accessToken, err := GetOfficialAccountAccessToken(appCode)
	if err != nil {
		return "", err
	}
	config := map[string]interface{}{
		"action_info": map[string]interface{}{
			"scene": map[string]interface{}{
				"scene_str": sceneStr,
			},
		},
	}
	// 判断二维码类型
	switch qrType {
	case QrcodeTypePermanent:
		config["action_name"] = "QR_LIMIT_STR_SCENE"
	default:
		config["action_name"] = "QR_STR_SCENE"
		config["expire_seconds"] = 2592000 // 30day
	}
	// get ticket
	body, err := jsonx.Marshal(config)
	if err != nil {
		return "", err
	}
	resp, err := http.Post("https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token="+accessToken,
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var qrTicket qrTicketResp
	err = jsonx.Unmarshal(bodyBytes, &qrTicket)
	if err != nil {
		return "", err
	}
	// get qrcode
	return "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=" + qrTicket.Ticket, nil
}
