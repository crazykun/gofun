package cache

import (
	"fmt"
	"gofun/conf"
	"log"

	"github.com/go-redis/redis/v8"
)

type RedisPool struct {
	Instance map[string]*redis.Client
}

func (this *RedisPool) addInstance(redis_name string, db *redis.Client) {
	log.Println("=============  初始化 redis " + redis_name + "  =============")
	this.Instance[redis_name] = db
}

func (this *RedisPool) GetInstance(redis_name string) *redis.Client {
	return this.Instance[redis_name]
}

var pool *RedisPool

func RegisterRedisPool(clientName string, cacheConfig conf.RedisConfig) {

	if pool == nil {
		pool = new(RedisPool)
		pool.Instance = make(map[string]*redis.Client)
	}
	addr := fmt.Sprintf("%s:%d", cacheConfig.Host, cacheConfig.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     cacheConfig.Auth,
		DB:           cacheConfig.Db,
		PoolSize:     cacheConfig.MaxNum, //
		MinIdleConns: cacheConfig.MinNum,
	})
	pool.addInstance(clientName, rdb)
}

func GetRedisInstance() *RedisPool {
	return pool
}

func CloseRedisPool() {
	for name, rdb := range pool.Instance {
		if rdb == nil {
			continue
		}
		rdb.Close()
		log.Println("关闭redis连接：", name)
	}
}
