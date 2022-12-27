package google_ency

import (
	"crypto/rand"
	"fmt"

	"github.com/jmolboy/googleauth"
)

type RandType int8

const (
	RandTypeAlphaNum RandType = 1
	RandTypeAlpha    RandType = 2
	RandTypeNum      RandType = 3
)

/**
fun:获取google验证密码
*/
func GetCode() string {

	//code := randoms.Krand(8, randoms.KC_RAND_KIND_ALL)

	//return string(code)
	return randSecret(16, RandTypeAlpha)
}
func CheckGoolgeCode(code string, secret string) bool {
	//secret := "dfdfdfdf"
	/*ga := googleAuthenticator.NewGAuth()
	value,err :=ga.GetCode(secret)
	if err!=nil{
		return false
	}*/
	value, err := googleauth.GetCode(secret)
	if err != nil {
		fmt.Println("code=" + code + ", secret=" + secret + ", googleauth.GetCode error, " + err.Error())
		return false
	}

	fmt.Println("googleauth.GetCode, value=" + value + ", code=" + code)
	return value == code
}

func randSecret(strSize int, randType RandType) string {
	var dictionary string

	if randType == RandTypeAlphaNum {
		dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	if randType == RandTypeAlpha {
		dictionary = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	if randType == RandTypeNum {
		dictionary = "0123456789"
	}

	var bytes = make([]byte, strSize)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	fmt.Println(string(bytes))
	//使用base32算法
	return string(bytes)
}
