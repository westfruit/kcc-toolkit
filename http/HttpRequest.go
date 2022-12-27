package http

import (
	"crypto/tls"
	"fmt"

	"github.com/astaxie/beego/logs"
	"golang.org/x/net/http2"

	"context"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func HttpPost(url string, data string) string {
	resp, err := http.Post(url,
		"application/json",
		strings.NewReader(data))
	if err != nil {
		fmt.Println(err)
		logs.Error("\r\n post error ", err)
		//defer resp.Body.Close()
		return ""
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return ""
	}

	//logs.Info(string(body))
	return string(body)
}

func HttpGet(url string) (string, error) {
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil) //建立一个请求
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
		return "", err
	}
	//Add 头协议
	//reqest.Header.Add("Accept", "text/html,application/json,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	reqest.Header.Add("Accept", "application/json")
	reqest.Header.Add("Accept-Language", "en,ja,zh-CN;q=0.8,zh;q=0.6")
	reqest.Header.Add("Connection", "keep-alive")
	reqest.Header.Set("Content-Type", "application/json")
	//reqest.Header.Add("Cookie", "cookie")
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0")
	response, err := client.Do(reqest) //提交
	defer response.Body.Close()
	/*cookies := response.Cookies() //遍历cookies
	for _, cookie := range cookies {
		fmt.Println("cookie:", cookie)
	}*/

	body, err1 := ioutil.ReadAll(response.Body)
	if err1 != nil {
		// handle error
		return "", err1
	}
	//fmt.Println(string(body))
	return string(body), nil
}

func HttpNewPost(url string, data string, headers map[string]string) (*http.Response, error) {
	tr := &http.Transport{ //解决x509: certificate signed by unknown authority
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		DisableKeepAlives: false,
	}
	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: tr, //解决x509: certificate signed by unknown authority
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Connection", "close")
	//defer req.Body.Close()

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	if err != nil {
		//log.Println(err.Error())
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		//log.Println(err.Error())
		return nil, err
	}
	return resp, nil
}

func HttpNewPostWithContext(ctx context.Context, url string, data string, headers map[string]string) (*http.Response, error) {
	tr := &http.Transport{ //解决x509: certificate signed by unknown authority
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		DisableKeepAlives: false,
	}
	client := &http.Client{
		Timeout:   15 * time.Second,
		Transport: tr, //解决x509: certificate signed by unknown authority
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, strings.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Add("Connection", "close")
	//defer req.Body.Close()

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	if err != nil {
		//log.Println(err.Error())
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		//log.Println(err.Error())
		return nil, err
	}
	return resp, nil
}

func HttpPostNoAuth(url string, data string) string {

	resp, err := HttpNewPost(url,
		data,
		nil)

	if err != nil {
		fmt.Println(err)
		logs.Error("\r\n post error ", err)
		//defer resp.Body.Close()
		return ""
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return ""
	}
	if resp != nil {
		resp.Close = true
	}
	resp.Body.Close()
	//resp.Close =true
	//logs.Info(string(body))
	return string(body)
}

func HttpPostNoAuthContext(ctx context.Context, url string, data string) string {

	resp, err := HttpNewPostWithContext(ctx, url,
		data,
		nil)

	if err != nil {
		fmt.Println(err)
		logs.Error("\r\n post error ", err)
		//defer resp.Body.Close()
		return ""
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return ""
	}
	resp.Close = true
	//logs.Info(string(body))
	return string(body)
}

func Http2PostNoAuth(url string, data string) string {

	resp, err := Http2NewPost(url,
		data,
		nil)

	if err != nil {
		fmt.Println(err)
		logs.Error("\r\n post error ", err)
		//defer resp.Body.Close()
		return ""
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return ""
	}
	if resp != nil {
		resp.Close = true
	}
	resp.Body.Close()
	//resp.Close =true
	//logs.Info(string(body))
	return string(body)
}

func Http2NewPost(url string, data string, headers map[string]string) (*http.Response, error) {
	tr := &http2.Transport{ //解决x509: certificate signed by unknown authority
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		//DisableKeepAlives:false,
	}
	client := &http.Client{
		Timeout:   15 * time.Second,
		Transport: tr, //解决x509: certificate signed by unknown authority
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Add("Connection", "close")
	//defer req.Body.Close()

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	if err != nil {
		//log.Println(err.Error())
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		//log.Println(err.Error())
		return nil, err
	}
	return resp, nil
}

func HttpGetWithHeader(url string, headers map[string]string) (string, error) {
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil) //建立一个请求
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
		return "", err
	}
	//Add 头协议
	//reqest.Header.Add("Accept", "text/html,application/json,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	//reqest.Header.Add("Accept", "application/json")
	//reqest.Header.Add("Accept-Language", "en,ja,zh-CN;q=0.8,zh;q=0.6")
	//reqest.Header.Add("Connection", "keep-alive")
	reqest.Header.Set("Content-Type", "application/json")
	//reqest.Header.Add("Cookie", "cookie")
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0")

	for k, v := range headers {
		reqest.Header.Add(k, v)
	}

	response, err := client.Do(reqest) //提交
	defer response.Body.Close()
	/*cookies := response.Cookies() //遍历cookies
	for _, cookie := range cookies {
		fmt.Println("cookie:", cookie)
	}*/

	body, err1 := ioutil.ReadAll(response.Body)
	if err1 != nil {
		// handle error
		return "", err1
	}
	//fmt.Println(string(body))
	return string(body), nil
}
