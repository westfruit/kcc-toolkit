package eth

import (
	"log"

	"context"

	_ "kcc/kcc-toolkit/conf"

	"math/big"
	"reflect"
	"regexp"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/shopspring/decimal"
	"github.com/willf/pad"
)

// 若在该地址存储了字节码，该地址是智能合约
// 当地址上没有字节码时，我们知道它不是一个智能合约，它是一个标准的以太坊账户。
func IsContract(addr string) bool {
	address := common.HexToAddress(addr)

	// nil is latest block
	bytecode, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}

	return len(bytecode) > 0
}

// IsValidAddress validate hex address
func IsValidAddress(addr interface{}) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	switch v := addr.(type) {
	case string:
		return re.MatchString(v)
	case common.Address:
		return re.MatchString(v.Hex())
	default:
		return false
	}
}

func IsValidTxHash(txHash string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{64}$")
	return re.MatchString(txHash)
}

// IsZeroAddress validate if it's a 0 address
func IsZeroAddress(addr interface{}) bool {
	var address common.Address
	switch v := addr.(type) {
	case string:
		address = common.HexToAddress(v)
	case common.Address:
		address = v
	default:
		return false
	}

	zeroAddressBytes := common.FromHex("0x0000000000000000000000000000000000000000")
	addressBytes := address.Bytes()
	return reflect.DeepEqual(addressBytes, zeroAddressBytes)
}

// ToDecimal wei to decimals
func ToDecimal(val interface{}, decimals int) decimal.Decimal {
	value := new(big.Int)
	switch v := val.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)

	return result
}

// ToWei decimals to wei
func ToWei(val interface{}, decimals int) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := val.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei
}

// CalcGasCost calculate gas cost given gas limit (units) and gas price (wei)
func CalcGasCost(gasLimit uint64, gasPrice *big.Int) *big.Int {
	gasLimitBI := big.NewInt(int64(gasLimit))
	return gasLimitBI.Mul(gasLimitBI, gasPrice)
}

// SigRSV signatures R S V returned as arrays
func SigRSV(isig interface{}) ([32]byte, [32]byte, uint8) {
	var sig []byte
	switch v := isig.(type) {
	case []byte:
		sig = v
	case string:
		sig, _ = hexutil.Decode(v)
	}

	sigstr := common.Bytes2Hex(sig)
	rS := sigstr[0:64]
	sS := sigstr[64:128]
	R := [32]byte{}
	S := [32]byte{}
	copy(R[:], common.FromHex(rS))
	copy(S[:], common.FromHex(sS))
	vStr := sigstr[128:130]
	vI, _ := strconv.Atoi(vStr)
	V := uint8(vI + 27)

	return R, S, V
}

// 格式化数量，根据少数位数
func FormatAmount(amount *big.Int, decimals int) float64 {
	ret := pad.Left(amount.String(), decimals, "0")

	if len := len(ret); len > decimals {
		ret = ret[:len-decimals] + "." + ret[len-decimals:]
	} else {
		ret = "0." + ret
	}

	val, _ := strconv.ParseFloat(ret, 64)

	return val
}
