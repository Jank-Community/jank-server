package mapper

import (
	"fmt"

	"github.com/labstack/echo/v4"

	account "jank.com/jank_blog/internal/model/account"
	"jank.com/jank_blog/internal/utils"
)

// AssignRolesToOneAccountByIDs 为用户分配角色
// 参数：
//   - c: Echo 上下文
//   - account: 用户角色信息
//
// 返回值：
//   - error: 操作过程中的错误
func AssignRolesToOneAccountByIDs(c echo.Context, account *account.AccountRole) error {
	db := utils.GetDBFromContext(c)
	if err := db.Create(account).Error; err != nil {
		return fmt.Errorf("为用户分配角色失败: %v", err)
	}
	return nil
}

// RevokeRoleFromAccount 从用户撤销角色
// 参数：
//   - c: Echo 上下文
//   - accountID: 用户 ID
//   - roleID: 角色 ID
//
// 返回值：
//   - error: 操作过程中的错误
func RevokeRoleFromAccount(c echo.Context, accountID int64, roleID int64) error {
	db := utils.GetDBFromContext(c)
	if err := db.Where("account_id = ? AND role_id = ?", accountID, roleID).Delete(&account.AccountRole{}).Error; err != nil {
		return fmt.Errorf("从用户撤销角色失败: %v", err)
	}
	return nil
}

// GetRolesByAccountID 获取用户的所有角色
// 参数：
//   - c: Echo 上下文
//   - accountID: 用户 ID
//
// 返回值：
//   - []*account.Role: 角色列表
//   - error: 操作过程中的错误
func GetRolesByAccountID(c echo.Context, accountID int64) ([]*account.Role, error) {
	db := utils.GetDBFromContext(c)
	var roles []*account.Role

	// 使用CAST确保布尔值和整数比较的兼容性
	query := `
		SELECT r.* 
		FROM roles r 
		JOIN account_roles ar ON r.id = ar.role_id 
		WHERE ar.account_id = ? 
		AND CAST(r.deleted AS INTEGER) = 0
	`

	err := db.Raw(query, accountID).Scan(&roles).Error

	if err != nil {
		return nil, fmt.Errorf("获取用户角色失败: %v", err)
	}
	return roles, nil
}
