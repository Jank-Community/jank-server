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

// CreateOneRole 创建角色
// 参数：
//   - c: Echo 上下文
//   - req: 创建角色请求
//
// 返回值：
//   - *account.RoleVO: 创建后的角色视图对象
//   - error: 操作过程中的错误
func CreateOneRole(c echo.Context, req *dto.CreateOneRoleRequest) (*account.RoleVO, error) {
	var roleVO *account.RoleVO

	err := utils.RunDBTransaction(c, func(tx error) error {
		role := &model.Role{
			Name:        req.Name,
			Description: req.Description,
			Status:      req.Status,
		}

		if err := mapper.CreateOneRole(c, role); err != nil {
			utils.BizLogger(c).Errorf("创建角色失败: %v", err)
			return fmt.Errorf("创建角色失败: %w", err)
		}

		vo, err := utils.MapModelToVO(role, &account.RoleVO{})
		if err != nil {
			utils.BizLogger(c).Errorf("角色创建时映射 VO 失败: %v", err)
			return fmt.Errorf("角色创建时映射 VO 失败: %w", err)
		}

		roleVO = vo.(*account.RoleVO)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return roleVO, nil
}

// UpdateOneRole 更新角色
// 参数：
//   - c: Echo 上下文
//   - req: 创建角色请求
//
// 返回值：
//   - *account.RoleVO: 更新后的角色视图对象
//   - error: 操作过程中的错误
func UpdateOneRole(c echo.Context, req *dto.UpdateOneRoleRequest) (*account.RoleVO, error) {
	var roleVO *account.RoleVO

	err := utils.RunDBTransaction(c, func(tx error) error {
		role, err := mapper.GetOneRoleByID(c, req.ID)
		if err != nil {
			utils.BizLogger(c).Errorf("获取角色失败: %v", err)
			return fmt.Errorf("获取角色失败: %w", err)
		}

		role.Name = req.Name
		role.Description = req.Description
		role.Status = req.Status

		if err := mapper.UpdateOneRole(c, role); err != nil {
			utils.BizLogger(c).Errorf("更新角色失败: %v", err)
			return fmt.Errorf("更新角色失败: %w", err)
		}

		vo, err := utils.MapModelToVO(role, &account.RoleVO{})
		if err != nil {
			utils.BizLogger(c).Errorf("角色更新时映射 VO 失败: %v", err)
			return fmt.Errorf("角色更新时映射 VO 失败: %w", err)
		}

		roleVO = vo.(*account.RoleVO)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return roleVO, nil
}

// DeleteOneRoleByID 删除角色
// 参数：
//   - c: Echo 上下文
//   - req: 删除角色请求
//
// 返回值：
//   - error: 操作过程中的错误
func DeleteOneRoleByID(c echo.Context, req *dto.DeleteOneRoleRequest) error {
	return utils.RunDBTransaction(c, func(tx error) error {
		if err := mapper.DeleteOneRoleByID(c, req.ID); err != nil {
			utils.BizLogger(c).Errorf("删除角色失败: %v", err)
			return fmt.Errorf("删除角色失败: %w", err)
		}

		return nil
	})
}

// ListAllRoles 获取所有角色
// 参数：
//   - c: Echo 上下文
//
// 返回值：
//   - *account.RoleVO: 更新后的角色视图对象
//   - error: 操作过程中的错误
func ListAllRoles(c echo.Context) ([]*account.RoleVO, error) {
	var roleVos []*account.RoleVO

	err := utils.RunDBTransaction(c, func(tx error) error {
		roles, err := mapper.GetAllRoles(c)
		if err != nil {
			utils.BizLogger(c).Errorf("获取所有角色失败: %v", err)
			return fmt.Errorf("获取所有角色失败: %w", err)
		}

		for _, role := range roles {
			roleVo, err := utils.MapModelToVO(role, &account.RoleVO{})
			if err != nil {
				utils.BizLogger(c).Errorf("获取角色列表时映射 VO 失败: %v", err)
				return fmt.Errorf("获取角色列表时映射 VO 失败: %w", err)
			}
			roleVos = append(roleVos, roleVo.(*account.RoleVO))
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return roleVos, nil
}
