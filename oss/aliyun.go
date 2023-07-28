package oss

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//说明文档：https://help.aliyun.com/document_detail/88601.html
//var datetime string = time.Now().UTC().Format(http.TimeFormat)

var (
	AccessKeyID     string
	AccessKeySecret string
	BucketName      string
	BucketUrl       string
	EndPoint        string
	PathPrefix      string
)

func init() {
	AccessKeySecret = viper.GetString("oss.aliyun.accessKeySecret")
	AccessKeyID = viper.GetString("oss.aliyun.accessKeyId")
	BucketName = viper.GetString("oss.aliyun.bucketName")
	BucketUrl = viper.GetString("oss.aliyun.domain")
	EndPoint = viper.GetString("oss.aliyun.endPoint")
	PathPrefix = viper.GetString("oss.aliyun.prefix")
}

//上传文件流
func FileUpload(fileName string, r io.Reader) (string, error) {
	prefix := "【上传文件流】"

	aliClient, err := oss.New(aliConf.EndPoint, aliConf.AccessKeyId, aliConf.AccessKeySecret)
	if err != nil {
		logrus.Error(prefix, "创建OSSClient实例异常：", err)
		return "", err
	}

	// 获取存储空间
	bucket, err := aliClient.Bucket(aliConf.BucketName)
	if err != nil {
		logrus.Error(prefix, "获取存储空间异常：", err)
		return "", err
	}

	// 上传文件
	filePath := aliConf.UploadPath + "/" + fileName
	err = bucket.PutObject(filePath, r)
	if err != nil {
		logrus.Error(prefix, "上传文件流异常：", err)
		return "", err
	}

	fileUrl := aliConf.Domain + "/" + filePath
	logrus.Info(prefix, "文件上传成功, 文件地址:", fileUrl)

	return fileUrl, nil
}

//上传字符串
func Base64Upload(key string, value string) (string, error) {
	// 创建OSSClient实例
	client, err := oss.New(EndPoint, AccessKeyID, AccessKeySecret)
	if err != nil {
		logrus.Error("创建OSSClient实例异常：", err)
		return "", err
	}

	// 获取存储空间
	bucket, err := client.Bucket(BucketName)
	if err != nil {
		logrus.Error("获取存储空间异常：", err)
		return "", err
	}

	// 上传字符串
	base64Bytes, _ := base64.StdEncoding.DecodeString(value) //成图片文件并把文件写入到buffer
	buffer := bytes.NewBuffer(base64Bytes)                   // 必须加一个buffer 不然没有read方法就会报错

	err = bucket.PutObject(key, buffer)
	if err != nil {
		logrus.Error("上传字符串异常：", err)
		return "", err
	}
	filePath := fmt.Sprintf("%s/%s", BucketUrl, key)
	return filePath, nil
}

//文件流方式上传（本地不存文件）
func StreamUpload(ossFileName string, r io.Reader) (string, error) {
	prefix := "【上传文件流】"
	// 创建OSSClient实例
	client, err := oss.New(EndPoint, AccessKeyID, AccessKeySecret)
	if err != nil {
		logrus.Error(prefix, "创建OSSClient实例异常：", err)
		return "", err
	}

	// 获取存储空间
	bucket, err := client.Bucket(BucketName)
	if err != nil {
		logrus.Error(prefix, "获取存储空间异常：", err)
		return "", err
	}

	// 上传文件
	filePath := PathPrefix + "/" + ossFileName
	err = bucket.PutObject(filePath, r)
	if err != nil {
		logrus.Error(prefix, "上传文件流异常：", err)
		return "", err
	}

	filePath = BucketUrl + "/" + filePath

	return filePath, nil
}

//上传字符串
func AliyunUploadString(key string, value string) error {
	// 创建OSSClient实例
	client, err := oss.New(EndPoint, AccessKeyID, AccessKeySecret)
	if err != nil {
		logrus.Error("创建OSSClient实例异常：", err)
		return err
	}

	// 获取存储空间
	bucket, err := client.Bucket(BucketName)
	if err != nil {
		logrus.Error("获取存储空间异常：", err)
		return err
	}

	// 上传字符串
	err = bucket.PutObject(key, strings.NewReader(value))
	if err != nil {
		logrus.Error("上传字符串异常：", err)
		return err
	}

	return nil
}

//上传文件流
func AliyunUploadStream(ossFileName string, r io.Reader) (string, error) {
	prefix := "【上传文件流】"
	// 创建OSSClient实例
	client, err := oss.New(EndPoint, AccessKeyID, AccessKeySecret)
	if err != nil {
		logrus.Error(prefix, "创建OSSClient实例异常：", err)
		return "", err
	}

	// 获取存储空间
	bucket, err := client.Bucket(BucketName)
	if err != nil {
		logrus.Error(prefix, "获取存储空间异常：", err)
		return "", err
	}

	// 上传文件
	filePath := PathPrefix + "/" + ossFileName
	err = bucket.PutObject(filePath, r)
	if err != nil {
		logrus.Error(prefix, "上传文件流异常：", err)
		return "", err
	}

	filePath = BucketUrl + "/" + filePath

	return filePath, nil
}
