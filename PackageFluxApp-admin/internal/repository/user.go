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
	"go.uber.org/zap"
)

var (
	ErrDuplicateEmail  = dao.ErrDuplicateEmail
	ErrUserIdNotFound  = dao.ErrUserNotFound
	ErrUserKeyNotExist = cache.ErrUserKeyNotExist
)

type UserRepository interface {
	Create(ctx context.Context, u domain.Register) error
	UpdateUsageNumber(ctx context.Context, userId string) error

	FindById(ctx context.Context, email string) (domain.User, error)
	FindByEmail(ctx context.Context, email string) (domain.User, error)

	ExistsByEmail(ctx context.Context, email string) (bool, error)
}

type userRepository struct {
	dao   dao.UserDAO
	cache cache.UserCache
}

func NewUserRepository(dao dao.UserDAO, cache cache.UserCache) UserRepository {
	return &userRepository{
		dao:   dao,
		cache: cache,
	}
}

func (repo *userRepository) Create(ctx context.Context, u domain.Register) error {
	return repo.dao.Insert(ctx, repo.toEntity(u))
}

func (repo *userRepository) UpdateUsageNumber(ctx context.Context, userId string) error {
	err := repo.dao.UpdateUsageNumber(ctx, userId)
	if err != nil {
		return err
	}

	err = repo.cache.Del(ctx, userId)
	if err != nil {
		// 网络崩了，也可能是 redis 崩了
		zap.L().Error("Redis异常", zap.Error(err))
		return err
	}

	return err
}

func (repo *userRepository) FindById(ctx context.Context, id string) (domain.User, error) {
	u, err := repo.cache.Get(ctx, id)

	if err == nil {
		return *u, nil
	}

	user, err := repo.dao.FindById(ctx, id)
	if err != nil {
		return domain.User{}, err
	}

	userinfo := repo.toDomain(user)

	err = repo.cache.Set(ctx, userinfo)
	if err != nil {
		// 网络崩了，也可能是 redis 崩了
		zap.L().Error("Redis异常", zap.Error(err))
		return domain.User{}, err
	}

	return userinfo, err
}

func (repo *userRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	user, err := repo.dao.FindByEmail(ctx, email)
	return repo.toDomain(user), err
}

func (repo *userRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	return repo.dao.ExistsByEmail(ctx, email)
}

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

func (repo *userRepository) toDomain(u *model.User) domain.User {
	return domain.User{
		User:       *u,
		Email:      u.Email.String,
		Phone:      u.Phone.String,
		CreateTime: u.CoreModels.CreateTime.Format("2006-01-02 15:04:05.000"),
		UpdateTime: u.CoreModels.UpdateTime.Format("2006-01-02 15:04:05.000"),
	}
}
