/**
 * Description：
 * FileName：detail.go
 * Author：CJiaの用心
 * Create：2025/2/27 11:29:43
 * Remark：
 */

package web

import (
	"errors"
	"github.com/carefuly/PackageFluxApp/config"
	"github.com/carefuly/PackageFluxApp/internal/domain"
	"github.com/carefuly/PackageFluxApp/internal/model"
	"github.com/carefuly/PackageFluxApp/internal/service"
	"github.com/carefuly/PackageFluxApp/pkg/models"
	"github.com/carefuly/PackageFluxApp/pkg/response"
	validate "github.com/carefuly/PackageFluxApp/pkg/validator"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type DetailHandler interface {
	RegisterRoutes(router *gin.RouterGroup)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
	GetById(ctx *gin.Context)
	GetListAll(ctx *gin.Context)
}

type detailHandler struct {
	rely config.RelyConfig
	svc  service.DetailService
}

func NewDetailHandler(rely config.RelyConfig, svc service.DetailService) DetailHandler {
	return &detailHandler{
		rely: rely,
		svc:  svc,
	}
}

type DetailRequest struct {
	LogoUrl     string   `json:"logoUrl"`                           // logo地址
	AppName     string   `json:"appName" binding:"required,max=50"` // 应用名称
	AppleId     string   `json:"appleId" binding:"max=50"`          // appleId
	Description string   `json:"description" binding:"max=255"`     // 应用描述
	Preview     []string `json:"preview"`                           // 预览图
	Remark      string   `json:"remark" binding:"max=255"`          // 备注
}

func (h *detailHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/detail/create", h.Create)
	router.DELETE("/detail/delete/:id", h.Delete)
	router.PUT("/detail/update/:id", h.Update)
	router.GET("/detail/getById/:id", h.GetById)
	router.GET("/detail/listAll", h.GetListAll)
}

func (h *detailHandler) Create(ctx *gin.Context) {
	uid, ok := ctx.MustGet("userId").(string)
	if !ok {
		ctx.Set("internal", uid)
		zap.S().Error("用户ID获取失败", uid)
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}

	var req DetailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validate.NewValidatorError(h.rely.Trans).HandleValidatorError(ctx, err)
		return
	}

	err := h.svc.Create(ctx, domain.Detail{
		Detail: model.Detail{
			CoreModels: models.CoreModels{
				Creator:  uid,
				Modifier: uid,
				Remark:   req.Remark,
			},
			LogoUrl:     req.LogoUrl,
			AppName:     req.AppName,
			AppleId:     req.AppleId,
			Description: req.Description,
			UserId:      uid,
		},
		Preview: req.Preview,
	})

	switch {
	case err == nil:
		response.NewResponse().SuccessResponse(ctx, "新增成功", nil)
	case errors.Is(err, service.ErrDuplicateUserIdAndAppName):
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "同一用户下应用名称不能重复，请重新输入", nil)
	default:
		ctx.Set("internal", err.Error())
		zap.L().Error("新增应用详情异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
	}
}

func (h *detailHandler) Delete(ctx *gin.Context) {
	uid, ok := ctx.MustGet("userId").(string)
	if !ok {
		ctx.Set("internal", uid)
		zap.S().Error("用户ID获取失败", uid)
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}

	id := ctx.Param("id")
	if id == "" || len(id) == 0 {
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "参数错误", nil)
		return
	}

	err := h.svc.Delete(ctx, id, uid)

	switch {
	case err == nil:
		response.NewResponse().SuccessResponse(ctx, "删除成功", nil)
	case errors.Is(err, service.ErrDetailNotFound):
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "记录不存在", nil)
	default:
		ctx.Set("internal", err.Error())
		zap.L().Error("删除应用详情异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
	}
}

func (h *detailHandler) Update(ctx *gin.Context) {
	uid, ok := ctx.MustGet("userId").(string)
	if !ok {
		ctx.Set("internal", uid)
		zap.S().Error("用户ID获取失败", uid)
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}

	id := ctx.Param("id")
	if id == "" || len(id) == 0 {
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "参数错误", nil)
		return
	}

	var req DetailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validate.NewValidatorError(h.rely.Trans).HandleValidatorError(ctx, err)
		return
	}

	err := h.svc.Update(ctx, id, uid, domain.Detail{
		Detail: model.Detail{
			CoreModels: models.CoreModels{
				Modifier: uid,
				Remark:   req.Remark,
			},
			LogoUrl:     req.LogoUrl,
			AppName:     req.AppName,
			AppleId:     req.AppleId,
			Description: req.Description,
			UserId:      uid,
		},
		Preview: req.Preview,
	})

	switch {
	case err == nil:
		response.NewResponse().SuccessResponse(ctx, "更新成功", nil)
	case errors.Is(err, service.ErrDetailNotFound):
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "记录不存在", nil)
	default:
		ctx.Set("internal", err.Error())
		zap.L().Error("更新应用详情异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
	}
}

func (h *detailHandler) GetById(ctx *gin.Context) {
	uid, ok := ctx.MustGet("userId").(string)
	if !ok {
		ctx.Set("internal", uid)
		zap.S().Error("用户ID获取失败", uid)
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}

	id := ctx.Param("id")
	if id == "" || len(id) == 0 {
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "参数错误", nil)
		return
	}

	detail, err := h.svc.FindById(ctx, id, uid)

	switch {
	case err == nil:
		response.NewResponse().SuccessResponse(ctx, "查询成功", detail)
	case errors.Is(err, service.ErrDetailRecordNotFound):
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "记录不存在", nil)
	default:
		ctx.Set("internal", uid)
		zap.S().Error("根据Id查询应用详情异常", uid)
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
	}
}

func (h *detailHandler) GetListAll(ctx *gin.Context) {
	uid, ok := ctx.MustGet("userId").(string)
	if !ok {
		ctx.Set("internal", uid)
		zap.S().Error("用户ID获取失败", uid)
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}

	appName := ctx.DefaultQuery("appName", "")

	list, err := h.svc.FindListAll(ctx, uid, appName)

	if err != nil {
		ctx.Set("internal", err.Error())
		zap.L().Error("查询应用详情列表异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}

	response.NewResponse().SuccessResponse(ctx, "查询成功", list)
}
