// Package utils 提供图形验证码生成工具
// 创建者：Done-0
// 创建时间：2025-05-10
package utils

import (
	"fmt"

	"github.com/mojocn/base64Captcha"
)

const (
	CAPTCHA_SOURCE = "346789ABCDEFGHJKLMNPQRTUVWXY" // 验证码字符源（移除易混淆字符：0O、1I、5S、2Z）
	FONT_FILE      = "wqy-microhei.ttc"             // 字体文件
	IMG_HEIGHT     = 80                             // 验证码图片高度
	IMG_WIDTH      = 200                            // 验证码图片宽度
	NOISE_COUNT    = 0                              // 干扰点数量
	CAPTCHA_LENGTH = 4                              // 验证码字符长度
)

var store = base64Captcha.DefaultMemStore

// GenImgVerificationCode 生成图形验证码
// 返回值：
//   - string: 图片验证码的 Base64 编码
//   - string: 验证码答案
//   - error: 生成过程中的错误
func GenImgVerificationCode() (string, string, error) {
	driver := createDriver()
	captcha := base64Captcha.NewCaptcha(driver, store)
	_, content, answer := captcha.Driver.GenerateIdQuestionAnswer()
	item, err := captcha.Driver.DrawCaptcha(content)
	if err != nil {
		return "", "", fmt.Errorf("生成图形验证码失败: %v", err)
	}
	return item.EncodeB64string(), answer, nil
}

// createDriver 创建验证码的驱动配置
// 返回值：
//   - *base64Captcha.DriverString: 验证码驱动对象
func createDriver() *base64Captcha.DriverString {
	return &base64Captcha.DriverString{
		Height:          IMG_HEIGHT,
		Width:           IMG_WIDTH,
		NoiseCount:      NOISE_COUNT,
		ShowLineOptions: base64Captcha.OptionShowSineLine,
		Length:          CAPTCHA_LENGTH,
		Source:          CAPTCHA_SOURCE,
		Fonts:           []string{FONT_FILE},
	}
}
