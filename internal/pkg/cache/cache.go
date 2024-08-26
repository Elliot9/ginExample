package cache

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	redisRepo "github/elliot9/ginExample/internal/repository/redis"
	"sort"
	"strings"

	"github.com/redis/go-redis/v9"
)

type Cache interface {
	Exists(ctx context.Context, key string) bool
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, value any) error
	Del(ctx context.Context, key string) bool
	Tags(tags []string) Cache
	Flush(ctx context.Context) error
	Remember(ctx context.Context, key string, callback func() (any, error)) ([]byte, error)
	Close() error
}

type cacheRepo struct {
	client *redis.Client
	tags   []string
}

func New(client redisRepo.Repo) Cache {
	return &cacheRepo{
		client: client.Get(),
	}
}

func (c *cacheRepo) Get(ctx context.Context, key string) ([]byte, error) {
	cacheKey := c.getKey(key)
	result, err := c.client.Get(ctx, cacheKey).Result()
	if err != nil {
		return nil, err
	}
	return []byte(result), nil
}

func (c *cacheRepo) getKey(key string) string {
	if len(c.tags) > 0 {
		combinedKey := fmt.Sprintf("%s:%s", strings.Join(c.tags, ":"), key)
		hashedKey := sha256.Sum256([]byte(combinedKey))
		return fmt.Sprintf("%x", hashedKey)
	}
	return key
}

func (c *cacheRepo) getGroupKey(key string) string {
	return fmt.Sprintf("%s:%s", "cache:group", key)
}

func (c *cacheRepo) getTagKey(tagName string) string {
	return fmt.Sprintf("%s:%s", "cache:tags", tagName)
}

func (c *cacheRepo) Set(ctx context.Context, key string, value any) error {
	var data []byte
	var err error

	// 檢查 value 是否已經是 []byte 類型
	if b, ok := value.([]byte); ok {
		data = b
	} else {
		// 如果不是，則進行 JSON 編碼
		data, err = json.Marshal(value)
		if err != nil {
			return fmt.Errorf("json marshal fail: %w", err)
		}
	}

	cacheKey := c.getKey(key)
	pipe := c.client.Pipeline()
	pipe.Set(ctx, cacheKey, data, 0)

	if len(c.tags) > 0 {
		groupKey := c.getGroupKey(cacheKey)
		for _, tag := range c.tags {
			pipe.SAdd(ctx, groupKey, tag)
			pipe.SAdd(ctx, c.getTagKey(tag), cacheKey)
		}
	}

	_, err = pipe.Exec(ctx)
	return err
}

func (c *cacheRepo) Del(ctx context.Context, key string) bool {
	cacheKey := c.getKey(key)
	c.client.Del(ctx, cacheKey)

	groupKey := c.getGroupKey(cacheKey)
	tags, _ := c.client.SMembers(ctx, groupKey).Result()

	pipe := c.client.Pipeline()
	for _, tagToRemove := range tags {
		pipe.SRem(ctx, c.getTagKey(tagToRemove), cacheKey)
	}
	pipe.Del(ctx, groupKey)
	_, err := pipe.Exec(ctx)

	return err == nil
}

func (c *cacheRepo) Close() error {
	return c.client.Close()
}

func (c *cacheRepo) Tags(tags []string) Cache {
	sort.Strings(tags)
	c.tags = tags
	return c
}

func (c *cacheRepo) Flush(ctx context.Context) error {
	if len(c.tags) == 0 {
		return fmt.Errorf("no tags specified")
	}

	for _, tag := range c.tags {
		// 獲取標籤對應的所有緩存鍵
		caches, err := c.client.SMembers(ctx, c.getTagKey(tag)).Result()
		if err != nil {
			return fmt.Errorf("failed to get caches for tag %s: %w", tag, err)
		}
		fmt.Printf("Caches for tag %s: %v\n", tag, caches)

		for _, cache := range caches {
			groupKey := c.getGroupKey(cache)

			// 獲取緩存對應的所有標籤
			tags, err := c.client.SMembers(ctx, groupKey).Result()
			if err != nil {
				return fmt.Errorf("failed to get tags for cache %s: %w", cache, err)
			}

			// 使用管道來執行批量操作
			pipe := c.client.Pipeline()
			for _, tagToRemove := range tags {
				if tagToRemove != tag { // 排除自身標籤
					pipe.SRem(ctx, c.getTagKey(tagToRemove), cache)
				}
			}
			pipe.Del(ctx, groupKey)
			pipe.Del(ctx, cache)

			// 執行管道操作
			_, err = pipe.Exec(ctx)
			if err != nil {
				return fmt.Errorf("failed to execute pipeline for cache %s: %w", cache, err)
			}
			fmt.Printf("Successfully removed cache %s and its associations\n", cache)
		}

		// 最後刪除標籤本身
		err = c.client.Del(ctx, c.getTagKey(tag)).Err()
		if err != nil {
			return fmt.Errorf("failed to delete tag %s: %w", tag, err)
		}
		fmt.Printf("Successfully deleted tag %s\n", tag)
	}

	return nil
}

func (c *cacheRepo) Remember(ctx context.Context, key string, callback func() (any, error)) ([]byte, error) {
	if c.Exists(ctx, key) {
		fmt.Println("hits cache")
		return c.Get(ctx, key)
	}

	value, err := callback()
	if err != nil {
		return nil, err
	}

	jsonValue, err := json.Marshal(value)
	if err != nil {
		return nil, fmt.Errorf("json marshal fail: %w", err)
	}

	err = c.Set(ctx, key, jsonValue)
	if err != nil {
		return nil, err
	}

	return jsonValue, nil
}

func (c *cacheRepo) Exists(ctx context.Context, key string) bool {
	cacheKey := c.getKey(key)
	return c.client.Exists(ctx, cacheKey).Val() > 0
}
