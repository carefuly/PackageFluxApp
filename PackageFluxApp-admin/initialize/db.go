/**
 * Description：
 * FileName：db.go
 * Author：CJiaの用心
 * Create：2025/2/19 15:42:59
 * Remark：
 */

package initialize

import (
	"fmt"
	"github.com/carefuly/PackageFluxApp/config"
	"github.com/carefuly/PackageFluxApp/model"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

type DbPool struct {
	db *gorm.DB
}

func NewDbPool() *DbPool {
	return &DbPool{
		db: new(gorm.DB),
	}
}

func (i DbPool) InitDb(database config.DatabaseConfig) *gorm.DB {
	// 配置日志
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)
	// dsn
	var dsn string
	// 全局模式
	var err error
	// 判断数据库类型
	if database.Type == "mysql" {
		// mysql
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
			database.Username, database.Password, database.Host, database.Port, database.Database, database.Charset)
		i.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: database.Prefix, // 表名前缀
			},
			Logger: newLogger,
		})
	}
	// 连接数据失败
	if err != nil {
		zap.L().Error("数据库连接失败", zap.Error(err))
	}
	// 迁移表
	// i.db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.User{})
	// i.db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Detail{})
	// i.db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Version{})
	// i.db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Logger{})
	i.db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.File{})
	// i.db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.OperateLogger{})

	return i.db
}
