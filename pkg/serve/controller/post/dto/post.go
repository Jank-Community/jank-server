// Package dto 提供文章相关的数据传输对象定义
// 创建者：Done-0
// 创建时间：2025-05-10
package dto

// CreateOnePostRequest         发布文章的请求结构体
// @Param	title				body	string	true	"文章标题"
// @Param	image				body	string	true	"文章图片(可选)"
// @Param	visibility			body	string	true	"文章可见性(可选,默认 private)"
// @Param	content_html	    body	string	true	"文章内容(markdown格式)"
// @Param	category_id			body	int64	true	"文章分类ID"
type CreateOnePostRequest struct {
	Title           string `json:"title" xml:"title" form:"title" query:"title" validate:"required,min=1,max=225"`
	Image           string `json:"image" xml:"image" form:"image" query:"image" default:""`
	Visibility      bool   `json:"visibility" xml:"visibility" form:"visibility" query:"visibility" validate:"omitempty,boolean" default:"false"`
	ContentMarkdown string `json:"content_markdown" xml:"content_markdown" form:"content_markdown" query:"content_markdown" default:""`
	CategoryID      int64  `json:"category_id" xml:"category_id" form:"category_id" query:"category_id" validate:"omitempty"`
}

// DeleteOnePostRequest    文章删除请求
// @Param id path int true "文章 ID"
type DeleteOnePostRequest struct {
	ID int64 `json:"id" xml:"id" form:"id" query:"id" validate:"required"`
}

// GetOnePostRequest        获取文章的请求结构体
// @Param	id		path	string	true	"文章 ID"
type GetOnePostRequest struct {
	ID int64 `json:"id" xml:"id" form:"id" query:"id" validate:"omitempty" default:"0"`
}

// UpdateOnePostRequest       更新文章请求参数结构体
// @Param   id   			  body    int	    	true      "文章 ID"
// @Param   title		      body    string        false	  "文章标题"
// @Param   image 		      body 	  string        false     "文章图片(可选)"
// @Param   visibility 	      body 	  string        false     "文章可见性(可选)"
// @Param   content_markdown  body    string 		false     "文章内容(markdown格式)"
// @Param   category_id 	  body    int64         false     "文章分类ID列表(可选)"
type UpdateOnePostRequest struct {
	ID              int64  `json:"id" xml:"id" form:"id" query:"id" validate:"required"`
	Title           string `json:"title" xml:"title" form:"title" query:"title" validate:"min=0,max=255" default:""`
	Image           string `json:"image" xml:"image" form:"image" query:"image" default:""`
	Visibility      bool   `json:"visibility" xml:"visibility" form:"visibility" query:"visibility" validate:"omitempty,boolean" default:"false"`
	ContentMarkdown string `json:"content_markdown" xml:"content_markdown" form:"content_markdown" query:"content_markdown" default:""`
	CategoryID      int64  `json:"category_id" xml:"category_id" form:"category_id" query:"category_id" validate:"omitempty"`
}
