package eth

import (
	"github.com/spf13/viper"

	_ "kcc/kcc-toolkit/conf"

	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	client    *ethclient.Client
	wssClient *ethclient.Client

	network = viper.GetString("chain.network")
	rpcUrl  = viper.GetString("chain." + network + ".rpcUrl")
	wssUrl  = viper.GetString("chain." + network + ".wssUrl")
)

type LogTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
}

func NewClient() (*ethclient.Client, error) {
	if client != nil {
		return client, nil
	}

	client, err := ethclient.Dial(rpcUrl)
	return client, err
}

func NewWssClient() (*ethclient.Client, error) {
	if wssClient != nil {
		return wssClient, nil
	}

	wssClient, err := ethclient.Dial(wssUrl)
	return wssClient, err
}
