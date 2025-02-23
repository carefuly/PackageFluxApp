/**
 * Description：
 * FileName：types.go
 * Author：CJiaの用心
 * Create：2025/2/22 21:08:06
 * Remark：
 */

package oss

import (
	"context"
	"mime/multipart"
)

// StorageOSS 是一个通用的 OSS 操作接口
type StorageOSS interface {
	BatchUploadFiles(ctx context.Context, path string, files []*multipart.FileHeader) ([]*multipart.FileHeader, error)

	// UploadFile(ctx context.Context, bucketName, objectKey, filePath string) error
	// BatchUploadFiles(ctx context.Context, bucketName string, files map[string]string) error // 新增批量上传接

	// DeleteFile(bucketName, objectKey string) error
}
