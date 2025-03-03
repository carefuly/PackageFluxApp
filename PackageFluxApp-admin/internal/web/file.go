/**
 * Description：
 * FileName：file.go
 * Author：CJiaの用心
 * Create：2025/2/27 22:15:53
 * Remark：
 */

package web

import (
	"errors"
	"fmt"
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
	"os"
	"path/filepath"
	"strconv"
)

type FileHandler interface {
	RegisterRoutes(router *gin.RouterGroup)
	BatchUpload(ctx *gin.Context)
	BatchDelete(ctx *gin.Context)
	GetById(ctx *gin.Context)
	GetListAllPage(ctx *gin.Context)
	GetListAll(ctx *gin.Context)
}

type fileHandler struct {
	rely   config.RelyConfig
	svc    service.FileService
	aliYun service.AliYunService
}

func NewFileHandler(rely config.RelyConfig, svc service.FileService, aliYun service.AliYunService) FileHandler {
	return &fileHandler{
		rely:   rely,
		svc:    svc,
		aliYun: aliYun,
	}
}

func (h *fileHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/file/batchUpload", h.BatchUpload)
	router.POST("/file/batchDelete", h.BatchDelete)
	router.GET("/file/getById/:id", h.GetById)
	router.GET("/file/listPage", h.GetListAllPage)
	router.GET("/file/listAll", h.GetListAll)
}

func (h *fileHandler) BatchUpload(ctx *gin.Context) {
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

	var files []domain.File
	for _, file := range uploadFiles {
		dir, _ := os.Getwd()
		filePath := path + "/" + file.Filename

		fileSuffix := filepath.Ext(file.Filename)

		files = append(files, domain.File{
			File: model.File{
				Name:    file.Filename,
				Size:    fmt.Sprintf("%.3f", float64(file.Size)/1024/1024),
				Type:    fileSuffix,
				Suffix:  file.Header.Get("Content-Type"),
				BaseDir: dir,
				FileDir: dir + filePath,
				Path:    filePath,
				Url:     h.rely.AliYun.OssPrefix + filePath,
				UserId:  uid,
			},
		})
	}

	result := h.svc.BatchCreate(ctx, files)

	response.NewResponse().SuccessResponse(ctx, fmt.Sprintf("【成功上传%d个文件】%v", len(result), result), nil)
}

func (h *fileHandler) BatchDelete(ctx *gin.Context) {
	type IdsRequest []string

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

	paths, errs := h.svc.BatchDelete(ctx, uid, req)
	fmt.Println("paths", paths)
	if len(errs) != 0 {
		ctx.Set("internal", errs)
		zap.S().Error("删除文件异常", errs)
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}

	err := h.aliYun.BatchDeleteFiles(ctx, paths)
	if err != nil {
		ctx.Set("internal", err)
		zap.L().Error("删除OSS文件异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}

	response.NewResponse().SuccessResponse(ctx, fmt.Sprintf("【成功删除%d个文件】%v", len(paths), paths), nil)
}

func (h *fileHandler) GetById(ctx *gin.Context) {
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

	file, err := h.svc.FindById(ctx, id, uid)

	switch {
	case err == nil:
		response.NewResponse().SuccessResponse(ctx, "查询成功", file)
	case errors.Is(err, service.ErrDetailRecordNotFound):
		response.NewResponse().ErrorResponse(ctx, http.StatusBadRequest, "记录不存在", nil)
	default:
		ctx.Set("internal", err)
		zap.S().Error("根据Id查询文件详情异常", err)
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
	}
}

func (h *fileHandler) GetListAllPage(ctx *gin.Context) {
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

	total, list, err := h.svc.FindListPage(ctx, uid, filters.Pagination{
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

func (h *fileHandler) GetListAll(ctx *gin.Context) {
	uid, ok := ctx.MustGet("userId").(string)
	if !ok {
		ctx.Set("internal", uid)
		zap.S().Error("用户ID获取失败", uid)
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}

	name := ctx.DefaultQuery("name", "")

	list, err := h.svc.FindListAll(ctx, uid, name)

	if err != nil {
		ctx.Set("internal", err.Error())
		zap.L().Error("查询文件列表异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
		return
	}

	response.NewResponse().SuccessResponse(ctx, "查询成功", list)
}
