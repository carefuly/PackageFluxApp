/**
 * Description：
 * FileName：user.go
 * Author：CJiaの用心
 * Create：2025/2/26 14:25:35
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

var ErrUserKeyNotExist = redis.Nil

type UserCache interface {
	Get(ctx context.Context, id string) (*domain.User, error)
	Set(ctx context.Context, d domain.User) error
	Del(ctx context.Context, id string) error
}

type RedisUserCache struct {
	cmd        redis.Cmdable
	expiration time.Duration
}

func NewRedisUserCache(cmd redis.Cmdable) UserCache {
	return &RedisUserCache{
		cmd:        cmd,
		expiration: time.Minute * 15,
	}
}

func (c *RedisUserCache) Get(ctx context.Context, id string) (*domain.User, error) {
	key := c.key(id)

	data, err := c.cmd.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var u domain.User
	err = json.Unmarshal([]byte(data), &u)
	return &u, err
}

func (c *RedisUserCache) Set(ctx context.Context, d domain.User) error {
	key := c.key(d.Id)
	data, err := json.Marshal(d)
	if err != nil {
		return err
	}
	return c.cmd.Set(ctx, key, data, c.expiration).Err()
}

func (c *RedisUserCache) Del(ctx context.Context, id string) error {
	key := c.key(id)
	return c.cmd.Del(ctx, key).Err()
}

func (c *RedisUserCache) key(id string) string {
	// user:info:
	return fmt.Sprintf("user:info:%s", id)
}
