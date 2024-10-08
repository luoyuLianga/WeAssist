// redis初始化连接

package redis

import (
	"WeAssist/common/config"
	"context"

	"github.com/go-redis/redis/v8"
)

var RedisDb *redis.Client

// SetupRedisDb 连接redis
func SetupRedisDb() error {
	var ctx = context.Background()
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Address,
		Password: config.Config.Redis.Password,
		DB:       5,
	})

	_, err := RedisDb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}

func ZScore(id int, score int) redis.Z {
	return redis.Z{Score: float64(score), Member: id}
}
