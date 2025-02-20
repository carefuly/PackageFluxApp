/**
 * Description：
 * FileName：logs.go
 * Author：CJiaの用心
 * Create：2025/2/20 10:03:59
 * Remark：
 */

package model

import "github.com/carefuly/PackageFluxApp/pkg/models"

type Logger struct {
	models.CoreModels
	LogoUrl      string `gorm:"type:text;comment:logo地址" json:"logo_url"`
	AppName      string `gorm:"type:varchar(50);not null;comment:应用名称" json:"app_name"`
	UniAppId     string `gorm:"type:varchar(50);not null;comment:应用标识AppID" json:"uni_app_id"`
	UniAppRemark string `gorm:"type:varchar(50);comment:应用标识备注" json:"uni_app_remark"`
	VersionName  string `gorm:"type:varchar(50);comment:版本名称" json:"version_name"`
	VersionCode  string `gorm:"type:varchar(50);comment:版本号" json:"version_code"`
	Brand        string `gorm:"type:varchar(50);comment:设备品牌" json:"brand"`
	SystemName   string `gorm:"type:varchar(50);comment:系统名称" json:"system_name"`
	DetailID     int    `gorm:"type:int;comment:应用详情ID" json:"detail_id"`
	Detail       Detail `gorm:"foreignKey:DetailID;constraint:OnDelete:CASCADE;"`
}
