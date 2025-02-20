/**
 * Description：
 * FileName：login.go
 * Author：CJiaの用心
 * Create：2025/2/19 21:12:47
 * Remark：
 */

package repository

import (
	"github.com/carefuly/PackageFluxApp/internal/domain"
	"github.com/carefuly/PackageFluxApp/internal/repository/dao"
	"github.com/carefuly/PackageFluxApp/model"
	"golang.org/x/net/context"
)

var ErrUserNotFound = dao.ErrUserNotFound

type LoginRepository interface {
	Login(ctx context.Context, email string) (domain.Login, error)
	UserInfo(ctx context.Context, uid string) (domain.UserInfo, error)
}

type loginRepository struct {
	dao dao.LoginAO
}

func NewLoginRepository(dao dao.LoginAO) LoginRepository {
	return &loginRepository{
		dao: dao,
	}
}

func (repo *loginRepository) Login(ctx context.Context, email string) (domain.Login, error) {
	u, err := repo.dao.Login(ctx, email)
	if err != nil {
		return domain.Login{}, err
	}
	return repo.toDomain(u), nil
}

func (repo *loginRepository) UserInfo(ctx context.Context, uid string) (domain.UserInfo, error) {
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
		Email: u.Email.String,
		CreateTime: u.CoreModels.CreateTime.Format("2006-01-02 15:04:05.000"),
		UpdateTime: u.CoreModels.UpdateTime.Format("2006-01-02 15:04:05.000"),
	}, nil
}

func (repo *loginRepository) toDomain(u model.User) domain.Login {
	return domain.Login{
		RecordId: u.CoreModels.RecordId,
		Email:    u.Email.String,
		Password: u.Password.String,
	}
}
