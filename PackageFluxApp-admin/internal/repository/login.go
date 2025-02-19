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

func (repo *loginRepository) toDomain(u model.User) domain.Login {
	return domain.Login{
		RecordId: u.CoreModels.RecordId,
		Email:    u.Email.String,
		Password: u.Password.String,
		// Id:       u.Id,
		// Email:    u.Email.String,
		// Phone:    u.Phone.String,
		// Password: u.Password,
		// AboutMe:  u.AboutMe,
		// Nickname: u.Nickname,
		// Birthday: time.UnixMilli(u.Birthday),
	}
}
