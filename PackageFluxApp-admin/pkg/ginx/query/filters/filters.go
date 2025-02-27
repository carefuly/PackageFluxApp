/**
 * Description：
 * FileName：filters.go
 * Author：CJiaの用心
 * Create：2025/2/21 11:19:37
 * Remark：
 */

package filters

type Filters struct {
	Creator    int  `gorm:"type:int;comment:创建人" json:"creator"`             // 创建人
	Modifier   int  `gorm:"type:int;comment:修改人" json:"modifier"`            // 修改人
	BelongDept int  `gorm:"type:int;comment:数据归属部门" json:"belongDept"`       // 数据归属部门
	Status     bool `gorm:"type:bool;default:true;comment:状态" json:"status"` // 状态
}

type Pagination struct {
	Page     int `json:"page"`     // 当前页
	PageSize int `json:"pageSize"` // 每页显示的条数
}
