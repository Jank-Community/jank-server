// Package dto 提供权限相关的数据传输对象定义
// 创建者：Done-0
// 创建时间：2025-06-12
package dto

// AssignRolesToAccountRequest 分配角色请求结构
// @Description 分配角色给账户的请求结构
// @Param   RoleIDs      []string "角色 ID 列表"
// @Param   AccountID 	 int64   "用户 ID"
type AssignRolesToAccountRequest struct {
	RoleIDs   []string `json:"role_ids" xml:"role_ids" form:"role_ids" query:"role_ids" validate:"required,dive,required"`
	AccountID int64    `json:"account_id,string" xml:"account_id" form:"account_id" validate:"required"`
}

// RevokeRolesFromOneAccountRequest 撤销角色请求结构
// @Description 撤销账户的角色请求结构
// @Param   RoleIDs      []string "角色 ID 列表"
// @Param   AccountID 	 int64   "用户 ID"
type RevokeRolesFromOneAccountRequest struct {
	RoleIDs   []string `json:"role_ids" xml:"role_ids" form:"role_ids" query:"role_ids" validate:"required,dive,required"`
	AccountID int64    `json:"account_id,string" xml:"account_id" form:"account_id" validate:"required"`
}

// GetAccountRolesRequest 获取账户角色请求结构
// @Description 获取账户角色的请求结构
// @Param   AccountID 	 int64   "用户 ID"
type GetAccountRolesRequest struct {
	AccountID int64 `json:"account_id,string" xml:"account_id" form:"account_id" query:"account_id" validate:"required"`
}
