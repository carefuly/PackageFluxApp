/**
 * Description：
 * FileName：register.go
 * Author：CJiaの用心
 * Create：2025/2/19 16:41:04
 * Remark：
 */

package repository

import (
	"context"
	"database/sql"
	"github.com/carefuly/PackageFluxApp/internal/domain"
	"github.com/carefuly/PackageFluxApp/internal/repository/dao"
	"github.com/carefuly/PackageFluxApp/model"
)

var (
	ErrDuplicateUser = dao.ErrDuplicateEmail
)

type RegisterRepository interface {
	Register(ctx context.Context, u domain.Register) error
}

type registerRepository struct {
	dao dao.RegisterDAO
}

func NewRegisterRepository(dao dao.RegisterDAO) RegisterRepository {
	return &registerRepository{
		dao: dao,
	}
}

func (repo *registerRepository) Register(ctx context.Context, u domain.Register) error {
	return repo.dao.Insert(ctx, repo.toEntity(u))
}

func (repo *registerRepository) toEntity(u domain.Register) model.User {
	return model.User{
		Username: "",
		Email: sql.NullString{
			String: u.Email,
			Valid:  u.Email != "",
		},
		Password: sql.NullString{
			String: u.Password,
			Valid:  u.Password != "",
		},
		PasswordStr: sql.NullString{
			String: u.PasswordStr,
			Valid:  u.PasswordStr != "",
		},
	}
}
