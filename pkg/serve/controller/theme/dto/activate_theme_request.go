// Package dto 提供主题相关的数据传输对象
// 创建者：Done-0
// 创建时间：2025-05-14
package dto

// ActivateThemeRequest 激活主题请求
type ActivateThemeRequest struct {
	ID int64 `json:"id" xml:"id" form:"id" query:"id" validate:"required"` // 主题ID
}
