package http

import (
	//"git.bemular.net/bemular2/bml-toolkit/convert"
	//"geos/geos-toolkit/convert"
	"io/ioutil"
	"net/http"
	"strings"

	"kcc/kcc-toolkit/convert"
)

func Get(url string) (string, error) {
	return doHttp(url, "Get", "", "", nil, nil)
}

func GetWithHeaders(url string, headers map[string]string) (string, error) {
	return doHttp(url, "Get", "", "", nil, headers)
}

//发送post请求
func Post(url string, data string) (string, error) {
	return doHttp(url, "POST", "", data, nil, nil)
}

func PostJson(url string, data interface{}) (string, error) {
	return doHttp(url, "POST", "application/json", convert.ObjToJson(data), nil, nil)
}

//发送post请求，获取商家token
func PostWithAuth(url string, data string, username string, password string) (string, error) {

	return doHttp(url, "POST", "application/json", data, &HttpBasicAuth{
		Username: username,
		Password: password,
	}, nil)
}

func PostWithAuth2(url string, data interface{}, username string, password string) (string, error) {
	return PostWithAuth(url, convert.ObjToJson(data), username, password)
}

func PostWithHeaders(url string, data string, headers map[string]string) (string, error) {
	return doHttp(url, "POST", "application/json", data, nil, headers)
}

func PostWithHeaders2(url string, data interface{}, headers map[string]string) (string, error) {
	return PostWithHeaders(url, convert.ObjToJson(data), headers)
}

func doHttp(
	url string,
	method string,
	contentType string,
	data string,
	auth *HttpBasicAuth,
	headers map[string]string) (string, error) {
	client := &http.Client{}

	var req *http.Request
	var err error

	if len(data) > 0 {
		req, err = http.NewRequest(method, url, strings.NewReader(data))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return "", err
	}

	if method == "POST" && len(contentType) > 0 {
		req.Header.Set("Content-Type", contentType)
	} else {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	if auth != nil {
		req.SetBasicAuth(auth.Username, auth.Password)
	}

	if headers != nil && len(headers) > 0 {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

type HttpBasicAuth struct {
	Username string
	Password string
}
