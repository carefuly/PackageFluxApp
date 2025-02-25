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
	Username    string         `gorm:"type:varchar(40);comment:用户名" json:"username"`         // 用户名
	Email       sql.NullString `gorm:"type:varchar(40);unique;comment:邮箱" json:"-"`          // 邮箱
	Password    sql.NullString `gorm:"type:varchar(255);comment:密码" json:"-"`                // 密码
	PasswordStr sql.NullString `gorm:"type:varchar(255);comment:明文密码" json:"-"`              // 明文密码
	Avatar      string         `gorm:"type:text;comment:头像" json:"avatar"`                   // 头像
	UsageNumber int            `gorm:"type:int;default:0;comment:接口调用次数" json:"usageNumber"` // 接口调用次数
	Total       int            `gorm:"type:int;default:1000;comment:接口总次数" json:"total"`     // 接口总次数
}
