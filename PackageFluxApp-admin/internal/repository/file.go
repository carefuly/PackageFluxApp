/**
 * Description：
 * FileName：file.go
 * Author：CJiaの用心
 * Create：2025/2/27 21:53:19
 * Remark：
 */

package repository

import (
	"context"
	"errors"
	"github.com/carefuly/PackageFluxApp/internal/domain"
	"github.com/carefuly/PackageFluxApp/internal/model"
	"github.com/carefuly/PackageFluxApp/internal/repository/cache"
	"github.com/carefuly/PackageFluxApp/internal/repository/dao"
	"github.com/carefuly/PackageFluxApp/pkg/ginx/query/filters"
	"github.com/carefuly/PackageFluxApp/pkg/models"
	"go.uber.org/zap"
)

var (
	ErrFileRecordNotFound = dao.ErrFileRecordNotFound
	ErrFileNotFound       = dao.ErrFileNotFound
)

type FileRepository interface {
	Create(ctx context.Context, f domain.File) error
	Delete(ctx context.Context, id, userId string) (int64, error)
	FindByIdAndUserId(ctx context.Context, id, userId string) (domain.File, error)
	FindListPage(ctx context.Context, userId string, page filters.Pagination) (int64, []domain.File, error)
	FindListAll(ctx context.Context, userId, name string) (int64, []domain.File, error)
	ExistsByUserIdAndName(ctx context.Context, userId, name string) (string, error)
}

type fileRepository struct {
	dao   dao.FileDAO
	cache cache.FileCache
}

func NewFileRepository(dao dao.FileDAO, cache cache.FileCache) FileRepository {
	return &fileRepository{
		dao:   dao,
		cache: cache,
	}
}

func (repo *fileRepository) Create(ctx context.Context, f domain.File) error {
	return repo.dao.Insert(ctx, repo.toEntity(f))
}

func (repo *fileRepository) Delete(ctx context.Context, id, userId string) (int64, error) {
	return repo.dao.Delete(ctx, id, userId)
}

func (repo *fileRepository) FindByIdAndUserId(ctx context.Context, id, userId string) (domain.File, error) {
	f, err := repo.cache.Get(ctx, id)
	switch {
	case err == nil:
		return *f, err
	case errors.Is(err, cache.ErrFileNotExist):
		file, err := repo.dao.FindByIdAndUserId(ctx, id, userId)
		if err != nil {
			return domain.File{}, err
		}

		fileInfo := repo.toDomain(file)

		err = repo.cache.Set(ctx, fileInfo)
		if err != nil {
			zap.L().Error("Redis异常", zap.Error(err))
			return domain.File{}, err
		}
		return fileInfo, err
	default:
		return domain.File{}, err
	}
}

func (repo *fileRepository) FindListPage(ctx context.Context, userId string, page filters.Pagination) (int64, []domain.File, error) {
	rows, list, err := repo.dao.FindListPage(ctx, userId, page)
	if err != nil {
		return rows, []domain.File{}, err
	}
	var files []domain.File
	for _, v := range list {
		files = append(files, repo.toDomain(v))
	}
	return rows, files, err
}

func (repo *fileRepository) FindListAll(ctx context.Context, userId, appName string) (int64, []domain.File, error) {
	rows, list, err := repo.dao.FindListAll(ctx, userId, appName)
	if err != nil {
		return rows, []domain.File{}, err
	}
	var files []domain.File
	for _, v := range list {
		files = append(files, repo.toDomain(v))
	}
	return rows, files, err
}

func (repo *fileRepository) ExistsByUserIdAndName(ctx context.Context, userId, name string) (string, error) {
	return repo.dao.ExistsByUserIdAndName(ctx, userId, name)
}

func (repo *fileRepository) toEntity(d domain.File) model.File {
	return model.File{
		CoreModels: models.CoreModels{
			Creator:  d.UserId,
			Modifier: d.UserId,
			Remark:   d.Remark,
		},
		Name:    d.Name,
		Size:    d.Size,
		Suffix:  d.Suffix,
		BaseDir: d.BaseDir,
		FileDir: d.FileDir,
		Path:    d.Path,
		Url:     d.Url,
		UserId:  d.UserId,
	}
}

func (repo *fileRepository) toDomain(f *model.File) domain.File {
	return domain.File{
		File:       *f,
		CreateTime: f.CoreModels.CreateTime.Format("2006-01-02 15:04:05.000"),
		UpdateTime: f.CoreModels.UpdateTime.Format("2006-01-02 15:04:05.000"),
	}
}
