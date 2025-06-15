// Package dto 提供类目相关的数据传输对象定义
// 创建者：Done-0
// 创建时间：2025-05-10
package dto

// CreateOneCategoryRequest       创建类目请求
// @Param Name        body string true  "类目名称"
// @Param Description body string false "类目描述"
// @Param ParentID    body int64  false "父类目 ID"
type CreateOneCategoryRequest struct {
	Name        string `json:"name" xml:"name" form:"name" query:"name" validate:"required,min=1"`
	Description string `json:"description" xml:"description" form:"description" query:"description" default:""`
	ParentID    int64  `json:"parent_id,string" xml:"parent_id" form:"parent_id" query:"parent_id" validate:"omitempty"`
}

// DeleteOneCategoryRequest  删除类目请求
// @Param ID path int64 true "类目 ID"
type DeleteOneCategoryRequest struct {
	ID int64 `json:"id,string" xml:"id" form:"id" query:"id" validate:"required"`
}

// GetOneCategoryRequest 获取类目请求
// @Param ID path int64 true "类目 ID"
type GetOneCategoryRequest struct {
	ID int64 `json:"id,string" xml:"id" form:"id" query:"id" validate:"required"`
}

// UpdateOneCategoryRequest    更新类目请求
// @Param ID          body     int64   true  "类目 ID"
// @Param Name        body     string  true  "类目名称"
// @Param Description body     string  false "类目描述"
// @Param ParentID    body     int64   false "父类目 ID"
type UpdateOneCategoryRequest struct {
	ID          int64  `json:"id,string" xml:"id" form:"id" query:"id" validate:"required"`
	Name        string `json:"name" xml:"name" form:"name" query:"name" validate:"required,min=1,max=255"`
	Description string `json:"description" xml:"description" form:"description" query:"description" default:""`
	ParentID    int64  `json:"parent_id,string" xml:"parent_id" form:"parent_id" query:"parent_id" validate:"omitempty"`
}
