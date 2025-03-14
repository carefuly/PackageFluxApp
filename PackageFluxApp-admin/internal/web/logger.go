/**
 * Description：
 * FileName：logger.go
 * Author：CJiaの用心
 * Create：2025/3/14 15:39:41
 * Remark：
 */

package web

import (
	"github.com/carefuly/PackageFluxApp/config"
	"github.com/carefuly/PackageFluxApp/internal/domain"
	"github.com/carefuly/PackageFluxApp/internal/model"
	"github.com/carefuly/PackageFluxApp/internal/service"
	"github.com/carefuly/PackageFluxApp/pkg/ginx/query/filters"
	"github.com/carefuly/PackageFluxApp/pkg/response"
	validate "github.com/carefuly/PackageFluxApp/pkg/validator"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type LoggerHandler interface {
	RegisterRoutes(router *gin.RouterGroup)
	Create(ctx *gin.Context)
	GetListAllPage(ctx *gin.Context)
}

type loggerHandler struct {
	rely config.RelyConfig
	svc  service.LoggerService
}

func NewLoggerHandler(rely config.RelyConfig, svc service.LoggerService) LoggerHandler {
	return &loggerHandler{
		rely: rely,
		svc:  svc,
	}
}

func (h *loggerHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/logger/create", h.Create)
	router.GET("/logger/listPage", h.GetListAllPage)
}

func (h *loggerHandler) Create(ctx *gin.Context) {
	type LoggerRequest struct {
		AppId      string `json:"appId"`       // 应用标识Id
		AppName    string ` json:"appName"`    // 应用名称
		AppVersion string ` json:"appVersion"` // 应用版本
		Brand      string `json:"brand"`       // 设备品牌
		SystemName string `json:"system"`      // 系统名称
		UniqueId   string `json:"uniqueId"`    // 应用唯一Id
		UniRemark  string `json:"uniRemark"`   // 应用标识备注
		DetailId   string `json:"detail_id"`   // 应用详情ID
	}

	var req LoggerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validate.NewValidatorError(h.rely.Trans).HandleValidatorError(ctx, err)
		return
	}

	err := h.svc.Create(ctx, domain.Logger{
		Logger: model.Logger{
			AppId:      req.AppId,
			AppName:    req.AppName,
			AppVersion: req.AppVersion,
			Brand:      req.Brand,
			SystemName: req.SystemName,
			UniqueId:   req.UniqueId,
			UniRemark:  req.UniRemark,
			DetailId:   req.DetailId,
		},
	})

	switch {
	case err == nil:
		response.NewResponse().SuccessResponse(ctx, "新增成功", nil)
	default:
		ctx.Set("internal", err.Error())
		zap.L().Error("新增应用检查日志异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
	}
}

func (h *loggerHandler) GetListAllPage(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)
	pageSizeStr := ctx.DefaultQuery("pageSize", "10")
	pageSize, _ := strconv.Atoi(pageSizeStr)

	detailId := ctx.DefaultQuery("detailId", "")
	if detailId == "" {
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "应用id不能为空", nil)
		return
	}

	total, list, err := h.svc.FindListPage(ctx, detailId, filters.Pagination{
		Page:     page,
		PageSize: pageSize,
	})

	if err != nil {
		ctx.Set("internal", err.Error())
		zap.L().Error("分页查询应用检查日志列表异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}

	response.NewResponse().SuccessResponse(ctx, "查询成功", gin.H{
		"list":     list,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})

}
