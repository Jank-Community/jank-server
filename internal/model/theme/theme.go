// Package model 提供主题数据模型定义
// 创建者：Done-0
// 创建时间：2025-05-14
package model

import (
	"jank.com/jank_blog/internal/model/base"
)

// Theme 主题模型
type Theme struct {
	base.Base
	Name        string `gorm:"type:varchar(255);not null" json:"name"`               // 主题名称
	Description string `gorm:"type:text" json:"description"`                         // 主题描述
	Logo        string `gorm:"type:varchar(1024)" json:"logo"`                       // 主题 Logo 图片路径
	Website     string `gorm:"type:varchar(1024)" json:"website"`                    // 主题官网 URL
	Version     string `gorm:"type:varchar(50)" json:"version"`                      // 主题版本
	Author      string `gorm:"type:varchar(255)" json:"author"`                      // 主题作者
	FolderName  string `gorm:"type:varchar(255);not null" json:"folder_name"`         // 主题文件夹名称
	Activated   bool   `gorm:"type:boolean;not null;default:false" json:"activated"` // 是否激活
	ThemePath   string `gorm:"type:varchar(1024);not null" json:"theme_path"`         // 主题路径
	Screenshots string `gorm:"type:varchar(1024)" json:"screenshots"`                // 主题截图
}

// TableName 指定表名
// 返回值：
//   - string: 表名
func (Theme) TableName() string {
	return "themes"
}

// ThemeSetting 主题设置模型
type ThemeSetting struct {
	base.Base
	ThemeID      int64  `gorm:"type:bigint;not null;index" json:"theme_id"`    // 主题ID
	SettingKey   string `gorm:"type:varchar(255);not null" json:"setting_key"` // 设置键
	SettingValue string `gorm:"type:text" json:"setting_value"`                // 设置值
}

// TableName 指定表名
// 返回值：
//   - string: 表名
func (ThemeSetting) TableName() string {
	return "theme_settings"
}

// ThemeFile 主题文件结构体
type ThemeFile struct {
	Name     string       // 文件名
	Path     string       // 文件路径
	IsFile   bool         // 是否为文件
	Editable bool         // 是否可编辑
	Children []*ThemeFile // 子文件或目录
}
