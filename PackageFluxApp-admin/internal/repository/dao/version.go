/**
 * Description：
 * FileName：version.go
 * Author：CJiaの用心
 * Create：2025/2/28 13:11:34
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
	ErrVersionRecordNotFound           = gorm.ErrRecordNotFound
	ErrVersionNotFound                 = errors.New("记录不存在")
	ErrDuplicateDetailIdAndVersionCode = errors.New("同一应用下应用版本号不能重复")
)

type VersionDAO interface {
	Insert(ctx context.Context, v model.Version) error
	Delete(ctx context.Context, id, detailId, userId string) (int64, error)
	Update(ctx context.Context, id, detailId, userId string, v model.Version) (int64, error)
	UpdateFormal(ctx context.Context, id, detailId, userId string) (int64, error)
	FindByIdAndDetailIdAndUserId(ctx context.Context, id, detailId, userId string) (*model.Version, error)
	FindListAll(ctx context.Context, detailId, userId string) (int64, []*model.Version, error)
	ExistsByDetailIdAndCode(ctx context.Context, userId, detailId, versionCode string) (bool, error)
}

type GORMVersionDAO struct {
	db *gorm.DB
}

func NewGORMVersionDAO(db *gorm.DB) VersionDAO {
	return &GORMVersionDAO{
		db: db,
	}
}

func (dao *GORMVersionDAO) Insert(ctx context.Context, v model.Version) error {
	return dao.db.WithContext(ctx).Create(&v).Error
}

func (dao *GORMVersionDAO) Delete(ctx context.Context, id, detailId, userId string) (int64, error) {
	result := dao.db.WithContext(ctx).
		Where("id = ? AND detail_id = ? AND user_id = ?", id, detailId, userId).
		Delete(&model.Version{})
	return result.RowsAffected, result.Error
}

func (dao *GORMVersionDAO) Update(ctx context.Context, id, detailId, userId string, v model.Version) (int64, error) {
	result := dao.db.WithContext(ctx).Model(&v).Where("id = ? AND detail_id = ? AND user_id = ?", id, detailId, userId).
		Updates(map[string]any{
			"description":    v.Description,
			"HBuilderUpdate": v.HBuilderUpdate,
			"forceUpdate":    v.ForceUpdate,
			"apkUrl":         v.ApkUrl,
			"wgtUrl":         v.WgtUrl,
			"modifier":       v.Modifier,
			"remark":         v.Remark,
		})
	return result.RowsAffected, result.Error
}

func (dao *GORMVersionDAO) UpdateFormal(ctx context.Context, id, detailId, userId string) (int64, error) {
	dao.db.WithContext(ctx).Model(&model.Version{}).
		Where("detail_id = ? AND user_id = ?", detailId, userId).
		Update("formalVersion", false)
	result := dao.db.WithContext(ctx).Model(&model.Version{}).
		Where("id = ? AND detail_id = ? AND user_id = ?", id, detailId, userId).
		Update("formalVersion", true)
	return result.RowsAffected, result.Error
}

func (dao *GORMVersionDAO) FindByIdAndDetailIdAndUserId(ctx context.Context, id, detailId, userId string) (*model.Version, error) {
	var version model.Version
	err := dao.db.WithContext(ctx).
		Where("id = ? AND detail_id = ? AND user_id = ?", id, detailId, userId).
		First(&version).Error
	return &version, err
}

func (dao *GORMVersionDAO) FindListAll(ctx context.Context, detailId, userId string) (int64, []*model.Version, error) {
	query := dao.db.WithContext(ctx).
		Where("status = ? AND detail_id = ? AND user_id = ?", true, detailId, userId).
		Order("update_time DESC, sort ASC")
	var versions []*model.Version
	result := query.Find(&versions)
	return result.RowsAffected, versions, result.Error
}

func (dao *GORMVersionDAO) ExistsByDetailIdAndCode(ctx context.Context, userId, detailId, versionCode string) (bool, error) {
	var count int64
	err := dao.db.WithContext(ctx).Model(&model.Version{}).
		Where("user_id = ? AND detail_id = ? AND versionCode = ?", userId, detailId, versionCode).
		Count(&count).Error
	return count > 0, err
}
