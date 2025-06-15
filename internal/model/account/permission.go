// Package model 提供角色权限数据模型定义
// 创建者：Done-0
// 创建时间：2025-06-12
package model

import "jank.com/jank_blog/internal/model/base"

// Permission 权限模型
type Permission struct {
	base.Base
	Key         string `gorm:"type:varchar(255);not null" json:"key"`              // 权限标识，例如 get:/user/{id} 或 custom:permission_name
	Name        string `gorm:"type:varchar(32);not null" json:"name"`              // 权限名称，如 '读取', '写入', '删除'
	Description string `gorm:"type:varchar(255);default: null" json:"description"` // 权限描述
	Status      bool   `gorm:"default:true" json:"status"`                         // 权限状态，true(1)表示启用，false(0)表示禁用
}

// TableName 指定表名
// 返回值：
//   - string: 表名
func (Permission) TableName() string {
	return "permissions"
}
