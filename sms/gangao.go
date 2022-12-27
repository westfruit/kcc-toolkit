package sms

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"kcc/kcc-toolkit/convert"
	"kcc/kcc-toolkit/encode"
	http "kcc/kcc-toolkit/webreq"

	"github.com/thoas/go-funk"

	"github.com/spf13/viper"

	"net/url"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/html/charset"
)

const (
	GangaoParamSpid     = "spid"     // 客户代码，不能为空
	GangaoParamPwd      = "pwd"      // 密码，不能为空
	GangaoParamId       = "id"       // 客户短信ID，int64型,用于返回对应回执，保证唯一值
	GangaoParamEntityId = "entityid" // 机构ID，int型，可为空
	GangaoParamOwnerId  = "ownerid"  // 用户ID，int型，可为空
	GangaoParamMobiles  = "mobiles"  // 手机号码列表，不能为空，长度不限制，多个手机号码用英文逗号分隔
	GangaoParamSms      = "sms"      // 短信内容，不能为空，最大支持700字节，内容进行Base64编码，再URLENCODE编码
	GangaoParamExt      = "ext"      // 特服号扩展码，数字，可为空，不要超过6位
	GangaoParamPri      = "pri"      // 发送优先级别，0-9，数字越大越优先发送
	GangaoParamChannel  = "channel"  // 通道代码，可为空
	GangaoParamSeq      = "seq"      // 流水号，int64型，要求唯一值，预防重复发送。如果不用，可为空
)

// 港澳短信接口
type GangaoSms struct {
}

// 短信调用响应
type GangaoRes struct {
	Result int32  `json:"result" xml:"Result"`
	Desc   string `json:"desc" xml:"Desc"` // base64编码
}

// 获取返回的明文信息
func (r *GangaoRes) GetDescText() string {
	val, err := encode.Base64Decode(r.Desc)
	if err != nil {
		logrus.Error("base64编码错误: ", err)
		return ""
	}

	return val
}

func (s *GangaoSms) Name() string {
	return "gangao"
}

func (s *GangaoSms) Send(msg *Message) (*SmsResult, error) {

	// 获取一个随机接口地址
	apiUrls := viper.GetStringSlice("sms.gangao.apiUrls")

	postUrl := apiUrls[funk.RandomInt(0, len(apiUrls))]
	logrus.Info("短信请求接口：", postUrl)

	postData := s.getPostData(msg)

	result := new(SmsResult)

	data, err := http.Post(postUrl, postData)
	logrus.Info("短信请求url:,", postUrl, ", 结果: ", data)

	if err != nil {
		logrus.Error("gangao 发送短信错误, ", err)

		// 原始返回错误
		result.ResData = data
		return result, err
	}

	res := new(GangaoRes)

	// 处理xml编码
	decoder := xml.NewDecoder(bytes.NewReader([]byte(data)))
	decoder.CharsetReader = charset.NewReaderLabel

	err = decoder.Decode(res)
	if err != nil {
		logrus.Error("短信请求返回结果，XML格式化错误, ", err)

		return result, err
	}

	result.ResData = data
	descText := res.GetDescText()

	if res.Result == 0 {
		// 短信发送成功
		result.Success = true
		result.Message = descText

		return result, nil
	}

	logrus.Error("短信发送失败, code=", res.Result, ", 错误信息：", descText)

	return result, errors.New(descText)
}

func (s *GangaoSms) getPostData(msg *Message) string {

	username := viper.GetString("sms.gangao.username")
	password := viper.GetString("sms.gangao.password")

	// 为了防止传输产生乱码, 先将字符串转化为gbk编码与服务商端保持一致，内容先进行BASE64编码，然后再URLENCODE编码。
	content := url.QueryEscape(encode.Base64Encode(convert.ConvertStr2GBK(msg.Content)))

	data := fmt.Sprintf("spid=%s&pwd=%s&mobiles=%s&sms=%s&seq=%s&id=%s", username, password, msg.Phone, content, msg.SourceId, msg.SourceId)

	return data
}
