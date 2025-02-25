/**
 * Description：
 * FileName：version.go
 * Author：CJiaの用心
 * Create：2025/2/25 15:49:49
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

type VersionHandler interface {
	RegisterRoutes(router *gin.RouterGroup)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
	Formal(ctx *gin.Context)
	FindById(ctx *gin.Context)
	FindListAll(ctx *gin.Context)
}

type versionHandler struct {
	rely config.RelyConfig
	svc  service.VersionService
}

func NewVersionHandler(rely config.RelyConfig, svc service.VersionService) VersionHandler {
	return &versionHandler{
		rely: rely,
		svc:  svc,
	}
}

type VersionFilters struct {
	filters.Filters
}

type VersionRequest struct {
	VersionCode    string `json:"versionCode" binding:"required"` // 版本号
	Description    string ` json:"description"`                   // 更新说明
	HBuilderUpdate bool   `json:"HBuilderUpdate"`                 // HBuilderX是否更新
	ForceUpdate    bool   `json:"forceUpdate"`                    // 是否强制更新
	FormalVersion  bool   `json:"formalVersion"`                  // 是否是正式版本
	ApkUrl         string `json:"apkUrl" binding:"required"`      // apk下载地址
	WgtUrl         string `json:"wgtUrl" binding:"required"`      // wgt下载地址
	DetailId       string `json:"detail_id"`                      // 应用详情ID
	Remark         string `json:"remark" binding:"max=255"`       // 备注
}

type FormalRequest struct {
	DetailId  string `json:"detailId" binding:"required"`  // 应用详情ID
	VersionId string `json:"versionId" binding:"required"` // 版本ID
}

type VersionListResponse []domain.Detail // 列表

func (h *versionHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/version/create", h.Create)
	router.DELETE("/version/delete/:id", h.Delete)
	router.PUT("/version/update/:id", h.Update)
	router.PUT("/version/appFormal", h.Formal)
	router.GET("/version/getById/:id", h.FindById)
	router.GET("/version/listAll", h.FindListAll)
}

func (h *versionHandler) Create(ctx *gin.Context) {
	uid, ok := ctx.MustGet("userId").(string)
	if !ok {
		ctx.Set("internal", uid)
		zap.S().Error("用户ID获取失败", uid)
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}

	var req VersionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validate.NewValidatorError(h.rely.Trans).HandleValidatorError(ctx, err)
		return
	}

	err := h.svc.Create(ctx, domain.Version{
		Version: model.Version{
			CoreModels: models.CoreModels{
				Creator:  uid,
				Modifier: uid,
				Remark:   req.Remark,
			},
			VersionCode:    req.VersionCode,
			Description:    req.Description,
			HBuilderUpdate: req.HBuilderUpdate,
			ForceUpdate:    req.ForceUpdate,
			FormalVersion:  req.FormalVersion,
			ApkUrl:         req.ApkUrl,
			WgtUrl:         req.WgtUrl,
			DetailID:       req.DetailId,
		},
	})

	switch {
	case err == nil:
		response.NewResponse().SuccessResponse(ctx, "新增成功", nil)
	case errors.Is(err, service.ErrDuplicateVersionCode):
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "相同的应用和版本号已存在，请重新输入", nil)
	default:
		ctx.Set("internal", err.Error())
		zap.L().Error("新增应用版本异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}
}

func (h *versionHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" || len(id) == 0 {
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "参数错误", nil)
		return
	}

	err := h.svc.Delete(ctx, id)

	if err != nil {
		ctx.Set("internal", err.Error())
		zap.L().Error("删除应用版本异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}

	response.NewResponse().SuccessResponse(ctx, "删除成功", nil)
}

func (h *versionHandler) Update(ctx *gin.Context) {
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

	var req VersionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validate.NewValidatorError(h.rely.Trans).HandleValidatorError(ctx, err)
		return
	}

	err := h.svc.Update(ctx, id, domain.Version{
		Version: model.Version{
			CoreModels: models.CoreModels{
				Modifier: uid,
				Remark:   req.Remark,
			},
			Description:    req.Description,
			HBuilderUpdate: req.HBuilderUpdate,
			ForceUpdate:    req.ForceUpdate,
			FormalVersion:  req.FormalVersion,
			ApkUrl:         req.ApkUrl,
			WgtUrl:         req.WgtUrl,
			DetailID:       req.DetailId,
		},
	})

	switch {
	case err == nil:
		response.NewResponse().SuccessResponse(ctx, "更新成功", nil)
	case errors.Is(err, service.ErrDuplicateVersionCode):
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "相同的应用和版本号已存在，请重新输入", nil)
	default:
		ctx.Set("internal", err.Error())
		zap.L().Error("更新应用版本异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}
}

func (h *versionHandler) Formal(ctx *gin.Context) {
	var req FormalRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validate.NewValidatorError(h.rely.Trans).HandleValidatorError(ctx, err)
		return
	}

	err := h.svc.Formal(ctx, domain.FormalVersion{
		DetailId:  req.DetailId,
		VersionId: req.VersionId,
	})

	if err != nil {
		ctx.Set("internal", err.Error())
		zap.L().Error("设置正式版异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}

	response.NewResponse().SuccessResponse(ctx, "设置成功", nil)
}

func (h *versionHandler) FindById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" || len(id) == 0 {
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "参数错误", nil)
		return
	}

	v, err := h.svc.FindById(ctx, id)

	switch {
	case err == nil:
		response.NewResponse().SuccessResponse(ctx, "查询成功", v)
	case errors.Is(err, service.ErrVersionNotFound):
		response.NewResponse().ErrorResponse(ctx, http.StatusNotFound, "应用版本不存在", nil)
	default:
		ctx.Set("internal", err.Error())
		zap.L().Error("根据ID查询应用详情异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}
}

func (h *versionHandler) FindListAll(ctx *gin.Context) {
	statusStr := ctx.DefaultQuery("status", "true")
	status, _ := strconv.ParseBool(statusStr)
	detailId := ctx.DefaultQuery("detail_id", "")

	if detailId == "" {
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "参数错误", nil)
		return
	}

	list, err := h.svc.FindListAll(ctx, domain.FiltersVersion{
		Filters: filters.Filters{
			Status: status,
		},
		DetailId: detailId,
	})

	if err != nil {
		ctx.Set("internal", err.Error())
		zap.L().Error("查询应用详情列表异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}

	response.NewResponse().SuccessResponse(ctx, "查询成功", list)
}
