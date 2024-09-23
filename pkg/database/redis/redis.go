package redis

import (
	"github.com/go-redis/redis/v7"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	URI string
}

func NewClient(cfg *Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.URI,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Error("failed to ping storage")
		return nil, err
	}

	return client, nil
}
