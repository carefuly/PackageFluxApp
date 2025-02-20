/**
 * Description：
 * FileName：login.go
 * Author：CJiaの用心
 * Create：2025/2/19 21:00:14
 * Remark：
 */

package service

import (
	"context"
	"errors"
	"github.com/carefuly/PackageFluxApp/internal/domain"
	"github.com/carefuly/PackageFluxApp/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidUserOrPassword = errors.New("用户不存在/密码错误")
)

type LoginService interface {
	Login(ctx context.Context, u domain.Login) (domain.Login, error)
	UserInfo(ctx context.Context, uid string) (domain.UserInfo, error)
}

type loginService struct {
	repo repository.LoginRepository
}

func NewLoginService(repo repository.LoginRepository) LoginService {
	return &loginService{
		repo: repo,
	}
}

func (svc *loginService) Login(ctx context.Context, user domain.Login) (domain.Login, error) {
	u, err := svc.repo.Login(ctx, user.Email)
	if errors.Is(err, repository.ErrUserNotFound) {
		return domain.Login{}, ErrInvalidUserOrPassword
	}
	if err != nil {
		return domain.Login{}, err
	}
	// 检查密码对不对
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))
	if err != nil {
		return domain.Login{}, ErrInvalidUserOrPassword
	}
	return u, nil
}

func (svc *loginService) UserInfo(ctx context.Context, uid string) (domain.UserInfo, error) {
	return svc.repo.UserInfo(ctx, uid)
}
