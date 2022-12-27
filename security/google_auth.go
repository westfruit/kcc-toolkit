package security

import (
	"bytes"
	"encoding/base64"
	"image/png"

	"github.com/pquerna/otp/totp"
)

type BindInfo struct {
	Url     string
	QrCode  string
	BindKey string
}

// 一、开启谷歌验证
// 基本工作流程：
// 1、为用户生成新的TOTP密钥。key,_ := totp.Generate(...)。
// 2、显示用户的密钥的密钥和QR码。key.Secret()和key.Image(...)。
// 3、测试用户是否可以成功使用他们的TOTP。totp.Validate(...)。
// 4、在后端存储用户的TOTP密码。 key.Secret()
// 5、为用户提供“恢复代码”。

// 二、验证
// 基本工作流程：
// 1、正常提示并验证用户密码。
// 2、如果用户启用了TOTP，则提示输入谷歌验证码。
// 3、从后端检索用户的TOTP密钥。
// 4、验证用户的密码。 totp.Validate(...)

//产生开启谷歌验证的信息
func GenerateBindInfo(account string, issuer string) *BindInfo {

	info := &BindInfo{}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: account,
	})

	// Convert TOTP key into a PNG
	var buf bytes.Buffer
	img, err := key.Image(200, 200)
	if err != nil {
		panic(err)
	}
	png.Encode(&buf, img)

	qrImg := base64.StdEncoding.EncodeToString(buf.Bytes())

	info.QrCode = qrImg
	info.BindKey = key.Secret()

	return info
}

//验证
func Validate(passcode string, secret string) bool {
	return totp.Validate(passcode, secret)
}
