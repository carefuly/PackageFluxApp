/**
 * Description：
 * FileName：file.go
 * Author：CJiaの用心
 * Create：2025/2/27 22:03:58
 * Remark：
 */

package service

import (
	"context"
	"errors"
	"github.com/carefuly/PackageFluxApp/internal/domain"
	"github.com/carefuly/PackageFluxApp/internal/repository"
	"github.com/carefuly/PackageFluxApp/pkg/ginx/query/filters"
)

var (
	ErrFileRecordNotFound = repository.ErrFileRecordNotFound
)

type FileService interface {
	BatchCreate(ctx context.Context, files []domain.File) []string
	BatchDelete(ctx context.Context, userId string, ids []string) ([]string, []error)
	FindListPage(ctx context.Context, userId string, page filters.Pagination) (int64, []domain.File, error)
	FindListAll(ctx context.Context, userId, name string) ([]domain.File, error)
}

type fileService struct {
	repo repository.FileRepository
}

func NewFileService(repo repository.FileRepository) FileService {
	return &fileService{repo: repo}
}

func (svc *fileService) Create(ctx context.Context, f domain.File) error {
	id, err := svc.repo.ExistsByUserIdAndName(ctx, f.UserId, f.Name)
	switch {
	case err == nil || errors.Is(err, repository.ErrFileRecordNotFound):
		if id != "" {
			_ = svc.Delete(ctx, id, f.UserId)
		}
		err = svc.repo.Create(ctx, f)
		return err
	default:
		return err
	}
}

func (svc *fileService) BatchCreate(ctx context.Context, files []domain.File) []string {
	var str []string
	for _, file := range files {
		err := svc.Create(ctx, file)
		if err == nil {
			str = append(str, file.Name)
		}
	}
	return str
}

func (svc *fileService) Delete(ctx context.Context, id, userId string) error {
	rowsAffected, err := svc.repo.Delete(ctx, id, userId)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return repository.ErrFileNotFound
	}
	return err
}

func (svc *fileService) BatchDelete(ctx context.Context, userId string, ids []string) ([]string, []error) {
	var str []string
	var errs []error
	for _, id := range ids {
		file, err := svc.repo.FindByIdAndUserId(ctx, id, userId)
		switch {
		case err == nil:
			err = svc.Delete(ctx, id, userId)
			switch {
			case err == nil:
				str = append(str, file.Path)
			default:
				errs = append(errs, err)
			}
		default:
			errs = append(errs, err)
		}
	}
	return str, errs
}

func (svc *fileService) FindById(ctx context.Context, id, userId string) (domain.File, error) {
	f, err := svc.repo.FindByIdAndUserId(ctx, id, userId)
	if err != nil {
		if errors.Is(err, repository.ErrFileRecordNotFound) {
			return domain.File{}, repository.ErrFileRecordNotFound
		}
		return domain.File{}, err
	}
	return f, nil
}

func (svc *fileService) FindListPage(ctx context.Context, userId string, page filters.Pagination) (int64, []domain.File, error) {
	total, list, err := svc.repo.FindListPage(ctx, userId, page)
	if err != nil {
		return 0, []domain.File{}, err
	}
	if total == 0 {
		return 0, []domain.File{}, err
	}
	return total, list, nil
}

func (svc *fileService) FindListAll(ctx context.Context, userId, name string) ([]domain.File, error) {
	rowsAffected, list, err := svc.repo.FindListAll(ctx, userId, name)
	if err != nil {
		return []domain.File{}, err
	}
	if rowsAffected == 0 {
		return []domain.File{}, err
	}
	return list, nil
}
