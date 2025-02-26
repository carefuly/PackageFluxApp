/**
 * Description：
 * FileName：version.go
 * Author：CJiaの用心
 * Create：2025/2/25 15:50:47
 * Remark：
 */

package repository

import (
	"context"
	"github.com/carefuly/PackageFluxApp/internal/domain"
	"github.com/carefuly/PackageFluxApp/internal/repository/dao"
	"github.com/carefuly/PackageFluxApp/model"
	"github.com/carefuly/PackageFluxApp/pkg/models"
)

var (
	ErrDuplicateVersionCode = dao.ErrDuplicateVersionCode
	ErrVersionNotFound      = dao.ErrVersionNotFound
	ErrFormalNotFound       = dao.ErrFormalNotFound
)

type VersionRepository interface {
	Create(ctx context.Context, version domain.Version) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string, version domain.Version) error
	FindLatestFormalVersion(ctx context.Context, detailId string) (domain.Version, error)
	Formal(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (domain.Version, error)
	FindListAll(ctx context.Context, f domain.FiltersVersion) ([]domain.Version, error)
}

type versionRepository struct {
	dao dao.VersionDAO
}

func NewVersionRepository(dao dao.VersionDAO) VersionRepository {
	return &versionRepository{
		dao: dao,
	}
}

func (repo *versionRepository) Create(ctx context.Context, version domain.Version) error {
	// 检查 DetailID 和 VersionCode 是否已存在
	exists, err := repo.dao.ExistsByDetailIdAndVersionCode(ctx, version.DetailID, version.VersionCode)
	if err != nil {
		return err
	}
	if exists {
		return dao.ErrDuplicateVersionCode
	}
	return repo.dao.Create(ctx, repo.toEntity(version))
}

func (repo *versionRepository) Delete(ctx context.Context, id string) error {
	return repo.dao.Delete(ctx, id)
}

func (repo *versionRepository) Update(ctx context.Context, id string, version domain.Version) error {
	// 检查 DetailID 和 VersionCode 是否已存在
	exists, err := repo.dao.ExistsByDetailIdAndVersionCodeExcludeCurrent(ctx, version.DetailID, version.VersionCode, version.ID)
	if err != nil {
		return err
	}
	if exists {
		return dao.ErrDuplicateVersionCode
	}
	return repo.dao.Update(ctx, id, repo.toEntity(version))
}

func (repo *versionRepository) FindLatestFormalVersion(ctx context.Context, detailId string) (domain.Version, error) {
	return repo.dao.FindLatestFormalVersion(ctx, detailId)
}

func (repo *versionRepository) Formal(ctx context.Context, id string) error  {
	return repo.dao.Formal(ctx, id)
}

func (repo *versionRepository) FindById(ctx context.Context, id string) (domain.Version, error) {
	d, err := repo.dao.FindById(ctx, id)
	if err != nil {
		return domain.Version{}, err
	}
	return repo.toDomain(d), err
}

func (repo *versionRepository) FindListAll(ctx context.Context, f domain.FiltersVersion) ([]domain.Version, error) {
	list, err := repo.dao.FindListAll(ctx, f)
	if err != nil {
		return []domain.Version{}, err
	}
	if len(list) == 0 {
		return []domain.Version{}, err
	}

	var versions []domain.Version
	for _, v := range list {
		versions = append(versions, repo.toDomain(v))
	}
	return versions, err
}

func (repo *versionRepository) toDomain(v model.Version) domain.Version {
	return domain.Version{
		Version:    v,
		CreateTime: v.CoreModels.CreateTime.Format("2006-01-02 15:04:05.000"),
		UpdateTime: v.CoreModels.UpdateTime.Format("2006-01-02 15:04:05.000"),
	}
}

func (repo *versionRepository) toEntity(v domain.Version) model.Version {
	return model.Version{
		CoreModels: models.CoreModels{
			Creator:  v.Creator,
			Modifier: v.Modifier,
			Remark:   v.Remark,
		},
		VersionCode:    v.VersionCode,
		Description:    v.Description,
		HBuilderUpdate: v.HBuilderUpdate,
		ForceUpdate:    v.ForceUpdate,
		FormalVersion:  v.FormalVersion,
		ApkUrl:         v.ApkUrl,
		WgtUrl:         v.WgtUrl,
		DetailID:       v.DetailID,
	}
}
