/**
 * Description：
 * FileName：register.go
 * Author：CJiaの用心
 * Create：2025/2/19 16:40:27
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
)

type RegisterDAO interface {
	Insert(ctx context.Context, u model.User) error
}

type registerDAO struct {
	db *gorm.DB
}

func NewRegisterDAO(db *gorm.DB) RegisterDAO {
	return &registerDAO{
		db: db,
	}
}

func (dao *registerDAO) Insert(ctx context.Context, u model.User) error {
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
