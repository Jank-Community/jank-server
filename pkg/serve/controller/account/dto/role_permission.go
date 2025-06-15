// Package dto 提供角色权限相关的数据传输对象定义
// 创建者：Done-0
// 创建时间：2025-06-15
package dto

// AssignPermissionsToRoleRequest 分配权限给角色请求结构
// @Description 分配权限给角色的请求结构
// @Param   PermissionIDs  []string "权限 ID 列表"
// @Param   RoleID         int64   "角色 ID"
type AssignPermissionsToRoleRequest struct {
	PermissionIDs []string `json:"permission_ids" xml:"permission_ids" form:"permission_ids" query:"permission_ids" validate:"required,dive,required"`
	RoleID        int64    `json:"role_id,string" xml:"role_id" form:"role_id" validate:"required"`
}

// RevokePermissionsFromRoleRequest 撤销角色权限请求结构
// @Description 从角色撤销权限的请求结构
// @Param   PermissionIDs  []string "权限 ID 列表"
// @Param   RoleID         int64   "角色 ID"
type RevokePermissionsFromRoleRequest struct {
	PermissionIDs []string `json:"permission_ids" xml:"permission_ids" form:"permission_ids" query:"permission_ids" validate:"required,dive,required"`
	RoleID        int64    `json:"role_id,string" xml:"role_id" form:"role_id" validate:"required"`
}

// GetRolePermissionsRequest 获取角色权限请求结构
// @Description 获取角色所有权限的请求结构
// @Param   RoleID   int64   "角色 ID"
type GetRolePermissionsRequest struct {
	RoleID int64 `json:"role_id,string" xml:"role_id" form:"role_id" query:"role_id" validate:"required"`
}
