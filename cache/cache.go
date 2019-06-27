package cache

import (
	"github.com/kuhufu/gcache"
)
import . "github.com/kuhufu/revdol/config"

var cacheSize = 100 * 1024 * 1024

var defaultCache gcache.CacheStore

const (
	MemCache   = "mem"
	RedisCache = "redis"
)

func init() {
	switch Config.Cache.Type {
	case MemCache:
		defaultCache = gcache.NewMemCache(cacheSize)
	case RedisCache:
		defaultCache = gcache.NewRedisCache(10, "tcp", "127.0.0.1:6379", "")
	default:
		defaultCache = gcache.NewMemCache(cacheSize)
	}
}

func Set(key string, val []byte, expireSeconds int) error {
	return defaultCache.Set(key, val, expireSeconds)

}

func Get(key string) (value []byte, err error) {
	return defaultCache.Get(key)
}

func GetUnmarshal(key string) (value interface{}, err error) {
	return defaultCache.GetUnmarshal(key)
}
