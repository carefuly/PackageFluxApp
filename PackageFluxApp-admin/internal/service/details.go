/**
 * Description：
 * FileName：details.go
 * Author：CJiaの用心
 * Create：2025/2/24 16:16:07
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
	Delete(ctx context.Context, id, userId string) error
	Update(ctx context.Context, id string, detail domain.Detail) error
	FindById(ctx context.Context, id, userId string) (domain.Detail, error)
	FindListAll(ctx context.Context, f domain.FiltersDetail) ([]domain.Detail, error)
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

func (svc *detailsService) Delete(ctx context.Context, id, userId string) error {
	return svc.repo.Delete(ctx, id, userId)
}

func (svc *detailsService) Update(ctx context.Context, id string, detail domain.Detail) error {
	return svc.repo.Update(ctx, id, detail)
}

func (svc *detailsService) FindById(ctx context.Context, id, userId string) (domain.Detail, error) {
	return svc.repo.FindById(ctx, id, userId)
}

func (svc *detailsService) FindListAll(ctx context.Context, f domain.FiltersDetail) ([]domain.Detail, error) {
	return svc.repo.FindListAll(ctx, f)
}
