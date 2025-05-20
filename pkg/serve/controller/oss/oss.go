// Package oss 提供对象存储相关的HTTP接口处理
// 创建者：Done-0
// 创建时间：2025-05-10
package oss

import (
	"net/http"

	"github.com/labstack/echo/v4"

	bizErr "jank.com/jank_blog/internal/error"
	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/serve/controller/oss/dto"
	service "jank.com/jank_blog/pkg/serve/service/oss"
	"jank.com/jank_blog/pkg/vo"
)

// UploadOneFile godoc
// @Summary      上传文件
// @Description  上传文件到 MinIO 对象存储
// @Tags         对象存储
// @Accept       multipart/form-data
// @Produce      json
// @Param        file        formData  file   true  "要上传的文件"
// @Param        bucket_name formData  string true  "存储桶名称"
// @Success      200     {object}   vo.Result{data=string}  "上传成功"
// @Failure      400     {object}   vo.Result              "请求参数错误"
// @Failure      500     {object}   vo.Result              "服务器错误"
// @Security     BearerAuth
// @Router       /oss/uploadFile [post]
func UploadOneFile(c echo.Context) error {
	req := new(dto.UploadOneFileRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, err, bizErr.New(bizErr.BAD_REQUEST, err.Error())))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, errors, bizErr.New(bizErr.BAD_REQUEST)))
	}

	objectName, err := service.UploadOneFile(c, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, objectName))
}

// DownloadOneFile godoc
// @Summary      下载文件
// @Description  从 MinIO 对象存储下载文件
// @Tags         对象存储
// @Accept       json
// @Produce      json
// @Param        request  body      dto.DownloadOneFileRequest  true  "下载文件请求参数"
// @Success      200     {object}   vo.Result{data=string}  "下载成功"
// @Failure      400     {object}   vo.Result              "请求参数错误"
// @Failure      404     {object}   vo.Result              "文件不存在"
// @Failure      500     {object}   vo.Result              "服务器错误"
// @Security     BearerAuth
// @Router       /oss/downloadFile [post]
func DownloadOneFile(c echo.Context) error {
	req := new(dto.DownloadOneFileRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, err, bizErr.New(bizErr.BAD_REQUEST, err.Error())))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, errors, bizErr.New(bizErr.BAD_REQUEST)))
	}

	result, err := service.DownloadOneFile(c, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, result))
}

// DeleteOneFile godoc
// @Summary      删除文件
// @Description  从 MinIO 对象存储删除文件
// @Tags         对象存储
// @Accept       json
// @Produce      json
// @Param        request  body      dto.DeleteOneFileRequest  true  "删除文件请求参数"
// @Success      200     {object}   vo.Result{data=string}  "删除成功"
// @Failure      400     {object}   vo.Result              "请求参数错误"
// @Failure      404     {object}   vo.Result              "文件不存在"
// @Failure      500     {object}   vo.Result              "服务器错误"
// @Security     BearerAuth
// @Router       /oss/deleteFile [post]
func DeleteOneFile(c echo.Context) error {
	req := new(dto.DeleteOneFileRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, err, bizErr.New(bizErr.BAD_REQUEST, err.Error())))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, errors, bizErr.New(bizErr.BAD_REQUEST)))
	}

	if err := service.DeleteOneFile(c, req); err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, "文件删除成功"))
}

// ListAllObjects godoc
// @Summary      列出对象
// @Description  列出 MinIO 对象存储中的文件
// @Tags         对象存储
// @Accept       json
// @Produce      json
// @Param        request  body      dto.ListAllObjectsRequest  true  "列出对象请求参数"
// @Success      200     {object}   vo.Result{data=[]string}  "获取成功"
// @Failure      400     {object}   vo.Result              "请求参数错误"
// @Failure      500     {object}   vo.Result              "服务器错误"
// @Security     BearerAuth
// @Router       /oss/listObjects [post]
func ListAllObjects(c echo.Context) error {
	req := new(dto.ListAllObjectsRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, err, bizErr.New(bizErr.BAD_REQUEST, err.Error())))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, errors, bizErr.New(bizErr.BAD_REQUEST)))
	}

	objects, err := service.ListAllObjects(c, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, objects))
}
