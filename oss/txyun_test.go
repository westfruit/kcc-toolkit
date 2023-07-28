package oss

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	txyunConf = TxyunConf{
		SecretId:   "",
		SecretKey:  "",
		BucketUrl:  "",
		UploadPath: "upload",
	}

	initTxyunClient(txyunConf)

	os.Exit(m.Run())
}

// 测试上传文件到aws
func TestUploadFileToTxyun(t *testing.T) {
	fileUrl, err := UploadFileToTxyun(`D:\download\images\test.jpg`)
	if err != nil {
		t.Errorf("UploadFileToTxyun error, %s", err.Error())
		return
	}

	fmt.Println("UploadFileToTxyun success, fileUrl=", fileUrl)
}

func TestUploadToTxyun(t *testing.T) {
	byteData, _ := ioutil.ReadFile(`D:\download\images\test.jpg`)

	fileUrl, err := UploadToTxyun("test.jpg", bytes.NewReader(byteData))
	if err != nil {
		t.Errorf("UploadFileToTxyun error, %s", err.Error())
		return
	}

	fmt.Println("UploadFileToTxyun success, fileUrl=", fileUrl)
}
