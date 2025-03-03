/**
 * Description：
 * FileName：check.go
 * Author：CJiaの用心
 * Create：2025/3/1 00:16:36
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

type CheckService interface {
	CheckLatest(ctx context.Context, detailId, versionCode string) (domain.Check, error)
}

type checkService struct {
	userRepo    repository.UserRepository
	detailRepo  repository.DetailRepository
	versionRepo repository.VersionRepository
}

func NewCheckService(userRepo repository.UserRepository, detailRepo repository.DetailRepository, versionRepo repository.VersionRepository) CheckService {
	return &checkService{
		userRepo:    userRepo,
		detailRepo:  detailRepo,
		versionRepo: versionRepo,
	}
}

func (svc *checkService) CheckLatest(ctx context.Context, detailId, versionCode string) (domain.Check, error) {
	version, err := svc.versionRepo.FindByDetailIdAndFormalVersion(ctx, detailId)
	if err != nil {
		if errors.Is(err, repository.ErrVersionRecordNotFound) {
			return domain.Check{
				Tip:  "当前暂未发布正式版本",
				Type: -1,
			}, nil
		}
		return domain.Check{}, err
	}
	if version.Id == "" {
		return domain.Check{
			Tip:  "当前暂未发布正式版本",
			Type: -1,
		}, nil
	}

	compareVersion, err := svc.CompareVersion(versionCode, version.VersionCode)
	if err != nil {
		return domain.Check{}, err
	}

	err = svc.userRepo.UpdateUsageNumber(ctx, version.UserId)
	if err != nil {
		return domain.Check{}, err
	}

	switch compareVersion {
	case 0:
		return domain.Check{
			NeedUpdate:     false,
			Version:        version.VersionCode,
			ApkUrl:         version.ApkUrl,
			WgtUrl:         version.WgtUrl,
			HBuilderUpdate: version.HBuilderUpdate,
			ForceUpdate:    version.ForceUpdate,
			Description:    version.Description,
			Tip:            "当前已是最新版本",
			Type:           1,
		}, nil
	case -1:
		return domain.Check{
			NeedUpdate:     true,
			Version:        version.VersionCode,
			ApkUrl:         version.ApkUrl,
			WgtUrl:         version.WgtUrl,
			HBuilderUpdate: version.HBuilderUpdate,
			ForceUpdate:    version.ForceUpdate,
			Description:    version.Description,
			Tip:            "当前低于正式版",
			Type:           0,
		}, nil
	case 1:
		return domain.Check{
			Tip:  "当前大于正式版",
			Type: 2,
		}, nil
	}
	return domain.Check{}, nil
}

// CompareVersion 比较两个版本号
// 返回值：
//
//	-1: v1 < v2
//	 0: v1 == v2
//	 1: v1 > v2
func (svc *checkService) CompareVersion(v1, v2 string) (int, error) {
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
func (svc *checkService) parseInt(s string) int {
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
