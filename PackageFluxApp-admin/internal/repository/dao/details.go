/**
 * Description：
 * FileName：details.go
 * Author：CJiaの用心
 * Create：2025/2/20 21:02:32
 * Remark：
 */

package dao

import (
	"context"
	"errors"
	"github.com/carefuly/PackageFluxApp/model"
	"gorm.io/gorm"
)

var (
	ErrDuplicateAppName = errors.New("应用名冲突")
	ErrRecordNotFound   = gorm.ErrRecordNotFound
)

type DetailsDAO interface {
	Create(ctx context.Context, detail model.Detail) error
	Delete(ctx context.Context, userId uint, recordId string) error
	Update(ctx context.Context, detail model.Detail) error
	FindById(ctx context.Context, userId uint, recordId string) (model.Detail, error)
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

func (dao *detailsDAO) Delete(ctx context.Context, userId uint, recordId string) error {
	return dao.db.WithContext(ctx).Where("user_id = ? AND record_id = ?", userId, recordId).Delete(&model.Detail{}).Error
}

func (dao *detailsDAO) Update(ctx context.Context, detail model.Detail) error {
	return dao.db.WithContext(ctx).Model(&detail).Where("record_id = ?", detail.RecordId).
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

func (dao *detailsDAO) FindById(ctx context.Context, userId uint, recordId string) (model.Detail, error) {
	var res model.Detail
	err := dao.db.WithContext(ctx).Where("user_id = ? AND record_id = ?", userId, recordId).First(&res).Error
	if err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			return model.Detail{}, ErrRecordNotFound
		} else {
			return model.Detail{}, err
		}
	}
	return res, err
}
