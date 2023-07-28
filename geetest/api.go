package geetest

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"github.com/westfruit/kcc-toolkit/convert"
	"github.com/westfruit/kcc-toolkit/webreq"
	"strings"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

type GeetestClient struct {
	AppId      string
	AppKey     string
	ClientType string
}

// 验证初始化
func Register(param *RegisterParam) *GeetestResult {
	logrus.Info("开始注册geetest")

	result := new(GeetestResult)

	// user_id作为终端用户的唯一标识，确定用户的唯一性；作用于提供进阶数据分析服务，可在api1 或 api2 接口传入，
	// 不传入也不影响验证服务的使用；若担心用户信息风险，可作预处理(如哈希处理)再提供到极验
	params := map[string]interface{}{
		"gt":          AppId,
		"user_id":     param.UserId,
		"client_type": param.ClientType,
		"ip_address":  param.IPAddress, // 终端客户ip地址
	}

	resData, err := call(HttpMethodGet, RegisterUrl, params)
	if err != nil {
		logrus.Error("向geetest注册错误, ", err)
		result.Msg = "请求极验服务器错误"

		return result
	}

	logrus.Info("geetest注册结束, data=", resData)
	challenge := gjson.Get(resData, "challenge").String()

	// 返回数据
	d := map[string]interface{}{
		"success":     ResultStatusSuccess,
		"gt":          AppId,
		"challenge":   challenge,
		"new_captcha": NewCaptcha,
	}

	// 请求geetest服务器失败
	if challenge == "" || challenge == "0" {
		d["success"] = ResultStatusFailure
		d["challenge"] = GenerateUUID()

		result.Status = ResultStatusFailure
		result.Msg = "请求极验register接口失败，后续流程走宕机模式"
	} else {
		hash := ""

		if DigestMode == DigestModeSHA256 {
			hash = encodeBySHA256(challenge + AppKey)
		} else if DigestMode == DigestModeHMACSHA256 {
			hash = encodeByHMACSHA256(challenge, AppKey)
		} else {
			hash = encodeByMD5(challenge + AppKey)
		}

		d["challenge"] = hash
		result.Status = ResultStatusSuccess
	}

	// json化数据
	result.Data = convert.ObjToJson(d)

	return result
}

/**
 * 正常流程下（即验证初始化成功），二次验证
 */
func SuccessValidate(param *SuccessValidateParam) *GeetestResult {
	result := new(GeetestResult)

	logrus.Infof("SuccessValidate(): 开始二次验证 正常模式, challenge=%s, validate=%s, seccode=%s.", param.Challenge, param.Validate, param.Seccode)
	if !checkParam(param.Challenge, param.Validate, param.Seccode) {
		result.Msg = "正常模式，本地校验，参数challenge、validate、seccode不可为空"
		return result
	}

	// 二次验证
	params := map[string]interface{}{
		"captchaid":   AppId,
		"seccode":     param.Seccode,
		"challenge":   param.Challenge,
		"user_id":     param.UserId,
		"client_type": param.ClientType,
		"ip_address":  param.IPAddress, // 终端客户ip地址
	}

	logrus.Infof("SuccessValidate(): 二次验证 正常模式, 向极验发送请求, url=%s, params=%s.", ValidateUrl, convert.ObjToJson(params))
	resData, err := call(HttpMethodPost, ValidateUrl, params)
	if err != nil {
		logrus.Errorf("SuccessValidate(): 二次验证 正常模式, 请求异常, %s", err.Error())
		result.Msg = "请求极验服务器异常"
		return result
	}

	logrus.Infof("SuccessValidate(): 二次验证 正常模式, 与极验网络交互正常, 返回数据=%s.", resData)

	// 验证结果
	seccode := gjson.Get(resData, "seccode").String()
	if seccode == "" {
		result.Msg = "请求极验validate接口失败"
	} else if seccode == "false" {
		result.Msg = "极验二次验证不通过"
	} else {
		result.Status = ResultStatusSuccess
	}

	return result
}

/**
 * 异常流程下（即验证初始化失败，宕机模式），二次验证
 * 注意：由于是宕机模式，初衷是保证验证业务不会中断正常业务，所以此处只作简单的参数校验，可自行设计逻辑。
 */
func FailValidate(param *FailValidateParam) *GeetestResult {
	result := new(GeetestResult)

	logrus.Infof("FailValidate(): 开始二次验证 宕机模式, challenge=%s, validate=%s, seccode=%s.", param.Challenge, param.Validate, param.Seccode)

	if !checkParam(param.Challenge, param.Validate, param.Seccode) {
		result.Msg = "宕机模式，本地校验，参数challenge、validate、seccode不可为空."
	} else {
		result.Status = ResultStatusSuccess
	}

	logrus.Infof("FailValidate(): 二次验证 宕机模式, lib包返回信息=%s.", convert.ObjToJson(result))

	return result
}

// 发送请求
func call(method string, url string, params map[string]interface{}) (string, error) {

	// 基础参数
	params["json_format"] = JsonFormat
	params["sdk"] = SdkVersion
	params["digestmod"] = DigestMode

	requestURL := ApiUrl + url

	// 组装请求参数
	data := ""
	for k, v := range params {
		data += fmt.Sprintf("%s=%v&", k, v)
	}

	result := ""
	var err error

	if strings.EqualFold(method, "POST") {
		logrus.Info("向geetest发起POST请求, requestURL=", requestURL, ", data=", data)
		result, err = webreq.Post(requestURL, data)
	} else {
		if len(data) > 0 {
			requestURL = requestURL + "?" + data
		}

		logrus.Info("向geetest发起GET请求, url=", requestURL)
		result, err = webreq.Get(requestURL)
	}

	logrus.Info("geetest返回结=", result)
	if err != nil {
		logrus.Error("请求geetest错误, 请求地址: ", requestURL, ", 错误: ", err)
		return "", err
	}

	return result, err
}

/**
 * md5 加密
 */
func encodeByMD5(value string) string {
	h := md5.New()
	h.Write([]byte(value))
	return fmt.Sprintf("%x", h.Sum(nil))
}

/**
 * sha256加密
 */
func encodeBySHA256(value string) string {
	h := sha256.New()
	h.Write([]byte(value))
	return fmt.Sprintf("%x", h.Sum(nil))
}

/**
 * hmac-sha256 加密
 */
func encodeByHMACSHA256(value string, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(value))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// 生成一个不带减号的uuid
func GenerateUUID() string {
	id := strings.Replace(uuid.New().String(), "-", "", -1)
	return id
}

/**
 * 校验二次验证的三个参数，校验通过返回true，校验失败返回false
 */
func checkParam(challenge string, validate string, seccode string) bool {
	return !(strings.TrimSpace(challenge) == "" || strings.TrimSpace(validate) == "" || strings.TrimSpace(seccode) == "")
}
