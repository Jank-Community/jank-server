// Package routes 提供路由注册功能
// 创建者：Done-0
// 创建时间：2025-05-10
package routes

import (
	"github.com/labstack/echo/v4"

	auth_middleware "jank.com/jank_blog/internal/middleware/auth"
	"jank.com/jank_blog/pkg/serve/controller/oss"
)

// RegisterOssRoutes 注册对象存储相关路由
// 参数：
//   - r: Echo 路由组数组，r[0] 为 API v1 版本组
func RegisterOssRoutes(r ...*echo.Group) {
	// api v1 group
	apiV1 := r[0]
	ossGroupV1 := apiV1.Group("/oss")
	ossGroupV1.POST("/uploadOneFile", oss.UploadOneFile, auth_middleware.AuthMiddleware())
	ossGroupV1.GET("/downloadOneFile", oss.DownloadOneFile, auth_middleware.AuthMiddleware())
	ossGroupV1.POST("/deleteOneFile", oss.DeleteOneFile, auth_middleware.AuthMiddleware())
	ossGroupV1.GET("/listAllObjects", oss.ListAllObjects, auth_middleware.AuthMiddleware())
}
