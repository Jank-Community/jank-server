// Package account 提供角色权限相关的视图对象定义
// 创建者：Done-0
// 创建时间：2025-06-12
package account

// RolePermissionVO 角色权限返回值对象
// @Description 角色权限信息的返回结构
// @Property   role_id        int64            "角色ID"
// @Property   role_name      string           "角色名称"
// @Property   permissions    []PermissionVO   "权限列表"
type RolePermissionVO struct {
	RoleID      int64          `json:"role_id"`
	RoleName    string         `json:"role_name"`
	Permissions []PermissionVO `json:"permissions"`
}
