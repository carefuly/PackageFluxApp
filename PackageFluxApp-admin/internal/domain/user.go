/**
 * Description：
 * FileName：user.go
 * Author：CJiaの用心
 * Create：2025/2/26 11:35:00
 * Remark：
 */

package domain

import "github.com/carefuly/PackageFluxApp/internal/model"

type Register struct {
	Username    string `json:"username"`    // 用户名
	Email       string `json:"email"`       // 邮箱
	Password    string `json:"password"`    // 密码
	PasswordStr string `json:"passwordStr"` // 明文密码
}

type Login struct {
	Email    string `json:"email"`    // 邮箱
	Password string `json:"password"` // 密码
}

type User struct {
	model.User
	Email      string `json:"email"`      // 邮箱
	Phone      string `json:"phone"`      // 手机号
	CreateTime string `json:"createTime"` // 创建时间
	UpdateTime string `json:"updateTime"` // 更新时间
}
