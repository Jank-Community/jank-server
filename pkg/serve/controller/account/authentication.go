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

// CheckAccountPermission 检查用户权限
// @Summary      检查用户是否拥有特定权限
// @Description  根据用户 ID 和权限标识检查用户是否拥有该权限
// @Tags         权限验证
// @Accept       json
// @Produce      json
// @Param        request  body      dto.CheckAccountPermissionRequest  true  "检查权限请求"
// @Success      200     {object}  vo.Result{data=account.CheckPermissionVO}  "检查成功"
// @Failure      400     {object}  vo.Result{message=string} "参数错误"
// @Failure      500     {object}  vo.Result{message=string} "服务器错误"
// @Security     BearerAuth
// @Router       /authentication/checkPermission [post]
func CheckAccountPermission(c echo.Context) error {
	req := new(dto.CheckAccountPermissionRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, err, bizErr.New(bizErr.BAD_REQUEST, err.Error())))
	}

	errors := utils.Validator(req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, errors, bizErr.New(bizErr.BAD_REQUEST, "请求参数校验失败")))
	}

	response, err := service.CheckPermissionFromOneAccount(c, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, response))
}

// GetAccountPermissions 获取用户所有权限
// @Summary      获取用户所有权限
// @Description  根据用户 ID 获取该用户的所有权限列表
// @Tags         权限验证
// @Accept       json
// @Produce      json
// @Param        request  query     dto.GetAccountPermissionsRequest  true  "获取权限请求"
// @Success      200     {object}  vo.Result{data=account.AccountPermissionsVO}  "获取成功"
// @Failure      400     {object}  vo.Result{message=string} "参数错误"
// @Failure      500     {object}  vo.Result{message=string} "服务器错误"
// @Security     BearerAuth
// @Router       /authentication/getAccountPermissions [get]
func GetAccountPermissions(c echo.Context) error {
	req := new(dto.GetAccountPermissionsRequest)
	if err := (&echo.DefaultBinder{}).BindQueryParams(c, req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, err, bizErr.New(bizErr.BAD_REQUEST, err.Error())))
	}

	errors := utils.Validator(req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, errors, bizErr.New(bizErr.BAD_REQUEST, "请求参数校验失败")))
	}

	response, err := service.GetPermissionsFromOneAccount(c, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, response))
}
