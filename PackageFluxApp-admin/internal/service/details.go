/**
 * Description：
 * FileName：details.go
 * Author：CJiaの用心
 * Create：2025/2/20 20:47:33
 * Remark：
 */

package service

import (
	"context"
	"github.com/carefuly/PackageFluxApp/internal/domain"
	"github.com/carefuly/PackageFluxApp/internal/repository"
)

var (
	ErrDuplicateAppName = repository.ErrDuplicateAppName
	ErrRecordNotFound   = repository.ErrRecordNotFound
)

type DetailsService interface {
	Create(ctx context.Context, detail domain.Detail) error
	Delete(ctx context.Context, userId uint, recordId string) error
	Update(ctx context.Context, detail domain.Detail) error
	FindById(ctx context.Context, userId uint, recordId string) (domain.Detail, error)
	ListAll(ctx context.Context, f domain.FiltersDetail) ([]domain.Detail, error)
}

type detailsService struct {
	repo repository.DetailsRepository
}

func NewDetailsService(repo repository.DetailsRepository) DetailsService {
	return &detailsService{
		repo: repo,
	}
}

func (svc *detailsService) Create(ctx context.Context, detail domain.Detail) error {
	return svc.repo.Create(ctx, detail)
}

func (svc *detailsService) Delete(ctx context.Context, userId uint, recordId string) error {
	return svc.repo.Delete(ctx, userId, recordId)
}

func (svc *detailsService) Update(ctx context.Context, detail domain.Detail) error {
	return svc.repo.Update(ctx, detail)
}

func (svc *detailsService) FindById(ctx context.Context, userId uint, recordId string) (domain.Detail, error) {
	return svc.repo.FindById(ctx, userId, recordId)
}

func (svc *detailsService) ListAll(ctx context.Context, f domain.FiltersDetail) ([]domain.Detail, error) {
	return svc.repo.ListAll(ctx, f)
}
