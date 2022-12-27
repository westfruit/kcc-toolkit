package sms

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/spf13/viper"
)

var (
	awsConf *aws.Config
)

func init() {

	accessKeyId := viper.GetString("sms.awssns.accessKeyId")
	accessKeySecret := viper.GetString("sms.awssns.accessKeySecret")
	region := viper.GetString("sms.awssns.region")

	awsConf = &aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKeyId, accessKeySecret, ""),
	}
}

// aws sms
type AwsSms struct {
}

func (s *AwsSms) Name() string {
	return "awssns"
}

func (s *AwsSms) Send(msg *Message) (*SmsResult, error) {
	result := new(SmsResult)

	sess := session.Must(session.NewSession(awsConf))

	snsClient := sns.New(sess)

	params := &sns.PublishInput{
		Message:     aws.String(msg.Signature + msg.Content),
		PhoneNumber: aws.String(msg.Phone),
	}

	res, err := snsClient.Publish(params)
	if err != nil {
		return nil, err
	}

	result.Success = true
	result.MessageId = *res.MessageId
	result.ResData = res.String()

	return result, nil
}
