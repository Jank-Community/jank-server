// Package account 提供账户相关的视图对象定义
// 创建者：Done-0
// 创建时间：2025-05-10
package account

// GetAccountVO     获取账户信息请求体
// @Description	请求获取账户信息时所需参数
// @Property			email	    body	string	true	"用户邮箱"
// @Property			nickname	body	string	true	"用户昵称"
// @Property			phone	    body	string	true	"用户手机号"
type GetAccountVO struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

// LoginVO           返回给前端的登录信息
// @Description	登录成功后返回的访问令牌和刷新令牌
// @Property			access_token	body	string	true	"访问令牌"
// @Property			refresh_token	body	string	true	"刷新令牌"
type LoginVO struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// RegisterAccountVO     获取账户信息请求体
// @Description	请求获取账户信息时所需参数
// @Property			email	    body	string	true	"用户邮箱"
// @Property			nickname	body	string	true	"用户昵称"
// @Property			role_code	body	string	true	"用户角色编码"
type RegisterAccountVO struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}
