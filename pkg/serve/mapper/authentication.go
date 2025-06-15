package mapper

import (
	"fmt"

	"github.com/labstack/echo/v4"

	account "jank.com/jank_blog/internal/model/account"
	"jank.com/jank_blog/internal/utils"
)

// GetOnePermissionByKey 根据权限标识获取权限
// 参数：
//   - c: Echo 上下文
//   - key: 权限标识
//
// 返回值：
//   - *account.Permission: 权限对象
//   - error: 操作过程中的错误
func GetOnePermissionByKey(c echo.Context, key string) (*account.Permission, error) {
	db := utils.GetDBFromContext(c)
	var permission account.Permission
	if err := db.Where("key = ?", key).First(&permission).Error; err != nil {
		return nil, fmt.Errorf("获取权限失败: %v", err)
	}
	return &permission, nil
}

// CheckUserHasPermission 检查用户是否拥有指定权限
// 参数：
//   - c: Echo 上下文
//   - accountID: 用户 ID
//   - permissionID: 权限 ID
//
// 返回值：
//   - bool: 是否拥有权限
//   - error: 操作过程中的错误
func CheckUserHasPermission(c echo.Context, accountID int64, permissionID int64) (bool, error) {
	db := utils.GetDBFromContext(c)

	query := `
		SELECT COUNT(*) 
		FROM account_roles ar 
		JOIN role_permissions rp ON ar.role_id = rp.role_id 
		WHERE ar.account_id = ? 
		AND rp.permission_id = ?
	`

	var count int64
	err := db.Raw(query, accountID, permissionID).Count(&count).Error

	if err != nil {
		return false, fmt.Errorf("检查用户权限失败: %v", err)
	}

	return count > 0, nil
}

// GetPermissionsByAccountID 获取用户的所有权限
// 参数：
//   - c: Echo 上下文
//   - accountID: 用户 ID
//
// 返回值：
//   - []*account.Permission: 权限列表
//   - error: 操作过程中的错误
func GetPermissionsByAccountID(c echo.Context, accountID int64) ([]*account.Permission, error) {
	db := utils.GetDBFromContext(c)
	var permissions []*account.Permission

	// 通过用户-角色-权限关系查询用户的所有权限
	query := `
		SELECT DISTINCT p.* 
		FROM permissions p 
		JOIN role_permissions rp ON p.id = rp.permission_id 
		JOIN account_roles ar ON rp.role_id = ar.role_id 
		WHERE ar.account_id = ? 
		AND CAST(p.status AS INTEGER) = 1
		AND CAST(p.deleted AS INTEGER) = 0
	`

	err := db.Raw(query, accountID).Scan(&permissions).Error

	if err != nil {
		return nil, fmt.Errorf("获取用户权限失败: %v", err)
	}

	return permissions, nil
}
