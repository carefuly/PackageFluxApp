/**
 * Description：
 * FileName：user.go
 * Author：CJiaの用心
 * Create：2025/2/24 15:41:42
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
	ErrUserNotFound = dao.ErrUserNotFound
)

type UserRepository interface {
	Register(ctx context.Context, u domain.Register) error
	Login(ctx context.Context, email string) (domain.Login, error)
	UserInfo(ctx context.Context, uid string) (domain.UserInfo, error)
}

type userRepository struct {
	dao dao.UserDAO
}

func NewUserRepository(dao dao.UserDAO) UserRepository {
	return &userRepository{
		dao: dao,
	}
}

func (repo *userRepository) Register(ctx context.Context, u domain.Register) error {
	return repo.dao.Insert(ctx, repo.toEntity(u))
}

func (repo *userRepository) Login(ctx context.Context, email string) (domain.Login, error) {
	u, err := repo.dao.Login(ctx, email)
	if err != nil {
		return domain.Login{}, err
	}
	return repo.toDomain(u), nil
}

func (repo *userRepository) UserInfo(ctx context.Context, uid string) (domain.UserInfo, error) {
	u, err := repo.dao.UserInfo(ctx, uid)
	if err != nil {
		return domain.UserInfo{}, err
	}
	return domain.UserInfo{
		User: model.User{
			CoreModels:  u.CoreModels,
			Username:    u.Username,
			Avatar:      u.Avatar,
			UsageNumber: u.UsageNumber,
			Total:       u.Total,
		},
		Email:      u.Email.String,
		CreateTime: u.CoreModels.CreateTime.Format("2006-01-02 15:04:05.000"),
		UpdateTime: u.CoreModels.UpdateTime.Format("2006-01-02 15:04:05.000"),
	}, nil
}

func (repo *userRepository) toDomain(u model.User) domain.Login {
	return domain.Login{
		ID:       u.CoreModels.ID,
		Email:    u.Email.String,
		Password: u.Password.String,
	}
}


func (repo *userRepository) toEntity(u domain.Register) model.User {
	return model.User{
		Username: u.Username,
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
