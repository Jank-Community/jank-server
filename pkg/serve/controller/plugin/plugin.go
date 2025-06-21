// Package post 提供插件相关的HTTP接口处理
// 创建者：ixuemy
// 创建时间：2025-06-21
package plugin

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	bizErr "jank.com/jank_blog/internal/error"
	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/serve/controller/plugin/dto"
	service "jank.com/jank_blog/pkg/serve/service/plugin"
	"jank.com/jank_blog/pkg/vo"
)

// GetOnePlugin 获取插件详情
// @Summary 获取插件详情
// @Description 根据插件ID获取插件的详细信息
// @Tags 插件管理
// @Accept json
// @Produce json
// @Param id query string true "插件ID"
// @Success 200 {object} vo.Response "成功返回插件详情"
// @Failure 400 {object} vo.Response "请求参数错误"
// @Failure 500 {object} vo.Response "服务器内部错误"
// @Router /plugin/getOnePlugin [get]
func GetOnePlugin(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, fmt.Errorf("插件ID错误"), bizErr.New(bizErr.BAD_REQUEST)))
	}

	pos, err := service.GetPlugin(c, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, pos))
}

// GetAllPlugins   插件列表
// @Summary 获取插件列表
// @Description 获取插件的分页列表，可以根据分类、搜索关键词、排序方式和排序顺序进行筛选
// @Tags Plugin
// @Accept json
// @Produce json
// @Param page query int false "页码，默认1"
// @Param page_size query int false "每页数量，默认5"
// @Param category query string false "插件分类"
// @Param search query string false "搜索关键词"
// @Param sort_by query string false "排序字段"
// @Param ascending query bool false "是否升序，默认false"
// @Success 200 {object} vo.Response
// @Failure 500 {object} vo.Response
// @Router /plugin/getAllPlugins [get]
func GetAllPlugins(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil || pageSize < 1 {
		pageSize = 5
	}
	category := c.QueryParam("category")
	search := c.QueryParam("search")
	sortBy := c.QueryParam("sort_by")
	ascending, err := strconv.ParseBool(c.QueryParam("ascending"))
	if err != nil {
		ascending = false
	}

	plugins, err := service.ListPlugins(c, page, pageSize, category, search, sortBy, ascending)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, plugins))
}

// RegisterPlugin 注册插件
// @Summary 注册新插件
// @Description 根据请求体内容注册一个新的插件
// @Tags Plugin
// @Accept json
// @Produce json
// @Param data body dto.RegisterPluginRequest true "插件注册信息"
// @Success 200 {object} vo.Response
// @Failure 400 {object} vo.Response
// @Failure 500 {object} vo.Response
// @Router /plugin/register [post]
func RegisterPlugin(c echo.Context) error {
	req := new(dto.RegisterPluginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, err, bizErr.New(bizErr.BAD_REQUEST, err.Error())))
	}

	errors := utils.Validator(req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, errors, bizErr.New(bizErr.BAD_REQUEST)))
	}

	createdPost, err := service.RegisterPlugin(c, *req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, createdPost))
}

// UpdatePlugin 更新插件
// @Summary 更新插件信息
// @Description 根据请求体内容更新指定插件的信息
// @Tags Plugin
// @Accept json
// @Produce json
// @Param data body dto.UpdatePluginRequest true "插件更新信息"
// @Success 200 {object} vo.Response
// @Failure 400 {object} vo.Response
// @Failure 500 {object} vo.Response
// @Router /plugin/update [put]
func UpdatePlugin(c echo.Context) error {
	req := new(dto.UpdatePluginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, err, bizErr.New(bizErr.BAD_REQUEST, err.Error())))
	}

	errors := utils.Validator(req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, errors, bizErr.New(bizErr.BAD_REQUEST)))
	}

	err := service.UpdatePlugin(c, *req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, ""))
}

// DeletePlugin 删除插件
// @Summary 删除插件
// @Description 根据ID删除指定的插件
// @Tags Plugin
// @Accept json
// @Produce json
// @Param data body dto.DeletePluginRequest true "插件删除请求"
// @Success 200 {object} vo.Response
// @Failure 400 {object} vo.Response
// @Failure 500 {object} vo.Response
// @Router /plugin/delete [delete]
func DeletePlugin(c echo.Context) error {
	req := new(dto.DeletePluginRequest)
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, err, bizErr.New(bizErr.BAD_REQUEST, err.Error())))
	}

	errors := utils.Validator(req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, errors, bizErr.New(bizErr.BAD_REQUEST)))
	}

	err := service.DeletePlugin(c, req.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, "插件删除成功"))
}

// DownloadPlugin 下载插件
// @Summary 下载插件
// @Description 根据ID下载指定的插件
// @Tags Plugin
// @Accept json
// @Produce json
// @Param data body dto.DownloadPluginRequest true "插件删除请求"
// @Success 200 {object} vo.Response
// @Failure 400 {object} vo.Response
// @Failure 500 {object} vo.Response
// @Router /plugin/delete [delete]
func DownloadPlugin(c echo.Context) error {
	req := new(dto.DownloadPluginRequest)
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, err, bizErr.New(bizErr.BAD_REQUEST, err.Error())))
	}

	errors := utils.Validator(req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, errors, bizErr.New(bizErr.BAD_REQUEST)))
	}

	_, err := service.DownloadPlugin(c, req.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, "插件下载成功"))
}

// UploadPlugin 上传插件包
// @Summary 上传插件包
// @Description 上传一个插件zip文件并保存到服务器
// @Tags Plugin
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "插件zip文件"
// @Success 200 {object} vo.Response
// @Failure 400 {object} vo.Response
// @Failure 500 {object} vo.Response
// @Router /plugin/upload [post]
func UploadPlugin(c echo.Context) error {
	path, err := service.UploadPluginPackage(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}
	return c.JSON(http.StatusOK, vo.Success(c, path))
}

// DownloadPluginFile 下载插件包
// @Summary 下载插件包
// @Description 根据文件名下载插件zip文件
// @Tags Plugin
// @Accept json
// @Produce application/zip
// @Param filename query string true "插件zip文件名"
// @Success 200 {file} file
// @Failure 400 {object} vo.Response
// @Failure 404 {object} vo.Response
// @Router /plugin/download [get]
func DownloadPluginFile(c echo.Context) error {
	return service.DownloadPluginPackage(c)
}
