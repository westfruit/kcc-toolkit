package sms

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Sms interface {
	// 短信接口名称
	Name() string
	Send(msg *Message) (*SmsResult, error)
}

type Message struct {
	SourceId  string `json:"sourceId"`  // 短信的源Id, 用于追踪
	Phone     string `json:"phone"`     // 接收短信的电话
	Content   string `json:"content"`   // 内容
	Signature string `json:"signature"` // 短信签名
}

// 短信发送接口
type SmsResult struct {
	MessageId string `json:"messageId"` // 短信平台商返回的消息Id
	Message   string `json:"message"`
	Success   bool   `json:"success"`
	ResData   string `json:"resData"` // 原始返回数据
}

func init() {

}

func New() Sms {

	var sms Sms

	provider := viper.GetString("sms.provider")
	logrus.Info("当前使用短信网关: ", provider)

	switch provider {
	case "gangao":
		sms = new(GangaoSms)
	case "awssns":
		sms = new(AwsSms)
	}

	return sms
}
