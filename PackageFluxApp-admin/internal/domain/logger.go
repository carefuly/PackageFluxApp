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

type LoggerStatistics struct {
	OpenNum      int64 `json:"openNum"`      // 打开数
	OpenTotalNum int64 `json:"openTotalNum"` // 累计打开数
	UserNum      int64 `json:"userNum"`      // 访问人数
	UserTotalNum int64 `json:"userTotalNum"` // 累计访问人数
}
