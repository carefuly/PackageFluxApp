/**
 * Description：
 * FileName：user.go
 * Author：CJiaの用心
 * Create：2025/2/24 15:40:06
 * Remark：
 */

package dao

import (
	"context"
	"errors"
	"github.com/carefuly/PackageFluxApp/model"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var (
	ErrDuplicateEmail = errors.New("邮箱冲突")
	ErrUserNotFound   = gorm.ErrRecordNotFound
)

type UserDAO interface {
	Insert(ctx context.Context, u model.User) error
	Login(ctx context.Context, email string) (model.User, error)
	UserInfo(ctx context.Context, uid string) (model.User, error)
}

type userDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) UserDAO {
	return &userDAO{
		db: db,
	}
}

func (dao *userDAO) Insert(ctx context.Context, u model.User) error {
	err := dao.db.WithContext(ctx).Create(&u).Error
	var me *mysql.MySQLError
	if errors.As(err, &me) {
		const duplicateErr uint16 = 1062
		if me.Number == duplicateErr {
			// 用户冲突，邮箱冲突
			return ErrDuplicateEmail
		}
	}
	return err
}

func (dao *userDAO) Login(ctx context.Context, email string) (model.User, error) {
	var u model.User
	err := dao.db.WithContext(ctx).Where("email=?", email).First(&u).Error
	return u, err
}

func (dao *userDAO) UserInfo(ctx context.Context, uid string) (model.User, error) {
	var u model.User
	err := dao.db.WithContext(ctx).Where("id=?", uid).First(&u).Error
	return u, err
}
