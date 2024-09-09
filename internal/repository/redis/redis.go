package redis

import (
	"context"
	"fmt"

	"github.com/elliot9/gin-example/config"

	"github.com/redis/go-redis/v9"
)

type Repo interface {
	Get() *redis.Client
	Close() error
}

type redisRepo struct {
	client *redis.Client
}

func (r *redisRepo) Get() *redis.Client {
	return r.client
}

func (r *redisRepo) Close() error {
	return r.client.Close()
}

func New() (Repo, error) {
	client, err := redisConnect()
	if err != nil {
		return nil, err
	}

	return &redisRepo{
		client: client,
	}, nil
}

func redisConnect() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.RedisSetting.Host, config.RedisSetting.Port),
		Password: config.RedisSetting.Password,
		DB:       0,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("[redis connection fail]")
	}

	return client, nil
}
