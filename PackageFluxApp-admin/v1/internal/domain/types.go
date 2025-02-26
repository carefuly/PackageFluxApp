/**
 * Description：
 * FileName：types.go
 * Author：CJiaの用心
 * Create：2025/2/24 21:53:20
 * Remark：
 */

package domain

type Pagination struct {
	Page     int `json:"page"`     // 当前页
	PageSize int `json:"pageSize"` // 每页显示的条数
}
