/**
 * Description：
 * FileName：operate.go
 * Author：CJiaの用心
 * Create：2025/2/18 01:53:44
 * Remark：
 */

package model

import (
	"context"
	"github.com/carefuly/PackageFluxApp/pkg/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type OperateLogger struct {
	models.CoreModels
	RequestUsername string `gorm:"type:varchar(40);comment:请求用户名" json:"requestUsername"`
	RequestTime     string `gorm:"type:varchar(40);comment:请求耗时" json:"requestTime"`
	RequestStatus   int    `gorm:"type:int;comment:响应状态码" json:"requestStatus"`
	RequestMethod   string `gorm:"type:varchar(8);comment:请求方式" json:"requestMethod"`
	RequestIp       string `gorm:"type:varchar(20);comment:请求IP地址" json:"requestIp"`
	RequestPath     string `gorm:"type:varchar(255);comment:请求地址" json:"requestPath"`
	RequestQuery    string `gorm:"type:text;comment:请求查询参数" json:"requestQuery"`
	RequestBody     any    `gorm:"type:text;comment:请求参数" json:"requestBody"`
	RequestOs       string `gorm:"type:varchar(40);comment:操作系统" json:"requestOs"`
	RequestBrowser  string `gorm:"type:varchar(64);comment:操作浏览器" json:"requestBrowser"`
	UserAgent       string `gorm:"type:varchar(255);comment:用户代理" json:"userAgent"`
	RequestCode     int    `gorm:"type:int;comment:自定义状态码" json:"requestCode"`
	RequestResult   string `gorm:"type:text;comment:响应信息" json:"requestResult"`
	Errors          string `gorm:"type:text;comment:错误信息" json:"errors"`
	Internal        string `gorm:"type:text;comment:系统错误" json:"internal"`
}

func (o *OperateLogger) Insert(ctx context.Context, db *gorm.DB, op OperateLogger) {
	currentLogger := db.Config.Logger
	// 临时禁用日志
	db.Config.Logger = logger.Default.LogMode(logger.Silent)

	err := db.WithContext(ctx).Create(&op).Error
	if err != nil {
		zap.L().Error("日志记录异常", zap.String("err", err.Error()))
	}

	// 恢复日志级别
	db.Config.Logger = currentLogger
}
