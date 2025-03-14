/**
 * Description：
 * FileName：logger.go
 * Author：CJiaの用心
 * Create：2025/3/14 15:26:21
 * Remark：
 */

package dao

import (
	"context"
	"github.com/carefuly/PackageFluxApp/internal/model"
	"github.com/carefuly/PackageFluxApp/pkg/ginx/query/filters"
	"gorm.io/gorm"
)

type LoggerDAO interface {
	Insert(ctx context.Context, l model.Logger) error
	FindListPage(ctx context.Context, detailId string, page filters.Pagination) (int64, []*model.Logger, error)
}

type GORMLoggerDAO struct {
	db *gorm.DB
}

func NewGORMLoggerDAO(db *gorm.DB) LoggerDAO {
	return &GORMLoggerDAO{
		db: db,
	}
}

func (dao *GORMLoggerDAO) Insert(ctx context.Context, l model.Logger) error {
	return dao.db.WithContext(ctx).Create(&l).Error
}

func (dao *GORMLoggerDAO) FindListPage(ctx context.Context, detailId string, page filters.Pagination) (int64, []*model.Logger, error) {
	var total int64
	var loggers []*model.Logger

	if err := dao.db.WithContext(ctx).Model(&model.Logger{}).
		Where("status = ? AND detail_id = ?", true, detailId).
		Count(&total).Error; err != nil {
		return 0, nil, err
	}

	offset := (page.Page - 1) * page.PageSize
	if err := dao.db.WithContext(ctx).
		Where("status = ? AND detail_id = ?", true, detailId).
		Offset(offset).Limit(page.PageSize).
		Order("update_time DESC, sort ASC").
		Find(&loggers).Error; err != nil {
		return 0, nil, err
	}

	return total, loggers, nil
}
