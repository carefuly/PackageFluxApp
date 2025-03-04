/**
 * Description：
 * FileName：version.go
 * Author：CJiaの用心
 * Create：2025/2/28 14:43:53
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

type VersionHandler interface {
	RegisterRoutes(router *gin.RouterGroup)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
	GetById(ctx *gin.Context)
	GetListAll(ctx *gin.Context)
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

type VersionRequest struct {
	VersionCode    string `json:"versionCode" binding:"required"` // 版本号
	Description    string ` json:"description"`                   // 更新说明
	HBuilderUpdate bool   `json:"HBuilderUpdate"`                 // HBuilderX是否更新
	ForceUpdate    bool   `json:"forceUpdate"`                    // 是否强制更新
	ApkUrl         string `json:"apkUrl" binding:"required"`      // apk下载地址
	WgtUrl         string `json:"wgtUrl" binding:"required"`      // wgt下载地址
	DetailId       string `json:"detail_id"`                      // 应用详情ID
	Remark         string `json:"remark" binding:"max=255"`       // 备注
}

func (h *versionHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/version/create", h.Create)
	router.DELETE("/version/delete/:id", h.Delete)
	router.PUT("/version/update/:id", h.Update)
	router.PUT("/version/updateFormal/:id", h.UpdateFormal)
	router.GET("/version/getById/:id", h.GetById)
	router.GET("/version/listAll", h.GetListAll)
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
			ApkUrl:         req.ApkUrl,
			WgtUrl:         req.WgtUrl,
			DetailId:       req.DetailId,
			UserId:         uid,
		},
	})

	switch {
	case err == nil:
		response.NewResponse().SuccessResponse(ctx, "新增成功", nil)
	case errors.Is(err, service.ErrDuplicateDetailIdAndVersionCode):
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "同一应用下应用版本号不能重复，请重新输入", nil)
	default:
		ctx.Set("internal", err.Error())
		zap.L().Error("新增版本异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
	}
}

func (h *versionHandler) Delete(ctx *gin.Context) {
	uid, ok := ctx.MustGet("userId").(string)
	if !ok {
		ctx.Set("internal", uid)
		zap.S().Error("用户ID获取失败", uid)
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}

	id := ctx.Param("id")
	detailId := ctx.DefaultQuery("detailId", "")
	if id == "" || len(id) == 0 {
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "参数错误", nil)
		return
	}

	err := h.svc.Delete(ctx, id, detailId, uid)

	switch {
	case err == nil:
		response.NewResponse().SuccessResponse(ctx, "删除成功", nil)
	case errors.Is(err, service.ErrVersionNotFound):
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "记录不存在", nil)
	default:
		ctx.Set("internal", err.Error())
		zap.L().Error("删除版本异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
	}
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

	err := h.svc.Update(ctx, id, req.DetailId, uid, domain.Version{
		Version: model.Version{
			CoreModels: models.CoreModels{
				Modifier: uid,
				Remark:   req.Remark,
			},
			Description:    req.Description,
			HBuilderUpdate: req.HBuilderUpdate,
			ForceUpdate:    req.ForceUpdate,
			ApkUrl:         req.ApkUrl,
			WgtUrl:         req.WgtUrl,
			DetailId:       req.DetailId,
			UserId:         uid,
		},
	})

	switch {
	case err == nil:
		response.NewResponse().SuccessResponse(ctx, "更新成功", nil)
	case errors.Is(err, service.ErrDuplicateDetailIdAndVersionCode):
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "同一应用下应用版本号不能重复，请重新输入", nil)
	case errors.Is(err, service.ErrVersionNotFound):
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "记录不存在", nil)
	default:
		ctx.Set("internal", err.Error())
		zap.L().Error("更新版本异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
	}
}

func (h *versionHandler) UpdateFormal(ctx *gin.Context) {
	uid, ok := ctx.MustGet("userId").(string)
	if !ok {
		ctx.Set("internal", uid)
		zap.S().Error("用户ID获取失败", uid)
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}

	id := ctx.Param("id")
	detailId := ctx.DefaultQuery("detailId", "")
	if id == "" || len(id) == 0 {
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "参数错误", nil)
		return
	}

	err := h.svc.UpdateFormal(ctx, id, detailId, uid)

	switch {
	case err == nil:
		response.NewResponse().SuccessResponse(ctx, "设置正式版成功", nil)
	case errors.Is(err, service.ErrVersionNotFound):
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "记录不存在", nil)
	case errors.Is(err, service.ErrVersionLowerInvalid):
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "暂不支持降低更新，请设置高于当前正式版的版本号更新", nil)
	default:
		ctx.Set("internal", err.Error())
		zap.L().Error("设置正式版异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
	}
}

func (h *versionHandler) GetById(ctx *gin.Context) {
	uid, ok := ctx.MustGet("userId").(string)
	if !ok {
		ctx.Set("internal", uid)
		zap.S().Error("用户ID获取失败", uid)
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}

	id := ctx.Param("id")
	detailId := ctx.DefaultQuery("detailId", "")
	if id == "" || len(id) == 0 {
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "参数错误", nil)
		return
	}

	detail, err := h.svc.FindById(ctx, id, detailId, uid)

	switch {
	case err == nil:
		response.NewResponse().SuccessResponse(ctx, "查询成功", detail)
	case errors.Is(err, service.ErrVersionRecordNotFound):
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "记录不存在", nil)
	default:
		ctx.Set("internal", err.Error())
		zap.S().Error("根据Id查询版本异常", err)
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
	}
}

func (h *versionHandler) GetListAll(ctx *gin.Context) {
	uid, ok := ctx.MustGet("userId").(string)
	if !ok {
		ctx.Set("internal", uid)
		zap.S().Error("用户ID获取失败", uid)
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}

	detailId := ctx.DefaultQuery("detailId", "")

	list, err := h.svc.FindListAll(ctx, detailId, uid)

	if err != nil {
		ctx.Set("internal", err.Error())
		zap.L().Error("查询版本列表异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}

	response.NewResponse().SuccessResponse(ctx, "查询成功", list)
}
