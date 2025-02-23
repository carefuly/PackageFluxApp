/**
 * Description：
 * FileName：files.go
 * Author：CJiaの用心
 * Create：2025/2/22 22:44:57
 * Remark：
 */

package domain

import "github.com/carefuly/PackageFluxApp/model"

type Files struct {
	model.File
	CreateTime string `json:"createTime"` // 创建时间
	UpdateTime string `json:"updateTime"` // 更新时间
}
