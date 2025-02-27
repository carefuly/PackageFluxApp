/**
 * Description：
 * FileName：file.go
 * Author：CJiaの用心
 * Create：2025/2/24 16:50:49
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
	BatchDelete(ctx context.Context, uid string, ids []string) ([]string, error)
	FindListAllWithPage(ctx context.Context, f domain.FiltersFile, p domain.Pagination) ([]domain.Files, int64, error)
	FindListAll(ctx context.Context, f domain.FiltersFile) ([]domain.Files, error)
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

func (repo *filesRepository) BatchDelete(ctx context.Context, uid string, ids []string) ([]string, error) {
	return repo.dao.BatchDelete(ctx, uid, ids)
}

func (repo *filesRepository) FindListAllWithPage(ctx context.Context, f domain.FiltersFile, p domain.Pagination) ([]domain.Files, int64, error) {
	list, total, err := repo.dao.FindListAllWithPage(ctx, f, p)
	var files []domain.Files
	for _, v := range list {
		files = append(files, repo.toDomain(v))
	}
	return files, total, err
}

func (repo *filesRepository) FindListAll(ctx context.Context, f domain.FiltersFile) ([]domain.Files, error) {
	var files []domain.Files
	list, err := repo.dao.FindListAll(ctx, f)
	for _, v := range list {
		files = append(files, repo.toDomain(v))
	}
	return files, err
}

func (repo *filesRepository) toDomain(f model.File) domain.Files {
	return domain.Files{
		File:       f,
		CreateTime: f.CoreModels.CreateTime.Format("2006-01-02 15:04:05.000"),
		UpdateTime: f.CoreModels.UpdateTime.Format("2006-01-02 15:04:05.000"),
	}
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
