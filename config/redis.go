package config

import (
	"github.com/go-redis/redis/v8"
	"github.com/vodeacloud/hr-api/pkg/logger"
)

func GetRedisCli() *redis.Client {
	return redis.NewClient(GetRedisOpt())
}

func GetRedisOpt() *redis.Options {
	config := GetConfig()
	return &redis.Options{
		Addr:     config.RedisHost + ":" + config.RedisPort,
		Password: config.RedisPassword,
	}
}

func CloseRedis(cli *redis.Client) {
	if err := cli.Close(); err != nil {
		logger.Fatalf("failed close connection redis: %v", err)
	}
}
