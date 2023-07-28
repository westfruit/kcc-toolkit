package udun

import (
	"fmt"
	"gitee.com/westfruit/kcc-toolkit/convert"
	"gitee.com/westfruit/kcc-toolkit/security"
	"gitee.com/westfruit/kcc-toolkit/webreq"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// 提币
func Withdraw(param *WithdrawParam) error {

	// 提币地址
	url := "/mch/withdraw"
	params := map[string]string{
		"address":      param.Address,
		"amount":       param.Amount,
		"merchantId":   viper.GetString("udun.merchantId"),
		"mainCoinType": fmt.Sprintf("%d", param.MainCoinType),
		"coinType":     fmt.Sprintf("%d", param.CoinType),
		"callUrl":      viper.GetString("udun.callbackUrl"),
		"businessId":   param.BusinessId,
		"memo":         param.Memo,
	}

	result, err := call(url, params)
	if err != nil {
		logrus.Error("申请提币错误, ", err)
		return err
	}

	logrus.Info("申请提币返回结果, result=", convert.ObjToJson(result))
	if result.Code != ApiResultCodeSuccess {
		return fmt.Errorf("提币失败, 错误信息: %s", result.Message)
	}

	return nil
}

// 获取主币币种
func getMainCoinType(token string) string {
	return viper.GetString("udun." + strings.ToLower(token) + ".mainCoinType")
}

// 获取子币币种
func getCoinType(token string) string {
	return viper.GetString("udun." + strings.ToLower(token) + ".coinType")
}

// 获取时间戳
func getTimestmap() int64 {
	return time.Now().Unix()
}

// 获取随机数
func getNonce() int64 {
	return 0
}

// 发送请求
func call(url string, params map[string]string) (*ApiResult, error) {
	apiUrl := viper.GetString("udun.apiUrl")

	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	nonce := fmt.Sprintf("%d", time.Now().UnixNano())

	// 增加时间戳, 毫秒
	data := make(map[string]string, 4)
	data["timestamp"] = timestamp
	data["nonce"] = nonce

	body := convert.ObjToJson([]interface{}{params})
	data["body"] = body
	data["sign"] = Sign(body, nonce, timestamp)

	// 拼接url时，将过滤掉特殊字符, 钱包会先转换再验证签名
	requestUrl := apiUrl + url
	logrus.Info("向钱包发起请求, url=", requestUrl)

	res, err := webreq.PostJson(requestUrl, data)
	logrus.Info("钱包服务返回结果, ", data)
	if err != nil {
		logrus.Error("请求钱包服务错误, 请求地址: ", requestUrl, ", 错误: ", err)
		return nil, err
	}

	var result ApiResult
	if err := convert.JsonToObj(res, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// 签名
func Sign(body string, nonce string, timestamp string) string {
	signature := security.Md5(body + viper.GetString("udun.apiKey") + nonce + timestamp)
	return signature
}
