package cache

import "github.com/kuhufu/gcache"

var cacheSize = 100 * 1024 * 1024

//var defaultCache = NewRedisCache(10, "tcp", "127.0.0.1:6379", "")
var defaultCache = gcache.NewMemCache(cacheSize)

func Set(key string, val []byte, expireSeconds int) error {
	return defaultCache.Set(key, val, expireSeconds)

}

func Get(key string) (value []byte, err error) {
	return defaultCache.Get(key)
}

func GetUnmarshal(key string) (value interface{}, err error) {
	return defaultCache.GetUnmarshal(key)
}
