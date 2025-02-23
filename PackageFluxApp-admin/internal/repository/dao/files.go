/**
 * Description：
 * FileName：files.go
 * Author：CJiaの用心
 * Create：2025/2/23 19:50:22
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
	ErrCreateFile  = errors.New("上传文件失败")
	ErrTransaction = errors.New("上传文件失败")
)

type FilesDAO interface {
	BatchCreate(ctx context.Context, files []model.File) error
	BatchDelete(ctx context.Context, uid uint, ids []string) ([]string, error)
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

		if existingFile.ID != 0 {
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

func (dao *filesDAO) BatchDelete(ctx context.Context, uid uint, ids []string) ([]string, error) {
	var pathsToDelete []string
	// 批量处理
	for _, id := range ids {
		// 检查文件是否存在
		var existingFile model.File
		if err := dao.db.WithContext(ctx).Where("record_id = ? AND user_id = ?", id, uid).First(&existingFile).Error; err != nil {
		}

		pathsToDelete = append(pathsToDelete, existingFile.Path)

		if existingFile.ID != 0 {
			if err := dao.db.Delete(&existingFile).Error; err != nil {
			}
		}
	}
	return pathsToDelete, nil
}
