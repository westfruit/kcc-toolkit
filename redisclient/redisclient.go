package redisclient

import (
	"context"
	_ "kcc/kcc-toolkit/conf"

	redis "github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	DefaultClient *redis.Client
)

func init() {

	addr := viper.GetString("redis.addr")
	password := viper.GetString("redis.password")
	db := viper.GetInt("redis.db")

	DefaultClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	err := DefaultClient.Ping(context.Background()).Err()
	if err != nil {
		logrus.Info("连接redis失败, ", err)
		panic("连接redis失败, redis地址=" + addr + ", error: " + err.Error())
	}

	logrus.Info("连接redis成功, 地址=", addr)
}
