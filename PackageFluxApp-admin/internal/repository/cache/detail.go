/**
 * Description：
 * FileName：detail.go
 * Author：CJiaの用心
 * Create：2025/2/27 10:51:09
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

var ErrDetailNotExist = redis.Nil

type DetailCache interface {
	Get(ctx context.Context, id string) (*domain.Detail, error)
	Set(ctx context.Context, d domain.Detail) error
}

type RedisDetailCache struct {
	cmd        redis.Cmdable
	expiration time.Duration
}

func NewRedisDetailCache(cmd redis.Cmdable) DetailCache {
	return &RedisDetailCache{
		cmd:        cmd,
		expiration: time.Minute * 15,
	}
}

func (c *RedisDetailCache) Get(ctx context.Context, id string) (*domain.Detail, error) {
	key := c.key(id)

	data, err := c.cmd.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var d domain.Detail
	err = json.Unmarshal([]byte(data), &d)
	return &d, err
}

func (c *RedisDetailCache) Set(ctx context.Context, d domain.Detail) error {
	key := c.key(d.Id)
	data, err := json.Marshal(d)
	if err != nil {
		return err
	}
	return c.cmd.Set(ctx, key, data, c.expiration).Err()
}

func (c *RedisDetailCache) key(id string) string {
	return fmt.Sprintf("detail:info:%s", id)
}
