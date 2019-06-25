package cache

import "testing"

func TestNewRedisCache(t *testing.T) {
	c := NewRedisCache(10, "tcp", "127.0.0.1:6379", "")
	c.Set("test_key", []byte("test_data"), -1)
}
