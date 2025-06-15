// Package mapper 提供数据模型与数据库交互的映射层，处理用户相关数据操作
// 创建者：Done-0
// 创建时间：2025-05-10
package mapper

import (
	"fmt"

	"github.com/labstack/echo/v4"

	account "jank.com/jank_blog/internal/model/account"
	"jank.com/jank_blog/internal/utils"
)

// CheckFirstUserExists 检查系统中是否已存在用户
// 参数：
//   - c: Echo 上下文
//
// 返回值：
//   - bool: 是否存在用户
//   - error: 操作过程中的错误
func CheckFirstUserExists(c echo.Context) (bool, error) {
	db := utils.GetDBFromContext(c)

	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM accounts WHERE CAST(deleted AS INTEGER) = 0 LIMIT 1)`

	err := db.Raw(query).Scan(&exists).Error
	if err != nil {
		return false, fmt.Errorf("检查用户存在性失败: %w", err)
	}

	return exists, nil
}

// GetOneAccountByEmail 根据邮箱获取用户账号信息
// 参数：
//   - c: Echo 上下文
//   - email: 用户邮箱
//
// 返回值：
//   - *account.Account: 用户信息
//   - error: 操作过程中的错误
func GetOneAccountByEmail(c echo.Context, email string) (*account.Account, error) {
	var user account.Account
	db := utils.GetDBFromContext(c)
	if err := db.Where("email = ? AND deleted = ?", email, false).First(&user).Error; err != nil {
		return nil, fmt.Errorf("获取用户失败: %w", err)
	}
	return &user, nil
}

// GetOneAccountByID 根据用户 ID 获取用户账号信息
// 参数：
//   - c: Echo 上下文
//   - accountID: 用户 ID
//
// 返回值：
//   - *account.Account: 账户信息
//   - error: 操作过程中的错误
func GetOneAccountByID(c echo.Context, accountID int64) (*account.Account, error) {
	var user account.Account
	db := utils.GetDBFromContext(c)
	if err := db.Where("id = ? AND deleted = ?", accountID, false).First(&user).Error; err != nil {
		return nil, fmt.Errorf("获取用户失败: %w", err)
	}
	return &user, nil
}

// CreateOneAccount 创建新用户
// 参数：
//   - c: Echo 上下文
//   - acc: 账户信息
//
// 返回值：
//   - error: 操作过程中的错误
func CreateOneAccount(c echo.Context, account *account.Account) error {
	db := utils.GetDBFromContext(c)
	if err := db.Create(account).Error; err != nil {
		return fmt.Errorf("创建用户失败: %w", err)
	}
	return nil
}

// UpdateOneAccountByID 更新账户信息
// 参数：
//   - c: Echo 上下文
//   - acc: 账户信息
//
// 返回值：
//   - error: 操作过程中的错误
func UpdateOneAccountByID(c echo.Context, account *account.Account) error {
	db := utils.GetDBFromContext(c)
	if err := db.Where("id = ? AND deleted = ?", account.ID, false).Updates(account).Error; err != nil {
		return fmt.Errorf("更新账户失败: %w", err)
	}
	return nil
}
