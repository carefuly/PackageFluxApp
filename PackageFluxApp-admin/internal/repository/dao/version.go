/**
 * Description：
 * FileName：version.go
 * Author：CJiaの用心
 * Create：2025/2/25 15:51:11
 * Remark：
 */

package dao

import (
	"context"
	"errors"
	"github.com/carefuly/PackageFluxApp/internal/domain"
	"github.com/carefuly/PackageFluxApp/model"
	"gorm.io/gorm"
)

var (
	ErrDuplicateVersionCode = errors.New("相同的详情ID和版本号已存在")
	ErrVersionNotFound      = gorm.ErrRecordNotFound
	ErrFormalNotFound       = errors.New("没有找到正式版")
)

type VersionDAO interface {
	Create(ctx context.Context, version model.Version) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string, version model.Version) error
	FindLatestFormalVersion(ctx context.Context, detailId string) (domain.Version, error)
	Formal(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (model.Version, error)
	FindListAll(ctx context.Context, f domain.FiltersVersion) ([]model.Version, error)
	ExistsByDetailIdAndVersionCode(ctx context.Context, detailId, versionCode string) (bool, error)
	ExistsByDetailIdAndVersionCodeExcludeCurrent(ctx context.Context, detailID, versionCode, id string) (bool, error)
}

type versionDAO struct {
	db *gorm.DB
}

func NewVersionDAO(db *gorm.DB) VersionDAO {
	return &versionDAO{
		db: db,
	}
}

func (dao *versionDAO) Create(ctx context.Context, version model.Version) error {
	return dao.db.WithContext(ctx).Create(&version).Error
}

func (dao *versionDAO) Delete(ctx context.Context, id string) error {
	return dao.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Version{}).Error
}

func (dao *versionDAO) Update(ctx context.Context, id string, version model.Version) error {
	return dao.db.WithContext(ctx).Model(&version).Where("id = ?", id).
		Updates(map[string]any{
			"description":      version.Description,
			"h_builder_update": version.HBuilderUpdate,
			"force_update":     version.ForceUpdate,
			"formal_version":   version.FormalVersion,
			"apk_url":          version.ApkUrl,
			"wgt_url":          version.WgtUrl,
			"modifier":         version.Modifier,
			"remark":           version.Remark,
		}).Error
}

// FindLatestFormalVersion 查询最新的正式版
func (dao *versionDAO) FindLatestFormalVersion(ctx context.Context, detailId string) (domain.Version, error) {
	var version domain.Version
	if err := dao.db.WithContext(ctx).Where("formal_version = ? AND detail_id = ?", true, detailId).
		Order("version_code DESC").First(&version).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Version{}, nil // 没有正式版
		}
		return domain.Version{}, err
	}
	return version, nil
}

func (dao *versionDAO) Formal(ctx context.Context, id string) error {
	return dao.db.WithContext(ctx).Model(&model.Version{}).Where("id = ?", id).Update("formal_version", true).Error
}

func (dao *versionDAO) FindById(ctx context.Context, id string) (model.Version, error) {
	var version model.Version
	err := dao.db.WithContext(ctx).Where("id = ?", id).First(&version).Error
	if err != nil {
		if errors.Is(err, ErrVersionNotFound) {
			return model.Version{}, ErrRecordNotFound
		} else {
			return model.Version{}, err
		}
	}
	return version, err
}

func (dao *versionDAO) FindListAll(ctx context.Context, f domain.FiltersVersion) ([]model.Version, error) {
	query := dao.db.WithContext(ctx).Where("status = ? AND detail_id = ?", f.Status, f.DetailId).
		Order("create_time DESC, sort ASC")
	var vs []model.Version
	if err := query.Find(&vs).Error; err != nil {
		return []model.Version{}, err
	}
	return vs, nil
}

func (dao *versionDAO) ExistsByDetailIdAndVersionCode(ctx context.Context, detailId, versionCode string) (bool, error) {
	var count int64
	err := dao.db.WithContext(ctx).Model(&domain.Version{}).
		Where("detail_id = ? AND version_code = ?", detailId, versionCode).Count(&count).Error
	return count > 0, err
}

func (dao *versionDAO) ExistsByDetailIdAndVersionCodeExcludeCurrent(ctx context.Context, detailID, versionCode, id string) (bool, error) {
	var count int64
	err := dao.db.WithContext(ctx).Model(&domain.Version{}).
		Where("detail_id = ? AND version_code = ? AND id != ?", detailID, versionCode, id).
		Count(&count).
		Error
	return count > 0, err
}
