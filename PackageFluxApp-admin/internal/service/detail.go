/**
 * Description：
 * FileName：detail.go
 * Author：CJiaの用心
 * Create：2025/2/27 11:21:42
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
	ErrDetailRecordNotFound      = repository.ErrDetailRecordNotFound
	ErrDetailNotFound            = repository.ErrDetailNotFound
	ErrDuplicateUserIdAndAppName = repository.ErrDuplicateUserIdAndAppName
)

type DetailService interface {
	Create(ctx context.Context, d domain.Detail) error
	Delete(ctx context.Context, id, userId string) error
	Update(ctx context.Context, id, userId string, d domain.Detail) error
	FindById(ctx context.Context, id, userId string) (domain.Detail, error)
	FindListAll(ctx context.Context, userId, appName string) ([]domain.Detail, error)
}

type detailService struct {
	repo repository.DetailRepository
}

func NewDetailService(repo repository.DetailRepository) DetailService {
	return &detailService{
		repo: repo,
	}
}

func (svc *detailService) Create(ctx context.Context, d domain.Detail) error {
	exists, err := svc.repo.ExistsByUserIdAndAppName(ctx, d.UserId, d.AppName)
	if err != nil {
		return err
	}
	if exists {
		return repository.ErrDuplicateUserIdAndAppName
	}

	err = svc.repo.Create(ctx, d)
	if svc.IsDuplicateEntryError(err) {
		return repository.ErrDuplicateUserIdAndAppName
	}

	return err
}

func (svc *detailService) Delete(ctx context.Context, id, userId string) error {
	rowsAffected, err := svc.repo.Delete(ctx, id, userId)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return repository.ErrDetailNotFound
	}
	return err
}

func (svc *detailService) Update(ctx context.Context, id, userId string, d domain.Detail) error {
	rowsAffected, err := svc.repo.Update(ctx, id, userId, d)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return repository.ErrDetailNotFound
	}
	return err
}

func (svc *detailService) FindById(ctx context.Context, id, userId string) (domain.Detail, error) {
	d, err := svc.repo.FindByIdAndUserId(ctx, id, userId)
	if err != nil {
		if errors.Is(err, ErrDetailRecordNotFound) {
			return domain.Detail{}, repository.ErrDetailRecordNotFound
		}
		return domain.Detail{}, err
	}
	return d, nil
}

func (svc *detailService) FindListAll(ctx context.Context, userId, appName string) ([]domain.Detail, error) {
	rowsAffected, list, err := svc.repo.FindListAll(ctx, userId, appName)
	if err != nil {
		return []domain.Detail{}, err
	}
	if rowsAffected == 0 {
		return []domain.Detail{}, err
	}
	return list, nil
}

func (svc *detailService) IsDuplicateEntryError(err error) bool {
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		// MySQL 错误码 1062 表示唯一冲突
		return mysqlErr.Number == 1062
	}
	return false
}
