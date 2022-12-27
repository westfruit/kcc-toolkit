package wallet

import (
	"regexp"
	"strings"
)

var (
	avMap map[string]AddressValidator
)

func init() {
	btcValidator := NewBTCAddressValidator()
	ethValidator := NewETHAddressValidator()
	trxValidator := NewTRXAddressValidator()

	avMap = map[string]AddressValidator{
		btcValidator.Symbol(): btcValidator,
		ethValidator.Symbol(): ethValidator,
		trxValidator.Symbol(): trxValidator,
	}
}

// 验证地址格式是否正确, 默认为true
func ValidateAddress(symbol string, address string) bool {
	if v, ok := avMap[strings.ToUpper(symbol)]; ok {
		return v.Validate(address)
	}

	return true
}

// 地址验证
type AddressValidator interface {

	// 返回地址代币符号
	Symbol() string

	// 验证地址格式是否正确
	Validate(address string) bool
}

type AddressValidatorBase struct {
	re *regexp.Regexp
}

func (v *AddressValidatorBase) Validate(address string) bool {
	return true
}

func (v *AddressValidatorBase) Symbol() string {
	return ""
}

// BTC
type BTCAddressValidator struct {
	AddressValidatorBase
}

func NewBTCAddressValidator() *BTCAddressValidator {
	v := &BTCAddressValidator{}
	v.re = regexp.MustCompile("^(bc1|[13])[a-zA-HJ-NP-Z0-9]{25,39}$")

	return v
}

func (v *BTCAddressValidator) Validate(address string) bool {
	valid := v.re.MatchString(address)
	return valid
}

func (v *BTCAddressValidator) Symbol() string {
	return "BTC"
}

// ETH
type ETHAddressValidator struct {
	AddressValidatorBase
}

func NewETHAddressValidator() *ETHAddressValidator {
	v := &ETHAddressValidator{}
	v.re = regexp.MustCompile("^(0x)[a-zA-Z0-9]{40}$")

	return v
}

func (v *ETHAddressValidator) Validate(address string) bool {
	valid := v.re.MatchString(address)
	return valid
}

func (v *ETHAddressValidator) Symbol() string {
	return "ETH"
}

// TRX
type TRXAddressValidator struct {
	AddressValidatorBase
}

func NewTRXAddressValidator() *TRXAddressValidator {
	v := &TRXAddressValidator{}
	v.re = regexp.MustCompile("^(T|t)$")

	return v
}

func (v *TRXAddressValidator) Validate(address string) bool {
	valid := strings.HasPrefix(address, "T")
	return valid
}

func (v *TRXAddressValidator) Symbol() string {
	return "TRX"
}
