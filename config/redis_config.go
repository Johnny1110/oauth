package config

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var (
	rdb *redis.Client
	ctx = context.Background()
)

func init() {
	//properties := GetProperties()
	//redisAddr := fmt.Sprintf("%s:%v", properties.Redis.IP, properties.Redis.Port)
	//sys.Logger().Infof("redis addr: %v", redisAddr)
	//rdb = redis.NewClient(&redis.Options{
	//	Addr:     redisAddr,
	//	Password: properties.Redis.Password, // 沒有密碼的話，設置為空字串
	//	DB:       0,                         // 使用默認的DB
	//	PoolSize: 10,                        // 連接池大小
	//})
	//
	//_, err := rdb.Ping(ctx).Result()
	//if err != nil {
	//	log.Fatalf("Could not connect to Redis: %v", err)
	//}
}

func GetRedisClient() *redis.Client {
	return rdb
}

func GetRedisContext() context.Context {
	return ctx
}
