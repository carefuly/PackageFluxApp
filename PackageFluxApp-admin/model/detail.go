/**
 * Description：
 * FileName：detail.go
 * Author：CJiaの用心
 * Create：2025/2/20 10:03:48
 * Remark：
 */

package model

import (
	"github.com/carefuly/PackageFluxApp/pkg/models"
)

type Detail struct {
	models.CoreModels
	LogoUrl     string `gorm:"type:text;comment:logo地址" json:"logoUrl"`                         // logo地址
	AppName     string `gorm:"type:varchar(50);not null;comment:应用名称" json:"appName"`           // 应用名称
	AppleId     string `gorm:"type:varchar(50);comment:appleId" json:"appleId"`                 // appleId
	Description string `gorm:"type:varchar(255);comment:应用描述" json:"description"`               // 应用描述
	Type        int    `gorm:"type:int;default:1;comment:1:APP" json:"type"`                    // 1:APP
	Preview     string `gorm:"type:json;comment:预览图" json:"preview"`                            // 预览图
	UserID      string `gorm:"type:varchar(100);comment:用户ID" json:"user_id"`                   // 用户ID
	User        *User  `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // 用户详情
	// VersionID   *string  `gorm:"type:varchar(100);comment:版本ID" json:"version_id"`                                      // 版本ID
	// Version     *Version `gorm:"foreignKey:VersionID;references:RecordId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 版本
}
