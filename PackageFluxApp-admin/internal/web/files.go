/**
 * Description：
 * FileName：files.go
 * Author：CJiaの用心
 * Create：2025/2/24 16:53:12
 * Remark：
 */

package web

import (
	"errors"
	"fmt"
	"github.com/carefuly/PackageFluxApp/config"
	"github.com/carefuly/PackageFluxApp/internal/domain"
	"github.com/carefuly/PackageFluxApp/internal/service"
	"github.com/carefuly/PackageFluxApp/model"
	"github.com/carefuly/PackageFluxApp/pkg/ginx/query/filters"
	"github.com/carefuly/PackageFluxApp/pkg/response"
	validate "github.com/carefuly/PackageFluxApp/pkg/validator"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
	"strconv"
)

type FilesHandler interface {
	RegisterRoutes(router *gin.RouterGroup)
	BatchUpload(ctx *gin.Context)
	BatchDelete(ctx *gin.Context)
	FindListAllWithPage(ctx *gin.Context)
	FindListAll(ctx *gin.Context)
}

type filesHandler struct {
	rely   config.RelyConfig
	svc    service.FilesService
	aliYun service.AliYunService
}

type IdsRequest []string // 文件ID列表

type FilesListPageResponse struct {
	List     []domain.Files
	Total    int64
	Page     int64
	PageSize int64
}

type FilesListResponse []domain.Files // 列表

func NewFilesHandler(rely config.RelyConfig, svc service.FilesService, aliYun service.AliYunService) FilesHandler {
	return &filesHandler{
		rely:   rely,
		svc:    svc,
		aliYun: aliYun,
	}
}

func (h *filesHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/files/batchUpload", h.BatchUpload)
	router.POST("/files/batchDelete", h.BatchDelete)
	router.GET("/files/listPage", h.FindListAllWithPage)
	router.GET("/files/listAll", h.FindListAll)
}

// BatchUpload
// @id 批量上传文件
// @Summary 批量上传文件
// @Description 批量上传文件
// @Tags 文件
// @Accept multipart/form-data
// @Produce application/json
// @Param files formData []file true "文件列表"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /v1/application/files/batchUpload [post]
// @Security ApiKeyAuth
func (h *filesHandler) BatchUpload(ctx *gin.Context) {
	uid, ok := ctx.MustGet("userId").(string)
	if !ok {
		ctx.Set("internal", uid)
		zap.S().Error("用户ID获取失败", uid)
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}

	email, ok := ctx.MustGet("email").(string)
	if !ok {
		ctx.Set("internal", email)
		zap.S().Error("用户邮箱获取失败", email)
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}

	form, _ := ctx.MultipartForm()

	if form == nil || len(form.File["files"]) == 0 {
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "请选择上传的文件", nil)
		return
	}

	uploadFiles, path, err := h.aliYun.BatchUploadFiles(ctx, email, form.File["files"])
	if err != nil {
		ctx.Set("internal", err)
		zap.L().Error("文件上传至OSS异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}

	var files []domain.Files
	for _, file := range uploadFiles {
		dir, _ := os.Getwd()
		filepath := path + "/" + file.Filename
		files = append(files, domain.Files{
			File: model.File{
				Name:    file.Filename,
				Size:    fmt.Sprintf("%.3f", float64(file.Size)/1024/1024),
				Suffix:  file.Header.Get("Content-Type"),
				BaseDir: dir,
				FileDir: dir + filepath,
				Path:    filepath,
				Url:     h.rely.AliYun.OssPrefix + filepath,
				UserID:  uid,
			},
		})
	}

	err = h.svc.BatchCreate(ctx, files)

	switch {
	case err == nil:
		response.NewResponse().SuccessResponse(ctx, "上传成功", nil)
	case errors.Is(err, service.ErrCreateFile):
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "上传文件失败", nil)
	default:
		ctx.Set("internal", err)
		zap.L().Error("文件入库异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
	}
}

// BatchDelete
// @id 批量删除文件
// @Summary 批量删除文件
// @Description 批量删除文件
// @Tags 文件
// @Accept application/json
// @Produce application/json
// @Param IdsRequest body IdsRequest true "参数"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /v1/application/files/batchDelete [post]
// @Security ApiKeyAuth
func (h *filesHandler) BatchDelete(ctx *gin.Context) {
	uid, ok := ctx.MustGet("userId").(string)
	if !ok {
		ctx.Set("internal", uid)
		zap.S().Error("用户ID获取失败", uid)
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}

	var req IdsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validate.NewValidatorError(h.rely.Trans).HandleValidatorError(ctx, err)
		return
	}
	paths, err := h.svc.BatchDelete(ctx, uid, req)
	if err != nil {
		ctx.Set("internal", err.Error())
		zap.L().Error("删除文件异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}

	err = h.aliYun.BatchDeleteFiles(ctx, paths)
	if err != nil {
		ctx.Set("internal", err)
		zap.L().Error("删除OSS文件异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}

	response.NewResponse().SuccessResponse(ctx, "删除成功", nil)
}

// FindListAllWithPage
// @id 分页查询文件列表
// @Summary 分页查询文件列表
// @Description 分页查询文件列表
// @Tags 文件
// @Accept application/json
// @Produce application/json
// @Param status query string true "状态" default(true)
// @Param page query int true "当前页" default(1)
// @Param pageSize query int true "每页显示的条数" default(10)
// @Success 200 {object} web.FilesListPageResponse
// @Failure 400 {object} response.Response
// @Router /v1/application/files/listPage [get]
// @Security ApiKeyAuth
func (h *filesHandler) FindListAllWithPage(ctx *gin.Context) {
	uid, ok := ctx.MustGet("userId").(string)
	if !ok {
		ctx.Set("internal", uid)
		zap.S().Error("用户ID获取失败", uid)
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}

	pageStr := ctx.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)
	pageSizeStr := ctx.DefaultQuery("pageSize", "10")
	pageSize, _ := strconv.Atoi(pageSizeStr)
	statusStr := ctx.DefaultQuery("status", "true")
	status, _ := strconv.ParseBool(statusStr)

	list, total, err := h.svc.FindListAllWithPage(ctx, domain.FiltersFile{
		Filters: filters.Filters{
			Status: status,
		},
		UserId: uid,
	}, domain.Pagination{
		Page:     page,
		PageSize: pageSize,
	})

	if err != nil {
		ctx.Set("internal", err.Error())
		zap.L().Error("分页查询文件列表异常", zap.Error(err))
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

// FindListAll
// @id 查询文件列表
// @Summary 查询文件列表
// @Description 查询文件列表
// @Tags 文件
// @Accept application/json
// @Produce application/json
// @Param status query string true "状态" default(true)
// @Success 200 {object} web.FilesListResponse
// @Failure 400 {object} response.Response
// @Router /v1/application/files/listAll [get]
// @Security ApiKeyAuth
func (h *filesHandler) FindListAll(ctx *gin.Context) {
	uid, ok := ctx.MustGet("userId").(string)
	if !ok {
		ctx.Set("internal", uid)
		zap.S().Error("用户ID获取失败", uid)
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}

	statusStr := ctx.DefaultQuery("status", "true")
	status, _ := strconv.ParseBool(statusStr)

	list, err := h.svc.FindListAll(ctx, domain.FiltersFile{
		Filters: filters.Filters{
			Status: status,
		},
		UserId: uid,
	})

	if err != nil {
		ctx.Set("internal", err.Error())
		zap.L().Error("查询文件列表异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}

	response.NewResponse().SuccessResponse(ctx, "查询成功", list)
}
