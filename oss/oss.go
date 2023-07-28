package oss

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
	// 阿里云
	OSSProviderTypeAliyun = 1

	// 亚马逊
	OSSProviderTypeAmazon = 2

	// 配置Key
	OSSConfKeyAliyun = "oss.aliyun"

	OSSConfKeyAws = "oss.aws"
)

// 上传配置
type UploadConf struct {
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
	Domain          string // 域名
	EndPoint        string
	UploadPath      string // oss上传路径
	Timeout         int32  // 上传超时
	Region          string // 地区
}

var (
	// aliyun
	aliConf UploadConf

	// aws
	awsConf  UploadConf
	uploader *s3manager.Uploader
)

func init() {
	// 初始化阿里云
	if viper.IsSet(OSSConfKeyAliyun) {
		if err := viper.UnmarshalKey(OSSConfKeyAliyun, &aliConf); err != nil {
			logrus.Error("初始化阿里云OSS错误, ", err)
		} else {
			logrus.Info("初始化阿里云OSS成功")
		}
	} else {
		logrus.Info("未配置阿里云OSS, 无需初始化阿里云OSS")
	}

	// 初始化aws
	if viper.IsSet(OSSConfKeyAws) {
		if err := viper.UnmarshalKey(OSSConfKeyAws, &awsConf); err != nil {
			logrus.Error("初始化亚马逊OSS错误, ", err)
		} else {
			sess := session.Must(session.NewSession(&aws.Config{
				Endpoint:    aws.String(awsConf.EndPoint),
				Region:      aws.String(awsConf.Region),
				Credentials: credentials.NewStaticCredentials(awsConf.AccessKeyId, awsConf.AccessKeySecret, ""),
			}))

			// S3 service client the Upload manager will use.
			svc := s3.New(sess)

			// Create an uploader with S3 client and default options
			uploader = s3manager.NewUploaderWithClient(svc)
			logrus.Info("初始化亚马逊OSS成功")
		}
	} else {
		logrus.Info("未配置亚马逊OSS, 无需初始化亚马逊OSS")
	}
}


