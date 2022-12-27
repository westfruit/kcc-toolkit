package wallet

// 提币参数
type ApplyTransactionParam struct {
	Symbol     string  `json:"symbol"`
	ContractId string  `json:"contractId"`
	Amount     float64 `json:"amount"`
	Usid       string  `json:"usid"`
	Memo       string  `json:"memo"`
	ToAddress  string  `json:"to"`
}
