/**
 * Description：
 * FileName：details.go
 * Author：CJiaの用心
 * Create：2025/2/24 16:14:24
 * Remark：
 */

package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/carefuly/PackageFluxApp/internal/domain"
	"github.com/carefuly/PackageFluxApp/internal/repository/dao"
	"github.com/carefuly/PackageFluxApp/model"
	"github.com/carefuly/PackageFluxApp/pkg/models"
)

var (
	ErrDuplicateAppName = dao.ErrDuplicateAppName
	ErrRecordNotFound   = dao.ErrRecordNotFound
)

type DetailsRepository interface {
	Create(ctx context.Context, detail domain.Detail) error
	Delete(ctx context.Context, id, userId string) error
	Update(ctx context.Context, id string, detail domain.Detail) error
	FindById(ctx context.Context, id, userId string) (domain.Detail, error)
	FindListAll(ctx context.Context, f domain.FiltersDetail) ([]domain.Detail, error)
}

type detailsRepository struct {
	dao dao.DetailsDAO
}

func NewDetailsRepository(dao dao.DetailsDAO) DetailsRepository {
	return &detailsRepository{
		dao: dao,
	}
}

func (repo *detailsRepository) Create(ctx context.Context, details domain.Detail) error {
	return repo.dao.Create(ctx, repo.toEntity(details))
}

func (repo *detailsRepository) Delete(ctx context.Context, id, userId string) error {
	return repo.dao.Delete(ctx, id, userId)
}

func (repo *detailsRepository) Update(ctx context.Context, id string, detail domain.Detail) error {
	return repo.dao.Update(ctx, id, repo.toEntity(detail))
}

func (repo *detailsRepository) FindById(ctx context.Context, id, userId string) (domain.Detail, error) {
	d, err := repo.dao.FindById(ctx, id, userId)
	if err != nil {
		return domain.Detail{}, err
	}
	return repo.toDomain(d), err
}

func (repo *detailsRepository) FindListAll(ctx context.Context, f domain.FiltersDetail) ([]domain.Detail, error) {
	list, err := repo.dao.FindListAll(ctx, f)
	if err != nil {
		return []domain.Detail{}, err
	}
	var details []domain.Detail

	if len(list) == 0 {
		return []domain.Detail{}, err
	}

	for _, v := range list {
		details = append(details, repo.toDomain(v))
	}
	return details, err
}

func (repo *detailsRepository) toDomain(d model.Detail) domain.Detail {
	var decodedPreview []string
	_ = json.Unmarshal([]byte(d.Preview), &decodedPreview)
	return domain.Detail{
		Detail:     d,
		Preview:    decodedPreview,
		CreateTime: d.CoreModels.CreateTime.Format("2006-01-02 15:04:05.000"),
		UpdateTime: d.CoreModels.UpdateTime.Format("2006-01-02 15:04:05.000"),
	}
}

func (repo *detailsRepository) toEntity(d domain.Detail) model.Detail {
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
		UserID:      d.UserID,
	}
}
