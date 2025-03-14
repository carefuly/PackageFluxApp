/**
 * Description：
 * FileName：check.go
 * Author：CJiaの用心
 * Create：2025/3/1 00:21:13
 * Remark：
 */

package domain

type Check struct {
	NeedUpdate     bool   `json:"needUpdate"`     // 是否需要更新
	Version        string `json:"version"`        // 版本号
	ApkUrl         string `json:"pkgUrl"`         // apk下载地址
	WgtUrl         string `json:"wgtUrl"`         // wgt下载地址
	HBuilderUpdate bool   `json:"isHBuilderUpdate"` // 是否需要更新HBuilder
	ForceUpdate    bool   `json:"isForceUpdate"`    // 是否需要更新HBuilder
	Description    string `json:"description"`    // 更新描述
	Tip            string `json:"tip"`            // 更新提示
	Type           int    `json:"type"`           // 更新类型
}
