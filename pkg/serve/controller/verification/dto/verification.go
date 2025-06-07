// Package dto 提供文章相关的数据传输对象定义
// 创建者：Done-0
// 创建时间：2025-06-07
package dto

// GetOneVerificationCode            获取图形验证码
// @Param			email	    body	string	true	"用户邮箱"
type GetOneVerificationCode struct {
	Email string `json:"email" xml:"email" form:"email" query:"email" validate:"required,email"`
}
