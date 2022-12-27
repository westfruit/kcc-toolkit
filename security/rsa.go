package security

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

//在线网站：http://web.chacuo.net/netrsakeypair

// 公钥和私钥可以从文件中读取
var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICWwIBAAKBgQCHN26n8EQr//gD9T8Keee/7303oPUuV2fE+HAx//DOgWzplc11
rZfpuFEerjEhPy62o7SdzWSLvt8bySDWey1PLFj1pzBaTsgNwvB+sbVIzrW9JHw5
BUTcZZD0G740ygF4OzKghMGvH5fOP4OuLGrAZ4pQ1j0KAq3VwiCSGH3a3wIDAQAB
AoGAMOkSM9krL6dFdVkO1qFF/R2J88dbKMohFRSwsMVdu7UBSnUPftOuMbKkVS65
QsdyBEqvGK2lAw+l8I0OPccMmhyTD94eIQePrbfCQrMv93qTPhXGIdES82CMBInR
d9lw2t/k6GdRASh8CdAXlwlNhDUokq0pwUVxS/TjGo5uR5ECQQCK5zDkSNS2bAq9
sNgI4H1UErudF1cYDOl9+d/BkvU+l6aGTSDAGAHIgy5UYysccw1/NeeF7U233hrc
Ld2yJDh1AkEA+TSkSXWaI1scyvpLkMOS+Rs2edmPEGSsWYaRJY5t7VSmWf5EyCsF
mEAIhAZo1TcdmAq4mjyHwND3t0j+jOo7gwJAGRgrfRKrW0mppxuL7A6ilc3Ml1Tg
Jzgt9tRt2Er0g5piO2EzyYCM0ezhMd24rCsdBR539Xop0E4QDDc9FjH1AQJAcFhs
C/XphDDqHROTh3KzynubJnmPBnvsOvvwaXb5VboIuJ9pkLfccNPFQwfEsQKE4+tJ
Z7cMGD/cvRM6NnrOUwJAb/9BSzGdQRwy+98H93APFEU/NRyFQ/0Q4R5DGAcUqiy1
ENAVVopSljU8ZaZmhvAeMszG0TwePPxv6+mH+0Ofgg==
-----END RSA PRIVATE KEY-----
`)

var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCHN26n8EQr//gD9T8Keee/7303
oPUuV2fE+HAx//DOgWzplc11rZfpuFEerjEhPy62o7SdzWSLvt8bySDWey1PLFj1
pzBaTsgNwvB+sbVIzrW9JHw5BUTcZZD0G740ygF4OzKghMGvH5fOP4OuLGrAZ4pQ
1j0KAq3VwiCSGH3a3wIDAQAB
-----END PUBLIC KEY-----
`)

// 加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

//加密为base64字符串
func RsaEncryptToBase64String(text string) string {

	data, _ := RsaEncrypt([]byte(text))

	s := base64.StdEncoding.EncodeToString(data)

	return s
}

//解密base64字符串为原始文本
func RsaDecryptToText(base64Str string) string {
	b, _ := base64.StdEncoding.DecodeString(base64Str)

	origData, _ := RsaDecrypt(b)

	return string(origData)
}
