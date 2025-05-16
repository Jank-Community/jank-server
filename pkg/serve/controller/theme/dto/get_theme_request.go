// Package dto 提供主题相关的数据传输对象
// 创建者：Done-0
// 创建时间：2025-05-14
package dto

// GetThemeRequest 获取主题请求
type GetThemeRequest struct {
	ID int64 `json:"id" xml:"id" form:"id" query:"id" validate:"required"` // 主题ID
}

// GetThemeFileRequest 获取主题文件请求
type GetThemeFileRequest struct {
	ID   int64  `json:"id" xml:"id" form:"id" query:"id" validate:"required"`         // 主题ID，为空时使用当前激活的主题
	Path string `json:"path" xml:"path" form:"path" query:"path" validate:"required"` // 文件路径
}
