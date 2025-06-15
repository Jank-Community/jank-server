// Package routes 提供路由注册功能
// 创建者：Done-0
// 创建时间：2025-05-10
package routes

import (
	"github.com/labstack/echo/v4"

	auth_middleware "jank.com/jank_blog/internal/middleware/auth"
	"jank.com/jank_blog/pkg/serve/controller/account"
)

// RegisterAccountRoutes 注册账户相关路由
// 参数：
//   - r: Echo 路由组数组，r[0] 为 API v1 版本组
func RegisterAccountRoutes(r ...*echo.Group) {
	// api v1 group
	apiV1 := r[0]
	accountGroupV1 := apiV1.Group("/account")
	accountGroupV1.POST("/registerAccount", account.RegisterAcc)
	accountGroupV1.POST("/loginAccount", account.LoginAccount)
	accountGroupV1.GET("/getAccount", account.GetAccount, auth_middleware.AuthMiddleware())
	accountGroupV1.POST("/updateAccount", account.UpdateAccount, auth_middleware.AuthMiddleware())
	accountGroupV1.POST("/logoutAccount", account.LogoutAccount, auth_middleware.AuthMiddleware())
	accountGroupV1.POST("/resetPassword", account.ResetPassword, auth_middleware.AuthMiddleware())

	// 角色管理
	roleGroupV1 := apiV1.Group("/role", auth_middleware.AuthMiddleware())
	roleGroupV1.POST("/createOneRole", account.CreateOneRole)
	roleGroupV1.POST("/updateOneRole", account.UpdateOneRole)
	roleGroupV1.POST("/deleteOneRole", account.DeleteOneRole)
	roleGroupV1.GET("/listAllRoles", account.ListAllRoles)

	// 权限管理
	permissionGroup := apiV1.Group("/permission", auth_middleware.AuthMiddleware())
	permissionGroup.POST("/createOnePermission", account.CreateOnePermission)
	permissionGroup.POST("/updateOnePermission", account.UpdateOnePermission)
	permissionGroup.POST("/deleteOnePermission", account.DeleteOnePermission)
	permissionGroup.GET("/listAllPermissions", account.ListAllPermissions)

	// 用户角色管理
	accountRoleGroup := apiV1.Group("/account-role", auth_middleware.AuthMiddleware())
	accountRoleGroup.POST("/assignRoles", account.AssignRolesToOneAccount)
	accountRoleGroup.POST("/revokeRoles", account.RevokeRolesFromOneAccount)
	accountRoleGroup.GET("/getAccountRoles", account.GetAccountRolesFromOneAccount)

	// 角色权限关联
	RolePermissionGroup := apiV1.Group("/role-permission", auth_middleware.AuthMiddleware())
	RolePermissionGroup.POST("/assignPermissions", account.AssignPermissionsToOneRole)
	RolePermissionGroup.POST("/revokePermissions", account.RevokePermissionsFromOneRole)
	RolePermissionGroup.GET("/getRolePermissions", account.GetRolePermissionsFromOneRole)

	// 权限验证
	authenticationGroup := apiV1.Group("/authentication", auth_middleware.AuthMiddleware())
	authenticationGroup.POST("/checkPermission", account.CheckAccountPermission)
	authenticationGroup.GET("/getAccountPermissions", account.GetAccountPermissions)
}
