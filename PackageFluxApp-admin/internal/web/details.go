/**
 * Description：
 * FileName：details.go
 * Author：CJiaの用心
 * Create：2025/2/20 20:39:52
 * Remark：
 */

package web

import (
	"errors"
	"github.com/carefuly/PackageFluxApp/config"
	"github.com/carefuly/PackageFluxApp/internal/domain"
	"github.com/carefuly/PackageFluxApp/internal/service"
	"github.com/carefuly/PackageFluxApp/model"
	"github.com/carefuly/PackageFluxApp/pkg/ginx/query/filters"
	"github.com/carefuly/PackageFluxApp/pkg/models"
	"github.com/carefuly/PackageFluxApp/pkg/response"
	validate "github.com/carefuly/PackageFluxApp/pkg/validator"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type DetailsHandler interface {
	RegisterRoutes(router *gin.RouterGroup)
	Create(ctx *gin.Context)
	// Delete(ctx *gin.Context)
	// FindById(ctx *gin.Context)
}

type detailsHandler struct {
	rely config.RelyConfig
	svc  service.DetailsService
}

func NewDetailsHandler(rely config.RelyConfig, svc service.DetailsService) DetailsHandler {
	return &detailsHandler{
		rely: rely,
		svc:  svc,
	}
}

type Filters struct {
	filters.Filters
	AppName string `json:"appName" binding:"max=50"` // 应用名称
}

type AppDetailsRequest struct {
	LogoUrl     string   `json:"logoUrl"`                           // logo地址
	AppName     string   `json:"appName" binding:"required,max=50"` // 应用名称
	AppleId     string   `json:"appleId" binding:"max=50"`          // appleId
	Description string   `json:"description" binding:"max=255"`     // 应用描述
	Preview     []string `json:"preview"`                           // 预览图
	Remark      string   `json:"remark" binding:"max=255"`          // 备注
}

func (h *detailsHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/details/create", h.Create)
	router.DELETE("/details/delete/:id", h.Delete)
	router.PUT("/details/update/:id", h.Update)
	router.GET("/details/getById/:id", h.FindById)
	router.GET("/details/listAll", h.ListAll)
}

func (h *detailsHandler) Create(ctx *gin.Context) {
	uid, ok := ctx.MustGet("userId").(uint)
	if !ok {
		ctx.Set("internal", uid)
		zap.S().Error("用户ID获取失败", uid)
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}

	var req AppDetailsRequest
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
			UserID:      uid,
		},
		Preview: req.Preview,
	})

	switch {
	case err == nil:
		response.NewResponse().SuccessResponse(ctx, "新增成功", nil)
	case errors.Is(err, service.ErrDuplicateAppName):
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "应用名冲突", nil)
	default:
		ctx.Set("internal", err.Error())
		zap.L().Error("新增应用详情异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}
}

func (h *detailsHandler) Delete(ctx *gin.Context) {
	uid, ok := ctx.MustGet("userId").(uint)
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

	err := h.svc.Delete(ctx, uid, id)

	if err != nil {
		ctx.Set("internal", err.Error())
		zap.L().Error("删除应用详情异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}

	response.NewResponse().SuccessResponse(ctx, "删除成功", nil)
}

func (h *detailsHandler) Update(ctx *gin.Context) {
	uid, ok := ctx.MustGet("userId").(uint)
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

	var req AppDetailsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validate.NewValidatorError(h.rely.Trans).HandleValidatorError(ctx, err)
		return
	}

	err := h.svc.Update(ctx, domain.Detail{
		Detail: model.Detail{
			CoreModels: models.CoreModels{
				RecordId: id,
				Modifier: uid,
				Remark:   req.Remark,
			},
			LogoUrl:     req.LogoUrl,
			AppName:     req.AppName,
			AppleId:     req.AppleId,
			Description: req.Description,
			UserID:      uid,
		},
		Preview: req.Preview,
	})

	if err != nil {
		ctx.Set("internal", err.Error())
		zap.L().Error("更新应用详情异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}

	response.NewResponse().SuccessResponse(ctx, "更新成功", nil)
}

func (h *detailsHandler) FindById(ctx *gin.Context) {
	uid, ok := ctx.MustGet("userId").(uint)
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

	d, err := h.svc.FindById(ctx, uid, id)

	switch {
	case err == nil:
		response.NewResponse().SuccessResponse(ctx, "查询成功", d)
	case errors.Is(err, service.ErrRecordNotFound):
		response.NewResponse().ErrorResponse(ctx, http.StatusNotFound, "应用详情不存在", nil)
	default:
		ctx.Set("internal", err.Error())
		zap.L().Error("根据ID查询应用详情异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}
}

func (h *detailsHandler) ListAll(ctx *gin.Context) {
	uid, ok := ctx.MustGet("userId").(uint)
	if !ok {
		ctx.Set("internal", uid)
		zap.S().Error("用户ID获取失败", uid)
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}

	statusStr := ctx.DefaultQuery("status", "true")
	status, _ := strconv.ParseBool(statusStr)
	appName := ctx.DefaultQuery("appName", "")

	list, err := h.svc.ListAll(ctx, domain.FiltersDetail{
		Filters: filters.Filters{
			Status: status,
		},
		UserId:  uid,
		AppName: appName,
	})

	if err != nil {
		ctx.Set("internal", err.Error())
		zap.L().Error("查询应用详情列表异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}

	response.NewResponse().SuccessResponse(ctx, "查询成功", list)
}
