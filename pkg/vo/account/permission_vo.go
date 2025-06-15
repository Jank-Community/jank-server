// Package account 提供权限相关的视图对象定义
// 创建者：Done-0
// 创建时间：2025-06-12
package account

// PermissionVO 权限返回值对象
// @Description 权限信息的返回结构
// @Property   id           int64  "角色 ID"
// @Property   key          string "权限标识，如 'custom:export_data', 'get:/user/profile'"
// @Property   name         string "角色名称"
// @Property   description  string "角色描述"
// @Property   status       bool "角色状态"
type PermissionVO struct {
	ID          int64  `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}
