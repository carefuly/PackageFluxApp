/**
 * Description：
 * FileName：version.go
 * Author：CJiaの用心
 * Create：2025/2/25 15:47:27
 * Remark：
 */

package domain

import (
	"github.com/carefuly/PackageFluxApp/model"
	"github.com/carefuly/PackageFluxApp/pkg/ginx/query/filters"
)

type FiltersVersion struct {
	filters.Filters
	DetailId string `json:"detail_id"` // 应用详情ID
}

type FormalVersion struct {
	DetailId  string `json:"detail_id"`  // 应用详情ID
	VersionId string `json:"version_id"` // 版本ID
}

type Version struct {
	model.Version
	CreateTime string `json:"createTime"` // 创建时间
	UpdateTime string `json:"updateTime"` // 更新时间
}
