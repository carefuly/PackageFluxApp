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

var (
	bizRegister = "register"
)

type RegisterHandler interface {
	RegisterRoutes(router *gin.RouterGroup)
	// SendEmailRegisterCaptchaHandler(ctx *gin.Context)
	// EmailRegisterHandler(ctx *gin.Context)
}

type registerHandler struct {
	rely    config.RelyConfig
	svc     service.UserService
	codeSvc service.CodeService
}

func NewRegisterHandler(rely config.RelyConfig, svc service.UserService, codeSvc service.CodeService) RegisterHandler {
	return &registerHandler{
		rely:    rely,
		svc:     svc,
		codeSvc: codeSvc,
	}
}

func (h *registerHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/send-register-captcha", h.SendEmailCaptchaRegisterHandler)
	router.POST("/email-register", h.EmailCaptchaRegisterHandler)
	router.POST("/password-register", h.PassWordRegisterHandler)
}

func (h *registerHandler) SendEmailCaptchaRegisterHandler(ctx *gin.Context) {
	type SendRegisterRequest struct {
		Email string `json:"email" binding:"required,email"` // 邮箱
	}

	var req SendRegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validate.NewValidatorError(h.rely.Trans).HandleValidatorError(ctx, err)
		return
	}

	code, err := h.codeSvc.Send(ctx, req.Email, "Package注册", bizRegister)
	// 不管成功还是失败, 都要返回验证码
	zap.L().Info("验证码发送成功", zap.String("email", req.Email), zap.String("code", code))

	switch {
	case err == nil:
		response.NewResponse().SuccessResponse(ctx, "发送成功, 验证码10分钟有效", nil)
	case errors.Is(err, service.ErrCodeSendTooMany):
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "邮箱发送太频繁，请稍后再试", nil)
	default:
		ctx.Set("internal", err.Error())
		zap.L().Error("短信发送异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
	}
}

func (h *registerHandler) EmailCaptchaRegisterHandler(ctx *gin.Context) {
	type RegisterRequest struct {
		Username string `json:"username" binding:"required,min=3,max=20"` // 用户名
		Email    string `json:"email" binding:"required,email"`           // 邮箱
		Password string `json:"password" binding:"required,min=6,max=50"` // 密码
		Code     string `json:"code" binding:"required,min=6,max=6"`      // 验证码
	}

	var req RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validate.NewValidatorError(h.rely.Trans).HandleValidatorError(ctx, err)
		return
	}

	ok, err := h.codeSvc.Verify(ctx, bizRegister, req.Email, req.Code)

	switch {
	case ok:
		// 验证成功，继续注册流程
		err = h.svc.Register(ctx, domain.Register{
			Username: req.Username,
			Email:    req.Email,
			Password: req.Password,
			PasswordStr: req.Password,
		})
		switch {
		case err == nil:
			response.NewResponse().SuccessResponse(ctx, "注册成功", nil)
		case errors.Is(err, service.ErrDuplicateEmail):
			response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "邮箱已存在，请重新输入", nil)
		default:
			ctx.Set("internal", err.Error())
			zap.L().Error("邮箱验证码注册异常", zap.Error(err))
			response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		}
	case errors.Is(err, service.ErrUserBlocked):
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "操作过于频繁，请10分钟后再试", nil)
	case errors.Is(err, service.ErrCodeNotFound):
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "请先发送注册验证码", nil)
	case errors.Is(err, service.ErrCodeVerifyTooMany):
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "验证次数已耗尽，请10分钟后再试", nil)
	case errors.Is(err, service.ErrCodeIncorrect):
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "验证码错误，请重新输入", nil)
	default: // 其他错误
		ctx.Set("internal", err.Error())
		zap.L().Error("验证验证码异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
	}
}

func (h *registerHandler) PassWordRegisterHandler(ctx *gin.Context) {
	type PassWordRegisterRequest struct {
		Username string `json:"username" binding:"required,min=3,max=20"` // 用户名
		Email    string `json:"email" binding:"required,email"`           // 邮箱
		Password string `json:"password" binding:"required,min=6,max=50"` // 密码
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
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "邮箱已存在，请重新输入", nil)
	default:
		ctx.Set("internal", err.Error())
		zap.L().Error("密码注册异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
	}
}
