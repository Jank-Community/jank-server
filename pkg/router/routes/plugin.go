// Package routes 提供路由注册功能
// 创建者：ixuemy
// 创建时间：2025-06-21
package routes

import (
	"github.com/labstack/echo/v4"
	auth_middleware "jank.com/jank_blog/internal/middleware/auth"
	"jank.com/jank_blog/pkg/serve/controller/plugin"
)

// RegisterPluginRoutes 注册插件相关路由
// 参数：
//   - r: Echo 路由组数组，r[0] 为 API v1 版本组
func RegisterPluginRoutes(r ...*echo.Group) {
	// api v1 groupdownload
	apiV1 := r[0]
	pluginGroupV1 := apiV1.Group("/plugin")
	pluginGroupV1.GET("/getOnePlugin", plugin.GetOnePlugin)
	pluginGroupV1.GET("/getAllPlugins", plugin.GetAllPlugins)
	pluginGroupV1.POST("/upload", plugin.UploadPlugin, auth_middleware.AuthMiddleware())
	pluginGroupV1.GET("/download", plugin.DownloadPluginFile)
	pluginGroupV1.POST("/registerPlugin", plugin.RegisterPlugin)
	pluginGroupV1.POST("/updatePlugin", plugin.UpdatePlugin, auth_middleware.AuthMiddleware())
	pluginGroupV1.POST("/deletePlugin", plugin.DeletePlugin, auth_middleware.AuthMiddleware())
	pluginGroupV1.POST("/downloadPlugin", plugin.DownloadPlugin)
}
