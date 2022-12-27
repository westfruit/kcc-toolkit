package security

import (
	"github.com/afocus/captcha"
)

var (
	// 验证码
	cap *captcha.Captcha
)

func init() {
	// 构建验证码
	cap = captcha.New()
	// 设置验证码大小
	cap.SetSize(80, 36)
	// 设置字体
	cap.SetFont("./conf/comic.ttf")
}

// 获取数字验证码图片
func GetNumCaptcha() (*captcha.Image, string) {

	// 创建验证码 4个字符 captcha.NUM 字符模式数字类型
	// 返回验证码图像对象、验证码字符串，后期可以对字符串进行对比 判断验证
	return cap.Create(4, captcha.NUM)
}
