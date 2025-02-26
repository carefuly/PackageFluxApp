/**
 * Description：
 * FileName：files.go
 * Author：CJiaの用心
 * Create：2025/2/24 16:52:29
 * Remark：
 */

package service

import (
	"context"
	"github.com/carefuly/PackageFluxApp/internal/domain"
	"github.com/carefuly/PackageFluxApp/internal/repository"
)

var (
	ErrCreateFile = repository.ErrCreateFile
)

type FilesService interface {
	BatchCreate(ctx context.Context, files []domain.Files) error
	BatchDelete(ctx context.Context, uid string, ids []string) ([]string, error)
	FindListAllWithPage(ctx context.Context, f domain.FiltersFile, p domain.Pagination) ([]domain.Files, int64, error)
	FindListAll(ctx context.Context, f domain.FiltersFile) ([]domain.Files, error)
}

type filesService struct {
	repo repository.FilesRepository
}

func NewFilesService(repo repository.FilesRepository) FilesService {
	return &filesService{
		repo: repo,
	}
}

func (svc *filesService) BatchCreate(ctx context.Context, files []domain.Files) error {
	return svc.repo.BatchCreate(ctx, files)
}

func (svc *filesService) BatchDelete(ctx context.Context, uid string, ids []string) ([]string, error) {
	return svc.repo.BatchDelete(ctx, uid, ids)
}

func (svc *filesService) FindListAllWithPage(ctx context.Context, f domain.FiltersFile, p domain.Pagination) ([]domain.Files, int64, error) {
	return svc.repo.FindListAllWithPage(ctx, f, p)
}

func (svc *filesService) FindListAll(ctx context.Context, f domain.FiltersFile) ([]domain.Files, error) {
	return svc.repo.FindListAll(ctx, f)
}
