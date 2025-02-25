/**
 * Description：
 * FileName：aliYun.go
 * Author：CJiaの用心
 * Create：2025/2/22 21:43:45
 * Remark：
 */

package service

import (
	"context"
	"github.com/carefuly/PackageFluxApp/internal/service/oss"
	"mime/multipart"
)

type AliYunService interface {
	BatchUploadFiles(ctx context.Context, email string, files []*multipart.FileHeader) ([]*multipart.FileHeader, string, error)
	BatchDeleteFiles(ctx context.Context, paths []string) error
}

type aliYunService struct {
	oss oss.StorageOSS
}

func NewAliYunService(oss oss.StorageOSS) AliYunService {
	return &aliYunService{
		oss: oss,
	}
}

func (svc *aliYunService) BatchUploadFiles(ctx context.Context, email string, files []*multipart.FileHeader) ([]*multipart.FileHeader, string, error) {
	path := "packageFluxApp" + "/" + email
	uploadFiles, err := svc.oss.BatchUploadFiles(ctx, path, files)
	return uploadFiles, path, err
}

func (svc *aliYunService) BatchDeleteFiles(ctx context.Context, paths []string) error {
	return svc.oss.BatchDeleteFiles(ctx, paths)
}
