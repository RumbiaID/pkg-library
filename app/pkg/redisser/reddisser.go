package redisser

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Redisdatabase int    `validate:"number,min=0" name:"REDIS_DB"`
	Redishost     string `validate:"required" name:"REDIS_HOST"`
	Redisport     int    `validate:"required,number" name:"REDIS_PORT"`
	Redispassword string `name:"REDIS_PASSWORD"`
}

func NewRedis(config *Config) *redis.Client {
	var ctx = context.Background()
	r := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Redishost, config.Redisport),
		Password: config.Redispassword,
		DB:       config.Redisdatabase,
	})

	err := r.Ping(ctx).Err()
	if err != nil {
		logrus.Fatal(err)
	}
	return r
}
