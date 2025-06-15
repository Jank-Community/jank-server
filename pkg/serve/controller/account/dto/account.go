// Package dto 提供账户相关的数据传输对象定义
// 创建者：Done-0
// 创建时间：2025-05-10
package dto

// RegisterOneAccountRequest  用户注册请求体
// @Description	用户注册所需参数
// @Param			Email		body	string	true	"用户邮箱"
// @Param			Phone		body	string	true	"用户手机号"
// @Param			Nickname	body	string	true	"用户昵称"
// @Param			Password	body	string	true	"用户密码"
// @Param			EmailVerificationCode	body	string	true	"用户邮箱验证码"
// @Param			ImgVerificationCode	body	string	true	"用户图片验证码"
type RegisterOneAccountRequest struct {
	Email                 string `json:"email" xml:"email" form:"email" query:"email" validate:"required"`
	Phone                 string `json:"phone" xml:"phone" form:"phone" query:"phone" default:""`
	Nickname              string `json:"nickname" xml:"nickname" form:"nickname" query:"nickname" validate:"required,min=1,max=20"`
	Password              string `json:"password" xml:"password" form:"password" query:"password" validate:"required,min=6,max=20"`
	EmailVerificationCode string `json:"email_verification_code" xml:"email_verification_code" form:"email_verification_code" query:"email_verification_code" validate:"required"`
	ImgVerificationCode   string `json:"img_verification_code" xml:"img_verification_code" form:"img_verification_code" query:"img_verification_code" validate:"required"`
}

// LoginOneAccountRequest 用户登录请求体
// @Description	用户登录请求所需参数
// @Param			Email		body	string	true	"用户邮箱"
// @Param			Password	body	string	true	"用户密码"
// @Param			ImgVerificationCode	body	string	true	"图片验证码"
type LoginOneAccountRequest struct {
	Email               string `json:"email" xml:"email" form:"email" query:"email" validate:"required,email"`
	Password            string `json:"password" xml:"password" form:"password" query:"password" validate:"required"`
	ImgVerificationCode string `json:"img_verification_code" xml:"img_verification_code" form:"img_verification_code" query:"img_verification_code" validate:"required"`
}

// GetOneAccountRequest            获取账户信息请求体
// @Description	请求获取账户信息时所需参数
// @Param			Email	    body	string	true	"用户邮箱"
type GetOneAccountRequest struct {
	Email string `json:"email" xml:"email" form:"email" query:"email" validate:"required,email"`
}

// UpdateOneAccountRequest  更新账户信息请求体
// @Description	用户更新账户信息所需参数
// @Param			Nickname	body	string	true	"用户昵称"
// @Param			Phone		body	string	true	"用户手机号"
// @Param			Avatar		body	string	true	"用户头像"
type UpdateOneAccountRequest struct {
	Nickname string `json:"nickname" xml:"nickname" form:"nickname" query:"nickname" validate:"required,min=1,max=20"`
	Phone    string `json:"phone" xml:"phone" form:"phone" query:"phone"`
	Avatar   string `json:"avatar" xml:"avatar" form:"avatar" query:"avatar"`
}

// ResetPwdRequest  重置密码请求体
// @Description	用户重置密码所需参数
// @Param			Email					body	string	true	"用户邮箱"
// @Param			NewPassword			body	string	true	"新密码"
// @Param			AgainNewPassword		body	string	true	"再次输入新密码"
// @Param			EmailVerificationCode	body	string	true	"邮箱验证码"
type ResetPwdRequest struct {
	Email                 string `json:"email" xml:"email" form:"email" query:"email" validate:"required,email"`
	NewPassword           string `json:"new_password" xml:"new_password" form:"new_password" query:"new_password" validate:"required,min=6,max=20"`
	AgainNewPassword      string `json:"again_new_password" xml:"again_new_password" form:"again_new_password" query:"again_new_password" validate:"required,min=6,max=20"`
	EmailVerificationCode string `json:"email_verification_code" xml:"email_verification_code" form:"email_verification_code" query:"email_verification_code" validate:"required"`
}
