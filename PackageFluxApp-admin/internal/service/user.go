/**
 * Description：
 * FileName：user.go
 * Author：CJiaの用心
 * Create：2025/2/26 14:36:28
 * Remark：
 */

package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/carefuly/PackageFluxApp/config"
	"github.com/carefuly/PackageFluxApp/internal/domain"
	"github.com/carefuly/PackageFluxApp/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var (
	ErrDuplicateEmail        = repository.ErrDuplicateEmail
	ErrUserIdNotFound        = repository.ErrUserIdNotFound
	ErrInvalidUserOrPassword = errors.New("邮箱或密码错误")
	ErrGenerateTokenError    = errors.New("生成Token异常")
)

type UserService interface {
	Register(ctx context.Context, u domain.Register) error
	Login(ctx *gin.Context, rely config.RelyConfig, u domain.Login) (string, error)
	UserInfo(ctx context.Context, id string) (*domain.User, error)
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
	exists, err := svc.repo.ExistsByEmail(ctx, u.Email)
	if err != nil {
		return err
	}
	if exists {
		return repository.ErrDuplicateEmail
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)

	err = svc.repo.Create(ctx, u)
	if svc.IsDuplicateEntryError(err) {
		return repository.ErrDuplicateEmail
	}

	return err
}

func (svc *userService) Login(ctx *gin.Context, rely config.RelyConfig, u domain.Login) (string, error) {
	user, err := svc.repo.FindByEmail(ctx, u.Email)
	fmt.Println(user)
	fmt.Println(err)
	if err != nil {
		return "", ErrInvalidUserOrPassword
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if err != nil {
		return "", ErrInvalidUserOrPassword
	}

	return svc.setJWTToken(ctx, rely, user)
}

func (svc *userService) UserInfo(ctx context.Context, id string) (*domain.User, error) {
	user, err := svc.repo.FindById(ctx, id)
	switch {
	case err == nil:
		return &user, nil
	case errors.Is(err, repository.ErrUserIdNotFound):
		return nil, repository.ErrUserIdNotFound
	default:
		return nil, err
	}
}

func (svc *userService) Logout(ctx *gin.Context, rely config.RelyConfig, u domain.Login) (string, error) {
	return "", nil
}

func (svc *userService) IsDuplicateEntryError(err error) bool {
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		// MySQL 错误码 1062 表示唯一冲突
		return mysqlErr.Number == 1062
	}
	return false
}

func (svc *userService) setJWTToken(ctx *gin.Context, rely config.RelyConfig, u domain.User) (string, error) {
	uc := UserClaims{
		UId:       u.Id,
		Email:     u.Email,
		UserAgent: ctx.GetHeader("User-Agent"),
		RegisteredClaims: jwt.RegisteredClaims{
			// 一小时过期
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 60 * 1)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, uc)
	tokenStr, err := token.SignedString([]byte(rely.Token.ApiKeyAuth))

	if err != nil {
		return "", ErrGenerateTokenError
	}
	return tokenStr, err
}
