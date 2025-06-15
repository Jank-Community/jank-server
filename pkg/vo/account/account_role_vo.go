// Package account 提供账户角色相关的视图对象定义
// 创建者：Done-0
// 创建时间：2025-06-12
package account

// AccountRoleVO 用户角色返回值对象
// @Description 用户角色信息的返回结构
// @Property   account_id        int64  "用户ID"
// @Property   roles            []RoleVO "角色列表"
type AccountRoleVO struct {
	AccountID int64    `json:"account_id"`
	Roles     []RoleVO `json:"roles"`
}
