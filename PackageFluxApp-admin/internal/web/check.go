/**
 * Description：
 * FileName：check.go
 * Author：CJiaの用心
 * Create：2025/3/1 00:15:56
 * Remark：
 */

package web

import (
	"github.com/carefuly/PackageFluxApp/config"
	"github.com/carefuly/PackageFluxApp/internal/service"
	"github.com/carefuly/PackageFluxApp/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type CheckHandler interface {
	RegisterRoutes(router *gin.RouterGroup)
	CheckLatest(ctx *gin.Context)
}

type checkHandler struct {
	rely config.RelyConfig
	svc  service.CheckService
}

func NewCheckHandler(rely config.RelyConfig, svc service.CheckService) CheckHandler {
	return &checkHandler{
		rely: rely,
		svc:  svc,
	}
}

type CheckRequest struct {
	DetailId    string `json:"detailId" binding:"required"`    // 详情id
	VersionCode string `json:"versionCode" binding:"required"` // 版本号
}

func (h *checkHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/check/latest", h.CheckLatest)
}

func (h *checkHandler) CheckLatest(ctx *gin.Context) {
	detailId := ctx.DefaultQuery("detailId", "")
	if detailId == "" {
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "应用id不能为空", nil)
		return
	}
	versionCode := ctx.DefaultQuery("versionCode", "")
	if versionCode == "" {
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "版本号不能为空", nil)
		return
	}

	latest, err := h.svc.CheckLatest(ctx, detailId, versionCode)

	if err != nil {
		zap.L().Error("检查最新版异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
	}

	response.NewResponse().SuccessResponse(ctx, "检查成功", latest)
}
