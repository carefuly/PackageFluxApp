/**
 * Description：
 * FileName：filters.go
 * Author：CJiaの用心
 * Create：2025/2/21 11:19:37
 * Remark：
 */

package filters

type Filters struct {
	Creator    uint `gorm:"type:int;comment:创建人" json:"creator"`             // 创建人
	Modifier   uint `gorm:"type:int;comment:修改人" json:"modifier"`            // 修改人
	BelongDept uint `gorm:"type:int;comment:数据归属部门" json:"belongDept"`       // 数据归属部门
	Status     bool `gorm:"type:bool;default:true;comment:状态" json:"status"` // 状态
}
