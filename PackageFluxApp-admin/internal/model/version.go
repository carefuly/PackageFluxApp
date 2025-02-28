/**
 * Description：
 * FileName：version.go
 * Author：CJiaの用心
 * Create：2025/2/28 11:49:38
 * Remark：
 */

package model

import (
	"github.com/carefuly/PackageFluxApp/pkg/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Version struct {
	models.CoreModels
	VersionCode    string  `gorm:"type:varchar(20);index:idx_detail_version_code,unique;not null;column:versionCode;comment:版本号" json:"versionCode"` // 版本号
	Description    string  `gorm:"type:varchar(255);column:description;comment:更新说明" json:"description"`                                             // 更新说明
	HBuilderUpdate bool    `gorm:"type:bool;default:false;column:HBuilderUpdate;comment:HBuilderX是否更新" json:"HBuilderUpdate"`                        // HBuilderX是否更新
	ForceUpdate    bool    `gorm:"type:bool;default:false;column:forceUpdate;comment:是否强制更新" json:"forceUpdate"`                                     // 是否强制更新
	FormalVersion  bool    `gorm:"type:bool;default:false;column:formalVersion;comment:是否是正式版本" json:"formalVersion"`                                // 是否是正式版本
	ApkUrl         string  `gorm:"type:varchar(512);not null;column:apkUrl;comment:apk下载地址" json:"apkUrl"`                                           // apk下载地址
	WgtUrl         string  `gorm:"type:varchar(512);not null;column:wgtUrl;comment:wgt下载地址" json:"wgtUrl"`                                           // wgt下载地址
	DetailId       string  `gorm:"type:varchar(100);index:idx_detail_version_code,unique;column:detail_id;comment:应用详情ID" json:"detail_id"`          // 应用详情ID
	Detail         *Detail `gorm:"foreignKey:DetailId;constraint:OnDelete:CASCADE;"`                                                                 // 应用详情
	UserId         string  `gorm:"type:varchar(100);column:user_id;comment:用户ID" json:"user_id"`                                                     // 用户ID
	User           *User   `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`                                                  // 用户详情
}

func NewVersion() *Version {
	return &Version{}
}

func (u *Version) AutoMigrate(db *gorm.DB) {
	err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Version{})
	if err != nil {
		zap.L().Error("Version表模型迁移失败", zap.Error(err))
	}
}
