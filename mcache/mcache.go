package mcache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

const ()

var (
	// 默认的本地缓存
	DefaultCache *cache.Cache
)

func init() {
	DefaultCache = cache.New(5*time.Minute, 10*time.Minute)
}
