package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github/elliot9/ginExample/config"
)

type Repo interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key, value string) error
	Del(ctx context.Context, key string) bool
	Close() error
}

type cacheRepo struct {
	client *redis.Client
}

func New() (Repo, error) {
	client, err := redisConnect()
	if err != nil {
		return nil, err
	}

	return &cacheRepo{
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

func (c cacheRepo) Get(ctx context.Context, key string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (c cacheRepo) Set(ctx context.Context, key, value string) error {
	//TODO implement me
	panic("implement me")
}

func (c cacheRepo) Del(ctx context.Context, key string) bool {
	//TODO implement me
	panic("implement me")
}

func (c cacheRepo) Close() error {
	return c.client.Close()
}
