/**
 * Description：
 * FileName：server.go
 * Author：CJiaの用心
 * Create：2025/2/19 14:10:00
 * Remark：
 */

package config

type ServerConfig struct {
	Host string `yaml:"host" json:"host"`
	Port int    `yaml:"port" json:"port"`
}
