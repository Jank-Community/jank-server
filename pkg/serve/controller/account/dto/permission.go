// Package dto 提供权限相关的数据传输对象定义
// 创建者：Done-0
// 创建时间：2025-06-12
package dto

// CreateOnePermissionRequest 创建权限请求结构
// @Description 创建权限时的请求结构
// @Param   Key         string "权限标识，例如 get:/user/{id} 或 custom:permission_name"
// @Param   Name        string "权限名称
// @Param   Description string "权限描述"
// @Param   Status      bool   "权限状态，true(1)表示启用，false(0)表示禁用"
type CreateOnePermissionRequest struct {
	Key         string `json:"key" xml:"key" form:"key" query:"key" validate:"required"`
	Name        string `json:"name" xml:"name" form:"name" query:"name" validate:"required"`
	Description string `json:"description" xml:"description" form:"description" query:"description" validate:"required"`
	Status      bool   `json:"status" xml:"status" form:"status" query:"status" validate:"required"`
}

// UpdateOnePermissionRequest 更新权限请求结构
// @Description 更新权限时的请求结构
// @Param   ID          int64  "权限 ID"
// @Param   Key         string "权限标识，例如 get:/user/{id} 或 custom:permission_name"
// @Param   Name        string "权限名称
// @Param   Description string "权限描述"
// @Param   Status      bool   "权限状态，true(1)表示启用，false(0)表示禁用"
type UpdateOnePermissionRequest struct {
	ID          int64  `json:"id,string" xml:"id" form:"id" query:"id" validate:"required"`
	Key         string `json:"key" xml:"key" form:"key" query:"key" validate:"required"`
	Name        string `json:"name" xml:"name" form:"name" query:"name" validate:"required"`
	Description string `json:"description" xml:"description" form:"description" query:"description" validate:"required"`
	Status      bool   `json:"status" xml:"status" form:"status" query:"status" validate:"required"`
}

// DeleteOnePermissionRequest 删除权限请求结构
// @Description 删除权限时的请求结构
// @Param   ID int64 "权限 ID"
type DeleteOnePermissionRequest struct {
	ID int64 `json:"id,string" xml:"id" form:"id" query:"id" validate:"required"`
}
