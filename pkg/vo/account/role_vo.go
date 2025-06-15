package account

// RoleVO 角色返回值对象
// @Description 角色信息的返回结构
// @Property   id           int64  "角色 ID"
// @Property   name         string "角色名称"
// @Property   description  string "角色描述"
// @Property   status       bool "角色状态"
type RoleVO struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}
