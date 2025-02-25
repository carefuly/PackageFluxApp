/**
 * Description：
 * FileName：dao.go
 * Author：CJiaの用心
 * Create：2025/2/17 00:06:56
 * Remark：
 */

package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

// CoreModels 公共模型
// 核心标准抽象模型模型,可直接继承使用
// 增加审计字段, 覆盖字段时, 字段名称请勿修改, 必须统一审计字段名称
type CoreModels struct {
	ID string `gorm:"type:varchar(100);primary_key;comment:主键ID" json:"id"` // 主键ID
	// RecordId   string     `gorm:"type:varchar(100);uniqueIndex;comment:记录ID" json:"recordId"` // 记录ID
	Sort       int        `gorm:"type:int;default:1;comment:显示排序" json:"sort"`     // 显示排序
	Creator    string     `gorm:"type:varchar(100);comment:创建人" json:"creator"`    // 创建人
	Modifier   string     `gorm:"type:varchar(100);comment:修改人" json:"modifier"`   // 修改人
	BelongDept int        `gorm:"type:int;comment:数据归属部门" json:"belongDept"`       // 数据归属部门
	Status     bool       `gorm:"type:bool;default:true;comment:状态" json:"status"` // 状态
	CreateTime *time.Time `gorm:"autoCreateTime;comment:创建时间" json:"createTime"`   // 创建时间
	UpdateTime *time.Time `gorm:"autoUpdateTime;comment:修改时间" json:"updateTime"`   // 修改时间
	Remark     string     `gorm:"type:varchar(255);comment:备注" json:"remark"`      // 备注
}

func (c *CoreModels) BeforeCreate(tx *gorm.DB) (err error) {
	// 记录ID
	c.ID = uuid.NewV4().String()
	// 返回异常
	return
}
