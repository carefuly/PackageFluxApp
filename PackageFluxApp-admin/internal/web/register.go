/**
 * Description：
 * FileName：register.go
 * Author：CJiaの用心
 * Create：2025/2/26 15:07:24
 * Remark：
 */

package web

import (
	"errors"
	"github.com/carefuly/PackageFluxApp/config"
	"github.com/carefuly/PackageFluxApp/internal/domain"
	"github.com/carefuly/PackageFluxApp/internal/service"
	"github.com/carefuly/PackageFluxApp/pkg/response"
	validate "github.com/carefuly/PackageFluxApp/pkg/validator"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type RegisterHandler interface {
	RegisterRoutes(router *gin.RouterGroup)
	// SendEmailRegisterCaptchaHandler(ctx *gin.Context)
	// EmailRegisterHandler(ctx *gin.Context)
}

type registerHandler struct {
	rely config.RelyConfig
	svc  service.UserService
	// codeSvc service.CodeService
}

func NewRegisterHandler(rely config.RelyConfig, svc service.UserService) RegisterHandler {
	return &registerHandler{
		rely: rely,
		svc:  svc,
		// codeSvc: codeSvc,
	}
}

func (h *registerHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/password-register", h.PassWordRegisterCaptchaHandler)
	// router.POST("/send-register-captcha", h.SendEmailRegisterCaptchaHandler)
	// router.POST("/email-register", h.EmailRegisterHandler)
}

func (h *registerHandler) PassWordRegisterCaptchaHandler(ctx *gin.Context) {
	type PassWordRegisterRequest struct {
		Username string `json:"username" binding:"required,min=3,max=20"` // 用户名
		Email    string `json:"email" binding:"required,email"`           // 邮箱
		Password string `json:"password" binding:"required,min=6,max=16"` // 密码
	}

	var req PassWordRegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validate.NewValidatorError(h.rely.Trans).HandleValidatorError(ctx, err)
		return
	}

	err := h.svc.Register(ctx, domain.Register{
		Username:    req.Username,
		Email:       req.Email,
		Password:    req.Password,
		PasswordStr: req.Password,
	})

	switch {
	case err == nil:
		response.NewResponse().SuccessResponse(ctx, "注册成功", nil)
	case errors.Is(err, service.ErrDuplicateEmail):
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "邮箱已存在,请重新输入", nil)
	default:
		ctx.Set("internal", err.Error())
		zap.L().Error("注册异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
	}
}
