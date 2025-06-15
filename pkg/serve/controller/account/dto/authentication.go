// Package dto 提供权限验证相关的数据传输对象定义
// 创建者：Done-0
// 创建时间：2025-06-16
package dto

// CheckAccountPermissionRequest 检查用户权限请求结构
// @Description 检查用户是否拥有特定权限的请求结构
// @Param   AccountID    int64  "用户 ID"
// @Param   PermissionKey string "权限标识，如 'custom:export_data', 'get:/user/profile'"
type CheckAccountPermissionRequest struct {
	AccountID     int64  `json:"account_id,string" xml:"account_id" form:"account_id" query:"account_id" validate:"required"`
	PermissionKey string `json:"permission_key" xml:"permission_key" form:"permission_key" query:"permission_key" validate:"required"`
}

// GetAccountPermissionsRequest 获取用户所有权限请求结构
// @Description 获取用户所有权限的请求结构
// @Param   AccountID    int64  "用户 ID"
type GetAccountPermissionsRequest struct {
	AccountID int64 `json:"account_id,string" xml:"account_id" form:"account_id" query:"account_id" validate:"required"`
}
