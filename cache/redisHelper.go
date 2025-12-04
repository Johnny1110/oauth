package cache

//
//import (
//	"errors"
//	"fmt"
//	"github.com/go-redis/redis/v8"
//	"oauth/config"
//	"oauth/enum"
//	"oauth/sys"
//	"time"
//)
//
//func Get(key string) string {
//	val, err := config.GetRedisClient().Get(config.GetRedisContext(), key).Result()
//	if errors.Is(err, redis.Nil) {
//		sys.Logger().Warningf("cache client not found for key: %s", key)
//		return ""
//	}
//	if err != nil {
//		sys.Logger().Warningf("cache client error: %s", err)
//		return ""
//	}
//	return val
//}
//
//func GenKeyString(key enum.RedisKey, content string) string {
//	if key.KeyPrefix != "" && key.KeySuffix == "" {
//		return fmt.Sprintf("%s:%s", key.KeyPrefix, content)
//	}
//	if key.KeyPrefix == "" && key.KeySuffix != "" {
//		return fmt.Sprintf("%s:%s", content, key.KeySuffix)
//	}
//	if key.KeyPrefix == "" && key.KeySuffix == "" {
//		return content
//	}
//	if key.KeyPrefix != "" && key.KeySuffix != "" {
//		return fmt.Sprintf("%s:%s:%s", key.KeyPrefix, content, key.KeySuffix)
//	}
//	panic("key prefix or key suffix not exist")
//}
//
//func Set(key string, val string, ttl int) bool {
//	err := config.GetRedisClient().Set(config.GetRedisContext(), key, val, time.Duration(ttl)*time.Minute).Err()
//	if err != nil {
//		sys.Logger().Warningf("cache client error: %s", err)
//		return false
//	}
//	return true
//}
//
//func ScanDelete(keyPrefix string) {
//	rdb := config.GetRedisClient()
//	ctx := config.GetRedisContext()
//	// scan
//	var cursor uint64
//	var keys []string
//	for {
//		var err error
//		var k []string
//		k, cursor, err = rdb.Scan(ctx, cursor, keyPrefix+"*", 0).Result()
//		if err != nil {
//			sys.Logger().Fatalf("Failed to scan keys: %v", err)
//		}
//		keys = append(keys, k...)
//		if cursor == 0 {
//			break
//		}
//	}
//
//	// delete
//	if len(keys) > 0 {
//		err := rdb.Del(ctx, keys...).Err()
//		if err != nil {
//			sys.Logger().Fatalf("Failed to delete keys: %v", err)
//		}
//		sys.Logger().Printf("Deleted %d keys\n", len(keys))
//	} else {
//		sys.Logger().Println("No keys found to delete")
//	}
//}
