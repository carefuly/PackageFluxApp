/**
 * Description：
 * FileName：user.go
 * Author：CJiaの用心
 * Create：2025/2/24 15:38:50
 * Remark：
 */

package domain

import "github.com/carefuly/PackageFluxApp/model"

type Register struct {
	Username    string `json:"username"`    // 用户名
	Email       string `json:"email"`       // 邮箱
	Password    string `json:"password"`    // 密码
	PasswordStr string `json:"passwordStr"` // 明文密码
}

type Login struct {
	ID       string `json:"id"`       // 主键id
	RecordId string `json:"recordId"` // 记录id
	Email    string `json:"email"`    // 邮箱
	Password string `json:"password"` // 密码
}

type UserInfo struct {
	model.User
	Email      string `json:"email"`      // 邮箱
	CreateTime string `json:"createTime"` // 创建时间
	UpdateTime string `json:"updateTime"` // 更新时间
}
