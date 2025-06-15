package service

import (
	"fmt"

	"github.com/labstack/echo/v4"

	model "jank.com/jank_blog/internal/model/account"
	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/serve/controller/account/dto"
	"jank.com/jank_blog/pkg/serve/mapper"
	accountVo "jank.com/jank_blog/pkg/vo/account"
)

// AssignRolesToOneAccount 分配角色给账户
// 参数：
//   - c: Echo 上下文
//   - req: 分配角色请求
//
// 返回值：
//   - error: 操作过程中的错误
func AssignRolesToOneAccount(c echo.Context, req *dto.AssignRolesToAccountRequest) error {
	return utils.RunDBTransaction(c, func(tx error) error {
		acc, err := mapper.GetOneAccountByID(c, req.AccountID)
		if err != nil {
			utils.BizLogger(c).Errorf("「%s」用户不存在: %v", acc.Email, err)
			return fmt.Errorf("「%s」用户不存在: %w", acc.Email, err)
		}

		// 遍历所有角色 ID 进行验证和分配
		for _, roleIDStr := range req.RoleIDs {
			var roleID int64
			if _, err := fmt.Sscanf(roleIDStr, "%d", &roleID); err != nil {
				utils.BizLogger(c).Errorf("角色ID格式错误: %v", err)
				return fmt.Errorf("角色ID格式错误: %w", err)
			}

			_, err = mapper.GetOneRoleByID(c, roleID)
			if err != nil {
				utils.BizLogger(c).Errorf("「%d」角色不存在: %v", roleID, err)
				return fmt.Errorf("「%d」角色不存在: %w", roleID, err)
			}

			accountRole := &model.AccountRole{
				AccountID: req.AccountID,
				RoleID:    roleID,
			}

			err = mapper.AssignRolesToOneAccountByIDs(c, accountRole)
			if err != nil {
				utils.BizLogger(c).Errorf("为用户分配角色失败: %v", err)
				return fmt.Errorf("为用户分配角色失败: %w", err)
			}
		}

		return nil
	})
}

// RevokeRolesFromOneAccount 从账户撤销角色
// 参数：
//   - c: Echo 上下文
//   - req: 撤销角色请求
//
// 返回值：
//   - error: 操作过程中的错误
func RevokeRolesFromOneAccount(c echo.Context, req *dto.RevokeRolesFromOneAccountRequest) error {
	return utils.RunDBTransaction(c, func(tx error) error {
		acc, err := mapper.GetOneAccountByID(c, req.AccountID)
		if err != nil {
			utils.BizLogger(c).Errorf("「%d」用户不存在: %v", req.AccountID, err)
			return fmt.Errorf("用户不存在: %w", err)
		}

		// 遍历所有角色 ID 进行验证和撤销
		for _, roleIDStr := range req.RoleIDs {
			var roleID int64
			if _, err := fmt.Sscanf(roleIDStr, "%d", &roleID); err != nil {
				utils.BizLogger(c).Errorf("角色ID格式错误: %v", err)
				return fmt.Errorf("角色ID格式错误: %w", err)
			}

			_, err = mapper.GetOneRoleByID(c, roleID)
			if err != nil {
				utils.BizLogger(c).Errorf("「%d」角色不存在: %v", roleID, err)
				return fmt.Errorf("「%d」角色不存在: %w", roleID, err)
			}

			err = mapper.RevokeRoleFromAccount(c, req.AccountID, roleID)
			if err != nil {
				utils.BizLogger(c).Errorf("从「%s」用户撤销角色失败: %v", acc.Email, err)
				return fmt.Errorf("从用户撤销角色失败: %w", err)
			}
		}

		return nil
	})
}

// GetRolesFromOneAccountFromOneAccount 获取账户的所有角色
// 参数：
//   - c: Echo 上下文
//   - req: 获取账户角色请求
//
// 返回值：
//   - *accountVo.AccountRoleVO: 账户角色视图对象
//   - error: 操作过程中的错误
func GetRolesFromOneAccountFromOneAccount(c echo.Context, req *dto.GetAccountRolesRequest) (*accountVo.AccountRoleVO, error) {
	acc, err := mapper.GetOneAccountByID(c, req.AccountID)
	if err != nil {
		utils.BizLogger(c).Errorf("「%d」用户不存在: %v", req.AccountID, err)
		return nil, fmt.Errorf("用户不存在: %w", err)
	}

	// 获取用户的所有角色
	roles, err := mapper.GetRolesByAccountID(c, req.AccountID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取「%s」用户的角色失败: %v", acc.Email, err)
		return nil, fmt.Errorf("获取用户角色失败: %w", err)
	}

	roleVOs := make([]accountVo.RoleVO, 0, len(roles))
	for _, role := range roles {
		roleVO, err := utils.MapModelToVO(role, &accountVo.RoleVO{})
		if err != nil {
			utils.BizLogger(c).Errorf("获取用户所有角色时映射 VO 失败: %v", err)
			return nil, fmt.Errorf("获取用户所有角色时映射 VO 失败: %w", err)
		}

		roleVOPtr := roleVO.(*accountVo.RoleVO)
		roleVOs = append(roleVOs, *roleVOPtr)
	}

	return &accountVo.AccountRoleVO{
		AccountID: req.AccountID,
		Roles:     roleVOs,
	}, nil
}
