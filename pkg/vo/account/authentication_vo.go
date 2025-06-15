// Package account 提供权限验证相关的视图对象定义
// 创建者：Done-0
// 创建时间：2025-06-16
package account

// CheckPermissionVO 权限检查返回值对象
// @Description 权限检查结果的返回结构
// @Property   has_permission  bool   "是否拥有权限"
// @Property   permission_key  string "权限标识"
type CheckPermissionVO struct {
	HasPermission bool   `json:"has_permission"`
	PermissionKey string `json:"permission_key"`
}

// AccountPermissionsVO 用户权限返回值对象
// @Description 用户权限信息的返回结构
// @Property   account_id    int64           "用户ID"
// @Property   permissions   []PermissionVO  "权限列表"
type AccountPermissionsVO struct {
	AccountID   int64          `json:"account_id"`
	Permissions []PermissionVO `json:"permissions"`
}
