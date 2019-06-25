package cache

import (
	"encoding/json"
	"github.com/coocood/freecache"
)

type MemCache struct {
	inner *freecache.Cache
}

func NewMemCache(size int) cacheStore {
	return &MemCache{
		inner: freecache.NewCache(size),
	}
}

func (c *MemCache) Set(key string, val []byte, expireSeconds int) error {
	return c.inner.Set([]byte(key), val, expireSeconds)
}

func (c *MemCache) Get(key string) (value []byte, err error) {
	return c.inner.Get([]byte(key))
}

func (c *MemCache) GetUnmarshal(key string) (value interface{}, err error) {
	data, err := c.inner.Get([]byte(key))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &value)
	return
}
