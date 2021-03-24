package wxtypes

import "time"

// ------ 宏变量 ------

const (
	MessageDynamicTagPrefix = "{#"
	MessageDynamicTagSuffix = "#}"
	MessageDynamicNickname  = "nickname"
)

// ------ 消息基础结构 ------

// 消息类型
const (
	MessageTypeText      = 1 // 文字消息
	MessageTypeMedia     = 2 // 语音消息
	MessageTypeVideo     = 3 // 视频消息
	MessageTypeMusic     = 4 // 音乐消息
	MessageTypeNews      = 5 // 网页图文消息
	MessageTypeMpNews    = 6 // 公众号图文消息
	MessageTypeMenu      = 7 // 菜单消息
	MessageTypeCard      = 8 // 卡券消息
	MessageTypeTextMedia = 9 // 文字+图片消息
)

// 消息状态
const (
	MessageStatusDisable = 0 // 关闭
	MessageStatusEnable  = 1 // 启动
)

// TextMessageContent 文本消息
type TextMessageContent struct {
	Content string `json:"content"`
}

// ------ 应用 ------

// AppConfig 应用配置
type AppConfig struct {
	AppCodes []string `json:"app_codes"`
}

// ------ 粉丝筛选 ------

// 消息粉丝筛选
const (
	MessageUserConfigAll            = 1 // 留存粉丝，需要选择群发
	MessageUserConfigActive         = 2 // 可触达粉丝
	MessageUserConfigBeforeUnactive = 3 // 临触达粉丝
	MessageUserConfigGender         = 4 // 按性别筛选
	MessageUserConfigRegion         = 5 // 按地区筛选
	MessageUserConfigSubscribe      = 6 // 按关注时间
	MessageUserConfigBirthday       = 7 // 按生日筛选
	MessageUserConfigConstellation  = 8 // 按星座筛选
)

// UserConfigAll 筛选粉丝配置，全部
type UserConfigAll struct {
}

// UserConfigBeforeUnactive 筛选粉丝配置，临触达
type UserConfigBeforeUnactive struct {
	StartTime int32 `json:"start_time"`
	EndTime   int32 `json:"end_time"`
}

// UserConfigGender 筛选粉丝配置，性别
type UserConfigGender struct {
	Gender string `json:"gender"`
}

// UserConfigSubscribe 筛选粉丝配置，关注时间
type UserConfigSubscribe struct {
	StartTime *time.Time `json:"start_time"`
	EndTime   *time.Time `json:"end_time"`
}

// UserConfigConstellation 筛选粉丝配置，星座
type UserConfigConstellation struct {
	Constellation string `json:"constellation"`
}

// ------ 发送时间 ------

// 消息发送时间
const (
	MessageTimeConfigTypeOnce   = 1 // 单次发送
	MessageTimeConfigTypeRepeat = 2 // 重复发送
)

// TimeConfig 发送时间配置
type TimeConfig struct {
	Type    int32       `json:"type"`
	Content interface{} `json:"content"`
}

// TimeConfigOnce 发送时间配置，单次发送
type TimeConfigOnce struct {
	Time *time.Time `json:"time"`
}

// TimeConfigRepeat 发送时间配置，重发发送
type TimeConfigRepeat struct {
	Type  int32       `json:"type"`
	Value interface{} `json:"value"`
}

// 消息发送时间类型
const (
	MessageTimeConfigSchedule       = 1 // 定时任务，cron格式，可以每天定时或者特定时间
	MessageTimeConfigAfterSubscribe = 2 // 新关注后
	MessageTimeConfigBeforeUnactive = 3 // 临触达前
)
