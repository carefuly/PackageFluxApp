/**
 * Description：
 * FileName：logger.go
 * Author：CJiaの用心
 * Create：2025/3/14 15:30:25
 * Remark：
 */

package repository

import (
	"context"
	"github.com/carefuly/PackageFluxApp/internal/domain"
	"github.com/carefuly/PackageFluxApp/internal/model"
	"github.com/carefuly/PackageFluxApp/internal/repository/dao"
	"github.com/carefuly/PackageFluxApp/pkg/ginx/query/filters"
	"github.com/carefuly/PackageFluxApp/pkg/models"
)

type LoggerRepository interface {
	Create(ctx context.Context, f domain.Logger) error
	FindListPage(ctx context.Context, detailId string, page filters.Pagination) (int64, []domain.Logger, error)
}

type loggerRepository struct {
	dao dao.LoggerDAO
}

func NewLoggerRepository(dao dao.LoggerDAO) LoggerRepository {
	return &loggerRepository{
		dao: dao,
	}
}

func (repo *loggerRepository) Create(ctx context.Context, l domain.Logger) error {
	return repo.dao.Insert(ctx, repo.toEntity(l))
}

func (repo *loggerRepository) FindListPage(ctx context.Context, detailId string, page filters.Pagination) (int64, []domain.Logger, error) {
	rows, list, err := repo.dao.FindListPage(ctx, detailId, page)
	if err != nil {
		return rows, []domain.Logger{}, err
	}
	var loggers []domain.Logger
	for _, v := range list {
		loggers = append(loggers, repo.toDomain(v))
	}
	return rows, loggers, err
}

func (repo *loggerRepository) toEntity(d domain.Logger) model.Logger {
	return model.Logger{
		CoreModels: models.CoreModels{
			Remark: d.Remark,
		},
		AppId:      d.AppId,
		AppName:    d.AppName,
		AppVersion: d.AppVersion,
		Brand:      d.Brand,
		SystemName: d.SystemName,
		UniqueId:   d.UniqueId,
		UniRemark:  d.UniRemark,
		DetailId:   d.DetailId,
	}
}

func (repo *loggerRepository) toDomain(l *model.Logger) domain.Logger {
	return domain.Logger{
		Logger:     *l,
		CreateTime: l.CoreModels.CreateTime.Format("2006-01-02 15:04:05.000"),
		UpdateTime: l.CoreModels.UpdateTime.Format("2006-01-02 15:04:05.000"),
	}
}
