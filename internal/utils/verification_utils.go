// Package utils 提供通用工具函数
// 创建者：Done-0
// 创建时间：2025-05-10
package utils

import (
	"context"
	"strings"

	"github.com/labstack/echo/v4"

	"jank.com/jank_blog/internal/global"
)

const (
	EMAIL_VERIFICATION_CODE_CACHE_KEY_PREFIX = "EMAIL:VERIFICATION:CODE:"     // 邮箱验证码缓存前缀
	IMG_VERIFICATION_CODE_CACHE_PREFIX       = "IMG:VERIFICATION:CODE:CACHE:" // 图形验证码缓存前缀
)

// VerifyEmailCode 校验邮箱验证码
// 参数：
//   - c: Echo 上下文
//   - code: 验证码
//   - email: 邮箱地址
//
// 返回值：
//   - bool: 验证成功返回 true，失败返回 false
func VerifyEmailCode(c echo.Context, code, email string) bool {
	return VerifyCode(c, code, email, EMAIL_VERIFICATION_CODE_CACHE_KEY_PREFIX)
}

// VerifyImgCode 校验图形验证码
// 参数：
//   - c: Echo 上下文
//   - code: 验证码
//   - email: 邮箱地址
//
// 返回值：
//   - bool: 验证成功返回 true，失败返回 false
func VerifyImgCode(c echo.Context, code, email string) bool {
	return VerifyCode(c, code, email, IMG_VERIFICATION_CODE_CACHE_PREFIX)
}

// VerifyCode 通用验证码校验
// 参数：
//   - c: Echo 上下文
//   - code: 验证码
//   - email: 邮箱地址
//   - prefix: 缓存键前缀
//
// 返回值：
//   - bool: 验证成功返回 true，失败返回 false
func VerifyCode(c echo.Context, code, email, prefix string) bool {
	key := prefix + email

	storedCode, err := global.RedisClient.Get(c.Request().Context(), key).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			BizLogger(c).Error("验证码不存在或已过期")
		} else {
			BizLogger(c).Errorf("验证码校验失败: %v", err)
		}
		return false
	}

	storedCode = strings.ToUpper(strings.TrimSpace(storedCode))
	code = strings.ToUpper(strings.TrimSpace(code))

	if storedCode != code {
		BizLogger(c).Error("用户验证码错误")
		return false
	}

	if err := global.RedisClient.Del(context.Background(), key).Err(); err != nil {
		BizLogger(c).Errorf("删除验证码缓存失败: %v", err)
	}

	return true
}
