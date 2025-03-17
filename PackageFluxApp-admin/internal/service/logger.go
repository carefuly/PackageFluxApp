/**
 * Description：
 * FileName：logger.go
 * Author：CJiaの用心
 * Create：2025/3/14 15:36:22
 * Remark：
 */

package service

import (
	"context"
	"github.com/carefuly/PackageFluxApp/internal/domain"
	"github.com/carefuly/PackageFluxApp/internal/repository"
	"github.com/carefuly/PackageFluxApp/pkg/ginx/query/filters"
)

type LoggerService interface {
	Create(ctx context.Context, l domain.Logger) error
	FindListPage(ctx context.Context, detailId, createTime string, page filters.Pagination) (int64, []domain.Logger, error)
	Statistics(ctx context.Context, detailId, createTime string) (domain.LoggerStatistics, error)
}

type loggerService struct {
	repo repository.LoggerRepository
}

func NewLoggerService(repo repository.LoggerRepository) LoggerService {
	return &loggerService{
		repo: repo,
	}
}

func (svc *loggerService) Create(ctx context.Context, l domain.Logger) error {
	return svc.repo.Create(ctx, l)
}

func (svc *loggerService) FindListPage(ctx context.Context, detailId, createTime string, page filters.Pagination) (int64, []domain.Logger, error) {
	total, list, err := svc.repo.FindListPage(ctx, detailId, createTime, page)
	if err != nil {
		return 0, []domain.Logger{}, err
	}
	if total == 0 {
		return 0, []domain.Logger{}, err
	}
	return total, list, nil
}

func (svc *loggerService) Statistics(ctx context.Context, detailId, createTime string) (domain.LoggerStatistics, error) {
	openNum, err := svc.repo.CountByDetailIdAndCreateTime(ctx, detailId, createTime)
	if err != nil {
		return domain.LoggerStatistics{}, err
	}

	userNum, err := svc.repo.CountDistinctUniqueIdByCreateTime(ctx, detailId, createTime)
	if err != nil {
		return domain.LoggerStatistics{}, err
	}

	openTotalNum, err := svc.repo.CountByDetailId(ctx, detailId)
	if err != nil {
		return domain.LoggerStatistics{}, err
	}

	userTotalNum, err := svc.repo.CountDistinctUniqueIdByDetailId(ctx, detailId)
	if err != nil {
		return domain.LoggerStatistics{}, err
	}

	return domain.LoggerStatistics{
		OpenNum:      openNum,
		OpenTotalNum: openTotalNum,
		UserNum:      userNum,
		UserTotalNum: userTotalNum,
	}, nil
}
