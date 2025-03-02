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
	"strings"
)

var (
	ErrVersionRecordNotFound           = repository.ErrVersionRecordNotFound
	ErrVersionNotFound                 = repository.ErrVersionNotFound
	ErrDuplicateDetailIdAndVersionCode = repository.ErrDuplicateDetailIdAndVersionCode
	ErrVersionLowerInvalid             = errors.New("正式版版本号不能低于当前正式版")
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
	version, err := svc.repo.FindByFormalVersionCode(ctx, detailId, userId)
	v, err := svc.FindById(ctx, id, detailId, userId)
	if err != nil {
		return err
	}
	if version.Id != "" {
		compareResult, err := svc.CompareVersion(v.VersionCode, version.VersionCode)
		if err != nil {
			return err
		}
		if compareResult <= 0 {
			return ErrVersionLowerInvalid
		}
	}
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
	d, err := svc.repo.FindByIdAndDetailIdAndUserId(ctx, id, detailId, userId)
	if err != nil {
		if errors.Is(err, repository.ErrVersionRecordNotFound) {
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

// CompareVersion 比较两个版本号
// 返回值：
//
//	-1: v1 < v2
//	 0: v1 == v2
//	 1: v1 > v2
func (svc *versionService) CompareVersion(v1, v2 string) (int, error) {
	v1Parts := strings.Split(v1, ".")
	v2Parts := strings.Split(v2, ".")

	for i := 0; i < len(v1Parts) || i < len(v2Parts); i++ {
		var v1Num, v2Num int
		if i < len(v1Parts) {
			v1Num = svc.parseInt(v1Parts[i])
		}
		if i < len(v2Parts) {
			v2Num = svc.parseInt(v2Parts[i])
		}

		if v1Num < v2Num {
			return -1, nil
		} else if v1Num > v2Num {
			return 1, nil
		}
	}

	return 0, nil
}

// parseInt 将字符串转换为整数
func (svc *versionService) parseInt(s string) int {
	var num int
	for _, c := range s {
		if c >= '0' && c <= '9' {
			num = num*10 + int(c-'0')
		} else {
			break
		}
	}
	return num
}
