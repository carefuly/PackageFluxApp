/**
 * Description：
 * FileName：database.go
 * Author：CJiaの用心
 * Create：2025/2/19 14:11:25
 * Remark：
 */

package config

type DatabaseConfig struct {
	Type     string `yaml:"type" json:"type"`
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
	Database string `yaml:"dbname" json:"dbname"`
	Charset  string `yaml:"charset" json:"charset"`
	Prefix   string `yaml:"prefix" json:"prefix"`
}
