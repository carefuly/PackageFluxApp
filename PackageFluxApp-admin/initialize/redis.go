/**
 * Description：
 * FileName：redis.go
 * Author：CJiaの用心
 * Create：2025/2/19 16:02:18
 * Remark：
 */

package initialize

import (
	"fmt"
	"github.com/carefuly/PackageFluxApp/config"
	"github.com/redis/go-redis/v9"
)

func InitRedis(redisClient config.RedisConfig) redis.Cmdable {
	// 初始化Redis客户端
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisClient.Host, redisClient.Port), // Redis地址
		Password: redisClient.Password,                                     // Redis密码
		DB:       redisClient.Db,                                           // 使用的数据库
	})
}
