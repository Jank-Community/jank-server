package mapper

import (
	"fmt"

	"github.com/labstack/echo/v4"

	account "jank.com/jank_blog/internal/model/account"
	"jank.com/jank_blog/internal/utils"
)

// CreateOnePermission 创建权限
// 参数：
//   - c: Echo 上下文
//   - permission: 权限信息
//
// 返回值：
//   - error: 操作过程中的错误
func CreateOnePermission(c echo.Context, permission *account.Permission) error {
	db := utils.GetDBFromContext(c)
	if err := db.Create(permission).Error; err != nil {
		return fmt.Errorf("创建权限失败: %v", err)
	}
	return nil
}

// UpdateOnePermissionByID 更新权限
// 参数：
//   - c: Echo 上下文
//   - permission: 权限信息
//
// 返回值：
//   - error: 操作过程中的错误
func UpdateOnePermissionByID(c echo.Context, permission *account.Permission) error {
	db := utils.GetDBFromContext(c)
	if err := db.Where("id = ? AND deleted = ?", permission.ID, false).Updates(permission).Error; err != nil {
		return fmt.Errorf("更新权限失败: %w", err)
	}
	return nil
}

// DeleteOnePermissionByID 删除权限
// 参数：
//   - c: Echo 上下文
//   - permissionID: 权限 ID
//
// 返回值：
//   - error: 操作过程中的错误
func DeleteOnePermissionByID(c echo.Context, permissionID int64) error {
	db := utils.GetDBFromContext(c)
	if err := db.Model(&account.Permission{}).Where("id = ?", permissionID).Update("deleted", true).Error; err != nil {
		return fmt.Errorf("删除角色失败： %v", err)
	}
	return nil
}

// GetOnePermissionByID 根据权限 ID 获取权限
// 参数：
//   - c: Echo 上下文
//   - permissionID: 权限 ID
//
// 返回值：
//   - *account.Permission: 权限信息
//   - error: 操作过程中的错误
func GetOnePermissionByID(c echo.Context, permissionID int64) (*account.Permission, error) {
	var permission account.Permission
	db := utils.GetDBFromContext(c)
	if err := db.Where("id = ? AND deleted = ?", permissionID, false).First(&permission).Error; err != nil {
		return nil, fmt.Errorf("获取权限失败: %v", err)
	}
	return &permission, nil
}

// GetAllPermissions 获取所有权限
// 参数：
//   - c: Echo 上下文
//
// 返回值：
//   - []*account.Permission: 权限列表
//   - error: 操作过程中的错误
func GetAllPermissions(c echo.Context) ([]*account.Permission, error) {
	var permissions []*account.Permission
	db := utils.GetDBFromContext(c)
	if err := db.Where("deleted = ?", false).Find(&permissions).Error; err != nil {
		return nil, fmt.Errorf("获取权限列表失败: %v", err)
	}
	return permissions, nil
}
