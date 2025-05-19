// Package utils 提供邮件操作相关工具
// 创建者：Done-0
// 创建时间：2025-05-10
package utils

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"regexp"
	"time"

	"gopkg.in/gomail.v2"

	"jank.com/jank_blog/configs"
	"jank.com/jank_blog/internal/global"
)

// 邮件相关常量
const (
	EMAIL_SUBJECT = "【Jank Blog】注册验证码"
)

// 邮箱服务器配置
var emailServers = map[string]struct {
	Server string
	Port   int
	SSL    bool
}{
	"qq":      {"smtp.qq.com", 465, true},         // QQ 邮箱使用 SSL 加密
	"gmail":   {"smtp.gmail.com", 465, true},      // Gmail 使用 SSL 加密
	"outlook": {"smtp.office365.com", 587, false}, // Outlook 使用 TLS 加密
}

// SendEmail 发送邮件到指定邮箱
func SendEmail(content string, toEmails []string) (bool, error) {
	config, err := configs.LoadConfig()
	if err != nil {
		global.SysLog.Errorf("加载邮件配置失败: %v", err)
		return false, fmt.Errorf("加载邮件配置失败: %v", err)
	}

	// 获取SMTP配置
	emailType := config.AppConfig.Email.EmailType
	serverConfig := emailServers[emailType]

	// 创建邮件
	m := gomail.NewMessage()
	m.SetHeader("From", config.AppConfig.Email.FromEmail)
	m.SetHeader("To", toEmails...)
	m.SetHeader("Subject", EMAIL_SUBJECT)
	m.SetBody("text/plain", content)

	// 配置发送器
	d := gomail.NewDialer(
		serverConfig.Server,
		serverConfig.Port,
		config.AppConfig.Email.FromEmail,
		config.AppConfig.Email.EmailSmtp,
	)

	// 根据端口配置安全选项
	if serverConfig.SSL {
		d.SSL = true
	} else {
		d.TLSConfig = &tls.Config{
			ServerName: serverConfig.Server,
			MinVersion: tls.VersionTLS12,
		}
	}

	if err := d.DialAndSend(m); err != nil {
		global.SysLog.Errorf("发送邮件失败: %v", err)
		return false, fmt.Errorf("发送邮件失败: %v", err)
	}

	return true, nil
}

// NewRand 生成六位数随机验证码
func NewRand() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(900000) + 100000
}

// ValidEmail 检查邮箱格式是否有效
func ValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(pattern).MatchString(email)
}
