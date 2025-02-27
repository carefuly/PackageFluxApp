/**
 * Description：
 * FileName：detail.go
 * Author：CJiaの用心
 * Create：2025/2/27 10:35:33
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
	ErrDetailRecordNotFound      = gorm.ErrRecordNotFound
	ErrDetailNotFound            = errors.New("记录不存在")
	ErrDuplicateUserIdAndAppName = errors.New("同一用户下应用名称不能重复")
)

type DetailDAO interface {
	Insert(ctx context.Context, d model.Detail) error
	Delete(ctx context.Context, id, userId string) (int64, error)
	Update(ctx context.Context, id, userId string, d model.Detail) (int64, error)
	FindByIdAndUserId(ctx context.Context, id, userId string) (*model.Detail, error)
	FindListAll(ctx context.Context, userId, appName string) (int64, []*model.Detail, error)
	ExistsByUserIdAndAppName(ctx context.Context, userId, appName string) (bool, error)
}

type GORMDetailDAO struct {
	db *gorm.DB
}

func NewGORMDetailDAO(db *gorm.DB) DetailDAO {
	return &GORMDetailDAO{
		db: db,
	}
}

func (dao *GORMDetailDAO) Insert(ctx context.Context, d model.Detail) error {
	return dao.db.WithContext(ctx).Create(&d).Error
}

func (dao *GORMDetailDAO) Delete(ctx context.Context, id, userId string) (int64, error) {
	result := dao.db.WithContext(ctx).Where("id = ? AND user_id = ?", id, userId).Delete(&model.Detail{})
	return result.RowsAffected, result.Error
}

func (dao *GORMDetailDAO) Update(ctx context.Context, id, userId string, d model.Detail) (int64, error) {
	result := dao.db.WithContext(ctx).Model(&d).Where("id = ? AND user_id = ?", id, userId).
		Updates(map[string]any{
			"logoUrl":     d.LogoUrl,
			"appName":     d.AppName,
			"appleId":     d.AppleId,
			"description": d.Description,
			"preview":     d.Preview,
			"modifier":    d.Modifier,
			"remark":      d.Remark,
		})
	return result.RowsAffected, result.Error
}

func (dao *GORMDetailDAO) FindByIdAndUserId(ctx context.Context, id, userId string) (*model.Detail, error) {
	var detail model.Detail
	err := dao.db.WithContext(ctx).Where("id = ? AND user_id = ?", id, userId).First(&detail).Error
	return &detail, err
}

func (dao *GORMDetailDAO) FindListAll(ctx context.Context, userId, appName string) (int64, []*model.Detail, error) {
	query := dao.db.WithContext(ctx).Where("status = ? AND user_id = ?", true, userId).
		Order("update_time DESC, sort ASC")
	if appName != "" {
		query = query.Where("app_name LIKE ?", "%"+appName+"%")
	}
	var details []*model.Detail
	result := query.Find(&details)

	return result.RowsAffected, details, result.Error
}

func (dao *GORMDetailDAO) ExistsByUserIdAndAppName(ctx context.Context, userId, appName string) (bool, error) {
	var count int64
	err := dao.db.WithContext(ctx).Model(&model.Detail{}).
		Where("user_id = ? AND appName = ?", userId, appName).
		Count(&count).Error
	return count > 0, err
}
