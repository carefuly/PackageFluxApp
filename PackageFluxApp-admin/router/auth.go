/**
 * Description：
 * FileName：auth.go
 * Author：CJiaの用心
 * Create：2025/2/19 16:39:53
 * Remark：
 */

package router

import (
	"github.com/carefuly/PackageFluxApp/config"
	"github.com/carefuly/PackageFluxApp/internal/repository"
	"github.com/carefuly/PackageFluxApp/internal/repository/cache"
	"github.com/carefuly/PackageFluxApp/internal/repository/dao"
	"github.com/carefuly/PackageFluxApp/internal/service"
	"github.com/carefuly/PackageFluxApp/internal/service/mail"
	"github.com/carefuly/PackageFluxApp/internal/web"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	rely config.RelyConfig
}

func NewAuthRouter(rely config.RelyConfig) *AuthRouter {
	return &AuthRouter{
		rely: rely,
	}
}

func (r *AuthRouter) RegisterAuthRouter(router *gin.RouterGroup) {
	authRouter := router.Group("/auth")

	codeCache := cache.NewCodeCache(r.rely.Redis)
	codeRepository := repository.NewCodeRepository(codeCache)
	qqMailSender := mail.NewQQMailSender(r.rely.Mail.Username, r.rely.Mail.Password)
	codeService := service.NewCodeService(codeRepository, qqMailSender)

	registerDAO := dao.NewRegisterDAO(r.rely.Db)
	registerRepository := repository.NewRegisterRepository(registerDAO)
	registerService := service.NewRegisterService(registerRepository)
	registerHandler := web.NewRegisterHandler(r.rely, registerService, codeService)
	registerHandler.RegisterRoutes(authRouter)

	loginDao := dao.NewLoginDAO(r.rely.Db)
	loginRepository := repository.NewLoginRepository(loginDao)
	loginService := service.NewLoginService(loginRepository)
	loginHandler := web.NewLoginHandler(r.rely, loginService, codeService)
	loginHandler.RegisterRoutes(authRouter)
}
