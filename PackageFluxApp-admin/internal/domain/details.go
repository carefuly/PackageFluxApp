/**
 * Description：
 * FileName：details.go
 * Author：CJiaの用心
 * Create：2025/2/24 16:10:03
 * Remark：
 */

package domain

import (
	"github.com/carefuly/PackageFluxApp/model"
	"github.com/carefuly/PackageFluxApp/pkg/ginx/query/filters"
)

type FiltersDetail struct {
	filters.Filters
	UserId  string `json:"user_id"` // 用户ID
	AppName string `json:"appName"` // 应用名称
}

type Detail struct {
	model.Detail
	Preview    []string `json:"preview"`    // 预览图
	CreateTime string   `json:"createTime"` // 创建时间
	UpdateTime string   `json:"updateTime"` // 更新时间
}
