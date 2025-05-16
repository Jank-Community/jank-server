// Package theme 提供主题相关视图对象
// 创建者：Done-0
// 创建时间：2025-05-14
package theme

import "time"

// ThemeVO 主题视图对象
type ThemeVO struct {
	ID          int64     `json:"id"`          // 主题ID
	Name        string    `json:"name"`        // 主题名称
	Description string    `json:"description"` // 主题描述
	Logo        string    `json:"logo"`        // 主题Logo图片路径
	Website     string    `json:"website"`     // 主题官网URL
	Version     string    `json:"version"`     // 主题版本
	Author      string    `json:"author"`      // 主题作者
	FolderName  string    `json:"folderName"`  // 主题文件夹名称
	Activated   bool      `json:"activated"`   // 是否激活
	ThemePath   string    `json:"themePath"`   // 主题路径
	Screenshots string    `json:"screenshots"` // 主题截图
	CreatedAt   time.Time `json:"createdAt"`   // 创建时间
	UpdatedAt   time.Time `json:"updatedAt"`   // 更新时间
}

// ThemeFileVO 主题文件视图对象
type ThemeFileVO struct {
	Name     string         `json:"name"`     // 文件名
	Path     string         `json:"path"`     // 文件路径
	IsFile   bool           `json:"isFile"`   // 是否为文件
	Editable bool           `json:"editable"` // 是否可编辑
	Children []*ThemeFileVO `json:"children"` // 子文件/文件夹
}

// ThemeConfigItemVO 主题配置项视图对象
type ThemeConfigItemVO struct {
	Name         string      `json:"name"`         // 配置项名称
	Label        string      `json:"label"`        // 配置项标签
	Type         string      `json:"type"`         // 配置项类型
	DataType     string      `json:"dataType"`     // 数据类型
	DefaultValue interface{} `json:"defaultValue"` // 默认值
	Options      []string    `json:"options"`      // 可选项（当类型为select时）
	Description  string      `json:"description"`  // 描述
}

// ThemeConfigGroupVO 主题配置组视图对象
type ThemeConfigGroupVO struct {
	Name  string               `json:"name"`  // 配置组名称
	Label string               `json:"label"` // 配置组标签
	Items []*ThemeConfigItemVO `json:"items"` // 配置项列表
}
