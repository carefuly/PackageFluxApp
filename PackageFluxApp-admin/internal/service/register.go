/**
 * Description：
 * FileName：register.go
 * Author：CJiaの用心
 * Create：2025/2/19 16:30:51
 * Remark：
 */

package service

import (
	"context"
	"github.com/carefuly/PackageFluxApp/internal/domain"
	"github.com/carefuly/PackageFluxApp/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrDuplicateEmail = repository.ErrDuplicateUser
)

type RegisterService interface {
	Register(ctx context.Context, u domain.Register) error
}

type registerService struct {
	repo repository.RegisterRepository
}

func NewRegisterService(repo repository.RegisterRepository) RegisterService {
	return &registerService{
		repo: repo,
	}
}

func (svc *registerService) Register(ctx context.Context, u domain.Register) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordStr = u.Password
	u.Password = string(hash)
	return svc.repo.Register(ctx, u)
}
