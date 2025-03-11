/**
 * Description：
 * FileName：user.go
 * Author：CJiaの用心
 * Create：2025/2/26 11:35:37
 * Remark：
 */

package dao

import (
	"context"
	"errors"
	"github.com/carefuly/PackageFluxApp/internal/model"
	"gorm.io/gorm"
)

var (
	ErrDuplicateEmail = errors.New("邮箱已存在")
	ErrUserNotFound   = gorm.ErrRecordNotFound
)

type UserDAO interface {
	Insert(ctx context.Context, u model.User) error
	UpdateUsageNumber(ctx context.Context, userId string) error

	FindById(ctx context.Context, id string) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)

	ExistsByEmail(ctx context.Context, email string) (bool, error)
}

type GORMUserDAO struct {
	db *gorm.DB
}

func NewGORMUserDAO(db *gorm.DB) UserDAO {
	return &GORMUserDAO{
		db: db,
	}
}

func (dao *GORMUserDAO) Insert(ctx context.Context, u model.User) error {
	return dao.db.WithContext(ctx).Create(&u).Error
}

func (dao *GORMUserDAO) UpdateUsageNumber(ctx context.Context, userId string) error {
	return dao.db.WithContext(ctx).Model(&model.User{}).
		Where("id = ?", userId).
		Update("usage_number", gorm.Expr("usage_number + ?", 1)).
		Error
}

func (dao *GORMUserDAO) FindById(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	err := dao.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	return &user, err
}

func (dao *GORMUserDAO) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := dao.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	return &user, err
}

func (dao *GORMUserDAO) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	err := dao.db.WithContext(ctx).Model(&model.User{}).
		Where("email = ?", email).
		Count(&count).Error
	return count > 0, err
}
