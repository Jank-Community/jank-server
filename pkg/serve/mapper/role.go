package mapper

import (
	"fmt"

	"github.com/labstack/echo/v4"

	account "jank.com/jank_blog/internal/model/account"
	"jank.com/jank_blog/internal/utils"
)

// CreateOneRole 创建角色
// 参数：
//   - c: Echo 上下文
//   - role: 角色信息
//
// 返回值：
//   - error: 操作过程中的错误
func CreateOneRole(c echo.Context, role *account.Role) error {
	db := utils.GetDBFromContext(c)
	if err := db.Create(role).Error; err != nil {
		return fmt.Errorf("创建角色失败: %v", err)
	}
	return nil
}

// UpdateOneRole 更新角色
// 参数：
//   - c: Echo 上下文
//   - role: 权限信息
//
// 返回值：
//   - error: 操作过程中的错误
func UpdateOneRole(c echo.Context, role *account.Role) error {
	db := utils.GetDBFromContext(c)
	if err := db.Where("id = ? AND deleted = ?", role.ID, false).Updates(role).Error; err != nil {
		return fmt.Errorf("更新角色失败: %w", err)
	}
	return nil
}

// DeleteOneRoleByID 删除角色
// 参数：
//   - c: Echo 上下文
//   - roleID: 角色 ID
//
// 返回值：
//   - error: 操作过程中的错误
func DeleteOneRoleByID(c echo.Context, roleID int64) error {
	db := utils.GetDBFromContext(c)
	if err := db.Model(&account.Role{}).Where("id = ?", roleID).Update("deleted", true).Error; err != nil {
		return fmt.Errorf("删除角色失败： %v", err)
	}
	return nil
}

// GetOneRoleByID 根据角色 ID 获取角色
// 参数：
//   - c: Echo 上下文
//   - roleID: 角色 ID
//
// 返回值：
//   - *account.Role: 角色信息
//   - error: 操作过程中的错误
func GetOneRoleByID(c echo.Context, roleID int64) (*account.Role, error) {
	var role account.Role
	db := utils.GetDBFromContext(c)
	if err := db.Where("id = ? AND deleted = ?", roleID, false).First(&role).Error; err != nil {
		return nil, fmt.Errorf("获取角色失败: %v", err)
	}
	return &role, nil
}

// GetAllRoles 获取所有角色
// 参数：
//   - c: Echo 上下文
//
// 返回值：
//   - []*account.Role: 角色列表
//   - error: 操作过程中的错误
func GetAllRoles(c echo.Context) ([]*account.Role, error) {
	var roles []*account.Role
	db := utils.GetDBFromContext(c)
	if err := db.Where("deleted = ?", false).Find(&roles).Error; err != nil {
		return nil, fmt.Errorf("获取角色列表失败: %v", err)
	}
	return roles, nil
}

// GetRoleByName 根据角色名称获取角色
// 参数：
//   - c: Echo 上下文
//   - roleName: 角色名称
//
// 返回值：
//   - *account.Role: 角色信息
//   - error: 操作过程中的错误
func GetRoleByName(c echo.Context, roleName string) (*account.Role, error) {
	var role account.Role
	db := utils.GetDBFromContext(c)
	if err := db.Where("name = ? AND deleted = ?", roleName, false).First(&role).Error; err != nil {
		return nil, fmt.Errorf("获取角色失败: %v", err)
	}
	return &role, nil
}
