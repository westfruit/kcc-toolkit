package sms

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

/*
 * 阿里大鱼短信
 */
type DySms struct {
}

// SendSmsReply 发送短信返回
type SendSmsReply struct {
	Code    string `json:"Code,omitempty"`
	Message string `json:"Message,omitempty"`
}

func (d DySms) replace(in string) string {
	rep := strings.NewReplacer("+", "%20", "*", "%2A", "%7E", "~")
	return rep.Replace(url.QueryEscape(in))
}

// SendSms 发送短信
func (d DySms) SendSms(accessKeyID, accessSecret, phoneNumbers, signName, templateParam, templateCode string) (*SendSmsReply, error) {
	paras := map[string]string{
		"SignatureMethod":  "HMAC-SHA1",
		"SignatureNonce":   fmt.Sprintf("%d", rand.Int63()),
		"AccessKeyId":      accessKeyID,
		"SignatureVersion": "1.0",
		"Timestamp":        time.Now().UTC().Format("2006-01-02T15:04:05Z"),
		"Format":           "JSON",
		"Action":           "SendSms",
		"Version":          "2017-05-25",
		"RegionId":         "cn-hangzhou",
		"PhoneNumbers":     phoneNumbers,
		"SignName":         signName,
		"TemplateParam":    templateParam,
		"TemplateCode":     templateCode,
	}

	reply := &SendSmsReply{}

	var keys []string

	for k := range paras {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	var sortQueryString string

	for _, v := range keys {
		sortQueryString = fmt.Sprintf("%s&%s=%s", sortQueryString, d.replace(v), d.replace(paras[v]))
	}

	stringToSign := fmt.Sprintf("GET&%s&%s", d.replace("/"), d.replace(sortQueryString[1:]))

	mac := hmac.New(sha1.New, []byte(fmt.Sprintf("%s&", accessSecret)))
	mac.Write([]byte(stringToSign))
	sign := d.replace(base64.StdEncoding.EncodeToString(mac.Sum(nil)))

	str := fmt.Sprintf("http://dysmsapi.aliyuncs.com/?Signature=%s%s", sign, sortQueryString)

	resp, err := http.Get(str)
	if err != nil {
		return reply, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return reply, err
	}

	if err := json.Unmarshal(body, reply); err != nil {
		return reply, err
	}

	if reply.Code == "SignatureNonceUsed" {
		return d.SendSms(accessKeyID, accessSecret, phoneNumbers, signName, templateParam, templateCode)
	} else if reply.Code != "OK" {
		return reply, errors.New(reply.Code)
	}

	return reply, nil
}
