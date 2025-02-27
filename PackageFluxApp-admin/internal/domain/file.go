/**
 * Description：
 * FileName：file.go
 * Author：CJiaの用心
 * Create：2025/2/27 20:58:04
 * Remark：
 */

package domain

import "github.com/carefuly/PackageFluxApp/internal/model"

type File struct {
	model.File
	CreateTime string `json:"createTime"` // 创建时间
	UpdateTime string `json:"updateTime"` // 更新时间
}
