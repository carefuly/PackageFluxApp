/**
 * Description：
 * FileName：user.go
 * Author：CJiaの用心
 * Create：2025/2/26 14:25:35
 * Remark：
 */

package cache

import (
	"github.com/redis/go-redis/v9"
	"time"
)

type UserCache interface {
	// Set(ctx context.Context, du domain.User) error
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
