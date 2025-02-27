/**
 * Description：
 * FileName：file.go
 * Author：CJiaの用心
 * Create：2025/2/27 23:28:15
 * Remark：
 */

package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/carefuly/PackageFluxApp/internal/domain"
	"github.com/redis/go-redis/v9"
	"time"
)

var ErrFileNotExist = redis.Nil

type FileCache interface {
	Get(ctx context.Context, id string) (*domain.File, error)
	Set(ctx context.Context, d domain.File) error
}

type RedisFileCache struct {
	cmd        redis.Cmdable
	expiration time.Duration
}

func NewRedisFileCache(cmd redis.Cmdable) FileCache {
	return &RedisFileCache{
		cmd:        cmd,
		expiration: time.Minute * 15,
	}
}

func (c *RedisFileCache) Get(ctx context.Context, id string) (*domain.File, error) {
	key := c.key(id)

	data, err := c.cmd.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var f domain.File
	err = json.Unmarshal([]byte(data), &f)
	return &f, err
}

func (c *RedisFileCache) Set(ctx context.Context, f domain.File) error {
	key := c.key(f.Id)
	data, err := json.Marshal(f)
	if err != nil {
		return err
	}
	return c.cmd.Set(ctx, key, data, c.expiration).Err()
}

func (c *RedisFileCache) key(id string) string {
	return fmt.Sprintf("file:info:%s", id)
}
