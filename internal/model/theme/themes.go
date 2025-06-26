// Package model 提供主题模型定义
// 创建者：Pey121
// 创建时间：2025-05-30
package model

import (
	"jank.com/jank_blog/internal/model/base"
)

// Theme 主题模型
type Theme struct {
	base.Base          // 继承基础模型
	ThemeID     string `gorm:"unique;not null;size:100;index" json:"theme_id"` // 主题唯一标识
	Name        string `gorm:"not null;size:100" json:"name"`                  // 主题名称
	Description string `gorm:"type:text" json:"description"`                   // 主题描述
	Version     string `gorm:"size:20" json:"version"`                         // 主题版本
	Author      string `gorm:"size:100" json:"author"`                         // 主题作者
	AuthorURI   string `gorm:"size:255" json:"author_uri"`                     // 作者网址
	ThemeURI    string `gorm:"size:255" json:"theme_uri"`                      // 主题网址
	Screenshot  string `gorm:"size:255" json:"screenshot"`                     // 主题截图
	Tags        string `gorm:"size:500" json:"tags"`                           // 主题标签(JSON数组)
	IsActive    bool   `gorm:"default:false;index" json:"is_active"`           // 是否为当前激活主题
	IsInstalled bool   `gorm:"default:false" json:"is_installed"`              // 是否已安装
	InstallPath string `gorm:"size:255" json:"install_path"`                   // 安装路径
	Status      string `gorm:"size:20;default:'inactive';index" json:"status"` // 主题状态: active, inactive, error
	SortOrder   int    `gorm:"default:0" json:"sort_order"`                    // 排序顺序
}

// ThemeConfig 主题配置
type ThemeConfig struct {
	base.Base          // 继承基础模型
	ThemeID     string `gorm:"not null;size:100;index" json:"theme_id"` // 主题ID
	ConfigKey   string `gorm:"not null;size:100" json:"config_key"`     // 配置键名
	ConfigValue string `gorm:"type:text" json:"config_value"`           // 配置值(JSON格式)
	ConfigType  string `gorm:"size:50" json:"config_type"`              // 配置类型: string, number, boolean, array, object
	IsDefault   bool   `gorm:"default:false" json:"is_default"`         // 是否为默认配置
	Description string `gorm:"type:text" json:"description"`            // 配置描述
	SortOrder   int    `gorm:"default:0" json:"sort_order"`             // 排序顺序
}

// ThemeOption 主题选项 - 存储用户自定义的主题配置
type ThemeOption struct {
	base.Base          // 继承基础模型
	ThemeID     string `gorm:"not null;size:100;index" json:"theme_id"` // 主题ID
	UserID      int64  `gorm:"index" json:"user_id"`                    // 用户ID (0表示全局配置)
	OptionName  string `gorm:"not null;size:100" json:"option_name"`    // 选项名称
	OptionValue string `gorm:"type:text" json:"option_value"`           // 选项值(JSON格式)
	OptionType  string `gorm:"size:50" json:"option_type"`              // 选项类型
	IsGlobal    bool   `gorm:"default:true" json:"is_global"`           // 是否为全局配置
}

// TableName 指定表名
// 返回值：
//   - string: 表名
func (Theme) TableName() string {
	return "themes"
}

// TableName 指定表名
// 返回值：
//   - string: 表名
func (ThemeConfig) TableName() string {
	return "theme_configs"
}

// TableName 指定表名
// 返回值：
//   - string: 表名
func (ThemeOption) TableName() string {
	return "theme_options"
}
