/**
 * Description：
 * FileName：files.go
 * Author：CJiaの用心
 * Create：2025/2/22 21:45:49
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
	BatchDelete(ctx context.Context, uid uint, ids []string) ([]string, error)
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

func (svc *filesService) BatchDelete(ctx context.Context, uid uint, ids []string) ([]string, error) {
	return svc.repo.BatchDelete(ctx, uid, ids)
}
