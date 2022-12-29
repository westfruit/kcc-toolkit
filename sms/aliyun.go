package sms

import (
"os"
dysmsapi20170525  "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
openapi  "github.com/alibabacloud-go/darabonba-openapi/client"
"github.com/alibabacloud-go/tea/tea"
)


/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient (accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func _main (args []*string) (_err error) {
	client, _err := CreateClient(tea.String("accessKeyId"), tea.String("accessKeySecret"))
	if _err != nil {
		return _err
	}

	querySendDetailsRequest := &dysmsapi20170525.QuerySendDetailsRequest{
		ResourceOwnerAccount: tea.String("test"),
		ResourceOwnerId: tea.Int64(1),
		PhoneNumber: tea.String("test"),
		BizId: tea.String("test"),
		SendDate: tea.String("test"),
	}
	// 复制代码运行请自行打印 API 的返回值
	_, _err = client.QuerySendDetails(querySendDetailsRequest)
	if _err != nil {
		return _err
	}
	return _err
}


func main() {
	err := _main(tea.StringSlice(os.Args[1:]))
	if err != nil {
		panic(err)
	}
}