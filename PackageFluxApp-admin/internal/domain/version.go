/**
 * Description：
 * FileName：version.go
 * Author：CJiaの用心
 * Create：2025/2/28 12:34:00
 * Remark：
 */

package domain

import "github.com/carefuly/PackageFluxApp/internal/model"

type Version struct {
	model.Version
	CreateTime string `json:"createTime"` // 创建时间
	UpdateTime string `json:"updateTime"` // 更新时间
}
