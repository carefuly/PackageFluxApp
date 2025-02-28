/**
 * Description：
 * FileName：application.go
 * Author：CJiaの用心
 * Create：2025/2/20 21:12:31
 * Remark：
 */

package router

import (
	"github.com/carefuly/PackageFluxApp/config"
	"github.com/carefuly/PackageFluxApp/internal/repository"
	"github.com/carefuly/PackageFluxApp/internal/repository/cache"
	"github.com/carefuly/PackageFluxApp/internal/repository/dao"
	"github.com/carefuly/PackageFluxApp/internal/service"
	"github.com/carefuly/PackageFluxApp/internal/service/oss"
	"github.com/carefuly/PackageFluxApp/internal/web"
	"github.com/gin-gonic/gin"
)

type ApplicationRouter struct {
	rely config.RelyConfig
}

func NewApplicationRouter(rely config.RelyConfig) *ApplicationRouter {
	return &ApplicationRouter{
		rely: rely,
	}
}

func (r *ApplicationRouter) RegisterAuthRouter(router *gin.RouterGroup) {
	applicationRouter := router.Group("/application")

	aliYunOSS := oss.NewAliYunOSS(r.rely.AliYun.Endpoint, r.rely.AliYun.BucketName,
		r.rely.AliYun.AccessKeyID, r.rely.AliYun.AccessKeyID)
	aliYunService := service.NewAliYunService(aliYunOSS)

	redisDetailCache := cache.NewRedisDetailCache(r.rely.Redis)
	detailDAO := dao.NewGORMDetailDAO(r.rely.Db)
	detailRepository := repository.NewDetailRepository(detailDAO, redisDetailCache)
	detailService := service.NewDetailService(detailRepository)
	detailHandler := web.NewDetailHandler(r.rely, detailService)
	detailHandler.RegisterRoutes(applicationRouter)

	redisVersionCache := cache.NewRedisVersionCache(r.rely.Redis)
	versionDAO := dao.NewGORMVersionDAO(r.rely.Db)
	versionRepository := repository.NewVersionRepository(versionDAO, redisVersionCache)
	versionService := service.NewVersionService(versionRepository, detailRepository)
	versionHandler := web.NewVersionHandler(r.rely, versionService)
	versionHandler.RegisterRoutes(applicationRouter)

	redisFileCache := cache.NewRedisFileCache(r.rely.Redis)
	fileDAO := dao.NewGORMFileDAO(r.rely.Db)
	fileRepository := repository.NewFileRepository(fileDAO, redisFileCache)
	fileService := service.NewFileService(fileRepository)
	fileHandler := web.NewFileHandler(r.rely, fileService, aliYunService)
	fileHandler.RegisterRoutes(applicationRouter)
}
