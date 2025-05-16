// Package routes 提供路由注册功能
// 创建者：Done-0
// 创建时间：2025-05-10
package routes

import (
	"github.com/labstack/echo/v4"

	"jank.com/jank_blog/pkg/serve/controller/theme"
)

// RegisterThemeRoutes 注册主题相关路由
// 参数：
//   - r: Echo 路由组数组，r[0] 为 API v1 版本组
func RegisterThemeRoutes(r ...*echo.Group) {
	// API v1 路由组
	apiV1 := r[0]
	themeGroupV1 := apiV1.Group("/theme")

	// 主题基本信息API，包括获取和列出主题
	themeGroupV1.GET("/getActivatedTheme", theme.GetActivatedTheme) // 获取当前激活的主题
	themeGroupV1.GET("/getOneTheme", theme.GetOneTheme)             // 根据ID获取主题
	themeGroupV1.GET("/listAllThemes", theme.ListAllThemes)         // 列出所有主题
	themeGroupV1.POST("/activateOneTheme", theme.ActivateOneTheme)  // 激活指定主题
	themeGroupV1.POST("/deleteOneTheme", theme.DeleteOneTheme)      // 删除指定主题
}
