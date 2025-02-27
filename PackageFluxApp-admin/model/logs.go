/**
 * Description：
 * FileName：logs.go
 * Author：CJiaの用心
 * Create：2025/2/20 10:03:59
 * Remark：
 */

package model

import (
	"github.com/carefuly/PackageFluxApp/internal/model"
	"github.com/carefuly/PackageFluxApp/pkg/models"
)

type Logger struct {
	models.CoreModels
	LogoUrl      string  `gorm:"type:text;comment:logo地址" json:"logo_url"`                           // logo地址
	AppName      string  `gorm:"type:varchar(50);not null;comment:应用名称" json:"app_name"`           // 应用名称
	UniAppId     string  `gorm:"type:varchar(50);not null;comment:应用标识AppID" json:"uni_app_id"`    // 应用标识AppID
	UniAppRemark string  `gorm:"type:varchar(50);comment:应用标识备注" json:"uni_app_remark"`          // 应用标识备注
	VersionName  string  `gorm:"type:varchar(50);comment:版本名称" json:"version_name"`                // 版本名称
	VersionCode  string  `gorm:"type:varchar(50);comment:版本号" json:"version_code"`                  // 版本号
	Brand        string  `gorm:"type:varchar(50);comment:设备品牌" json:"brand"`                       // 设备品牌
	SystemName   string  `gorm:"type:varchar(50);comment:系统名称" json:"system_name"`                 // 系统名称
	DetailId     string  `gorm:"type:varchar(100);comment:应用详情ID" json:"detail_id"`                // 应用详情ID
	Detail       *model.Detail `gorm:"foreignKey:DetailID;references:RecordId;constraint:OnDelete:CASCADE;"` // 应用详情
}
