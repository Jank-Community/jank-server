// Package model 提供用户与角色数据模型定义
// 创建者：Done-0
// 创建时间：2025-06-12
package model

import "jank.com/jank_blog/internal/model/base"

// AccountRole 用户与角色关联表
type AccountRole struct {
	base.Base
	AccountID int64 `gorm:"index" json:"account_id"` // 用户ID
	RoleID    int64 `gorm:"index" json:"role_id"`    // 角色ID
}

// TableName 指定表名
// 返回值：
//   - string: 表名
func (AccountRole) TableName() string {
	return "account_roles"
}
