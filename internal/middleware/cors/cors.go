// Package cors_middleware 提供跨域资源共享中间件
// 创建者：Done-0
// 创建时间：2025-05-10
package cors_middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"jank.com/jank_blog/internal/global"
)

// InitCORS 初始化 CORS 中间件
// 返回值：
//   - echo.MiddlewareFunc: Echo 框架中间件函数
func InitCORS() echo.MiddlewareFunc {
	return corsWithConfig(defaultCORSConfig())
}

// CORSConfig 定义 CORS 中间件的配置
type corsConfig struct {
	AllowedOrigins   []string // 允许的源
	AllowedMethods   []string // 允许的方法
	AllowedHeaders   []string // 允许的头部
	AllowCredentials bool     // 是否允许携带证书
}

// DefaultCORSConfig 提供了默认的 CORS 配置
// 返回值：
//   - corsConfig: CORS 配置
func defaultCORSConfig() corsConfig {
	return corsConfig{
		AllowedOrigins: []string{"*"},                                       // 默认允许所有域名
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // 默认允许的请求方法
		AllowedHeaders: []string{
			// 标准 HTTP 请求头
			"Accept",          // 接受的响应类型
			"Accept-Encoding", // 接受的编码方式
			"Accept-Language", // 接受的语言
			"Connection",      // 连接类型
			"Content-Length",  // 内容长度
			"Content-Type",    // 内容类型
			"Host",            // 主机信息
			"Origin",          // 请求源
			"Referer",         // 来源页面
			"User-Agent",      // 用户代理

			// 认证与安全
			"Authorization", // 认证信息
			"Refresh-Token", // 刷新令牌
			"Cookie",        // Cookie 信息

			// 缓存控制
			"Cache-Control",       // 缓存控制
			"Pragma",              // 缓存指令
			"If-Match",            // 条件请求
			"If-None-Match",       // 条件请求
			"If-Modified-Since",   // 条件请求
			"If-Unmodified-Since", // 条件请求

			// 业务与追踪
			"Language",         // 语言偏好
			"Timezone",         // 时区信息
			"Platform",         // 平台信息
			"App-Version",      // 应用版本
			"Request-Id",       // 请求追踪
			"X-Requested-With", // AJAX 请求标识
		},
		AllowCredentials: false, // 默认不允许携带证书
	}
}

// corsWithConfig 返回一个 CORS 中间件函数
// 参数：
//   - config: CORS配置
//
// 返回值：
//   - echo.MiddlewareFunc: Echo 框架中间件函数
func corsWithConfig(config corsConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Access-Control-Allow-Origin", strings.Join(config.AllowedOrigins, ","))
			c.Response().Header().Set("Access-Control-Allow-Methods", strings.Join(config.AllowedMethods, ","))
			c.Response().Header().Set("Access-Control-Allow-Headers", strings.Join(config.AllowedHeaders, ","))

			if config.AllowCredentials {
				c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
			}

			// 处理预检请求，缓存预检请求结果 24 小时
			if c.Request().Method == "OPTIONS" {
				c.Set("Access-Control-Max-Age", "86400")
				return c.NoContent(http.StatusNoContent)
			}

			// 记录 CORS 请求
			if global.BizLog != nil {
				global.BizLog.Info("CORS request",
					"method", c.Request().Method,
					"path", c.Request().URL.Path,
					"origin", c.Request().Header.Get("Origin"),
				)
			}

			return next(c)
		}
	}
}
