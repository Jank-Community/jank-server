// Package model 提供角色与权限数据模型定义
// 创建者：Done-0
// 创建时间：2025-06-12
package model

import "jank.com/jank_blog/internal/model/base"

// RolePermission 角色和权限关联表
type RolePermission struct {
	base.Base
	RoleID       int64 `gorm:"index" json:"role_id"`       // 角色ID
	PermissionID int64 `gorm:"index" json:"permission_id"` // 权限ID
}

// TableName 指定表名
// 返回值：
//   - string: 表名
func (RolePermission) TableName() string {
	return "role_permissions"
}
