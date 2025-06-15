// Package category 提供类目相关的视图对象定义
// 创建者：Done-0
// 创建时间：2025-05-10
package category

// CategoriesVO 获取类目响应
// @Description 获取类目响应
// @Property		ID			body	string	true	"类目唯一标识"
// @Property		Name		body	string	true	"类目名称"
// @Property		Description	body	string	true	"类目描述"
// @Property		ParentID	body	string	true	"父类目 ID"
// @Property		Path		body	string	true	"类目路径"
// @Property		Children	body	[]*CategoriesVO	true	"子类目列表"
type CategoriesVO struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	ParentID    string          `json:"parent_id"`
	Path        string          `json:"path"`
	Children    []*CategoriesVO `json:"children"`
}
