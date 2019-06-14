package redis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"github.com/kuhufu/flyredis"
	"log"
	. "revdol/config"
	"strconv"
	"time"
)

var pool = 	flyredis.NewPool(&redis.Pool{
	MaxIdle:     50,
	MaxActive:   1000,
	IdleTimeout: 30 * time.Second,
	Dial: func() (conn redis.Conn, err error) {
		return redis.Dial("tcp", Config.Redis.URL)
	},
})

type ForumInfo struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}


func GetForumCount(id string) interface{} {
	result, err := pool.HGETALL("statistics:forum_count:" + id).IntMap()
	if err != nil {
		log.Println(err)
		return nil
	}

	return result
}

func GetFansNumById(id string) interface{} {
	result, err := pool.HGETALL("statistics:idol:fans_num:" + id).IntMap()
	if err != nil {
		log.Println(err)
		return nil
	}
	return result
}

func GetPopularNumById(id string) interface{} {
	result, err := pool.HGETALL("statistics:idol:popular_num:" + id).IntMap()
	if err != nil {
		log.Println(err)
		return nil
	}
	return result
}

func GetAllIdolMeta() interface{} {
	ids := []int{1, 2, 3, 4, 5, 6}
	result := map[int]interface{}{}
	for _, id := range ids {
		result[id] = GetIdolMetaById(strconv.Itoa(id))
	}
	return result
}

func GetIdolMetaById(id string) interface{} {
	result := map[string]interface{}{}

	result["popular_num"] = GetPopularNumById(id)
	result["fans_num"] = GetFansNumById(id)

	return result
}

func GetForumById(id string) []byte {
	result, err := pool.HGET("revdol:forum:detail", id).Bytes()
	if err != nil {
		log.Println(err)
		return nil
	}
	return result
}

func GetAllForum() []string {
	result, err := pool.HVALS("revdol:forum:detail").Strings()
	if err != nil {
		log.Println(err)
		return nil
	}
	return result
}

func GetAllUser() []string {
	result, err := pool.HVALS("revdol:user:detail").Strings()
	if err != nil {
		log.Println(err)
		return nil
	}
	return result
}

func GetUserById(id string) []byte {
	result, err := pool.HGET("revdol:user:detail", id).Bytes()
	if err != nil {
		log.Println(err)
		return nil
	}
	return result
}

func GetUserContributeById(id string) []byte {
	result, err := pool.HVALS("revdol:user:contribute:" + id).Values()
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

func GetAllUserContribute() func() []string {
	keys, err := pool.Do("KEYS", "revdol:user:contribute:*").Strings()
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

func GetIdolById(id string) []byte {
	result, err := pool.HGET("revdol:idol:detail", id).Bytes()
	if err != nil {
		log.Println(err)
		return nil
	}
	return result
}

func GetAllIdol() []byte {
	result, err := pool.HVALS("revdol:idol:detail").Values()
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
