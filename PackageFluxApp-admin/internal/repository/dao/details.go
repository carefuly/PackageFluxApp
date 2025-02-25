/**
 * Description：
 * FileName：details.go
 * Author：CJiaの用心
 * Create：2025/2/24 16:10:31
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
	ErrDuplicateAppName = errors.New("应用名冲突")
	ErrRecordNotFound   = gorm.ErrRecordNotFound
)

type DetailsDAO interface {
	Create(ctx context.Context, detail model.Detail) error
	Delete(ctx context.Context, id, userId string) error
	Update(ctx context.Context, id string, detail model.Detail) error
	Formal(ctx context.Context, id, versionId string) error
	FindById(ctx context.Context, id, userId string) (model.Detail, error)
	FindListAll(ctx context.Context, f domain.FiltersDetail) ([]model.Detail, error)
}

type detailsDAO struct {
	db *gorm.DB
}

func NewDetailsDAO(db *gorm.DB) DetailsDAO {
	return &detailsDAO{
		db: db,
	}
}

func (dao *detailsDAO) Create(ctx context.Context, detail model.Detail) error {
	var d model.Detail
	err := dao.db.WithContext(ctx).Where("user_id = ? AND app_name = ?", detail.UserID, detail.AppName).First(&d).Error
	switch {
	case errors.Is(err, ErrRecordNotFound):
		err := dao.db.WithContext(ctx).Create(&detail).Error
		return err
	default:
		return ErrDuplicateAppName
	}
}

func (dao *detailsDAO) Delete(ctx context.Context, id, userId string) error {
	return dao.db.WithContext(ctx).Where("id = ? AND user_id = ?", userId, id).Delete(&model.Detail{}).Error
}

func (dao *detailsDAO) Update(ctx context.Context, id string, detail model.Detail) error {
	return dao.db.WithContext(ctx).Model(&detail).Where("id = ?", id).
		Updates(map[string]any{
			"logo_url":    detail.LogoUrl,
			"app_name":    detail.AppName,
			"apple_id":    detail.AppleId,
			"description": detail.Description,
			"preview":     detail.Preview,
			"modifier":    detail.Modifier,
			"remark":      detail.Remark,
		}).Error
}

func (dao *detailsDAO) Formal(ctx context.Context, id, versionId string) error {
	return dao.db.WithContext(ctx).Model(&model.Detail{}).Where("id = ?", id).
		Update("version_id", versionId).Error
}

func (dao *detailsDAO) FindById(ctx context.Context, id, userId string) (model.Detail, error) {
	var res model.Detail
	err := dao.db.WithContext(ctx).Preload("User").
		Where("id = ? AND user_id = ?", id, userId).
		First(&res).Error
	if err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			return model.Detail{}, ErrRecordNotFound
		} else {
			return model.Detail{}, err
		}
	}
	return res, err
}

func (dao *detailsDAO) FindListAll(ctx context.Context, f domain.FiltersDetail) ([]model.Detail, error) {
	query := dao.db.WithContext(ctx).Where("status = ? AND user_id = ?", f.Status, f.UserId).
		Order("update_time DESC, sort ASC")
	if f.AppName != "" {
		query = query.Where("app_name LIKE ?", "%"+f.AppName+"%")
	}
	var ds []model.Detail
	if err := query.Find(&ds).Error; err != nil {
		return []model.Detail{}, err
	}
	return ds, nil
}
