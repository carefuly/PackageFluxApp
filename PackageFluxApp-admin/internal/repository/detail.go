/**
 * Description：
 * FileName：detail.go
 * Author：CJiaの用心
 * Create：2025/2/27 10:42:15
 * Remark：
 */

package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/carefuly/PackageFluxApp/internal/domain"
	"github.com/carefuly/PackageFluxApp/internal/model"
	"github.com/carefuly/PackageFluxApp/internal/repository/cache"
	"github.com/carefuly/PackageFluxApp/internal/repository/dao"
	"github.com/carefuly/PackageFluxApp/pkg/models"
	"go.uber.org/zap"
)

var (
	ErrDetailRecordNotFound      = dao.ErrDetailRecordNotFound
	ErrDetailNotFound            = dao.ErrDetailNotFound
	ErrDuplicateUserIdAndAppName = dao.ErrDuplicateUserIdAndAppName
)

type DetailRepository interface {
	Create(ctx context.Context, d domain.Detail) error
	Delete(ctx context.Context, id, userId string) (int64, error)
	Update(ctx context.Context, id, userId string, d domain.Detail) (int64, error)
	UpdateFormal(ctx context.Context, id, userId, versionId string) (int64, error)
	FindByIdAndUserId(ctx context.Context, id, userId string) (domain.Detail, error)
	FindListAll(ctx context.Context, userId, appName string) (int64, []domain.Detail, error)
	ExistsByUserIdAndAppName(ctx context.Context, userId, appName string) (bool, error)
}

type detailRepository struct {
	dao   dao.DetailDAO
	cache cache.DetailCache
}

func NewDetailRepository(dao dao.DetailDAO, cache cache.DetailCache) DetailRepository {
	return &detailRepository{
		dao:   dao,
		cache: cache,
	}
}

func (repo *detailRepository) Create(ctx context.Context, d domain.Detail) error {
	return repo.dao.Insert(ctx, repo.toEntity(d))
}

func (repo *detailRepository) Delete(ctx context.Context, id, userId string) (int64, error) {
	return repo.dao.Delete(ctx, id, userId)
}

func (repo *detailRepository) Update(ctx context.Context, id, userId string, d domain.Detail) (int64, error) {
	rowsAffected, err := repo.dao.Update(ctx, id, userId, repo.toEntity(d))
	if err != nil {
		return 0, err
	}
	// 更新缓存
	detail, err := repo.dao.FindByIdAndUserId(ctx, id, userId)
	if err != nil {
		return rowsAffected, err
	}

	detailInfo := repo.toDomain(detail)
	err = repo.cache.Set(ctx, detailInfo)
	if err != nil {
		zap.L().Error("Redis异常", zap.Error(err))
		return rowsAffected, err
	}
	return rowsAffected, err
}

func (repo *detailRepository) UpdateFormal(ctx context.Context, id, userId, versionId string) (int64, error) {
	rowsAffected, err := repo.dao.UpdateFormal(ctx, id, userId, versionId)
	if err != nil {
		return 0, err
	}
	// 更新缓存
	detail, err := repo.dao.FindByIdAndUserId(ctx, id, userId)
	if err != nil {
		return rowsAffected, err
	}

	detailInfo := repo.toDomain(detail)
	err = repo.cache.Set(ctx, detailInfo)
	if err != nil {
		zap.L().Error("Redis异常", zap.Error(err))
		return rowsAffected, err
	}
	return rowsAffected, err
}

func (repo *detailRepository) FindByIdAndUserId(ctx context.Context, id, userId string) (domain.Detail, error) {
	d, err := repo.cache.Get(ctx, id)
	switch {
	case err == nil:
		return *d, err
	case errors.Is(err, cache.ErrDetailNotExist):
		detail, err := repo.dao.FindByIdAndUserId(ctx, id, userId)
		if err != nil {
			return domain.Detail{}, err
		}

		detailInfo := repo.toDomain(detail)

		err = repo.cache.Set(ctx, detailInfo)
		if err != nil {
			zap.L().Error("Redis异常", zap.Error(err))
			return domain.Detail{}, err
		}
		return detailInfo, err
	default:
		return domain.Detail{}, err
	}
}

func (repo *detailRepository) FindListAll(ctx context.Context, userId, appName string) (int64, []domain.Detail, error) {
	rows, list, err := repo.dao.FindListAll(ctx, userId, appName)
	if err != nil {
		return rows, []domain.Detail{}, err
	}
	var details []domain.Detail
	for _, v := range list {
		details = append(details, repo.toDomain(v))
	}
	return rows, details, err
}

func (repo *detailRepository) ExistsByUserIdAndAppName(ctx context.Context, userId, appName string) (bool, error) {
	return repo.dao.ExistsByUserIdAndAppName(ctx, userId, appName)
}

func (repo *detailRepository) toEntity(d domain.Detail) model.Detail {
	var preview bytes.Buffer
	_ = json.NewEncoder(&preview).Encode(d.Preview)
	return model.Detail{
		CoreModels: models.CoreModels{
			Creator:  d.Creator,
			Modifier: d.Modifier,
			Remark:   d.Remark,
		},
		LogoUrl:     d.LogoUrl,
		AppName:     d.AppName,
		AppleId:     d.AppleId,
		Description: d.Description,
		Preview:     preview.String(),
		UserId:      d.UserId,
		VersionId:   d.VersionId,
	}
}

func (repo *detailRepository) toDomain(d *model.Detail) domain.Detail {
	var decodedPreview []string
	_ = json.Unmarshal([]byte(d.Preview), &decodedPreview)
	return domain.Detail{
		Detail:     *d,
		Preview:    decodedPreview,
		CreateTime: d.CoreModels.CreateTime.Format("2006-01-02 15:04:05.000"),
		UpdateTime: d.CoreModels.UpdateTime.Format("2006-01-02 15:04:05.000"),
	}
}
