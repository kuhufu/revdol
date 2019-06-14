package redisSource

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"github.com/kuhufu/flyredis"
	"log"
	. "revdol/config"
	"strconv"
	"time"
)

var DB = New(
	flyredis.NewPool(&redis.Pool{
		MaxIdle:     50,
		MaxActive:   1000,
		IdleTimeout: 30 * time.Second,
		Dial: func() (conn redis.Conn, err error) {
			return redis.Dial("tcp", Config.Redis.URL)
		},
	}))

type ForumInfo struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

type RedisSource struct {
	pool *flyredis.Pool
}

func New(pool *flyredis.Pool) *RedisSource {
	return &RedisSource{
		pool: pool,
	}
}

func (s *RedisSource) GetForumCount(id string) interface{} {
	result, err := s.pool.HGETALL("statistics:forum_count:" + id).IntMap()
	if err != nil {
		log.Println(err)
		return nil
	}

	return result
}

func (s *RedisSource) GetFansNumById(id string) interface{} {
	result, err := s.pool.HGETALL("statistics:idol:fans_num:" + id).IntMap()
	if err != nil {
		log.Println(err)
		return nil
	}
	return result
}

func (s *RedisSource) GetPopularNumById(id string) interface{} {
	result, err := s.pool.HGETALL("statistics:idol:popular_num:" + id).IntMap()
	if err != nil {
		log.Println(err)
		return nil
	}
	return result
}

func (s *RedisSource) GetAllIdolMeta() interface{} {
	ids := []int{1, 2, 3, 4, 5, 6}
	result := map[int]interface{}{}
	for _, id := range ids {
		result[id] = s.GetIdolMetaById(strconv.Itoa(id))
	}
	return result
}

func (s *RedisSource) GetIdolMetaById(id string) interface{} {
	result := map[string]interface{}{}

	result["popular_num"] = s.GetPopularNumById(id)
	result["fans_num"] = s.GetFansNumById(id)

	return result
}

func (s *RedisSource) GetForumById(id string) []byte {
	result, err := s.pool.HGET("revdol:forum:detail", id).Bytes()
	if err != nil {
		log.Println(err)
		return nil
	}
	return result
}

func (s *RedisSource) GetAllForum() []string {
	result, err := s.pool.HVALS("revdol:forum:detail").Strings()
	if err != nil {
		log.Println(err)
		return nil
	}
	return result
}

func (s *RedisSource) GetAllUser() []string {
	result, err := s.pool.HVALS("revdol:user:detail").Strings()
	if err != nil {
		log.Println(err)
		return nil
	}
	return result
}

func (s *RedisSource) GetUserById(id string) []byte {
	result, err := s.pool.HGET("revdol:user:detail", id).Bytes()
	if err != nil {
		log.Println(err)
		return nil
	}
	return result
}

func (s *RedisSource) GetUserContributeById(id string) []byte {
	result, err := s.pool.HVALS("revdol:user:contribute:" + id).Values()
	if err != nil {
		log.Println(err)
		return nil
	}
	for i, v := range result {
		var obj interface{}
		json.Unmarshal(v.([]byte), &obj)
		result[i] = obj
	}
	data, _ := json.Marshal(result)
	return data
}

func (s *RedisSource) GetAllUserContribute() func() []string {
	keys, err := s.pool.Do("KEYS", "revdol:user:contribute:*").Strings()
	if err != nil {
		log.Println(err)
		return nil
	}
	i := 0
	return func() []string {
		if i >= len(keys) {
			return nil
		}
		vals, err := flyredis.HVALS(keys[i]).Strings()
		if err != nil {
			log.Println(err)
		}
		i++
		return vals
	}
}

func (s *RedisSource) GetIdolById(id string) []byte {
	result, err := s.pool.HGET("revdol:idol:detail", id).Bytes()
	if err != nil {
		log.Println(err)
		return nil
	}
	return result
}

func (s *RedisSource) GetAllIdol() []byte {
	result, err := s.pool.HVALS("revdol:idol:detail").Values()
	if err != nil {
		log.Println(err)
		return nil
	}
	for i, v := range result {
		var obj interface{}
		json.Unmarshal(v.([]byte), &obj)
		result[i] = obj
	}
	data, _ := json.Marshal(result)
	return data
}
