package service

import (
	"fmt"

	"github.com/labstack/echo/v4"

	model "jank.com/jank_blog/internal/model/account"
	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/serve/controller/account/dto"
	"jank.com/jank_blog/pkg/serve/mapper"
	"jank.com/jank_blog/pkg/vo/account"
)

// AssignPermissionsToOneRole 分配权限给角色
// 参数：
//   - c: Echo 上下文
//   - req: 分配权限请求
//
// 返回值：
//   - error: 操作过程中的错误
func AssignPermissionsToOneRole(c echo.Context, req *dto.AssignPermissionsToRoleRequest) error {
	return utils.RunDBTransaction(c, func(tx error) error {
		role, err := mapper.GetOneRoleByID(c, req.RoleID)
		if err != nil {
			utils.BizLogger(c).Errorf("「%d」角色不存在: %v", req.RoleID, err)
			return fmt.Errorf("角色不存在: %w", err)
		}

		// 遍历所有权限 ID 进行验证和分配
		for _, permissionIDStr := range req.PermissionIDs {
			var permissionID int64
			if _, err := fmt.Sscanf(permissionIDStr, "%d", &permissionID); err != nil {
				utils.BizLogger(c).Errorf("权限ID格式错误: %v", err)
				return fmt.Errorf("权限ID格式错误: %w", err)
			}

			_, err = mapper.GetOnePermissionByID(c, permissionID)
			if err != nil {
				utils.BizLogger(c).Errorf("「%d」权限不存在: %v", permissionID, err)
				return fmt.Errorf("「%d」权限不存在: %w", permissionID, err)
			}

			rolePermission := &model.RolePermission{
				RoleID:       req.RoleID,
				PermissionID: permissionID,
			}

			err = mapper.AssignPermissionsToOneRoleByIDs(c, rolePermission)
			if err != nil {
				utils.BizLogger(c).Errorf("为「%s」角色分配权限失败: %v", role.Name, err)
				return fmt.Errorf("为角色分配权限失败: %w", err)
			}
		}

		return nil
	})
}

// RevokePermissionsFromOneRole 从角色撤销权限
// 参数：
//   - c: Echo 上下文
//   - req: 撤销权限请求
//
// 返回值：
//   - error: 操作过程中的错误
func RevokePermissionsFromOneRole(c echo.Context, req *dto.RevokePermissionsFromRoleRequest) error {
	return utils.RunDBTransaction(c, func(tx error) error {
		role, err := mapper.GetOneRoleByID(c, req.RoleID)
		if err != nil {
			utils.BizLogger(c).Errorf("「%d」角色不存在: %v", req.RoleID, err)
			return fmt.Errorf("角色不存在: %w", err)
		}

		// 遍历所有权限 ID 进行验证和撤销
		for _, permissionIDStr := range req.PermissionIDs {
			// 将字符串转换为int64
			var permissionID int64
			if _, err := fmt.Sscanf(permissionIDStr, "%d", &permissionID); err != nil {
				utils.BizLogger(c).Errorf("权限ID格式错误: %v", err)
				return fmt.Errorf("权限ID格式错误: %w", err)
			}

			_, err = mapper.GetOnePermissionByID(c, permissionID)
			if err != nil {
				utils.BizLogger(c).Errorf("「%d」权限不存在: %v", permissionID, err)
				return fmt.Errorf("「%d」权限不存在: %w", permissionID, err)
			}

			err = mapper.RevokePermissionFromOneRole(c, req.RoleID, permissionID)
			if err != nil {
				utils.BizLogger(c).Errorf("从「%s」角色撤销权限失败: %v", role.Name, err)
				return fmt.Errorf("从角色撤销权限失败: %w", err)
			}
		}

		return nil
	})
}

// GetPermissionsFromOneRole 获取角色的所有权限
// 参数：
//   - c: Echo 上下文
//   - req: 获取角色权限请求
//
// 返回值：
//   - *account.RolePermissionVO: 角色权限视图对象
//   - error: 操作过程中的错误
func GetPermissionsFromOneRole(c echo.Context, req *dto.GetRolePermissionsRequest) (*account.RolePermissionVO, error) {
	role, err := mapper.GetOneRoleByID(c, req.RoleID)
	if err != nil {
		utils.BizLogger(c).Errorf("「%d」角色不存在: %v", req.RoleID, err)
		return nil, fmt.Errorf("角色不存在: %w", err)
	}

	// 获取角色的所有权限
	permissions, err := mapper.GetPermissionsByRoleID(c, req.RoleID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取「%s」角色的权限失败: %v", role.Name, err)
		return nil, fmt.Errorf("获取角色权限失败: %w", err)
	}

	permissionVOs := make([]account.PermissionVO, 0, len(permissions))
	for _, permission := range permissions {
		permissionVO, err := utils.MapModelToVO(permission, &account.PermissionVO{})
		if err != nil {
			utils.BizLogger(c).Errorf("获取角色权限时映射 VO 失败: %v", err)
			return nil, fmt.Errorf("获取角色权限时映射 VO 失败: %w", err)
		}

		permissionVOPtr := permissionVO.(*account.PermissionVO)
		permissionVOs = append(permissionVOs, *permissionVOPtr)
	}

	return &account.RolePermissionVO{
		RoleID:      req.RoleID,
		RoleName:    role.Name,
		Permissions: permissionVOs,
	}, nil
}
