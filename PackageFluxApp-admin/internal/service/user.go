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
	"github.com/carefuly/PackageFluxApp/internal/domain"
	"github.com/carefuly/PackageFluxApp/internal/repository"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrDuplicateEmail = repository.ErrDuplicateEmail
)

type UserService interface {
	Register(ctx context.Context, u domain.Register) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
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

	return nil
}

func (svc *userService) IsDuplicateEntryError(err error) bool {
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		// MySQL 错误码 1062 表示唯一冲突
		return mysqlErr.Number == 1062
	}
	return false
}