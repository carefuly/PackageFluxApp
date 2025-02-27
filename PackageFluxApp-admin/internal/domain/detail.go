/**
 * Description：
 * FileName：detail.go
 * Author：CJiaの用心
 * Create：2025/2/27 10:34:12
 * Remark：
 */

package domain

import "github.com/carefuly/PackageFluxApp/internal/model"

type Detail struct {
	model.Detail
	Preview    []string `json:"preview"`    // 预览图
	CreateTime string `json:"createTime"` // 创建时间
	UpdateTime string `json:"updateTime"` // 更新时间
}
