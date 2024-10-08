// redis存取验证码
// author daniel

package util

import (
	"WeAssist/pkg/redis"
	"context"
	"log"
	"strconv"
	"time"
)

var ctx = context.Background()

type RedisStore struct{}

// 存验证码
func (r RedisStore) Update(key string, id int) error {
	redis.RedisDb.ZIncrBy(ctx, key, 1, strconv.Itoa(id))
	return nil
}

// 存验证码
func (r RedisStore) Set(key string, id int, score int) error {
	z := redis.ZScore(id, score)
	err := redis.RedisDb.ZAdd(ctx, key, &z).Err()
	if err != nil {
		log.Panicln(err.Error())
		return err
	}
	redis.RedisDb.Expire(ctx, key, 24*time.Hour).Err()
	return nil
}

// 获取排行傍的值
func (r RedisStore) Get(key string) []string {
	val, err := redis.RedisDb.ZRevRange(ctx, key, 0, -1).Result()
	if err != nil {
		return nil
	}
	return val
}
