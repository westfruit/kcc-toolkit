package eth

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// 消息哈唏
func hashMsg(msg string, prefix bool) []byte {
	if prefix {
		msg = fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(msg), msg)
	}

	return crypto.Keccak256([]byte(msg))
}

// 签名消息
func sign(privateKeyHex string, msg string, prefix bool) string {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return ""
	}

	hash := hashMsg(msg, prefix)
	signature, err := crypto.Sign(hash, privateKey)
	if err != nil {
		return ""
	}

	return hexutil.Encode(signature)
}

// 验证签名
func verifyHash(signAddr, signHex string, hash []byte) bool {
	if !strings.HasPrefix(signHex, "0x") {
		signHex = "0x" + signHex
	}

	sign, err := hexutil.Decode(signHex)
	if err != nil {
		fmt.Println("签名解码失败, ", err)
		return false
	}

	if sign[64] == 27 || sign[64] == 28 {
		sign[64] -= 27
	}

	pubKey, err := crypto.SigToPub(hash, sign)
	if err != nil {
		return false
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey).Hex()
	return strings.EqualFold(recoveredAddr, signAddr)
}

// 使用私钥对消息签名
func Sign(privateKeyHex string, msg string) string {
	return sign(privateKeyHex, msg, false)
}

// 会对消息增加一个特定消息前缀，用于兼容metamask签名
func SignWithPrefix(privateKeyHex string, msg string) string {
	return sign(privateKeyHex, msg, true)
}

func Verify(fromAddr, signHex, msg string) bool {
	return verifyHash(fromAddr, signHex, hashMsg(msg, false))
}

func VerifyWithPrefix(fromAddr, signHex, msg string) bool {
	return verifyHash(fromAddr, signHex, hashMsg(msg, true))
}
