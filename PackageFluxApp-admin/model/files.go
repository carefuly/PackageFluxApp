/**
 * Description：
 * FileName：file.go
 * Author：CJiaの用心
 * Create：2025/2/20 10:04:12
 * Remark：
 */

package model

import (
	"github.com/carefuly/PackageFluxApp/internal/model"
	"github.com/carefuly/PackageFluxApp/pkg/models"
)

type File struct {
	models.CoreModels
	Name    string      `gorm:"type:varchar(50);not null;column:name;comment:原文件名" json:"name"`       // 原文件名
	Size    string      `gorm:"type:float;not null;column:size;comment:原文件大小" json:"size"`           // 原文件大小
	Type    string      `gorm:"type:varchar(50);not null;column:type;comment:原文件类型" json:"type"`     // 原文件大小
	Suffix  string      `gorm:"type:varchar(50);not null;column:suffix;comment:原文件后缀" json:"suffix"` // 原文件后缀
	BaseDir string      `gorm:"type:varchar(255);column:baseDir;comment:项目根目录" json:"baseDir"`       // 项目根目录
	FileDir string      `gorm:"type:text;column:fileDir;comment:文件绝对路径" json:"fileDir"`             // 文件绝对路径
	Path    string      `gorm:"type:varchar(255);column:path;comment:文件所在路径" json:"path"`           // 文件所在路径
	Url     string      `gorm:"type:varchar(255);column:url;comment:文件地址" json:"url"`                 // 文件地址
	UserId  string      `gorm:"type:varchar(100);column:user_id;comment:用户ID" json:"user_id"`           // 用户ID
	User    *model.User `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`          // 用户
}
