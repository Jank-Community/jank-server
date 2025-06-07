// Package verification 提供验证码相关的业务逻辑处理
// 创建者：Done-0
// 创建时间：2025-05-10
package verification

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"

	bizErr "jank.com/jank_blog/internal/error"
	"jank.com/jank_blog/internal/global"
	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/serve/controller/verification/dto"
	"jank.com/jank_blog/pkg/vo/verification"
)

const (
	EMAIL_VERIFICATION_CODE_CACHE_EXPIRATION = 3 * time.Minute // 邮箱验证码缓存过期时间
	IMG_VERIFICATION_CODE_CACHE_EXPIRATION   = 3 * time.Minute // 图形验证码缓存过期时间
)

// GenerateImgVerificationCode 生成图形验证码
// 参数：
//   - c: Echo 上下文
//   - email: 邮箱地址
//
// 返回值：
//   - *verification.ImgVerificationVO: 图形验证码VO
//   - error: 操作过程中的错误
func GenerateImgVerificationCode(c echo.Context, req *dto.GetOneVerificationCode) (*verification.ImgVerificationVO, error) {
	key := utils.IMG_VERIFICATION_CODE_CACHE_PREFIX + req.Email

	// 生成单个图形验证码
	imgBase64, answer, err := utils.GenImgVerificationCode()
	if err != nil {
		utils.BizLogger(c).Errorf("生成图片验证码失败: %v", err)
		return nil, err
	}

	err = global.RedisClient.Set(context.Background(), key, answer, IMG_VERIFICATION_CODE_CACHE_EXPIRATION).Err()
	if err != nil {
		utils.BizLogger(c).Errorf("图形验证码写入缓存失败，key: %v, 错误: %v", key, err)
		return nil, err
	}

	return &verification.ImgVerificationVO{ImgBase64: imgBase64}, nil
}

// SendEmailVerificationCode 发送邮箱验证码
// 参数：
//   - c: Echo 上下文
//   - email: 邮箱地址
//
// 返回值：
//   - error: 操作过程中的错误
func SendEmailVerificationCode(c echo.Context, req *dto.GetOneVerificationCode) error {
	key := utils.EMAIL_VERIFICATION_CODE_CACHE_KEY_PREFIX + req.Email

	// 检查验证码是否存在
	exists, err := global.RedisClient.Exists(context.Background(), key).Result()
	if err != nil {
		utils.BizLogger(c).Errorf("检查邮箱验证码是否有效失败: %v", err)
		return err
	}
	if exists > 0 {
		return bizErr.New(bizErr.SERVER_ERR, "邮箱验证码已存在")
	}

	// 生成并缓存验证码
	code := utils.NewRand()
	err = global.RedisClient.Set(context.Background(), key, strconv.Itoa(code), EMAIL_VERIFICATION_CODE_CACHE_EXPIRATION).Err()
	if err != nil {
		utils.BizLogger(c).Errorf("邮箱验证码写入缓存失败: %v", err)
		return err
	}

	// 发送验证码邮件
	expirationInMinutes := int(EMAIL_VERIFICATION_CODE_CACHE_EXPIRATION.Round(time.Minute).Minutes())
	emailContent := fmt.Sprintf("您的注册验证码是: %d , 有效期为 %d 分钟。", code, expirationInMinutes)
	success, err := utils.SendEmail(emailContent, []string{req.Email})
	if !success {
		utils.BizLogger(c).Errorf("邮箱验证码发送失败，邮箱地址: %s, 错误: %v", req.Email, err)
		global.RedisClient.Del(context.Background(), key)
		return bizErr.New(bizErr.SEND_EMAIL_VERIFICATION_CODE_FAIL, err.Error())
	}

	return nil
}
