/**
 * Description：
 * FileName：version.go
 * Author：CJiaの用心
 * Create：2025/2/28 14:43:45
 * Remark：
 */

package service

import (
	"context"
	"errors"
	"github.com/carefuly/PackageFluxApp/internal/domain"
	"github.com/carefuly/PackageFluxApp/internal/repository"
	"github.com/go-sql-driver/mysql"
)

var (
	ErrVersionRecordNotFound           = repository.ErrVersionRecordNotFound
	ErrVersionNotFound                 = repository.ErrVersionNotFound
	ErrDuplicateDetailIdAndVersionCode = repository.ErrDuplicateDetailIdAndVersionCode
)

type VersionService interface {
	Create(ctx context.Context, d domain.Version) error
	Delete(ctx context.Context, id, detailId, userId string) error
	Update(ctx context.Context, id, detailId, userId string, v domain.Version) error
	UpdateFormal(ctx context.Context, id, detailId, userId string) error
	FindById(ctx context.Context, id, detailId, userId string) (domain.Version, error)
	FindListAll(ctx context.Context, detailId, userId string) ([]domain.Version, error)
}

type versionService struct {
	repo       repository.VersionRepository
	detailRepo repository.DetailRepository
}

func NewVersionService(repo repository.VersionRepository, detailRepo repository.DetailRepository) VersionService {
	return &versionService{
		repo:       repo,
		detailRepo: detailRepo,
	}
}

func (svc *versionService) Create(ctx context.Context, d domain.Version) error {
	exists, err := svc.repo.ExistsByDetailIdAndCode(ctx, d.UserId, d.DetailId, d.VersionCode)
	if err != nil {
		return err
	}
	if exists {
		return repository.ErrDuplicateDetailIdAndVersionCode
	}

	err = svc.repo.Create(ctx, d)
	if svc.IsDuplicateEntryError(err) {
		return repository.ErrDuplicateDetailIdAndVersionCode
	}

	return err
}

func (svc *versionService) Delete(ctx context.Context, id, detailId, userId string) error {
	rowsAffected, err := svc.repo.Delete(ctx, id, detailId, userId)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return repository.ErrVersionNotFound
	}
	return err
}

func (svc *versionService) Update(ctx context.Context, id, detailId, userId string, v domain.Version) error {
	rowsAffected, err := svc.repo.Update(ctx, id, detailId, userId, v)
	if svc.IsDuplicateEntryError(err) {
		return repository.ErrDuplicateDetailIdAndVersionCode
	}
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return repository.ErrVersionNotFound
	}
	return err
}

func (svc *versionService) UpdateFormal(ctx context.Context, id, detailId, userId string) error {
	rowsAffected, err := svc.repo.UpdateFormal(ctx, id, detailId, userId)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return repository.ErrVersionNotFound
	}
	_, err = svc.detailRepo.UpdateFormal(ctx, detailId, userId, id)
	if err != nil {
		return err
	}
	return err
}

func (svc *versionService) FindById(ctx context.Context, id, detailId, userId string) (domain.Version, error) {
	d, err := svc.repo.FindByIdAndUserId(ctx, id, detailId, userId)
	if err != nil {
		if errors.Is(err, ErrVersionRecordNotFound) {
			return domain.Version{}, repository.ErrVersionRecordNotFound
		}
		return domain.Version{}, err
	}
	return d, nil
}

func (svc *versionService) FindListAll(ctx context.Context, detailId, userId string) ([]domain.Version, error) {
	rowsAffected, list, err := svc.repo.FindListAll(ctx, detailId, userId)
	if err != nil {
		return []domain.Version{}, err
	}
	if rowsAffected == 0 {
		return []domain.Version{}, err
	}
	return list, nil
}

func (svc *versionService) IsDuplicateEntryError(err error) bool {
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		// MySQL 错误码 1062 表示唯一冲突
		return mysqlErr.Number == 1062
	}
	return false
}
