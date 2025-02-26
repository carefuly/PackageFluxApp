/**
 * Description：
 * FileName：files.go
 * Author：CJiaの用心
 * Create：2025/2/24 16:48:32
 * Remark：
 */

package domain

import (
	"github.com/carefuly/PackageFluxApp/model"
	"github.com/carefuly/PackageFluxApp/pkg/ginx/query/filters"
)

type FiltersFile struct {
	filters.Filters
	UserId string `json:"user_id"` // 用户ID
}

type Files struct {
	model.File
	CreateTime string `json:"createTime"` // 创建时间
	UpdateTime string `json:"updateTime"` // 更新时间
}
