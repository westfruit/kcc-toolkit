package geetest

import "github.com/spf13/viper"

const (
	// 相关请求参数

	// 极验二次验证表单传参字段 chllenge
	GeetestParamChallenge = "geetest_challenge"

	// 极验二次验证表单传参字段 validate
	GeetestParamValidate = "geetest_validate"

	// 极验二次验证表单传参字段 seccode
	GeetestParamSeccode = "geetest_seccode"

	// 极验验证API服务状态Session
	GeetestParamServerStatus = "gt_server_status"

	// 是否json格式
	JsonFormat = 1

	SdkVersion = "1.0.0"
)

var (
	// 密钥
	AppId  = viper.GetString("geetest.appId")
	AppKey = viper.GetString("geetest.appKey")

	// 服务请求地址
	ApiUrl      = viper.GetString("geetest.apiUrl")
	RegisterUrl = viper.GetString("geetest.registerUrl")
	ValidateUrl = viper.GetString("geetest.validateUrl")

	// 是否启用新验证码
	NewCaptcha = true

	// 是否开启geetest
	Enabled = viper.GetBool("geetest.enabled")

	// 数据签名模式
	DigestMode = viper.GetString("geetest.digestmod")
	ClientType = viper.GetString("geetest.clientType")
)

const (
	DigestModeMD5        = "md5"
	DigestModeSHA256     = "sha256"
	DigestModeHMACSHA256 = "hmac-sha256"

	ResultStatusFailure = 0
	ResultStatusSuccess = 1

	ResultStatusFailureText = "fail"
	ResultStatusSuccessText = "success"

	HttpMethodGet  = "GET"
	HttpMethodPost = "POST"

	ClientTypeWeb     = "web"     // pc浏览器
	ClientTypeH5      = "h5"      // 手机浏览器，包括webview
	ClientTypeNative  = "native"  // 原生app
	ClientTypeUnknown = "unknown" // 未知

	// 极验服务器状态key, 用来存放到redis
	ServerStatusKey = "geetest:gt_server_status"
)

// 返回结果
type GeetestResult struct {
	Status int    `json:"status"` // 0=失败, 1=成功
	Data   string `json:"data"`
	Msg    string `json:"msg"`
}

type Param struct {
	// user_id作为终端用户的唯一标识，确定用户的唯一性；作用于提供进阶数据分析服务，
	// 可在api1 或 api2 接口传入，不传入也不影响验证服务的使用；若担心用户信息风险，可作预处理(如哈希处理)再提供到极验
	UserId string

	// 客户端请求SDK服务器的ip地址
	IPAddress string

	// 客户端类型，web（pc浏览器），h5（手机浏览器，包括webview），native（原生app），unknown（未知）
	ClientType string
}

// 注册步骤参数
type RegisterParam struct {
	Param
}

// 注册结果
type RegisterResult struct {
	Success    int32
	AppId      string
	Challenge  string
	NewCaptcha bool
}

// 成功验证步骤参数
type SuccessValidateParam struct {
	Param
	Seccode   string `json:"geetest_seccode" binding:"required"`   // 核心校验数据, 极验二次验证表单传参字段 seccode
	Challenge string `json:"geetest_challenge" binding:"required"` // 流水号，一次完整验证流程的唯一标识, 极验二次验证表单传参字段 chllenge
	Validate  string `json:"geetest_validate" binding:"required"`  // 待校验的核心数据 chllenge
}

// 失败验证参数
type FailValidateParam struct {
	Challenge string `json:"challenge"`
	Validate  string `json:"validate"`
	Seccode   string `json:"seccode"`
}
