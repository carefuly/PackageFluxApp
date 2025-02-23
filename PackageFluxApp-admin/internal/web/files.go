/**
 * Description：
 * FileName：files.go
 * Author：CJiaの用心
 * Create：2025/2/22 20:25:16
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
	"github.com/carefuly/PackageFluxApp/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
)

type FilesHandler interface {
	RegisterRoutes(router *gin.RouterGroup)
	BatchUpload(ctx *gin.Context)
	// EmailRegisterHandler(ctx *gin.Context)
}

type filesHandler struct {
	rely   config.RelyConfig
	svc    service.FilesService
	aliYun service.AliYunService
}

func NewFilesHandler(rely config.RelyConfig, svc service.FilesService, aliYun service.AliYunService) FilesHandler {
	return &filesHandler{
		rely:   rely,
		svc:    svc,
		aliYun: aliYun,
	}
}

func (h *filesHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/files/batchUpload", h.BatchUpload)
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
// @Security token
func (h *filesHandler) BatchUpload(ctx *gin.Context) {
	uid, ok := ctx.MustGet("userId").(uint)
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
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "上传文件失败", nil)
	default:
		ctx.Set("internal", err)
		zap.L().Error("文件入库异常", zap.Error(err))
		response.NewResponse().ErrorResponse(ctx, http.StatusInternalServerError, "服务器异常", nil)
	}
}
