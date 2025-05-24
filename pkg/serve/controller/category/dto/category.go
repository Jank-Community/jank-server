// Package dto 提供类目相关的数据传输对象定义
// 创建者：Done-0
// 创建时间：2025-05-10
package dto

// CreateOneCategoryRequest       创建类目请求
// @Param name        body string true  "类目名称"
// @Param description body string false "类目描述"
// @Param parent_id   body int64  false "父类目ID"
type CreateOneCategoryRequest struct {
	Name        string `json:"name" xml:"name" form:"name" query:"name" validate:"required,min=1"`
	Description string `json:"description" xml:"description" form:"description" query:"description" default:""`
	ParentID    int64  `json:"parent_id,string" xml:"parent_id,string" form:"parent_id,string" query:"parent_id,string" validate:"omitempty"`
}

// DeleteOneCategoryRequest  删除类目请求
// @Param id path int64 true "类目ID"
type DeleteOneCategoryRequest struct {
	ID int64 `json:"id,string" xml:"id,string" form:"id,string" query:"id,string" validate:"required"`
}

// GetOneCategoryRequest 更新类目请求
// @Param id path int64 true "类目ID"
type GetOneCategoryRequest struct {
	ID int64 `json:"id,string" xml:"id,string" form:"id,string" query:"id,string" validate:"required"`
}

// UpdateOneCategoryRequest    更新类目请求
// @Param id          path     int64   true  "类目ID"
// @Param name        body     string  true  "类目名称"
// @Param description body     string  false "类目描述"
// @Param parent_id   body     int64   false "父类目ID"
type UpdateOneCategoryRequest struct {
	ID          int64  `json:"id,string" xml:"id,string" form:"id,string" query:"id,string" validate:"required"`
	Name        string `json:"name" xml:"name" form:"name" query:"name" validate:"required,min=1,max=255"`
	Description string `json:"description" xml:"description" form:"description" query:"description" default:""`
	ParentID    int64  `json:"parent_id,string" xml:"parent_id,string" form:"parent_id,string" query:"parent_id,string" validate:"omitempty"`
}
