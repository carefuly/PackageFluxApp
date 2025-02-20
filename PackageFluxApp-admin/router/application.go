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
	"github.com/carefuly/PackageFluxApp/internal/repository/dao"
	"github.com/carefuly/PackageFluxApp/internal/service"
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

	detailsDAO := dao.NewDetailsDAO(r.rely.Db)
	detailsRepository := repository.NewDetailsRepository(detailsDAO)
	detailsService := service.NewDetailsService(detailsRepository)
	detailsHandler := web.NewDetailsHandler(r.rely, detailsService)
	detailsHandler.RegisterRoutes(applicationRouter)
}
