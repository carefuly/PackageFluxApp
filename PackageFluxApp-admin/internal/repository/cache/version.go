/**
 * Description：
 * FileName：version.go
 * Author：CJiaの用心
 * Create：2025/2/28 13:11:39
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

var ErrVersionNotExist = redis.Nil

type VersionCache interface {
	Get(ctx context.Context, id, userId string) (*domain.Version, error)
	Set(ctx context.Context, d domain.Version) error
	Del(ctx context.Context, id, userId string) error
}

type RedisVersionCache struct {
	cmd        redis.Cmdable
	expiration time.Duration
}

func NewRedisVersionCache(cmd redis.Cmdable) VersionCache {
	return &RedisVersionCache{
		cmd:        cmd,
		expiration: time.Minute * 15,
	}
}

func (c *RedisVersionCache) Get(ctx context.Context, id, userId string) (*domain.Version, error) {
	key := c.key(id, userId)

	data, err := c.cmd.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var v domain.Version
	err = json.Unmarshal([]byte(data), &v)
	return &v, err
}

func (c *RedisVersionCache) Set(ctx context.Context, v domain.Version) error {
	key := c.key(v.Id, v.UserId)
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return c.cmd.Set(ctx, key, data, c.expiration).Err()
}

func (c *RedisVersionCache) Del(ctx context.Context, id, userId string) error {
	key := c.key(id, userId)
	return c.cmd.Del(ctx, key).Err()
}

func (c *RedisVersionCache) key(id, userId string) string {
	return fmt.Sprintf("version:info:%s:%s", id, userId)
}
