/**
 * Description：
 * FileName：user.go
 * Author：CJiaの用心
 * Create：2025/2/26 11:38:14
 * Remark：
 */

package model

import (
	"github.com/carefuly/PackageFluxApp/pkg/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// User 用户表
type User struct {
	models.CoreModels
	Username    string `gorm:"type:varchar(40);column:username;comment:用户名" json:"username"`             // 用户名
	Email       string `gorm:"type:varchar(50);unique;column:email;comment:邮箱" json:"email"`             // 邮箱
	Password    string `gorm:"type:varchar(255);column:password;comment:密码" json:"-"`                    // 密码
	PasswordStr string `gorm:"type:varchar(255);column:password_str;comment:明文密码" json:"-"`              // 明文密码
	Phone       string `gorm:"type:varchar(20);unique;column:phone;comment:手机号" json:"phone"`            // 手机号
	Avatar      string `gorm:"type:text;column:avatar;comment:头像" json:"avatar"`                         // 头像
	UsageNumber int    `gorm:"type:int;default:0;column:usage_number;comment:接口调用次数" json:"usageNumber"` // 接口调用次数
	Total       int    `gorm:"type:int;default:1000;column:Total;comment:接口总次数" json:"total"`            // 接口总次数
}

func NewUser() *User {
	return &User{}
}

func (u *User) AutoMigrate(db *gorm.DB) {
	err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
	if err != nil {
		zap.L().Error("User表模型迁移失败", zap.Error(err))
	}
}
