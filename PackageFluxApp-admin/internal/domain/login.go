/**
 * Description：
 * FileName：login.go
 * Author：CJiaの用心
 * Create：2025/2/19 20:44:34
 * Remark：
 */

package domain

type Login struct {
	RecordId string `json:"recordId"` // 记录id
	Email    string `json:"email"`    // 邮箱
	Password string `json:"password"` // 密码
}
