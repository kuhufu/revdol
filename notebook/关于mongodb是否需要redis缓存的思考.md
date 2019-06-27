# 关于mongodb是否需要redis缓存的思考

mongodb 和 redis都是NoSQL非关系型数据库。

redis 是基于KV键值对的内存数据库，可通过RDB或AOF持久化。

mongodb是基于文档的数据库，在非关系型数据库中，mongodb是最像关系型数据库的，关系型数据库的SQL基本可以映射到mongodb。

## 对比

### 查询

redis不支持复杂查询

mongodb支持复杂查询



### 高可用

redis 支持 replication

mongodb支持 replica set



### 拓展性

redis 集群

mongodb 分片shard



#### 扩展难易度

redis 较复杂

mongodb 简单



## Benchmark

**mongo**(触发索引)

```
pkg: github.com/kuhufu/revdol/dao/mongo
BenchmarkGetForumById-8   	    5000	    292231 ns/op
```

**redis**

```
pkg: github.com/kuhufu/revdol/dao/redis
BenchmarkGetForumById-8   	   10000	    115235 ns/op
```

可以看到 redis 的速度大约是 mongo 的3倍



## 结论

需要redis做缓存