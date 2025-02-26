/**
 * Description：
 * FileName：login.go
 * Author：CJiaの用心
 * Create：2025/2/26 21:46:38
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
	bizLogin = "login"
)

type LoginHandler interface {
	RegisterRoutes(router *gin.RouterGroup)
	SendEmailLoginCaptchaHandler(ctx *gin.Context)
	PassWordLoginHandler(ctx *gin.Context)
	UserInfoHandler(ctx *gin.Context)
}

type loginHandler struct {
	rely    config.RelyConfig
	svc     service.UserService
	codeSvc service.CodeService
}

func NewLoginHandler(rely config.RelyConfig, svc service.UserService, codeSvc service.CodeService) RegisterHandler {
	return &loginHandler{
		rely:    rely,
		svc:     svc,
		codeSvc: codeSvc,
	}
}

func (h *loginHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/send-login-captcha", h.SendEmailCaptchaLoginHandler)
	router.POST("/email-login", h.EmailCaptchaLoginHandler)
	router.POST("/password-login", h.PassWordLoginHandler)
	router.GET("/userinfo", h.UserInfoHandler)
}

func (h *loginHandler) SendEmailCaptchaLoginHandler(ctx *gin.Context) {
	type SendLoginRequest struct {
		Email string `json:"email" binding:"required,email"` // 邮箱
	}

	var req SendLoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validate.NewValidatorError(h.rely.Trans).HandleValidatorError(ctx, err)
		return
	}

	code, err := h.codeSvc.Send(ctx, req.Email, "Package登录", bizLogin)
	zap.L().Info("验证码发送成功", zap.String("email", req.Email), zap.String("code", code))

	switch {
	case err == nil:
		response.NewResponse().SuccessResponse(ctx, "发送成功, 验证码10分钟有效", nil)
	case errors.Is(err, service.ErrCodeSendTooMany):
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "短信发送太频繁，请稍后再试", nil)
	default:
		ctx.Set("internal", err.Error())
		zap.L().Error("短信发送异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
	}
}

func (h *loginHandler) EmailCaptchaLoginHandler(ctx *gin.Context) {
	type LoginRequest struct {
		Email    string `json:"email" binding:"required,email"`           // 邮箱
		Password string `json:"password" binding:"required,min=6,max=50"` // 密码
		Code     string `json:"code" binding:"required,min=6,max=6"`      // 验证码
	}

	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validate.NewValidatorError(h.rely.Trans).HandleValidatorError(ctx, err)
		return
	}

	ok, err := h.codeSvc.Verify(ctx, bizLogin, req.Email, req.Code)

	switch {
	case ok:
		token, err := h.svc.Login(ctx, h.rely, domain.Login{
			Email:    req.Email,
			Password: req.Password,
		})

		switch {
		case err == nil:
			response.NewResponse().SuccessResponse(ctx, "登录成功", gin.H{
				"token": "Bearer " + token,
			})
		case errors.Is(err, service.ErrInvalidUserOrPassword):
			response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "用户邮箱或者密码不对", nil)
		default:
			ctx.Set("internal", err.Error())
			zap.L().Error("邮箱登录异常", zap.Error(err))
			response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		}
	case errors.Is(err, service.ErrUserBlocked):
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "操作过于频繁，请10分钟后再试", nil)
	case errors.Is(err, service.ErrCodeNotFound):
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "请先发送登录验证码", nil)
	case errors.Is(err, service.ErrCodeVerifyTooMany):
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "验证次数已耗尽，请10分钟后再试", nil)
	case errors.Is(err, service.ErrCodeIncorrect):
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "验证码错误，请重新输入", nil)
	default:
		ctx.Set("internal", err.Error())
		zap.L().Error("验证验证码异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
	}
}

func (h *loginHandler) PassWordLoginHandler(ctx *gin.Context) {
	type LoginRequest struct {
		Email    string `json:"email" binding:"required,email"`           // 邮箱
		Password string `json:"password" binding:"required,min=6,max=50"` // 密码
	}

	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validate.NewValidatorError(h.rely.Trans).HandleValidatorError(ctx, err)
		return
	}

	token, err := h.svc.Login(ctx, h.rely, domain.Login{
		Email:    req.Email,
		Password: req.Password,
	})

	switch {
	case err == nil:
		response.NewResponse().SuccessResponse(ctx, "登录成功", gin.H{
			"token": "Bearer " + token,
		})
	case errors.Is(err, service.ErrInvalidUserOrPassword):
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "邮箱或密码错误", nil)
	case errors.Is(err, service.ErrGenerateTokenError):
		ctx.Set("internal", err.Error())
		zap.L().Error("生成Token异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
	default:
		ctx.Set("internal", err.Error())
		zap.L().Error("密码登录异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
	}
}

func (h *loginHandler) UserInfoHandler(ctx *gin.Context) {
	id, ok := ctx.MustGet("userId").(string)
	if !ok {
		ctx.Set("internal", id)
		zap.L().Error("用户ID获取失败", zap.Error(errors.New(id)))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}

	user, err := h.svc.UserInfo(ctx, id)

	switch {
	case err == nil:
		response.NewResponse().SuccessResponse(ctx, "获取成功", user)
	case errors.Is(err, service.ErrUserIdNotFound):
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "用户信息不存在", nil)
	default:
		ctx.Set("internal", err.Error())
		zap.L().Error("获取用户失败", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
	}
}
