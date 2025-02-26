/**
 * Description：
 * FileName：user.go
 * Author：CJiaの用心
 * Create：2025/2/24 15:43:14
 * Remark：
 */

package service

import (
	"context"
	"errors"
	"github.com/carefuly/PackageFluxApp/internal/domain"
	"github.com/carefuly/PackageFluxApp/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrDuplicateEmail        = repository.ErrDuplicateUser
	ErrInvalidUserOrPassword = errors.New("用户不存在/密码错误")
)

type UserService interface {
	Register(ctx context.Context, u domain.Register) error
	Login(ctx context.Context, u domain.Login) (domain.Login, error)
	UserInfo(ctx context.Context, uid string) (domain.UserInfo, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

type UserClaims struct {
	jwt.RegisteredClaims
	UId       string
	Email     string
	UserAgent string
}

func (svc *userService) Register(ctx context.Context, u domain.Register) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordStr = u.Password
	u.Password = string(hash)
	return svc.repo.Register(ctx, u)
}

func (svc *userService) Login(ctx context.Context, user domain.Login) (domain.Login, error) {
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

func (svc *userService) UserInfo(ctx context.Context, uid string) (domain.UserInfo, error) {
	return svc.repo.UserInfo(ctx, uid)
}
