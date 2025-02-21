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
	LogoUrl     string   `gorm:"type:text;comment:logo地址" json:"logoUrl"`
	AppName     string   `gorm:"type:varchar(50);not null;comment:应用名称" json:"appName"`
	AppleId     string   `gorm:"type:varchar(50);comment:appleId" json:"appleId"`
	Description string   `gorm:"type:varchar(255);comment:应用描述" json:"description"`
	Type        string   `gorm:"type:int;default:1;comment:1:APP" json:"type"`
	Preview     string   `gorm:"type:json;comment:预览图" json:"preview"`
	UserID      uint     `gorm:"type:int;comment:用户ID" json:"user_id"`
	User        *User    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	VersionID   *uint    `gorm:"type:int;comment:版本ID" json:"version_id"`
	Version     *Version `gorm:"foreignKey:VersionID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// AfterFind Hook
// func (c *Detail) AfterFind(tx *gorm.DB) (err error) {
// 	// 默认排序规则：按创建时间降序，按 Sort 升序
// 	tx.Order("create_time DESC, sort ASC").Find(c)
// 	return
// }
