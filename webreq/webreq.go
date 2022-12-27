package webreq

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

func objToJson(obj interface{}) string {

	if obj == nil {
		return ""
	}

	b, err := json.Marshal(obj)
	if err != nil {
		log.Info("ObjToJson, error, ", err)
		return ""
	}

	return string(b)
}

func Get(url string) (string, error) {
	return doHttp(url, "GET", "", "", nil, nil)
}

func GetWithHeaders(url string, headers map[string]string) (string, error) {
	return doHttp(url, "GET", "", "", nil, headers)
}

func Post(url string, data string) (string, error) {
	return doHttp(url, "POST", "", data, nil, nil)
}

func PostJson(url string, data interface{}) (string, error) {
	return doHttp(url, "POST", "application/json", objToJson(data), nil, nil)
}

func PostWithAuth(url string, data string, username string, password string) (string, error) {
	return doHttp(url, "POST", "application/json", data, &HttpBasicAuth{
		Username: username,
		Password: password,
	}, nil)
}

func PostWithHeaders(url string, data string, headers map[string]string) (string, error) {
	return doHttp(url, "POST", "application/json", data, nil, headers)
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

	if headers != nil {
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
