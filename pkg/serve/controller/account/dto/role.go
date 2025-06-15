// Package dto 提供角色相关的数据传输对象定义
// 创建者：Done-0
// 创建时间：2025-06-12
package dto

// CreateOneRoleRequest 创建角色请求结构
// @Description 创建角色时的请求结构
// @Param   Name        string "角色名称"
// @Param   Description string "角色描述"
// @Param   Status      bool   "角色状态，true(1)表示启用，false(0)表示禁用"
type CreateOneRoleRequest struct {
	Name        string `json:"name" xml:"name" form:"name" query:"name" validate:"required"`
	Description string `json:"description" xml:"description" form:"description" query:"description" validate:"required"`
	Status      bool   `json:"status" xml:"status" form:"status" query:"status" validate:"required"`
}

// UpdateOneRoleRequest 更新角色请求结构
// @Description 更新角色时的请求结构
// @Param   ID          int64  "角色 ID"
// @Param   Name        string "角色名称"
// @Param   Description string "角色描述"
// @Param   Status      bool   "角色状态，true(1)表示启用，false(0)表示禁用"
type UpdateOneRoleRequest struct {
	ID          int64  `json:"id,string" xml:"id" form:"id" query:"id" validate:"required"`
	Name        string `json:"name" xml:"name" form:"name" query:"name" validate:"required"`
	Description string `json:"description" xml:"description" form:"description" query:"description" validate:"required"`
	Status      bool   `json:"status" xml:"status" form:"status" query:"status" validate:"required"`
}

// DeleteOneRoleRequest 删除角色请求结构
// @Description 删除角色时的请求结构
// @Param   ID   		int64 "角色 ID"
type DeleteOneRoleRequest struct {
	ID int64 `json:"id,string" xml:"id" form:"id" query:"id" validate:"required"`
}
