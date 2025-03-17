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
	FindListPage(ctx context.Context, detailId, createTime string, page filters.Pagination) (int64, []*model.Logger, error)
	CountByDetailIdAndCreateTime(ctx context.Context, detailId, createTime string) (int64, error)
	CountDistinctUniqueIdByCreateTime(ctx context.Context, detailId, createTime string) (int64, error)
	CountByDetailId(ctx context.Context, detailId string) (int64, error)
	CountDistinctUniqueIdByDetailId(ctx context.Context, detailId string) (int64, error)
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

func (dao *GORMLoggerDAO) FindListPage(ctx context.Context, detailId, createTime string, page filters.Pagination) (int64, []*model.Logger, error) {
	var total int64
	var loggers []*model.Logger

	query := dao.db.WithContext(ctx).Model(&model.Logger{}).
		Where("status = ? AND detail_id = ?", true, detailId)

	if createTime != "" {
		// 使用 GORM 的 Where 条件来查询年月日
		query = query.Where("DATE(create_time) LIKE ?", "%"+createTime+"%")
	}

	// 计算总数
	if err := query.Count(&total).Error; err != nil {
		return 0, nil, err
	}

	// 分页查询
	offset := (page.Page - 1) * page.PageSize
	if err := query.Offset(offset).Limit(page.PageSize).
		Order("update_time DESC, sort ASC").
		Find(&loggers).Error; err != nil {
		return 0, nil, err
	}

	return total, loggers, nil
}

func (dao *GORMLoggerDAO) CountByDetailIdAndCreateTime(ctx context.Context, detailId, createTime string) (int64, error) {
	var count int64
	query := dao.db.WithContext(ctx).Model(&model.Logger{}).
		Where("detail_id = ? AND DATE(create_time) = ?", detailId, createTime).
		Count(&count)
	if query.Error != nil {
		return 0, query.Error
	}
	return count, nil
}

func (dao *GORMLoggerDAO) CountDistinctUniqueIdByCreateTime(ctx context.Context, detailId, createTime string) (int64, error) {
	var count int64
	query := dao.db.WithContext(ctx).Model(&model.Logger{}).
		Where("detail_id = ? AND DATE(create_time) = ?", detailId, createTime).
		Distinct("uniqueId").
		Count(&count)
	if query.Error != nil {
		return 0, query.Error
	}
	return count, nil
}

func (dao *GORMLoggerDAO) CountByDetailId(ctx context.Context, detailId string) (int64, error) {
	var count int64
	query := dao.db.WithContext(ctx).Model(&model.Logger{}).
		Where("detail_id = ?", detailId).
		Count(&count)
	if query.Error != nil {
		return 0, query.Error
	}
	return count, nil
}

func (dao *GORMLoggerDAO) CountDistinctUniqueIdByDetailId(ctx context.Context, detailId string) (int64, error) {
	var count int64
	query := dao.db.WithContext(ctx).Model(&model.Logger{}).
		Where("detail_id = ?", detailId).
		Distinct("uniqueId").
		Count(&count)
	if query.Error != nil {
		return 0, query.Error
	}
	return count, nil
}
