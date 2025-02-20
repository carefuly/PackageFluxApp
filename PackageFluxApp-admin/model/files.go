/**
 * Description：
 * FileName：files.go
 * Author：CJiaの用心
 * Create：2025/2/20 10:04:12
 * Remark：
 */

package model

import "github.com/carefuly/PackageFluxApp/pkg/models"

type File struct {
	models.CoreModels
	Name    string `gorm:"type:varchar(50);not null;comment:原文件名" json:"name"`
	Size    int    `gorm:"type:int;not null;comment:原文件大小" json:"size"`
	Suffix  string `gorm:"type:varchar(50);comment:原文件后缀" json:"suffix"`
	BaseDir string `gorm:"type:varchar(255);comment:项目根目录" json:"base_dir"`
	FileDir string `gorm:"type:text;comment:文件绝对路径" json:"file_dir"`
	Path    string `gorm:"type:varchar(255);comment:文件所在路径" json:"path"`
	Url     string `gorm:"type:varchar(255);comment:文件地址" json:"url"`
	UserID  int    `gorm:"type:int;comment:用户ID" json:"user_id"`
	User    User   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}
