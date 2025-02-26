/**
 * Description：
 * FileName：code.go
 * Author：CJiaの用心
 * Create：2025/2/26 16:20:54
 * Remark：
 */

package repository

import (
	"context"
	"github.com/carefuly/PackageFluxApp/internal/repository/cache"
)

var (
	ErrCodeSendTooMany   = cache.ErrCodeSendTooMany
	ErrUserBlocked       = cache.ErrUserBlocked
	ErrCodeNotFound      = cache.ErrCodeNotFound
	ErrCodeVerifyTooMany = cache.ErrCodeVerifyTooMany
	ErrCodeIncorrect     = cache.ErrCodeIncorrect
)

type CodeRepository interface {
	Set(ctx context.Context, biz, to, code string) error
	Verify(ctx context.Context, biz, to, code string) (bool, error)
}

type codeRepository struct {
	cache cache.CodeCache
}

func NewCodeRepository(cache cache.CodeCache) CodeRepository {
	return &codeRepository{
		cache: cache,
	}
}

func (repo *codeRepository) Set(ctx context.Context, biz, to, code string) error {
	return repo.cache.Set(ctx, biz, to, code)
}

func (repo *codeRepository) Verify(ctx context.Context, biz, to, code string) (bool, error) {
	return repo.cache.Verify(ctx, biz, to, code)
}
