package service

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/serve/controller/account/dto"
	"jank.com/jank_blog/pkg/serve/mapper"
	"jank.com/jank_blog/pkg/vo/account"
)

// CheckPermissionFromOneAccount 检查用户是否拥有特定权限
// 参数：
//   - c: Echo 上下文
//   - req: 权限检查请求
//
// 返回值：
//   - *account.CheckPermissionVO: 权限检查结果
//   - error: 操作过程中的错误
func CheckPermissionFromOneAccount(c echo.Context, req *dto.CheckAccountPermissionRequest) (*account.CheckPermissionVO, error) {
	acc, err := mapper.GetOneAccountByID(c, req.AccountID)
	if err != nil {
		utils.BizLogger(c).Errorf("「%d」用户不存在: %v", req.AccountID, err)
		return nil, fmt.Errorf("用户不存在: %w", err)
	}

	// 检查权限标识是否存在
	permission, err := mapper.GetOnePermissionByKey(c, req.PermissionKey)
	if err != nil {
		utils.BizLogger(c).Errorf("「%s」权限标识不存在: %v", req.PermissionKey, err)
		return nil, fmt.Errorf("权限标识不存在: %w", err)
	}

	// 检查用户是否有权限
	hasPermission, err := mapper.CheckUserHasPermission(c, req.AccountID, permission.ID)
	if err != nil {
		utils.BizLogger(c).Errorf("检查「%s」用户「%s」权限失败: %v", acc.Email, req.PermissionKey, err)
		return nil, fmt.Errorf("检查用户权限失败: %w", err)
	}

	return &account.CheckPermissionVO{
		HasPermission: hasPermission,
		PermissionKey: req.PermissionKey,
	}, nil
}

// GetPermissionsFromOneAccount 获取用户所有权限
// 参数：
//   - c: Echo 上下文
//   - req: 获取用户权限请求
//
// 返回值：
//   - *account.AccountPermissionsVO: 用户权限对象
//   - error: 操作过程中的错误
func GetPermissionsFromOneAccount(c echo.Context, req *dto.GetAccountPermissionsRequest) (*account.AccountPermissionsVO, error) {
	acc, err := mapper.GetOneAccountByID(c, req.AccountID)
	if err != nil {
		utils.BizLogger(c).Errorf("「%d」用户不存在: %v", req.AccountID, err)
		return nil, fmt.Errorf("用户不存在: %w", err)
	}

	// 获取用户所有权限
	permissions, err := mapper.GetPermissionsByAccountID(c, req.AccountID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取「%s」用户权限失败: %v", acc.Email, err)
		return nil, fmt.Errorf("获取用户权限失败: %w", err)
	}

	permissionVOs := make([]account.PermissionVO, 0, len(permissions))
	for _, permission := range permissions {
		permissionVO, err := utils.MapModelToVO(permission, &account.PermissionVO{})
		if err != nil {
			utils.BizLogger(c).Errorf("获取用户权限时映射 VO 失败: %v", err)
			return nil, fmt.Errorf("获取用户权限时映射 VO 失败: %w", err)
		}

		permissionVOPtr := permissionVO.(*account.PermissionVO)
		permissionVOs = append(permissionVOs, *permissionVOPtr)
	}

	return &account.AccountPermissionsVO{
		AccountID:   req.AccountID,
		Permissions: permissionVOs,
	}, nil
}
