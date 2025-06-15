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

// CreateOneRole 创建角色
// @Summary      创建角色
// @Description  创建一个新的角色，包括角色名、描述和状态信息
// @Tags         角色管理
// @Accept       json
// @Produce      json
// @Param        request  body      dto.CreateOneRoleRequest  true  "角色信息"
// @Success      200     {object}  vo.Result{data=account.RoleVo}  "创建成功"
// @Failure      400     {object}  vo.Result{message=string}    "参数错误"
// @Failure      500     {object}  vo.Result{message=string}    "服务器错误"
// @Router       /role/createOneRole [post]
func CreateOneRole(c echo.Context) error {
	req := new(dto.CreateOneRoleRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, err, bizErr.New(bizErr.BAD_REQUEST, err.Error())))
	}

	errors := utils.Validator(req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, errors, bizErr.New(bizErr.BAD_REQUEST, "请求参数校验失败")))
	}

	role, err := service.CreateOneRole(c, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, role))
}

// UpdateOneRole 更新角色
// @Summary      更新角色
// @Description  更新一个角色
// @Tags         角色管理
// @Accept       json
// @Produce      json
// @Param        request  body      dto.UpdateOneRoleRequest  true  "角色信息"
// @Success      200     {object}  vo.Result{data=account.RoleVo}  "更新成功"
// @Failure      400     {object}  vo.Result{message=string}    "参数错误"
// @Failure      500     {object}  vo.Result{message=string}    "服务器错误"
// @Router       /role/updateOneRole [post]
func UpdateOneRole(c echo.Context) error {
	req := new(dto.UpdateOneRoleRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, err, bizErr.New(bizErr.BAD_REQUEST, err.Error())))
	}

	errors := utils.Validator(req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, errors, bizErr.New(bizErr.BAD_REQUEST, "请求参数校验失败")))
	}

	role, err := service.UpdateOneRole(c, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, role))
}

// DeleteOneRole 删除角色
// @Summary      删除角色
// @Description  根据角色 ID 删除角色
// @Tags         角色管理
// @Accept       json
// @Produce      json
// @Param        request  body      dto.DeleteOneRoleRequest  true  "角色ID"
// @Success      200     {object}  vo.Result{data=string}  "删除成功"
// @Failure      400     {object}  vo.Result{message=string} "参数错误"
// @Failure      500     {object}  vo.Result{message=string} "服务器错误"
// @Router       /role/deleteOneRole [post]
func DeleteOneRole(c echo.Context) error {
	req := new(dto.DeleteOneRoleRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, err, bizErr.New(bizErr.BAD_REQUEST, err.Error())))
	}

	errors := utils.Validator(req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, errors, bizErr.New(bizErr.BAD_REQUEST, "请求参数校验失败")))
	}

	err := service.DeleteOneRoleByID(c, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, "角色删除成功"))
}

// ListAllRoles 获取所有角色
// @Summary      获取所有角色
// @Description  获取系统中所有角色的信息
// @Tags         角色管理
// @Accept       json
// @Produce      json
// @Success      200     {object}  vo.Result{data=[]account.RoleVo}  "获取成功"
// @Failure      500     {object}  vo.Result{message=string}     "服务器错误"
// @Router       /role/listAllRoles [get]
func ListAllRoles(c echo.Context) error {
	roles, err := service.ListAllRoles(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, roles))
}
