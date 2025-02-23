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
		if err := dao.db.Where("user_id = ? AND name = ?", file.UserID, file.Name).First(&existingFile).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 插入新数据
				if err := dao.db.Create(&file).Error; err != nil {
					return ErrCreateFile
				}
			} else if existingFile.ID != 0 {
				if err := dao.db.Delete(&existingFile).Error; err != nil {
					return ErrCreateFile
				}
			} else {
				return err
			}
		}
	}
	return nil
}
