package xcode

import (
	"io/ioutil"

	qrcode "github.com/skip2/go-qrcode"
)

// barcode and qrcode tool

//GenerateQRCode 生成二维码图片
func GenerateQRCode(text string, size int) ([]byte, error) {
	var png []byte
	png, err := qrcode.Encode(text, qrcode.Medium, size)
	if err != nil {
		return png, err
	}

	return png, nil
}

//GenerateQRCodeFile 生成二维码图片，并保存到指定路径
func GenerateQRCodeFile(text string, size int, filePath string) error {
	err := qrcode.WriteFile(text, qrcode.Medium, size, filePath)
	return err
}

//DrawQRCode 将二维码画到指定图片的指定位置,并设定二维码大小
func DrawQRCode(imgPath string, text string, size int, x int32, y int32) ([]byte, error) {
	imgData, err := ioutil.ReadFile(imgPath)
	if err != nil {
		return nil, err
	}

	return DrawQRCodeToImage(imgData, text, size, x, y)
}

//DrawQRCode 将二维码画到指定图片的指定位置,并设定二维码大小
func DrawQRCodeToImage(imgData []byte, text string, size int, x int32, y int32) ([]byte, error) {
	var data []byte

	return data, nil
}
