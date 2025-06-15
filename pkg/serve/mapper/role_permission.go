package mapper

import (
	"fmt"

	"github.com/labstack/echo/v4"

	account "jank.com/jank_blog/internal/model/account"
	"jank.com/jank_blog/internal/utils"
)

// AssignPermissionsToOneRoleByIDs 为角色分配权限
// 参数：
//   - c: Echo 上下文
//   - rolePermission: 角色权限模型
//
// 返回值：
//   - error: 操作过程中的错误
func AssignPermissionsToOneRoleByIDs(c echo.Context, rolePermission *account.RolePermission) error {
	db := utils.GetDBFromContext(c)
	if err := db.Create(rolePermission).Error; err != nil {
		return fmt.Errorf("为角色分配权限失败: %v", err)
	}
	return nil
}

// RevokePermissionFromOneRole 从角色撤销权限
// 参数：
//   - c: Echo 上下文
//   - roleID: 角色 ID
//   - permissionID: 权限 ID
//
// 返回值：
//   - error: 操作过程中的错误
func RevokePermissionFromOneRole(c echo.Context, roleID int64, permissionID int64) error {
	db := utils.GetDBFromContext(c)
	if err := db.Where("role_id = ? AND permission_id = ?", roleID, permissionID).Delete(&account.RolePermission{}).Error; err != nil {
		return fmt.Errorf("从角色撤销权限失败: %v", err)
	}
	return nil
}

// GetPermissionsByRoleID 获取角色的所有权限
// 参数：
//   - c: Echo 上下文
//   - roleID: 角色 ID
//
// 返回值：
//   - []*account.Permission: 权限列表
//   - error: 操作过程中的错误
func GetPermissionsByRoleID(c echo.Context, roleID int64) ([]*account.Permission, error) {
	db := utils.GetDBFromContext(c)
	var permissions []*account.Permission

	query := `
		SELECT p.* 
		FROM permissions p 
		JOIN role_permissions rp ON p.id = rp.permission_id 
		WHERE rp.role_id = ? 
		AND CAST(p.deleted AS INTEGER) = 0
	`

	err := db.Raw(query, roleID).Scan(&permissions).Error

	if err != nil {
		return nil, fmt.Errorf("获取角色权限失败: %v", err)
	}
	return permissions, nil
}
