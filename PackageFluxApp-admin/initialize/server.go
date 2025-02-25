/**
 * Description：
 * FileName：server.go
 * Author：CJiaの用心
 * Create：2025/2/19 16:03:40
 * Remark：
 */

package initialize

import (
	"fmt"
	"github.com/carefuly/PackageFluxApp/config"
	"github.com/carefuly/PackageFluxApp/docs"
	middleware2 "github.com/carefuly/PackageFluxApp/middleware"
	"github.com/carefuly/PackageFluxApp/router"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"reflect"
	"strings"
)

func InitGinMiddlewares(rely config.RelyConfig) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware2.Cors(),
		middleware2.NewLoginJWTMiddlewareBuilder().
			IgnorePaths("/api/v1/auth/send-register-captcha").
			IgnorePaths("/api/v1/auth/email-register").
			IgnorePaths("/api/v1/auth/send-login-captcha").
			IgnorePaths("/api/v1/auth/email-login").
			Build(),
		middleware2.NewLogger(rely.Logger).Logger(),
		middleware2.NewStorage().StorageLogger(rely.Db),
	}
}

func InitGinTrans(locale string) (ut.Translator, error) {
	var trans ut.Translator
	// 修改gin框架中的validator引擎属性, 实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个获取json的tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器
		// 第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
		uni := ut.New(enT, zhT, enT)
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return trans, fmt.Errorf("uni.GetTranslator(%s)", locale)
		}

		switch locale {
		case "en":
			err := entranslations.RegisterDefaultTranslations(v, trans)
			if err != nil {
				return nil, err
			}
		case "zh":
			err := zhtranslations.RegisterDefaultTranslations(v, trans)
			if err != nil {
				return nil, err
			}
		default:
			err := entranslations.RegisterDefaultTranslations(v, trans)
			if err != nil {
				return nil, err
			}
		}
		return trans, nil
	}
	return trans, nil
}

func InitWebServer(middle []gin.HandlerFunc, rely config.RelyConfig) *gin.Engine {
	server := gin.Default()
	server.Use(middle...)

	docs.SwaggerInfo.BasePath = "/api"
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	ApiGroup := server.Group("/api")
	v1 := ApiGroup.Group("/v1")

	router.NewAuthRouter(rely).RegisterAuthRouter(v1)
	router.NewApplicationRouter(rely).RegisterAuthRouter(v1)

	return server
}
