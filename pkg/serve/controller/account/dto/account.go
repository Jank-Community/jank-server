// Package dto 提供账户相关的数据传输对象定义
// 创建者：Done-0
// 创建时间：2025-05-10
package dto

// RegisterRequest  用户注册请求体
// @Description	用户注册所需参数
// @Param			email		body	string	true	"用户邮箱"
// @Param			phone		body	string	true	"用户手机号"
// @Param			nickname	body	string	true	"用户昵称"
// @Param			password	body	string	true	"用户密码"
// @Param			email_verification_code	body	string	true	"用户邮箱验证码"
// @Param			img_verification_code	body	string	true	"用户图片验证码"
type RegisterRequest struct {
	Email                 string `json:"email" xml:"email" form:"email" query:"email" validate:"required"`
	Phone                 string `json:"phone" xml:"phone" form:"phone" query:"phone" default:""`
	Nickname              string `json:"nickname" xml:"nickname" form:"nickname" query:"nickname" validate:"required,min=1,max=20"`
	Password              string `json:"password" xml:"password" form:"password" query:"password" validate:"required,min=6,max=20"`
	EmailVerificationCode string `json:"email_verification_code" xml:"email_verification_code" form:"email_verification_code" query:"email_verification_code" validate:"required"`
	ImgVerificationCode   string `json:"img_verification_code" xml:"img_verification_code" form:"img_verification_code" query:"img_verification_code" validate:"required"`
}

// LoginRequest 用户登录请求体
// @Description	用户登录请求所需参数
// @Param			email		body	string	true	"用户邮箱"
// @Param			password	body	string	true	"用户密码"
// @Param			img_verification_code	body	string	true	"图片验证码"
type LoginRequest struct {
	Email               string `json:"email" xml:"email" form:"email" query:"email" validate:"required,email"`
	Password            string `json:"password" xml:"password" form:"password" query:"password" validate:"required"`
	ImgVerificationCode string `json:"img_verification_code" xml:"img_verification_code" form:"img_verification_code" query:"img_verification_code" validate:"required"`
}

// GetAccountRequest            获取账户信息请求体
// @Description	请求获取账户信息时所需参数
// @Param			email	    body	string	true	"用户邮箱"
type GetAccountRequest struct {
	Email string `json:"email" xml:"email" form:"email" query:"email" validate:"required,email"`
}

// UpdateAccountRequest  更新账户信息请求体
// @Description	用户更新账户信息所需参数
// @Param			nickname	body	string	true	"用户昵称"
// @Param			phone		body	string	true	"用户手机号"
// @Param			avatar		body	string	true	"用户头像"
type UpdateAccountRequest struct {
	Nickname string `json:"nickname" xml:"nickname" form:"nickname" query:"nickname" validate:"required,min=1,max=20"`
	Phone    string `json:"phone" xml:"phone" form:"phone" query:"phone" default:""`
	Avatar   string `json:"avatar" xml:"avatar" form:"avatar" query:"avatar" default:""`
}

// ResetPwdRequest  重置密码请求体
// @Description	用户重置密码所需参数
// @Param			email					body	string	true	"用户邮箱"
// @Param			new_password			body	string	true	"新密码"
// @Param			again_new_password		body	string	true	"再次输入新密码"
// @Param			email_verification_code	body	string	true	"邮箱验证码"
type ResetPwdRequest struct {
	Email                 string `json:"email" xml:"email" form:"email" query:"email" validate:"required,email"`
	NewPassword           string `json:"new_password" xml:"new_password" form:"new_password" query:"new_password" validate:"required,min=6,max=20"`
	AgainNewPassword      string `json:"again_new_password" xml:"again_new_password" form:"again_new_password" query:"again_new_password" validate:"required,min=6,max=20"`
	EmailVerificationCode string `json:"email_verification_code" xml:"email_verification_code" form:"email_verification_code" query:"email_verification_code" validate:"required"`
}
