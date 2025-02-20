/**
 * Description：
 * FileName：details.go
 * Author：CJiaの用心
 * Create：2025/2/20 20:48:14
 * Remark：
 */

package domain

import (
	"github.com/carefuly/PackageFluxApp/model"
)

type Detail struct {
	model.Detail
	Preview    []string `json:"preview"`
	CreateTime string   `json:"createTime"`
	UpdateTime string   `json:"updateTime"`
}

// type CreateDetail struct {
// 	LogoUrl     string `json:"logo_url"`
// 	AppName     string `json:"app_name"`
// 	AppleId     string `json:"appleId"`
// 	Description string `json:"description"`
// 	Preview     []string `json:"preview"`
// 	UserID      uint   `json:"user_id"`
// 	Creator     uint   `json:"creator"`
// 	Modifier    uint   `json:"modifier"`
// 	Remark      string `json:"remark"`
// }
//
// type Details struct {
// 	models.CoreModels
// 	RecordId    string `json:"recordId"`
// 	LogoUrl     string `json:"logo_url"`
// 	AppName     string `json:"app_name"`
// 	AppleId     string `json:"appleId"`
// 	Description string `json:"description"`
// 	Type        string `json:"type"`
// 	Preview     string `json:"preview"`
// 	UserID      uint   `json:"user_id"`
// 	VersionID   *uint  `json:"version_id"`
// 	Creator     uint   `json:"creator"`
// 	Modifier    uint   `json:"modifier"`
// 	Remark      string `json:"remark"`
// 	CreateTime string `json:"createTime"`
// 	UpdateTime string `json:"updateTime"`
// }
