/**
 * Description：
 * FileName：register.go
 * Author：CJiaの用心
 * Create：2025/2/19 16:37:16
 * Remark：
 */

package domain

type Register struct {
	Username    string `json:"username"`    // 用户名
	Email       string `json:"email"`       // 邮箱
	Password    string `json:"password"`    // 密码
	PasswordStr string `json:"passwordStr"` // 明文密码
}
