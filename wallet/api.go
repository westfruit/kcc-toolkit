package wallet

import (
	"fmt"
	"gitee.com/westfruit/kcc-toolkit/convert"
	"gitee.com/westfruit/kcc-toolkit/webreq"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
)

// 钱包工具类

const (
	// 回复1000表示成功
	WalletResCodeSuccess = 1000
)

// 申请地址
func ApplyAddress(symbol string) (string, string, error) {
	logrus.Info("开始申请地址, symbo=", symbol)

	url := "/open/ApplyAddress"
	params := map[string]string{
		"symbol": strings.ToUpper(symbol),
		"count":  "1",
	}

	data, err := call(url, params)
	logrus.Info("钱包服务申请地址返回结果, data=", data)

	if err != nil {
		logrus.Error("钱包服务申请地址错误, symbol=", symbol, ", err=", err)

		return "", "", err
	}

	// 获取地址
	addressArray := gjson.Get(data, "data").Array()
	if len(addressArray) == 0 {
		logrus.Error("请求地址返回数据错误, ", data)
		return "", "", fmt.Errorf("请求地址返回数据错误")
	}

	firstData := addressArray[0].Raw

	// 地址及memo
	address := gjson.Get(firstData, "address").String()
	memo := ""

	if gjson.Get(firstData, "isMemo").String() == "1" {
		memo = gjson.Get(firstData, "memo").String()
	}

	return address, memo, nil
}

// 申请提币
func ApplyTransaction(param *ApplyTransactionParam) error {
	url := "/open/ApplyTransaction"
	params := map[string]string{
		"amount":     fmt.Sprintf("%f", param.Amount),
		"symbol":     strings.ToUpper(param.Symbol),
		"contractId": param.ContractId,
		"to":         param.ToAddress,
		"memo":       param.Memo,
		"usid":       param.Usid,
	}

	data, err := call(url, params)
	logrus.Info("申请提币返回结果, data=", data)

	if err != nil {
		logrus.Error("申请提币错误, ", err)
		return err
	}

	return nil
}

// 发送请求
func call(url string, params map[string]string) (string, error) {
	apiUrl := viper.GetString("wallet.apiUrl")
	appKey := viper.GetString("wallet.appKey")
	appSecret := viper.GetString("wallet.appSecret")

	// 增加时间戳, 毫秒
	params["timestamp"] = fmt.Sprintf("%d", time.Now().UnixNano()/time.Millisecond.Nanoseconds())

	// 签名不转化值
	signStr := getParamString(params, false)
	sign := convert.GetHmacSHA256(appSecret, signStr)
	fmt.Println("url=" + url + ", signstr=" + signStr + ", sign=" + sign)

	// 请求
	headers := map[string]string{
		"appKey": appKey,
		"sign":   sign,
	}

	// 拼接url时，将过滤掉特殊字符, 钱包会先转换再验证签名
	requestUrl := apiUrl + url + "?" + getParamString(params, true)
	logrus.Info("向钱包发起请求, url=", requestUrl)

	result, err := webreq.GetWithHeaders(requestUrl, headers)
	logrus.Info("钱包服务返回结果, ", result)
	if err != nil {
		logrus.Error("请求钱包服务错误, 请求地址: ", requestUrl, ", 错误: ", err)
		return "", err
	}

	if !successResp(result) {
		return "", fmt.Errorf("请求钱包服务成功，返回数据错误, %s", gjson.Get(result, "message").String())
	}

	return result, err
}

// 取得排序后的参数
func getParamString(params map[string]string, escape bool) string {
	var out []string

	for k, v := range params {
		// 去掉空参数
		if strings.TrimSpace(k) == "" || strings.TrimSpace(v) == "" {
			continue
		}

		// 是否转化url里的值
		if escape {
			out = append(out, fmt.Sprintf("%s=%s", k, url.QueryEscape(v)))
		} else {
			out = append(out, fmt.Sprintf("%s=%s", k, v))
		}
	}

	// sort
	sort.Slice(out, func(i, j int) bool {
		return out[i] < out[j]
	})

	return strings.Join(out, "&")
}

// 是否是成功响应
func successResp(data string) bool {
	code := gjson.Get(data, "code").Int()
	return code == WalletResCodeSuccess
}
