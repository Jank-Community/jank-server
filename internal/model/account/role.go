// Package model 提供用户角色数据模型定义
// 创建者：Done-0
// 创建时间：2025-06-12
package model

import (
	"jank.com/jank_blog/internal/model/base"
)

// Role 用户角色模型
type Role struct {
	base.Base
	Name        string `gorm:"type:varchar(32);not null" json:"name"`              // 角色名称，如 '管理员', '用户'
	Description string `gorm:"type:varchar(255);default: null" json:"description"` // 角色描述
	Status      bool   `gorm:"default:true" json:"status"`                         // 角色状态，true(1)表示启用，false(0)表示禁用
}

// TableName 指定表名
// 返回值：
//   - string: 表名
func (Role) TableName() string {
	return "roles"
}
