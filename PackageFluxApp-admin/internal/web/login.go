/**
 * Description：
 * FileName：login.go
 * Author：CJiaの用心
 * Create：2025/2/24 15:19:41
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
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"net/http"
	"time"
)

var (
	bizLogin = "login"
	JWTKey   = []byte("k6CswdUm77WKcbM68UQUuxVsHSpTCwgK")
)

type LoginHandler interface {
	RegisterRoutes(router *gin.RouterGroup)
	SendEmailLoginCaptchaHandler(ctx *gin.Context)
	EmailLoginHandler(ctx *gin.Context)
	UserInfoHandler(ctx *gin.Context)
}

type loginHandler struct {
	rely    config.RelyConfig
	svc     service.UserService
	codeSvc service.CodeService
}

func NewLoginHandler(rely config.RelyConfig, svc service.UserService,
	codeSvc service.CodeService) LoginHandler {
	return &loginHandler{
		rely:    rely,
		svc:     svc,
		codeSvc: codeSvc,
	}
}

type SendLoginRequest struct {
	Email string `json:"email" binding:"required,email"` // 邮箱
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`           // 邮箱
	Password string `json:"password" binding:"required,min=6,max=16"` // 密码
	Code     string `json:"code" binding:"required,min=6,max=6"`      // 验证码
}

type UserClaims struct {
	jwt.RegisteredClaims
	Uid       string
	Email     string
	UserAgent string
}

func (h *loginHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/send-login-captcha", h.SendEmailLoginCaptchaHandler)
	router.POST("/email-login", h.EmailLoginHandler)
	router.GET("/userinfo", h.UserInfoHandler)
}

// SendEmailLoginCaptchaHandler
// @id SendEmailLoginCaptcha
// @Summary 发送邮箱登录验证码
// @Description 发送邮箱登录验证码
// @Tags 用户登录
// @Accept application/json
// @Produce application/json
// @Param SendLoginRequest body SendLoginRequest true "验证码参数"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /v1/auth/send-login-captcha  [post]
func (h *loginHandler) SendEmailLoginCaptchaHandler(ctx *gin.Context) {
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

// EmailLoginHandler
// @id EmailLogin
// @Summary 邮箱验证码登录
// @Description 邮箱验证码登录
// @Tags 用户登录
// @Accept application/json
// @Produce application/json
// @Param LoginRequest body LoginRequest true "登录参数"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /v1/auth/email-login [post]
func (h *loginHandler) EmailLoginHandler(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validate.NewValidatorError(h.rely.Trans).HandleValidatorError(ctx, err)
		return
	}

	ok, err := h.codeSvc.Verify(ctx, bizLogin, req.Email, req.Code)

	switch {
	case ok:
		// 验证成功，继续注册流程
		u, err := h.svc.Login(ctx, domain.Login{
			Email:    req.Email,
			Password: req.Password,
		})

		switch {
		case err == nil:
			token := h.setJWTToken(ctx, u)
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
	default: // 其他错误
		ctx.Set("internal", err.Error())
		zap.L().Error("验证验证码异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
	}
}

// UserInfoHandler
// @id UserInfo
// @Summary 获取用户信息
// @Description 获取用户信息
// @Tags 用户登录
// @Accept application/json
// @Produce application/json
// @Success 200 {object} domain.UserInfo
// @Failure 400 {object} response.Response
// @Router /v1/auth/userinfo [get]
// @Security ApiKeyAuth
func (h *loginHandler) UserInfoHandler(ctx *gin.Context) {
	uid, ok := ctx.MustGet("userId").(string)
	if !ok {
		ctx.Set("internal", uid)
		zap.L().Error("用户ID获取失败", zap.Error(errors.New(uid)))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}

	u, err := h.svc.UserInfo(ctx, uid)
	if err != nil {
		ctx.Set("internal", err.Error())
		zap.L().Error("获取用户失败", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}
	response.NewResponse().SuccessResponse(ctx, "获取成功", u)
}

func (h *loginHandler) setJWTToken(ctx *gin.Context, u domain.Login) string {
	uc := UserClaims{
		Uid:       u.ID,
		Email:     u.Email,
		UserAgent: ctx.GetHeader("User-Agent"),
		RegisteredClaims: jwt.RegisteredClaims{
			// 一小时过期
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 60 * 4)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, uc)
	tokenStr, err := token.SignedString(JWTKey)
	if err != nil {
		ctx.Set("internal", err.Error())
		zap.L().Error("生成Token异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
	}
	return tokenStr
}
