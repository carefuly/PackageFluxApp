/**
 * Description：
 * FileName：token.go
 * Author：CJiaの用心
 * Create：2025/2/22 21:23:55
 * Remark：
 */

package config

type TokenConfig struct {
	ApiKeyAuth string `yaml:"apiKeyAuth" json:"apiKeyAuth"`
	Expire     int    `yaml:"expire" json:"expire"`
}
