/**
 * Description：
 * FileName：aliYun.go
 * Author：CJiaの用心
 * Create：2025/2/22 21:25:02
 * Remark：
 */

package config

type AliYunConfig struct {
	Endpoint        string `yaml:"endpoint" json:"endpoint"`
	BucketName      string `yaml:"bucketName" json:"bucketName"`
	OssPrefix      string `yaml:"ossPrefix" json:"ossPrefix"`
	AccessKeyID     string `yaml:"accessKeyID" json:"accessKeyID"`
	AccessKeySecret string `yaml:"accessKeySecret" json:"accessKeySecret"`
}
