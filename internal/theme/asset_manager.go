// Package theme 提供主题的检索和切换功能
// 资源管理器
// 创建者：Pey121
// 创建时间：2025-06-09
package theme

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"

	"jank.com/jank_blog/internal/utils"
)

// AssetManager 负责静态资源的查找与响应
// 字段：
//   - ThemeManager: 主题管理器实例
//
// 查找顺序: current_theme → default → 404
type AssetManager struct {
	ThemeManager *ThemeManager // 主题管理器实例
}

// NewAssetManager 创建资源管理器实例
// 参数：
//   - tm: 主题管理器实例
//
// 返回值：
//   - *AssetManager: 资源管理器实例
func NewAssetManager(tm *ThemeManager) *AssetManager {
	return &AssetManager{
		ThemeManager: tm,
	}
}

// GetAssetURL 获取资源的 URL 路径
// 参数：
//   - c: Echo 上下文
//   - path: 资源相对路径
//
// 返回值：
//   - string: 资源的 URL 路径
func (am *AssetManager) GetAssetURL(c echo.Context, path string) string {
	themeName := am.ThemeManager.activeThemeName
	return fmt.Sprintf("/themes/%s/assets/%s", themeName, path)
}

// ServeAsset 处理资源请求，支持 fallback 到 default 主题
// 参数：
//   - c: Echo 上下文
//
// 返回值：
//   - error: 处理过程中的错误
func (am *AssetManager) ServeAsset(c echo.Context) error {
	requestPath := c.Request().URL.Path

	// 提取 assets 子路径
	parts := strings.SplitN(requestPath, "/assets/", 2)
	if len(parts) != 2 {
		return echo.NewHTTPError(http.StatusBadRequest, "无效资源路径")
	}
	assetPath := parts[1]

	themeName := am.ThemeManager.activeThemeName

	// 构建查找路径列表
	searchPaths := []string{
		filepath.Join("themes", themeName, "assets", assetPath),
		filepath.Join("themes", "default", "assets", assetPath),
	}

	for _, p := range searchPaths {
		if utils.FileExists(p) {
			return c.File(p)
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, "资源不存在")
}
