package oss

import (
	"bytes"
	"io/ioutil"
	"testing"
)

// 测试上传文件到aws
func TestUploadToAws(t *testing.T) {

	data, err := ioutil.ReadFile(`test.png`)
	if err != nil {
		t.Fatalf("reafile error, %s", err)
		return
	}

	fileUrl, err := UploadToAws("test.png", bytes.NewReader(data))
	if err != nil {
		t.Fatalf("UploadToAws error, %s", err)
		return
	}
	t.Logf("upload success, fileUrl=%s", fileUrl)
}
