// Package verification 提供验证码相关的HTTP接口处理
// 创建者：Done-0
// 创建时间：2025-05-10
package verification

import (
	"net/http"

	"github.com/labstack/echo/v4"

	bizErr "jank.com/jank_blog/internal/error"
	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/serve/controller/verification/dto"
	service "jank.com/jank_blog/pkg/serve/service/verification"
	"jank.com/jank_blog/pkg/vo"
)

// SendImgVerificationCode godoc
// @Summary      生成图形验证码并返回Base64编码
// @Description  生成单个图形验证码并将其返回为Base64编码字符串，用户可以用该验证码进行校验。
// @Tags         用户
// @Accept       json
// @Produce      json
// @Param        email  query   string  true  "邮箱地址，用于生成验证码"
// @Success      200   {object} vo.Result{data=map[string]string} "成功返回验证码的Base64编码"
// @Failure      400   {object} vo.Result{data=string} "请求参数错误，邮箱地址为空"
// @Failure      500   {object} vo.Result{data=string} "服务器错误，生成验证码失败"
// @Router       /verification/sendImgVerificationCode [get]
func SendImgVerificationCode(c echo.Context) error {
	req := new(dto.GetOneVerificationCode)
	if err := (&echo.DefaultBinder{}).BindQueryParams(c, req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, err, bizErr.New(bizErr.BAD_REQUEST, err.Error())))
	}

	errors := utils.Validator(req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, errors, bizErr.New(bizErr.BAD_REQUEST, "请求参数校验失败")))
	}

	result, err := service.GenerateImgVerificationCode(c, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR)))
	}

	return c.JSON(http.StatusOK, vo.Success(c, result))
}

// SendEmailVerificationCode godoc
// @Summary 	发送邮箱验证码
// @Description 向指定邮箱发送验证码，验证码有效期为3分钟
// @Tags 		用户
// @Accept 		json
// @Produce 	json
// @Param 		email  query	string  true  "邮箱地址，用于发送验证码"
// @Success 	200    {object} vo.Result "邮箱验证码发送成功, 请注意查收邮件"
// @Failure 	400    {object} vo.Result "请求参数错误，邮箱地址为空"
// @Failure 	500    {object} vo.Result "服务器错误，邮箱验证码发送失败"
// @Router /verification/sendEmailVerificationCode [get]
func SendEmailVerificationCode(c echo.Context) error {
	req := new(dto.GetOneVerificationCode)
	if err := (&echo.DefaultBinder{}).BindQueryParams(c, req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, err, bizErr.New(bizErr.BAD_REQUEST, err.Error())))
	}

	errors := utils.Validator(req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, errors, bizErr.New(bizErr.BAD_REQUEST, "请求参数校验失败")))
	}

	err := service.SendEmailVerificationCode(c, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR)))
	}

	return c.JSON(http.StatusOK, vo.Success(c, "邮箱验证码发送成功, 请注意查收！"))
}
