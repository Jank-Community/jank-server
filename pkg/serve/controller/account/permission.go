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

// CreateOnePermission 创建权限
// @Summary      创建权限
// @Description  创建新的权限，权限信息包括代码和描述
// @Tags         权限管理
// @Accept       json
// @Produce      json
// @Param        request  body      dto.CreateOnePermissionRequest  true  "权限信息"
// @Success      200     {object}  vo.Result{data=account.PermissionVo}  "创建成功"
// @Failure      400     {object}  vo.Result{message=string}    "参数错误"
// @Failure      500     {object}  vo.Result{message=string}    "服务器错误"
// @Router       /permission/createOnePermission [post]
// 参数：
//   - c: Echo 上下文
//
// 返回值：
//   - error: 操作过程中的错误
func CreateOnePermission(c echo.Context) error {
	req := new(dto.CreateOnePermissionRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, err, bizErr.New(bizErr.BAD_REQUEST, err.Error())))
	}

	errors := utils.Validator(req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, errors, bizErr.New(bizErr.BAD_REQUEST, "请求参数校验失败")))
	}

	permission, err := service.CreateOnePermission(c, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, permission))
}

// UpdateOnePermission 更新权限
// @Summary      更新权限
// @Description  更新权限的代码和描述
// @Tags         权限管理
// @Accept       json
// @Produce      json
// @Param        request  body      dto.UpdateOnePermissionRequest  true  "权限信息"
// @Success      200     {object}  vo.Result{data=account.PermissionVo}  "更新成功"
// @Failure      400     {object}  vo.Result{message=string}    "参数错误"
// @Failure      500     {object}  vo.Result{message=string}    "服务器错误"
// @Router       /permission/updateOnePermission [post]
// 参数：
//   - c: Echo 上下文
//
// 返回值：
//   - error: 操作过程中的错误
func UpdateOnePermission(c echo.Context) error {
	req := new(dto.UpdateOnePermissionRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, err, bizErr.New(bizErr.BAD_REQUEST, err.Error())))
	}

	errors := utils.Validator(req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, errors, bizErr.New(bizErr.BAD_REQUEST, "请求参数校验失败")))
	}

	permission, err := service.UpdateOnePermission(c, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, permission))
}

// DeleteOnePermission 删除权限
// @Summary      删除权限
// @Description  根据权限ID删除权限
// @Tags         权限管理
// @Accept       json
// @Produce      json
// @Param        request  body      dto.DeleteOnePermissionRequest  true  "权限ID"
// @Success      200     {object}  vo.Result{data=string}  "删除成功"
// @Failure      400     {object}  vo.Result{message=string} "参数错误"
// @Failure      500     {object}  vo.Result{message=string} "服务器错误"
// @Router       /permission/deleteOnePermission [post]
// 参数：
//   - c: Echo 上下文
//
// 返回值：
//   - error: 操作过程中的错误
func DeleteOnePermission(c echo.Context) error {
	req := new(dto.DeleteOnePermissionRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, err, bizErr.New(bizErr.BAD_REQUEST, err.Error())))
	}

	errors := utils.Validator(req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, errors, bizErr.New(bizErr.BAD_REQUEST, "请求参数校验失败")))
	}

	err := service.DeleteOnePermissionByID(c, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, "权限删除成功"))
}

// ListAllPermissions 获取所有权限
// @Summary      获取所有权限
// @Description  获取系统中所有权限的信息
// @Tags         权限管理
// @Accept       json
// @Produce      json
// @Success      200     {object}  vo.Result{data=[]account.PermissionVo}  "获取成功"
// @Failure      500     {object}  vo.Result{message=string}     "服务器错误"
// @Router       /permission/listAllPermissions [post]
// 参数：
//   - c: Echo 上下文
//
// 返回值：
//   - error: 操作过程中的错误
func ListAllPermissions(c echo.Context) error {
	permissions, err := service.ListAllPermissions(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, permissions))
}
