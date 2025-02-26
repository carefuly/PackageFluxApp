/**
 * Description：
 * FileName：code.go
 * Author：CJiaの用心
 * Create：2025/2/26 16:18:44
 * Remark：
 */

package cache

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var (
	//go:embed lua/set_code.lua
	luaSetCode string
	//go:embed lua/verify_code.lua
	luaVerifyCode string

	ErrCodeSendTooMany   = errors.New("发送太频繁")
	ErrUserBlocked       = errors.New("用户被限制")
	ErrCodeNotFound      = errors.New("未发送验证码")
	ErrCodeVerifyTooMany = errors.New("验证次数已耗尽")
	ErrCodeIncorrect     = errors.New("验证码错误")
)

type CodeCache interface {
	Set(ctx context.Context, biz, to, code string) error
	Verify(ctx context.Context, biz, to, code string) (bool, error)
}

type codeCache struct {
	cmd redis.Cmdable
}

func NewCodeCache(cmd redis.Cmdable) CodeCache {
	return &codeCache{
		cmd: cmd,
	}
}

func (c *codeCache) Set(ctx context.Context, biz, to, code string) error {
	res, err := c.cmd.Eval(ctx, luaSetCode, []string{c.key(biz, to)}, code).Int()
	if err != nil {
		// 调用 redis 出了问题
		return err
	}

	switch res {
	case -2:
		return errors.New("验证码存在，但是没有过期时间")
	case -1:
		return ErrCodeSendTooMany
	default:
		return nil
	}
}

func (c *codeCache) Verify(ctx context.Context, biz, to, code string) (bool, error) {
	res, err := c.cmd.Eval(ctx, luaVerifyCode, []string{c.key(biz, to)}, code).Int()
	if err != nil {
		// 调用 redis 出了问题
		return false, err
	}

	switch res {
	case -4:
		return false, ErrCodeNotFound
	case -3:
		return false, ErrUserBlocked
	case -2:
		return false, ErrCodeIncorrect
	case -1:
		return false, ErrCodeVerifyTooMany
	default:
		return true, nil
	}
}

func (c *codeCache) key(biz, to string) string {
	// package:2224693191@qq.com:register
	return fmt.Sprintf("package:%s:%s", to, biz)
}
