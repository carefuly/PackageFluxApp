/**
 * Description：
 * FileName：login.go
 * Author：CJiaの用心
 * Create：2025/2/19 21:16:32
 * Remark：
 */

package dao

import (
	"context"
	"github.com/carefuly/PackageFluxApp/model"
	"gorm.io/gorm"
)

var ErrUserNotFound = gorm.ErrRecordNotFound

type LoginAO interface {
	Login(ctx context.Context, email string) (model.User, error)
	UserInfo(ctx context.Context, uid string) (model.User, error)
}

type loginDAO struct {
	db *gorm.DB
}

func NewLoginDAO(db *gorm.DB) LoginAO {
	return &loginDAO{
		db: db,
	}
}

func (dao *loginDAO) Login(ctx context.Context, email string) (model.User, error) {
	var u model.User
	err := dao.db.WithContext(ctx).Where("email=?", email).First(&u).Error
	return u, err
}

func (dao *loginDAO) UserInfo(ctx context.Context, uid string) (model.User, error) {
	var u model.User
	err := dao.db.WithContext(ctx).Where("record_id=?", uid).First(&u).Error
	return u, err
}
