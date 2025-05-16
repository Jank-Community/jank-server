// Package theme 提供主题相关的HTTP接口处理
// 创建者：Done-0
// 创建时间：2025-05-14
package theme

import (
	"net/http"

	"github.com/labstack/echo/v4"

	bizErr "jank.com/jank_blog/internal/error"
	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/serve/controller/theme/dto"
	service "jank.com/jank_blog/pkg/serve/service/theme"
	"jank.com/jank_blog/pkg/vo"
)

// GetActivatedTheme godoc
// @Summary      获取当前激活的主题
// @Description  获取当前已激活的主题信息
// @Tags         主题
// @Accept       json
// @Produce      json
// @Success      200     {object}   vo.Result{data=theme.ThemeVO}  "获取成功"
// @Failure      500     {object}   vo.Result          "服务器错误"
// @Router       /theme/getActivatedTheme [get]
func GetActivatedTheme(c echo.Context) error {
	theme, err := service.GetActivatedTheme(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, theme))
}

// GetOneTheme godoc
// @Summary      获取指定主题信息
// @Description  根据主题ID获取主题详细信息
// @Tags         主题
// @Accept       json
// @Produce      json
// @Param        themeID   path      string  true  "主题ID"
// @Success      200     {object}   vo.Result{data=theme.ThemeVO}  "获取成功"
// @Failure      400     {object}   vo.Result          "请求参数错误"
// @Failure      404     {object}   vo.Result          "主题不存在"
// @Failure      500     {object}   vo.Result          "服务器错误"
// @Router       /theme/getOneTheme [get]
func GetOneTheme(c echo.Context) error {
	req := new(dto.GetThemeRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, err, bizErr.New(bizErr.BAD_REQUEST, err.Error())))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, errors, bizErr.New(bizErr.BAD_REQUEST)))
	}

	theme, err := service.GetThemeByID(c, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, theme))
}

// ListAllThemes godoc
// @Summary      获取所有主题列表
// @Description  获取系统中所有可用的主题
// @Tags         主题
// @Accept       json
// @Produce      json
// @Success      200     {object}   vo.Result{data=[]theme.ThemeVO}  "获取成功"
// @Failure      500     {object}   vo.Result          "服务器错误"
// @Router       /theme/listAllThemes [get]
func ListAllThemes(c echo.Context) error {
	themes, err := service.ListAllThemes(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, themes))
}

// ActivateOneTheme godoc
// @Summary      激活主题
// @Description  激活指定ID的主题
// @Tags         主题
// @Accept       json
// @Produce      json
// @Param        themeID   path      string  true  "主题ID"
// @Success      200     {object}   vo.Result{data=theme.ThemeVO}  "激活成功"
// @Failure      400     {object}   vo.Result          "请求参数错误"
// @Failure      404     {object}   vo.Result          "主题不存在"
// @Failure      500     {object}   vo.Result          "服务器错误"
// @Security     BearerAuth
// @Router       /theme/activateOneTheme [post]
func ActivateOneTheme(c echo.Context) error {
	req := new(dto.ActivateThemeRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, err, bizErr.New(bizErr.BAD_REQUEST, err.Error())))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, errors, bizErr.New(bizErr.BAD_REQUEST)))
	}

	theme, err := service.ActivateOneTheme(c, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, theme))
}

// DeleteOneTheme godoc
// @Summary      删除主题
// @Description  删除指定ID的主题及其文件
// @Tags         主题
// @Accept       json
// @Produce      json
// @Param        themeID        path      string   true  "主题ID"
// @Param        deleteSetting  query     boolean  false "是否删除设置" default(false)
// @Success      200     {object}   vo.Result  "删除成功"
// @Failure      400     {object}   vo.Result  "请求参数错误"
// @Failure      404     {object}   vo.Result  "主题不存在"
// @Failure      500     {object}   vo.Result  "服务器错误"
// @Security     BearerAuth
// @Router       /theme/deleteOneTheme [post]
func DeleteOneTheme(c echo.Context) error {
	req := new(dto.DeleteThemeRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, err, bizErr.New(bizErr.BAD_REQUEST, err.Error())))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, errors, bizErr.New(bizErr.BAD_REQUEST)))
	}

	err := service.DeleteThemeByID(c, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, "主题删除成功"))
}
