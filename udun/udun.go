package udun

const (
	ApiResultCodeSuccess          = 200  // 成功
	ApiResultCodeIllegalParam     = 4005 // 非法参数
	ApiResultCodeInvalidSignature = 4162 // 签名错误
	ApiResultCodeInvalidAddress   = 4165 // 地址不合法
)

type ApiResult struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func ApiSuccess() *ApiResult {
	return &ApiResult{
		Code:    ApiResultCodeSuccess,
		Message: "SUCCESS",
	}
}

// api请求参数
type ApiParam struct {
	Timestamp string `form:"timestamp"  binding:"required"` // 时间戳
	Nonce     string `form:"nonce" binding:"required"`      // 随机数
	Sign      string `form:"sign" binding:"required"`       // 签名
	Body      string `form:"body" binding:"required"`       // 消息内容	json字符串
}

// 回调参数
type CallbackParam struct {
	Address      string `json:"Address"`
	Amount       string `json:"amount"`       // 交易数量，根据币种精度获取实际金额，实际金额=amount/pow(10,decimals)，即实际金额等于amount除以10的decimals次方
	Fee          string `json:"fee"`          // 矿工费，根据币种精度获取实际金额，实际金额获取同上
	Decimals     string `json:"decimals"`     // 币种精度
	CoinType     string `json:"coinType"`     // 子币种编号
	MainCoinType string `json:"mainCoinType"` // 主币种编号
	BusinessId   string `json:"businessId"`   // 业务编号，提币回调时为提币请求时传入的，充币回调无值
	BlockHigh    string `json:"blockHigh"`    // 区块高度
	Status       int32  `json:"status"`       // 状态
	TradeId      string `json:"tradeId"`      // 交易流水号
	TradeType    int32  `json:"tradeType"`    // 交易类型
	TxId         string `json:"txid"`         // 区块链交易哈希
	Memo         string `json:"memo"`         // 备注，XRP和EOS, 这2种类型币的充提币可能有值
}

const (
	// 回调状态
	CallbackStatusWaitAudit = 0 // 待审核
	CallbackStatusApproved  = 1 // 审核成功
	CallbackStatusRejected  = 2 // 审核驳回
	CallbackStatusSuccess   = 3 // 交易成功
	CallbackStatusFailure   = 4 // 交易失败

	// 回调类型
	CallbackTradeTypeDeposit  = 1 // 充值
	CallbackTradeTypeWithdraw = 2 // 提币
)

// 提币参数
type WithdrawParam struct {
	Address      string `json:"address" binding:"required"`    // 提币地址
	Amount       string `json:"amount" binding:"required"`     // 提币数量
	Symbol       string `json:"symbol" binding:"required"`     // 主币
	Token        string `json:"token" binding:"required"`      // 代币
	MainCoinType int32  `json:"mainCoinType"`                  // 主币类型
	CoinType     int32  `json:"coinType"`                      // 代币类型
	BusinessId   string `json:"businessId" binding:"required"` // 业务id，必须保证该字段在系统内唯一，如果重复，则该笔审核钱包不会接收。
	Memo         string `json:"memo"`                          // 备注,XRP和EOS，这两种币的提币申请该字段可选，其他类型币种不填
}
