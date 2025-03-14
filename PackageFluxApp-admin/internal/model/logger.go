/**
 * Description：
 * FileName：logger.go
 * Author：CJiaの用心
 * Create：2025/3/14 15:13:41
 * Remark：
 */

package model

import (
	"github.com/carefuly/PackageFluxApp/pkg/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Logger struct {
	models.CoreModels
	AppId      string  `gorm:"type:varchar(100);column:appId;comment:应用标识Id" json:"appId"`                         // 应用标识Id
	AppName    string  `gorm:"type:varchar(100);column:appName;comment:应用名称" json:"appName"`                       // 应用名称
	AppVersion string  `gorm:"type:varchar(100);column:appVersion;comment:应用版本(检查的版本)" json:"appVersion"`          // 应用版本
	Brand      string  `gorm:"type:varchar(100);column:brand;comment:设备品牌" json:"brand"`                           // 设备品牌
	SystemName string  `gorm:"type:varchar(100);column:systemName;comment:系统名称" json:"systemName"`                 // 系统名称
	UniqueId   string  `gorm:"type:varchar(100);column:uniqueId;comment:应用唯一Id(可以是用户Id或用户名，方便统计)" json:"uniqueId"` // 应用唯一Id
	UniRemark  string  `gorm:"type:varchar(100);column:uniRemark;comment:应用标识备注" json:"uniRemark"`                 // 应用标识备注
	DetailId   string  `gorm:"type:varchar(100);column:detail_id;comment:应用详情ID" json:"detail_id"`                 // 应用详情ID
	Detail     *Detail `gorm:"foreignKey:DetailId;constraint:OnDelete:CASCADE;"`                                   // 应用详情
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) AutoMigrate(db *gorm.DB) {
	err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Logger{})
	if err != nil {
		zap.L().Error("Logger表模型迁移失败", zap.Error(err))
	}
}
