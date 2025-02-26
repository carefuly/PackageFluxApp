/**
 * Description：
 * FileName：user.go
 * Author：CJiaの用心
 * Create：2025/2/26 14:18:41
 * Remark：
 */

package repository

import (
	"context"
	"database/sql"
	"github.com/carefuly/PackageFluxApp/internal/domain"
	"github.com/carefuly/PackageFluxApp/internal/model"
	"github.com/carefuly/PackageFluxApp/internal/repository/cache"
	"github.com/carefuly/PackageFluxApp/internal/repository/dao"
)

var (
	ErrDuplicateEmail = dao.ErrDuplicateEmail
)

type UserRepository interface {
	Create(ctx context.Context, u domain.Register) error
	ExistsByEmail(ctx context.Context, email string) (bool, error)
}

type userRepository struct {
	dao   dao.UserDAO
	cache cache.UserCache
}

func NewUserRepository(dao dao.UserDAO) UserRepository {
	return &userRepository{
		dao: dao,
		// cache: cache,
	}
}

func (repo *userRepository) Create(ctx context.Context, u domain.Register) error {
	return repo.dao.Insert(ctx, repo.toEntity(u))
}

func (repo *userRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	return repo.dao.ExistsByEmail(ctx, email)
}

// func (repo *userRepository) Create(ctx context.Context, u domain.Register) error {
// 	return repo.dao.Insert(ctx, repo.toEntity(u))
// }
//
// func (repo *userRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
// 	return repo.dao.ExistsByEmail(ctx, email)
// }

func (repo *userRepository) toEntity(d domain.Register) model.User {
	return model.User{
		Username: d.Username,
		Email: sql.NullString{
			String: d.Email,
			Valid:  d.Email != "",
		},
		Password:    d.Password,
		PasswordStr: d.PasswordStr,
	}
}
