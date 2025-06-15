package service

import (
	"fmt"
	
	"github.com/labstack/echo/v4"

	model "jank.com/jank_blog/internal/model/account"
	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/serve/mapper"
	"jank.com/jank_blog/pkg/serve/controller/account/dto"
	"jank.com/jank_blog/pkg/vo/account"
)

// CreateOnePermission 创建权限
// 参数：
//   - c: Echo 上下文
//   - req: 创建权限请求
//
// 返回值：
//   - *account.PermissionVO: 创建后的权限视图对象
//   - error: 操作过程中的错误
func CreateOnePermission(c echo.Context, req *dto.CreateOnePermissionRequest) (*account.PermissionVO, error) {
	var permissionVO *account.PermissionVO

	err := utils.RunDBTransaction(c, func(tx error) error {
		permission := &model.Permission{
			Key:         req.Key,
			Name:        req.Name,
			Description: req.Description,
			Status:      req.Status,
		}

		if err := mapper.CreateOnePermission(c, permission); err != nil {
			utils.BizLogger(c).Errorf("创建权限失败: %v", err)
			return fmt.Errorf("创建权限失败: %w", err)
		}

		vo, err := utils.MapModelToVO(permission, &account.PermissionVO{})
		if err != nil {
			utils.BizLogger(c).Errorf("权限创建时映射 VO 失败: %v", err)
			return fmt.Errorf("权限创建时映射 VO 失败: %w", err)
		}

		permissionVO = vo.(*account.PermissionVO)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return permissionVO, nil
}

// UpdateOnePermission 更新权限
// 参数：
//   - c: Echo 上下文
//   - req: 更新权限请求
//
// 返回值：
//   - *account.PermissionVO: 更新后的权限视图对象
//   - error: 操作过程中的错误
func UpdateOnePermission(c echo.Context, req *dto.UpdateOnePermissionRequest) (*account.PermissionVO, error) {
	var permissionVO *account.PermissionVO

	err := utils.RunDBTransaction(c, func(tx error) error {
		permission, err := mapper.GetOnePermissionByID(c, req.ID)
		if err != nil {
			utils.BizLogger(c).Errorf("获取权限失败: %v", err)
			return fmt.Errorf("获取权限失败: %w", err)
		}

		permission.Name = req.Name
		permission.Description = req.Description

		if err := mapper.UpdateOnePermissionByID(c, permission); err != nil {
			utils.BizLogger(c).Errorf("更新权限失败: %v", err)
			return fmt.Errorf("更新权限失败: %w", err)
		}

		vo, err := utils.MapModelToVO(permission, &account.PermissionVO{})
		if err != nil {
			utils.BizLogger(c).Errorf("权限更新时映射 VO 失败: %v", err)
			return fmt.Errorf("权限更新时映射 VO 失败: %w", err)
		}

		permissionVO = vo.(*account.PermissionVO)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return permissionVO, nil
}

// DeleteOnePermissionByID 删除权限
// 参数：
//   - c: Echo 上下文
//   - req: 删除权限请求
//
// 返回值：
//   - error: 操作过程中的错误
func DeleteOnePermissionByID(c echo.Context, req *dto.DeleteOnePermissionRequest) error {
	return utils.RunDBTransaction(c, func(tx error) error {
		if err := mapper.DeleteOnePermissionByID(c, req.ID); err != nil {
			utils.BizLogger(c).Errorf("删除权限失败: %v", err)
			return fmt.Errorf("删除权限失败: %w", err)
		}

		return nil
	})
}

// ListAllPermissions 获取所有权限
// 参数：
//   - c: Echo 上下文
//
// 返回值：
//   - []*account.PermissionVO: 权限视图对象列表
//   - error: 操作过程中的错误
func ListAllPermissions(c echo.Context) ([]*account.PermissionVO, error) {
	var permissionVos []*account.PermissionVO

	err := utils.RunDBTransaction(c, func(tx error) error {
		permissions, err := mapper.GetAllPermissions(c)
		if err != nil {
			utils.BizLogger(c).Errorf("获取所有权限失败: %v", err)
			return fmt.Errorf("获取所有权限失败: %w", err)
		}

		for _, permission := range permissions {
			permissionVo, err := utils.MapModelToVO(permission, &account.PermissionVO{})
			if err != nil {
				utils.BizLogger(c).Errorf("获取权限列表时映射 VO 失败: %v", err)
				return fmt.Errorf("获取权限列表时映射 VO 失败: %w", err)
			}
			permissionVos = append(permissionVos, permissionVo.(*account.PermissionVO))
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return permissionVos, nil
}
