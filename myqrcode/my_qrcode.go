package myqrcode

import (
	"encoding/base64"

	"github.com/skip2/go-qrcode"
)

//type QrCode struct{}

func Create(content string) string {

	var png []byte
	png, err := qrcode.Encode(content, qrcode.Medium, 256)
	if err != nil {
		panic(err)
	}

	imgStr := base64.StdEncoding.EncodeToString(png)

	imgStr = "data:image/png;base64," + imgStr

	return imgStr
}
