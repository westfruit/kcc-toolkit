package rand

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestRandTokenId100(t *testing.T) {
	tokenIdMap := make(map[int64]int)
	for i := 1; i <= 10000000; i++ {
		tokenId := RandTokenId()
		tokenIdMap[tokenId]++
	}

	var tokenIdCount int32
	var repeatTimes int32

	for k, v := range tokenIdMap {
		if v > 1 {
			tokenIdCount++
			repeatTimes += int32(v)

			fmt.Println("tokenId:", k, "重复次数:", v)
		}
	}

	fmt.Println("tokenId重复数:", tokenIdCount)
	fmt.Println("tokenId重复次数总和:", repeatTimes)
}

func TestRand(t *testing.T) {
	a := rand.Int()
	b := rand.Intn(100) //生成0-99之间的随机数
	//fmt.Println(a)
	//fmt.Println(b)
	t.Log(a)
	t.Log(b)
}
