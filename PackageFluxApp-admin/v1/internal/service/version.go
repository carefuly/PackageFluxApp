/**
 * Description：
 * FileName：version.go
 * Author：CJiaの用心
 * Create：2025/2/25 15:50:18
 * Remark：
 */

package service

import (
	"context"
	"errors"
	"github.com/carefuly/PackageFluxApp/internal/domain"
	"github.com/carefuly/PackageFluxApp/internal/repository"
	"strings"
)

var (
	ErrDuplicateVersionCode = repository.ErrDuplicateVersionCode
	ErrVersionNotFound      = repository.ErrVersionNotFound
	ErrFormalNotFound       = repository.ErrFormalNotFound
	ErrVersionLowerInvalid  = errors.New("正式版版本号不能低于当前正式版")
)

type VersionService interface {
	Create(ctx context.Context, version domain.Version) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string, version domain.Version) error
	Formal(ctx context.Context, version domain.FormalVersion) error
	FindById(ctx context.Context, id string) (domain.Version, error)
	FindListAll(ctx context.Context, f domain.FiltersVersion) ([]domain.Version, error)
}

type versionService struct {
	repo       repository.VersionRepository
	detailRepo repository.DetailsRepository
}

func NewVersionService(repo repository.VersionRepository, detailRepo repository.DetailsRepository) VersionService {
	return &versionService{
		repo: repo,
		detailRepo: detailRepo,
	}
}

func (svc *versionService) Create(ctx context.Context, version domain.Version) error {
	return svc.repo.Create(ctx, version)
}

func (svc *versionService) Delete(ctx context.Context, id string) error {
	return svc.repo.Delete(ctx, id)
}

func (svc *versionService) Update(ctx context.Context, id string, version domain.Version) error {
	return svc.repo.Update(ctx, id, version)
}

func (svc *versionService) Formal(ctx context.Context, version domain.FormalVersion) error {
	// 获取最新的正式版
	latestFormalVersion, err := svc.repo.FindLatestFormalVersion(ctx, version.DetailId)
	if err != nil {
		return err
	}

	formalVersion, err := svc.repo.FindById(ctx, version.VersionId)
	if errors.Is(err, repository.ErrRecordNotFound) {
		return ErrVersionNotFound
	} else if err != nil {
		return err
	}

	// 如果存在最新的正式版，检查版本号是否合法
	if latestFormalVersion.ID != "" {
		compareResult, err := svc.CompareVersion(formalVersion.VersionCode, latestFormalVersion.VersionCode)
		if err != nil {
			return err
		}
		if compareResult <= 0 {
			return ErrVersionLowerInvalid
		}
	}

	err = svc.detailRepo.Formal(ctx, version.DetailId, version.VersionId)
	if err != nil {
		return err
	}

	return svc.repo.Formal(ctx, version.VersionId)
}

func (svc *versionService) FindById(ctx context.Context, id string) (domain.Version, error) {
	return svc.repo.FindById(ctx, id)
}

func (svc *versionService) FindListAll(ctx context.Context, f domain.FiltersVersion) ([]domain.Version, error) {
	return svc.repo.FindListAll(ctx, f)
}

// CompareVersion 比较两个版本号
// 返回值：
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
