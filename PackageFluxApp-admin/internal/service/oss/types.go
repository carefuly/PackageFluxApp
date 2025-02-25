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
	BatchDeleteFiles(ctx context.Context, paths []string) error
}

