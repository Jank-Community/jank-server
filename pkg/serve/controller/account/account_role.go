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

// AssignRolesToOneAccount 为账户分配角色
// @Summary      为用户分配角色
// @Description  根据角色 ID 和账户 ID 为账户分配角色
// @Tags         账户角色管理
// @Accept       json
// @Produce      json
// @Param        request  body      dto.AssignRolesToAccountRequest  true  "分配角色信息"
// @Success      200     {object}  vo.Result{data=string}  "用户角色分配成功"
// @Failure      400     {object}  vo.Result{message=string} "参数错误"
// @Failure      500     {object}  vo.Result{message=string} "服务器错误"
// @Router       /account-role/assignRoleToAcc [post]
func AssignRolesToOneAccount(c echo.Context) error {
	req := new(dto.AssignRolesToAccountRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, err, bizErr.New(bizErr.BAD_REQUEST, err.Error())))
	}

	errors := utils.Validator(req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, errors, bizErr.New(bizErr.BAD_REQUEST, "请求参数校验失败")))
	}

	err := service.AssignRolesToOneAccount(c, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, "角色分配成功！"))
}

// RevokeRolesFromOneAccount 撤销用户角色
// @Summary      撤销用户角色
// @Description  根据角色 ID 和账户 ID 撤销用户角色
// @Tags         账户角色管理
// @Accept       json
// @Produce      json
// @Param        request  body      dto.RevokeRolesFromOneAccountRequest  true  "撤销角色信息"
// @Success      200     {object}  vo.Result{data=string}  "用户角色撤销成功"
// @Failure      400     {object}  vo.Result{message=string} "参数错误"
// @Failure      500     {object}  vo.Result{message=string} "服务器错误"
// @Router       /account-role/revokeRoles [post]
func RevokeRolesFromOneAccount(c echo.Context) error {
	req := new(dto.RevokeRolesFromOneAccountRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, err, bizErr.New(bizErr.BAD_REQUEST, err.Error())))
	}

	errors := utils.Validator(req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, errors, bizErr.New(bizErr.BAD_REQUEST, "请求参数校验失败")))
	}

	err := service.RevokeRolesFromOneAccount(c, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, "角色撤销成功！"))
}

// GetAccountRolesFromOneAccount 获取用户角色
// @Summary      获取用户角色
// @Description  根据账户 ID 获取用户所有角色
// @Tags         账户角色管理
// @Accept       json
// @Produce      json
// @Param        request  query     dto.GetAccountRolesRequest  true  "获取角色请求"
// @Success      200     {object}  vo.Result{data=account.AccountRoleVO}  "获取成功"
// @Failure      400     {object}  vo.Result{message=string} "参数错误"
// @Failure      500     {object}  vo.Result{message=string} "服务器错误"
// @Router       /account-role/getAccountRoles [get]
func GetAccountRolesFromOneAccount(c echo.Context) error {
	req := new(dto.GetAccountRolesRequest)
	if err := (&echo.DefaultBinder{}).BindQueryParams(c, req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, err, bizErr.New(bizErr.BAD_REQUEST, err.Error())))
	}

	errors := utils.Validator(req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, errors, bizErr.New(bizErr.BAD_REQUEST, "请求参数校验失败")))
	}

	response, err := service.GetRolesFromOneAccountFromOneAccount(c, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, response))
}
