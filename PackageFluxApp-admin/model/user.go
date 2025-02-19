/**
 * Description：
 * FileName：user.go
 * Author：CJiaの用心
 * Create：2025/2/18 01:54:00
 * Remark：
 */

package model

import (
	"database/sql"
	"github.com/carefuly/PackageFluxApp/pkg/models"
)

// User 用户表
type User struct {
	models.CoreModels
	Username    string         `gorm:"type:varchar(40);comment:用户名" json:"username"`
	Email       sql.NullString `gorm:"type:varchar(40);unique;comment:邮箱" json:"Email"`
	Password    sql.NullString `gorm:"type:varchar(255);comment:密码" json:"password"`
	PasswordStr sql.NullString `gorm:"type:varchar(255);comment:明文密码" json:"passwordStr"`
	Avatar      string         `gorm:"type:text;comment:头像" json:"avatar"`
	UsageNumber uint           `gorm:"type:int;default:0;comment:接口调用次数" json:"usageNumber"`
	Total       uint           `gorm:"type:int;default:1000;comment:接口总次数" json:"total"`
}
