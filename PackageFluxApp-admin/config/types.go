/**
 * Description：
 * FileName：types.go
 * Author：CJiaの用心
 * Create：2025/2/19 14:12:10
 * Remark：
 */

package config

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type RelyConfig struct {
	Logger *zap.Logger
	Db     *gorm.DB
	Redis  redis.Cmdable
	Mail   EmailConfig
	Trans  ut.Translator
	Token  TokenConfig
	AliYun AliYunConfig
}
