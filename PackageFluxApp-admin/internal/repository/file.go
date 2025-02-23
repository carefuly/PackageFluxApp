/**
 * Description：
 * FileName：file.go
 * Author：CJiaの用心
 * Create：2025/2/23 20:13:50
 * Remark：
 */

package repository

import (
	"context"
	"github.com/carefuly/PackageFluxApp/internal/domain"
	"github.com/carefuly/PackageFluxApp/internal/repository/dao"
	"github.com/carefuly/PackageFluxApp/model"
	"github.com/carefuly/PackageFluxApp/pkg/models"
)

var (
	ErrCreateFile = dao.ErrCreateFile
)

type FilesRepository interface {
	BatchCreate(ctx context.Context, files []domain.Files) error
}

type filesRepository struct {
	dao dao.FilesDAO
}

func NewFilesRepository(dao dao.FilesDAO) FilesRepository {
	return &filesRepository{
		dao: dao,
	}
}

func (repo *filesRepository) BatchCreate(ctx context.Context, files []domain.Files) error {
	var filesEntity []model.File
	for _, file := range files {
		filesEntity = append(filesEntity, repo.toEntity(file))
	}
	return repo.dao.BatchCreate(ctx, filesEntity)
}

func (repo *filesRepository) toEntity(d domain.Files) model.File {
	return model.File{
		CoreModels: models.CoreModels{
			Creator:  d.UserID,
			Modifier: d.UserID,
		},
		Name:    d.Name,
		Size:    d.Size,
		Suffix:  d.Suffix,
		BaseDir: d.BaseDir,
		FileDir: d.FileDir,
		Path:    d.Path,
		Url:     d.Url,
		UserID:  d.UserID,
	}
}
