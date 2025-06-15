package account

import (
	"net/http"

	"github.com/labstack/echo/v4"

	bizErr "jank.com/jank_blog/internal/error"
	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/serve/controller/account/dto"
	service "jank.com/jank_blog/pkg/serve/service/account"
	"jank.com/jank_blog/pkg/vo"
)

// AssignPermissionsToOneRole 为角色分配权限
// @Summary      为角色分配权限
// @Description  根据权限 ID 和角色 ID 为角色分配权限
// @Tags         角色权限管理
// @Accept       json
// @Produce      json
// @Param        request  body      dto.AssignPermissionsToRoleRequest  true  "分配权限信息"
// @Success      200     {object}  vo.Result{data=string}  "角色权限分配成功"
// @Failure      400     {object}  vo.Result{message=string} "参数错误"
// @Failure      500     {object}  vo.Result{message=string} "服务器错误"
// @Router       /role-permission/assignPermissions [post]
func AssignPermissionsToOneRole(c echo.Context) error {
	req := new(dto.AssignPermissionsToRoleRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, err, bizErr.New(bizErr.BAD_REQUEST, err.Error())))
	}

	errors := utils.Validator(req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, errors, bizErr.New(bizErr.BAD_REQUEST, "请求参数校验失败")))
	}

	err := service.AssignPermissionsToOneRole(c, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, "权限分配成功！"))
}

// RevokePermissionsFromOneRole 撤销角色的权限
// @Summary      撤销角色权限
// @Description  根据权限 ID 和角色 ID 撤销角色权限
// @Tags         角色权限管理
// @Accept       json
// @Produce      json
// @Param        request  body      dto.RevokePermissionsFromRoleRequest  true  "撤销权限信息"
// @Success      200     {object}  vo.Result{data=string}  "角色权限撤销成功"
// @Failure      400     {object}  vo.Result{message=string} "参数错误"
// @Failure      500     {object}  vo.Result{message=string} "服务器错误"
// @Router       /role-permission/revokePermissions [post]
func RevokePermissionsFromOneRole(c echo.Context) error {
	req := new(dto.RevokePermissionsFromRoleRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, err, bizErr.New(bizErr.BAD_REQUEST, err.Error())))
	}

	errors := utils.Validator(req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, errors, bizErr.New(bizErr.BAD_REQUEST, "请求参数校验失败")))
	}

	err := service.RevokePermissionsFromOneRole(c, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, "权限撤销成功！"))
}

// GetRolePermissionsFromOneRole 获取角色权限
// @Summary      获取角色权限
// @Description  根据角色 ID 获取角色所有权限
// @Tags         角色权限管理
// @Accept       json
// @Produce      json
// @Param        request  query     dto.GetRolePermissionsRequest  true  "获取权限请求"
// @Success      200     {object}  vo.Result{data=account.RolePermissionVO}  "获取成功"
// @Failure      400     {object}  vo.Result{message=string} "参数错误"
// @Failure      500     {object}  vo.Result{message=string} "服务器错误"
// @Router       /role-permission/getRolePermissions [get]
func GetRolePermissionsFromOneRole(c echo.Context) error {
	req := new(dto.GetRolePermissionsRequest)
	if err := (&echo.DefaultBinder{}).BindQueryParams(c, req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, err, bizErr.New(bizErr.BAD_REQUEST, err.Error())))
	}

	errors := utils.Validator(req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, errors, bizErr.New(bizErr.BAD_REQUEST, "请求参数校验失败")))
	}

	response, err := service.GetPermissionsFromOneRole(c, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, response))
}
