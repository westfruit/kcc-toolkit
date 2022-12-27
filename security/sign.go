package security

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/sirupsen/logrus"
)

/*
检查签名
*/
func CheckSign(secret string, param interface{}, excludeField ...string) bool {

	//序列化
	b, err := json.Marshal(param)
	if err != nil {
		logrus.Error("检查签名,序列化异常:", err)
		return false
	}

	//反序列化
	mapResult := make(map[string]interface{})
	if err := json.Unmarshal(b, &mapResult); err != nil {
		logrus.Error("检查签名,反序列化异常:", err)
		return false
	}

	var oldSignStr string
	oldSign := mapResult["Sign"]
	if oldSign == nil {
		oldSign = mapResult["sign"]
	}

	if oldSign != nil {
		//oldSignStr = reflect.ValueOf(oldSign).String()
		oldSignStr = fmt.Sprintf("%v", oldSign)
	}
	newSign := Sign(secret, param, excludeField...)

	logrus.Info("输入签名:", oldSignStr)
	logrus.Info("验证签名:", newSign)

	return oldSignStr == newSign
}

/*
签名
*/
func Sign(secret string, param interface{}, excludeField ...string) string {

	var mapResult = make(map[string]interface{})

	jsonStr, ok := param.(string)
	if ok {

		err := json.Unmarshal([]byte(jsonStr), &mapResult)
		if err != nil {
			logrus.Error("签名,反序列化异常:", err)
			return ""
		}

	} else {

		//序列化
		b, err := json.Marshal(param)
		if err != nil {
			logrus.Error("签名,序列化异常:", err)
			return ""
		}

		//反序列化
		if err := json.Unmarshal(b, &mapResult); err != nil {
			logrus.Error("签名,反序列化异常:", err)
			return ""
		}
	}

	//参数排序
	var keys []string
	size := len(excludeField)
	for k := range mapResult {

		if strings.ToLower(k) == "sign" {
			continue
		}

		if size > 0 {
			for _, v := range excludeField {
				if k == v {
					continue
				}
			}
		}

		keys = append(keys, k)
	}
	sort.Strings(keys)

	//参数拼接
	var signString string
	signString += secret

	for _, vkey := range keys {
		value := mapResult[vkey]
		//valueStr := reflect.ValueOf(value).String()
		if value != nil {
			valueStr := fmt.Sprintf("%v", value)
			if valueStr != "" {
				signString += strings.ToLower(vkey)
				signString += valueStr
			}
		}
	}

	signString += secret

	return strings.ToUpper(Md5(signString))
}

/*
Md5加密
*/
func Md5(enstr string) string {
	h := md5.New()
	h.Write([]byte(enstr))

	return hex.EncodeToString(h.Sum(nil))
}
