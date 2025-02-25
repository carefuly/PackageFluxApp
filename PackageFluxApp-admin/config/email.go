/**
 * Description：
 * FileName：email.go
 * Author：CJiaの用心
 * Create：2025/2/19 14:11:46
 * Remark：
 */

package config

type EmailConfig struct {
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
	From     string `yaml:"from" json:"from"`
	To       string `yaml:"to" json:"to"`
}
