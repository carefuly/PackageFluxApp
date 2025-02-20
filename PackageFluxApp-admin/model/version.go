/**
 * Description：
 * FileName：version.go
 * Author：CJiaの用心
 * Create：2025/2/20 10:04:25
 * Remark：
 */

package model

import "github.com/carefuly/PackageFluxApp/pkg/models"

type Version struct {
	models.CoreModels
	VersionCode    string `gorm:"type:varchar(20);comment:版本号" json:"versionCode"`
	Description    string `gorm:"type:varchar(255);comment:更新说明" json:"description"`
	HBuilderUpdate bool   `gorm:"type:bool;default:false;comment:HBuilderX是否更新" json:"h_builder_update"`
	ForceUpdate    bool   `gorm:"type:bool;default:false;comment:是否强制更新" json:"force_update"`
	FormalVersion  bool   `gorm:"type:bool;default:false;comment:是否是正式版本" json:"formal_version"`
	ApkUrl         string `gorm:"type:varchar(255);comment:apk下载地址" json:"apk_url"`
	WgtUrl         string `gorm:"type:varchar(255);comment:wgt下载地址" json:"wgt_url"`
	// DetailID uint    `gorm:"type:int;comment:应用详情ID" json:"detail_id"`
	// Detail   *Detail `gorm:"foreignKey:DetailID;constraint:OnDelete:CASCADE;"`
}
