/**
 * Description：
 * FileName：main.go
 * Author：CJiaの用心
 * Create：2025/2/24 15:16:21
 * Remark：
 */

package main

import (
	"fmt"
	"github.com/carefuly/PackageFluxApp/config"
	"github.com/carefuly/PackageFluxApp/initialize"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

// @title PackageFluxApp-admin
// @version 1.0
// @description PackageFluxApp-admin在线接口文档

// @termsOfService http://swagger.io/terms

// @contact.name CJiaの用心
// @contact.url http://www.swagger.io/support
// @contact.email 2224693191@qq.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization

// @host 127.0.0.1:8080
// @BasePath /api
func main() {
	var relyConfig config.RelyConfig
	relyConfig.Logger = initialize.InitStdoutLogger()
	initConfig := initialize.InitConfig(true)
	relyConfig.Mail = initConfig.EmailConfig
	relyConfig.Db = initialize.NewDbPool().InitDb(initConfig.DatabaseConfig)
	relyConfig.Redis = initialize.InitRedis(initConfig.RedisConfig)
	relyConfig.Token = initConfig.TokenConfig
	relyConfig.AliYun = initConfig.AliYunConfig

	middlewares := initialize.InitGinMiddlewares(relyConfig)
	relyConfig.Trans, _ = initialize.InitGinTrans("zh")
	server := initialize.InitWebServer(middlewares, relyConfig)

	err := server.Run(fmt.Sprintf("%s:%d", initConfig.ServerConfig.Host, initConfig.ServerConfig.Port))
	if err != nil {
		zap.L().Error("启动失败", zap.Error(err))
	}
	// 接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
