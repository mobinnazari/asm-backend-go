package db

import (
	"log"

	"github.com/redis/go-redis/v9"
)

func InitRedis(addr string) *redis.Client {
	opt, err := redis.ParseURL(addr)
	if err != nil {
		log.Panicln(err.Error())
	}

	return redis.NewClient(opt)
}
