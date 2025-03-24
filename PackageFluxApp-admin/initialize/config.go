/**
 * Description：
 * FileName：config.go
 * Author：CJiaの用心
 * Create：2025/2/19 14:15:43
 * Remark：
 */

package initialize

import (
	"fmt"
	"github.com/carefuly/PackageFluxApp/config"
	"github.com/fsnotify/fsnotify"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

func InitConfig(debug bool) *config.Config {
	// 从配置文件中读取配置信息
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("%s-pro.yaml", configFilePrefix)
	if debug {
		configFileName = fmt.Sprintf("%s-dev.yaml", configFilePrefix)
	}
	conf := new(config.Nacos)
	// 配置
	v := viper.New()
	// 文件的路径如何设置
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		zap.L().Error("配置文件读取失败", zap.Error(err))
	}
	// 全局变量
	if err := v.Unmarshal(conf); err != nil {
		panic(err)
	}

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("[监听配置文件修改]")
		zap.L().Debug("配置文件发生变动")
		if err := v.Unmarshal(conf); err != nil {
			zap.L().Error("配置文件发生变动，解析失败", zap.Error(err))
		}
		zap.L().Debug("配置文件发生变动", zap.Any("conf", conf))
	})

	// 从NaCos中读取配置信息
	serverConfig := []constant.ServerConfig{
		{
			IpAddr: conf.Host,
			Port:   conf.Port,
		},
	}
	clientConfig := constant.ClientConfig{
		NamespaceId:         conf.Namespace, // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
		Username:            conf.User,
		Password:            conf.Password,
	}
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfig,
		"clientConfig":  clientConfig,
	})

	if err != nil {
		zap.L().Error("创建NaCos配置失败", zap.Error(err))
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: conf.DataId,
		Group:  conf.Group,
	})
	if err != nil {
		zap.L().Debug("读取NaCos配置失败", zap.Error(err))
	}

	var globalConfig = new(config.Config)
	// 将配置信息写入到全局变量中
	err = yaml.Unmarshal([]byte(content), &globalConfig)
	if err != nil {
		zap.L().Error("解析配置文件失败", zap.Error(err))
	}

	globalConfig.ServerConfig.Host = conf.ServerConfig.Host
	globalConfig.ServerConfig.Port = conf.ServerConfig.Port

	fmt.Println("globalConfig", globalConfig)

	// 将配置信息写入到全局变量中
	return globalConfig
}
