/**
 * Description：
 * FileName：redis.go
 * Author：CJiaの用心
 * Create：2025/2/19 14:11:35
 * Remark：
 */

package config

type RedisConfig struct {
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Password string `yaml:"password" json:"password"`
	Db       int    `yaml:"db" json:"db"`
}
