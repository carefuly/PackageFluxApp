/**
 * Description：
 * FileName：config.go
 * Author：CJiaの用心
 * Create：2025/2/19 14:11:57
 * Remark：
 */

package config

type Config struct {
	ServerConfig   `yaml:"server" json:"server"`
	DatabaseConfig `yaml:"database" json:"database"`
	RedisConfig    `yaml:"redis" json:"redis"`
	EmailConfig    `yaml:"email" json:"email"`
	TokenConfig    `yaml:"token" json:"token"`
	AliYunConfig   `yaml:"aliYun" json:"aliYun"`
}
