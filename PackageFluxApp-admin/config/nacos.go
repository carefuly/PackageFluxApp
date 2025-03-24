/**
 * Description：
 * FileName：nacos.go
 * Author：CJiaの用心
 * Create：2025/2/19 14:09:36
 * Remark：
 */

package config

type EnvironmentConfig struct {
	Live string `mapstructure:"live"`
}

type NaCosConfig struct {
	Host      string `mapstructure:"host"`
	Port      uint64 `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataId"`
	Group     string `mapstructure:"group"`
}

type Nacos struct {
	EnvironmentConfig `mapstructure:"environment"`
	ServerConfig      ServerConfig `mapstructure:"server"`
	NaCosConfig       `mapstructure:"nacos"`
}
