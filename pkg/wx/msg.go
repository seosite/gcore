package wx

import (
	"context"
	"errors"

	"github.com/seosite/gcore/pkg/app"
	"github.com/seosite/gcore/pkg/core/jsonx"
	"github.com/seosite/gcore/pkg/core/rpcx"
	"github.com/seosite/gcore/pkg/core/third/proto/growth/wx_notify"
	"github.com/silenceper/wechat/v2/officialaccount/message"
)

var notifyClient wx_notify.WxNotifyServiceClient

func getNotifyRPCClient(address string) (wx_notify.WxNotifyServiceClient, error) {
	if notifyClient == nil {
		grpcx, err := rpcx.NewGRPC(address)
		if err != nil {
			return nil, err
		}
		notifyClient = wx_notify.NewWxNotifyServiceClient(grpcx.Conn)
	}
	return notifyClient, nil
}

// GetCustomerMessageManagerByAppCode 根据appcode获取客服消息管理器
func GetCustomerMessageManagerByAppCode(appCode string) (*message.Manager, error) {
	oa, err := GetOfficialAccountByAppCode(appCode)
	if err != nil {
		return nil, err
	}
	return oa.GetCustomerMessageManager(), nil
}

// SendTextMsgAsync send custom text messag async
func SendTextMsgAsync(appCode, openID, msg string) (string, error) {
	content := map[string]interface{}{
		"touser":  openID,
		"msgtype": "text",
		"text": map[string]interface{}{
			"content": msg,
		},
	}
	return sendMsg(appCode, content, true)
}

// SendImageMsgAsync send custom image messag async
func SendImageMsgAsync(appCode, openID, mediaID string) (string, error) {
	content := map[string]interface{}{
		"touser":  openID,
		"msgtype": "image",
		"image": map[string]interface{}{
			"media_id": mediaID,
		},
	}
	return sendMsg(appCode, content, true)
}

func sendMsg(appCode string, content map[string]interface{}, async bool) (string, error) {
	address := app.Config.ThirdService.WxNotify.Address
	notifyClient, err := getNotifyRPCClient(address)
	if err != nil {
		return "", err
	}
	r, err := notifyClient.SendCustomMsg(context.Background(), &wx_notify.WxCustmoMsgRequest{
		AppCode: appCode,
		Message: jsonx.MarshalToString(content),
		Async:   async,
	})
	if err != nil {
		return "", err
	}
	if r.Code != 0 {
		return "", errors.New(r.Msg)
	}
	return r.Msg, nil
}
