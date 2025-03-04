/**
 * Description：
 * FileName：version.go
 * Author：CJiaの用心
 * Create：2025/2/28 14:31:12
 * Remark：
 */

package repository

import (
	"context"
	"errors"
	"github.com/carefuly/PackageFluxApp/internal/domain"
	"github.com/carefuly/PackageFluxApp/internal/model"
	"github.com/carefuly/PackageFluxApp/internal/repository/cache"
	"github.com/carefuly/PackageFluxApp/internal/repository/dao"
	"github.com/carefuly/PackageFluxApp/pkg/models"
	"go.uber.org/zap"
)

var (
	ErrVersionRecordNotFound           = dao.ErrVersionRecordNotFound
	ErrVersionNotFound                 = dao.ErrVersionNotFound
	ErrDuplicateDetailIdAndVersionCode = dao.ErrDuplicateDetailIdAndVersionCode
)

type VersionRepository interface {
	Create(ctx context.Context, v domain.Version) error
	Delete(ctx context.Context, id, detailId, userId string) (int64, error)
	Update(ctx context.Context, id, detailId, userId string, v domain.Version) (int64, error)
	UpdateFormal(ctx context.Context, id, detailId, userId string) (int64, error)
	FindByIdAndDetailIdAndUserId(ctx context.Context, id, detailId, userId string) (domain.Version, error)
	FindByFormalVersionCode(ctx context.Context, detailId, userId string) (domain.Version, error)
	FindByDetailIdAndFormalVersion(ctx context.Context, detailId string) (domain.Version, error)
	FindListAll(ctx context.Context, detailId, userId string) (int64, []domain.Version, error)
	ExistsByDetailIdAndFormalVersion(ctx context.Context, detailId string) (bool, error)
	ExistsByDetailIdAndCode(ctx context.Context, userId, detailId, versionCode string) (bool, error)
}

type versionRepository struct {
	dao   dao.VersionDAO
	cache cache.VersionCache
}

func NewVersionRepository(dao dao.VersionDAO, cache cache.VersionCache) VersionRepository {
	return &versionRepository{
		dao:   dao,
		cache: cache,
	}
}

func (repo *versionRepository) Create(ctx context.Context, v domain.Version) error {
	return repo.dao.Insert(ctx, repo.toEntity(v))
}

func (repo *versionRepository) Delete(ctx context.Context, id, detailId, userId string) (int64, error) {
	return repo.dao.Delete(ctx, id, detailId, userId)
}

func (repo *versionRepository) Update(ctx context.Context, id, detailId, userId string, v domain.Version) (int64, error) {
	rowsAffected, err := repo.dao.Update(ctx, id, detailId, userId, repo.toEntity(v))
	if err != nil {
		return 0, err
	}
	// 更新缓存
	version, err := repo.dao.FindByIdAndDetailIdAndUserId(ctx, id, detailId, userId)
	if err != nil {
		return rowsAffected, err
	}
	versionInfo := repo.toDomain(version)
	err = repo.cache.Set(ctx, versionInfo)
	if err != nil {
		zap.L().Error("Redis异常", zap.Error(err))
		return rowsAffected, err
	}
	return rowsAffected, err
}

func (repo *versionRepository) UpdateFormal(ctx context.Context, id, detailId, userId string) (int64, error) {
	rowsAffected, err := repo.dao.UpdateFormal(ctx, id, detailId, userId)
	if err != nil {
		return 0, err
	}
	// 更新缓存
	version, err := repo.dao.FindByIdAndDetailIdAndUserId(ctx, id, detailId, userId)
	if err != nil {
		return rowsAffected, err
	}
	versionInfo := repo.toDomain(version)
	err = repo.cache.Set(ctx, versionInfo)
	if err != nil {
		zap.L().Error("Redis异常", zap.Error(err))
		return rowsAffected, err
	}
	return rowsAffected, err
}

func (repo *versionRepository) FindByIdAndDetailIdAndUserId(ctx context.Context, id, detailId, userId string) (domain.Version, error) {
	v, err := repo.cache.Get(ctx, id, userId)
	switch {
	case err == nil:
		return *v, err
	case errors.Is(err, cache.ErrVersionNotExist):
		version, err := repo.dao.FindByIdAndDetailIdAndUserId(ctx, id, detailId, userId)
		if err != nil {
			return domain.Version{}, err
		}

		versionInfo := repo.toDomain(version)
		err = repo.cache.Set(ctx, versionInfo)
		if err != nil {
			zap.L().Error("Redis异常", zap.Error(err))
			return domain.Version{}, err
		}
		return versionInfo, err
	default:
		return domain.Version{}, err
	}
}

func (repo *versionRepository) FindByFormalVersionCode(ctx context.Context, detailId, userId string) (domain.Version, error) {
	version, err := repo.dao.FindByFormalVersionCode(ctx, detailId, userId)
	if err != nil {
		return domain.Version{}, err
	}
	versionInfo := repo.toDomain(version)
	return versionInfo, err
}

func (repo *versionRepository) FindByDetailIdAndFormalVersion(ctx context.Context, detailId string) (domain.Version, error) {
	version, err := repo.dao.FindByDetailIdAndFormalVersion(ctx, detailId)
	if err != nil {
		return domain.Version{}, err
	}
	versionInfo := repo.toDomain(version)
	return versionInfo, err
}

func (repo *versionRepository) FindListAll(ctx context.Context, detailId, userId string) (int64, []domain.Version, error) {
	rows, list, err := repo.dao.FindListAll(ctx, detailId, userId)
	if err != nil {
		return rows, []domain.Version{}, err
	}
	var versions []domain.Version
	for _, v := range list {
		versions = append(versions, repo.toDomain(v))
	}
	return rows, versions, err
}

func (repo *versionRepository) ExistsByDetailIdAndFormalVersion(ctx context.Context, detailId string) (bool, error) {
	return repo.dao.ExistsByDetailIdAndFormalVersion(ctx, detailId)
}

func (repo *versionRepository) ExistsByDetailIdAndCode(ctx context.Context, userId, detailId, versionCode string) (bool, error) {
	return repo.dao.ExistsByDetailIdAndCode(ctx, userId, detailId, versionCode)
}

func (repo *versionRepository) toEntity(d domain.Version) model.Version {
	return model.Version{
		CoreModels: models.CoreModels{
			Creator:  d.Creator,
			Modifier: d.Modifier,
			Remark:   d.Remark,
		},
		VersionCode:    d.VersionCode,
		Description:    d.Description,
		HBuilderUpdate: d.HBuilderUpdate,
		ForceUpdate:    d.ForceUpdate,
		ApkUrl:         d.ApkUrl,
		WgtUrl:         d.WgtUrl,
		DetailId:       d.DetailId,
		UserId:         d.UserId,
	}
}

func (repo *versionRepository) toDomain(v *model.Version) domain.Version {
	return domain.Version{
		Version:    *v,
		CreateTime: v.CoreModels.CreateTime.Format("2006-01-02 15:04:05.000"),
		UpdateTime: v.CoreModels.UpdateTime.Format("2006-01-02 15:04:05.000"),
	}
}
