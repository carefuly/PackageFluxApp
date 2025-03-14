/**
 * Description：
 * FileName：logger.go
 * Author：CJiaの用心
 * Create：2025/3/14 15:25:38
 * Remark：
 */

package domain

import "github.com/carefuly/PackageFluxApp/internal/model"

type Logger struct {
	model.Logger
	CreateTime string `json:"createTime"` // 创建时间
	UpdateTime string `json:"updateTime"` // 更新时间
}
