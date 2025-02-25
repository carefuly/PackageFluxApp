/**
 * Description：
 * FileName：aliYun.go
 * Author：CJiaの用心
 * Create：2025/2/22 21:39:01
 * Remark：
 */

package oss

import (
	"context"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mime/multipart"
	"net/http"
	"strings"
)

// AliYunOSS 是阿里云 OSS 的实现
type AliYunOSS struct {
	Endpoint        string // OSS 端点
	BucketName      string // 存储空间名称
	AccessKeyID     string // 访问密钥 ID
	AccessKeySecret string // 访问密钥 Secret
}

func NewAliYunOSS(Endpoint string, BucketName string, AccessKeyID string, AccessKeySecret string) StorageOSS {
	return &AliYunOSS{
		Endpoint:        Endpoint,
		BucketName:      BucketName,
		AccessKeyID:     AccessKeyID,
		AccessKeySecret: AccessKeySecret,
	}
}

func (a *AliYunOSS) BatchUploadFiles(ctx context.Context, path string, files []*multipart.FileHeader) ([]*multipart.FileHeader, error) {
	// 从环境变量中获取访问凭证。运行本代码示例之前，请确保已设置环境变量OSS_ACCESS_KEY_ID和OSS_ACCESS_KEY_SECRET
	provider, err := oss.NewEnvironmentVariableCredentialsProvider()
	if err != nil {
		return []*multipart.FileHeader{}, fmt.Errorf("无法创建凭证提供程序: %v", err)
	}

	clientOptions := []oss.ClientOption{oss.SetCredentialsProvider(&provider)}
	clientOptions = append(clientOptions, oss.Region("yourRegion"))

	// 创建 OSS 客户端
	client, err := oss.New(a.Endpoint, a.AccessKeyID, a.AccessKeySecret, clientOptions...)
	if err != nil {
		return []*multipart.FileHeader{}, fmt.Errorf("无法创建阿里云OSS客户端: %w", err)
	}

	// 获取存储空间
	bucket, err := client.Bucket(a.BucketName)
	if err != nil {
		return []*multipart.FileHeader{}, fmt.Errorf("无法获取 bucket: %w", err)
	}

	var responseHeader http.Header

	// 检查文件夹是否存在，不存在则创建
	isExist, err := bucket.IsObjectExist(path+"/", oss.GetResponseHeader(&responseHeader))
	if err != nil {
		return []*multipart.FileHeader{}, fmt.Errorf("无法获取bucket文件夹信息: %w", err)
	}

	if !isExist {
		err := bucket.PutObject(path+"/", strings.NewReader(""))
		if err != nil {
			return []*multipart.FileHeader{}, fmt.Errorf("无法创建 bucket 文件夹: %w", err)
		}
	}

	var fileNames []*multipart.FileHeader
	for _, file := range files {
		open, err := file.Open()
		if err != nil {
			return []*multipart.FileHeader{}, fmt.Errorf("无法打开文件: %w", err)
		}

		// 获取文件名
		objectKey := path + "/" + file.Filename

		// 上传文件
		err = bucket.PutObject(objectKey, open)
		if err != nil {

		} else {
			fileNames = append(fileNames, file)
		}
	}

	return fileNames, nil
}

func (a *AliYunOSS) BatchDeleteFiles(ctx context.Context, paths []string) error {
	// 从环境变量中获取访问凭证。运行本代码示例之前，请确保已设置环境变量OSS_ACCESS_KEY_ID和OSS_ACCESS_KEY_SECRET
	provider, err := oss.NewEnvironmentVariableCredentialsProvider()
	if err != nil {
		return fmt.Errorf("无法创建凭证提供程序: %v", err)
	}

	clientOptions := []oss.ClientOption{oss.SetCredentialsProvider(&provider)}
	clientOptions = append(clientOptions, oss.Region("yourRegion"))

	// 创建 OSS 客户端
	client, err := oss.New(a.Endpoint, a.AccessKeyID, a.AccessKeySecret, clientOptions...)
	if err != nil {
		return fmt.Errorf("无法创建阿里云OSS客户端: %w", err)
	}

	// 获取存储空间
	bucket, err := client.Bucket(a.BucketName)
	if err != nil {
		return fmt.Errorf("无法获取 bucket: %w", err)
	}

	for _, path := range paths {
		fmt.Println("path", path)
		// 删除文件
		err = bucket.DeleteObject(path)
		if err != nil {
			return fmt.Errorf("无法删除文件: %w", err)
		}
	}
	return nil
}
