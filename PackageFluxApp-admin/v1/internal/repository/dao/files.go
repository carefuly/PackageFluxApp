/**
 * Description：
 * FileName：file.go
 * Author：CJiaの用心
 * Create：2025/2/24 16:48:50
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
	ErrCreateFile  = errors.New("上传文件失败")
	ErrTransaction = errors.New("上传文件失败")
)

type FilesDAO interface {
	BatchCreate(ctx context.Context, files []model.File) error
	BatchDelete(ctx context.Context, uid string, ids []string) ([]string, error)
	FindListAllWithPage(ctx context.Context, f domain.FiltersFile, p domain.Pagination) ([]model.File, int64, error)
	FindListAll(ctx context.Context, f domain.FiltersFile) ([]model.File, error)
}

type filesDAO struct {
	db *gorm.DB
}

func NewFilesDAO(db *gorm.DB) FilesDAO {
	return &filesDAO{
		db: db,

	}
}

func (dao *filesDAO) BatchCreate(ctx context.Context, files []model.File) error {
	// 批量处理
	for _, file := range files {
		// 检查文件是否存在
		var existingFile model.File
		if err := dao.db.WithContext(ctx).Where("user_id = ? AND name = ?", file.UserID, file.Name).First(&existingFile).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {

			}
		}

		if existingFile.Id != "" {
			if err := dao.db.Delete(&existingFile).Error; err != nil {
				return ErrCreateFile
			}
		}

		// 插入新数据
		if err := dao.db.Create(&file).Error; err != nil {
			return ErrCreateFile
		}

	}
	return nil
}

func (dao *filesDAO) BatchDelete(ctx context.Context, uid string, ids []string) ([]string, error) {
	var pathsToDelete []string
	// 批量处理
	for _, id := range ids {
		// 检查文件是否存在
		var existingFile model.File
		if err := dao.db.WithContext(ctx).Where("id = ? AND user_id = ?", id, uid).First(&existingFile).Error; err != nil {
		}

		pathsToDelete = append(pathsToDelete, existingFile.Path)

		if existingFile.Id != "" {
			if err := dao.db.Delete(&existingFile).Error; err != nil {
			}
		}
	}
	return pathsToDelete, nil
}

func (dao *filesDAO) FindListAllWithPage(ctx context.Context, f domain.FiltersFile, p domain.Pagination) ([]model.File, int64, error) {
	var files []model.File
	var total int64

	if err := dao.db.WithContext(ctx).Model(&model.File{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (p.Page - 1) * p.PageSize
	if err := dao.db.WithContext(ctx).Where("status = ? AND user_id = ?", f.Status, f.UserId).
		Offset(offset).Limit(p.PageSize).Find(&files).Error; err != nil {
		return nil, 0, err
	}

	return files, total, nil
}

func (dao *filesDAO) FindListAll(ctx context.Context, f domain.FiltersFile) ([]model.File, error) {
	var files []model.File
	err := dao.db.WithContext(ctx).Where("status = ? AND user_id = ?", f.Status, f.UserId).Find(&files).Error
	return files, err
}
