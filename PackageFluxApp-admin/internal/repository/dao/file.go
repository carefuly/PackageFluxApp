/**
 * Description：
 * FileName：file.go
 * Author：CJiaの用心
 * Create：2025/2/27 21:01:41
 * Remark：
 */

package dao

import (
	"context"
	"errors"
	"github.com/carefuly/PackageFluxApp/internal/model"
	"github.com/carefuly/PackageFluxApp/pkg/ginx/query/filters"
	"gorm.io/gorm"
)

var (
	ErrFileRecordNotFound = gorm.ErrRecordNotFound
	ErrFileNotFound       = errors.New("记录不存在")
)

type FileDAO interface {
	Insert(ctx context.Context, f model.File) error
	Delete(ctx context.Context, id, userId string) (int64, error)
	FindByIdAndUserId(ctx context.Context, id, userId string) (*model.File, error)
	FindListPage(ctx context.Context, userId string, page filters.Pagination) (int64, []*model.File, error)
	FindListAll(ctx context.Context, userId, appName string) (int64, []*model.File, error)
	ExistsByUserIdAndName(ctx context.Context, userId, name string) (string, error)
}

type GORMFileDAO struct {
	db *gorm.DB
}

func NewGORMFileDAO(db *gorm.DB) FileDAO {
	return &GORMFileDAO{
		db: db,
	}
}

func (dao *GORMFileDAO) Insert(ctx context.Context, f model.File) error {
	return dao.db.WithContext(ctx).Create(&f).Error
}

func (dao *GORMFileDAO) Delete(ctx context.Context, id, userId string) (int64, error) {
	result := dao.db.WithContext(ctx).Where("id = ? AND user_id = ?", id, userId).Delete(&model.File{})
	return result.RowsAffected, result.Error
}

func (dao *GORMFileDAO) FindByIdAndUserId(ctx context.Context, id, userId string) (*model.File, error) {
	var file model.File
	err := dao.db.WithContext(ctx).Where("id = ? AND user_id = ?", id, userId).First(&file).Error
	return &file, err
}

func (dao *GORMFileDAO) FindListPage(ctx context.Context, userId string, page filters.Pagination) (int64, []*model.File, error) {
	var total int64
	var files []*model.File

	if err := dao.db.WithContext(ctx).Model(&model.File{}).
		Where("status = ? AND user_id = ?", true, userId).
		Count(&total).Error; err != nil {
		return 0, nil, err
	}

	offset := (page.Page - 1) * page.PageSize
	if err := dao.db.WithContext(ctx).Where("status = ? AND user_id = ?", true, userId).
		Offset(offset).Limit(page.PageSize).Order("update_time DESC, sort ASC").
		Find(&files).Error; err != nil {
		return 0, nil, err
	}

	return total, files, nil
}

func (dao *GORMFileDAO) FindListAll(ctx context.Context, userId, name string) (int64, []*model.File, error) {
	query := dao.db.WithContext(ctx).Where("status = ? AND user_id = ?", true, userId).
		Order("update_time DESC, sort ASC")
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	var files []*model.File
	result := query.Find(&files)
	return result.RowsAffected, files, result.Error
}

func (dao *GORMFileDAO) ExistsByUserIdAndName(ctx context.Context, userId, name string) (string, error) {
	var file model.File
	err := dao.db.WithContext(ctx).Where("user_id = ? AND name = ?", userId, name).First(&file).Error

	if errors.Is(err, ErrFileRecordNotFound) {
		return "", ErrFileRecordNotFound
	}
	if err != nil {
		return "", err
	}
	return file.Id, err
}
