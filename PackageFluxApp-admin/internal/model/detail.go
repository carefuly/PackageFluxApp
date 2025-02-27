/**
 * Description：
 * FileName：detail.go
 * Author：CJiaの用心
 * Create：2025/2/27 10:13:40
 * Remark：
 */

package model

import (
	"github.com/carefuly/PackageFluxApp/pkg/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Detail struct {
	models.CoreModels
	LogoUrl     string `gorm:"type:text;column:logoUrl;comment:logo地址" json:"logoUrl"`                                              // logo地址
	AppName     string `gorm:"type:varchar(50);index:idx_user_app_name,unique;not null;column:appName;comment:应用名称" json:"appName"` // 应用名称
	AppleId     string `gorm:"type:varchar(50);column:appleId;comment:appleId" json:"appleId"`                                      // appleId
	Description string `gorm:"type:varchar(255);column:description;comment:应用描述" json:"description"`                                // 应用描述
	Type        int    `gorm:"type:int;default:1;column:type;comment:1:APP" json:"type"`                                            // 1:APP
	Preview     string `gorm:"type:json;column:preview;comment:预览图" json:"preview"`                                                 // 预览图
	UserId      string `gorm:"type:varchar(100);index:idx_user_app_name,unique;column:user_id;comment:用户ID" json:"user_id"`         // 用户ID
	User        *User  `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`                                     // 用户详情
	// VersionID   *string  `gorm:"type:varchar(100);comment:版本ID" json:"version_id"`                                        // 版本ID
	// Version     *Version `gorm:"foreignKey:VersionID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 版本
}

func NewDetail() *Detail {
	return &Detail{}
}

func (u *Detail) AutoMigrate(db *gorm.DB) {
	err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Detail{})
	if err != nil {
		zap.L().Error("Detail表模型迁移失败", zap.Error(err))
	}
}
