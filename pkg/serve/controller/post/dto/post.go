// Package dto 提供文章相关的数据传输对象定义
// 创建者：Done-0
// 创建时间：2025-05-10
package dto

// CreateOnePostRequest         发布文章的请求结构体
// @Param	Title				body	string	true	"文章标题"
// @Param	Image				body	string	true	"文章图片(可选)"
// @Param	Visibility			body	string	true	"文章可见性(可选,默认 private)"
// @Param	ContentMarkdown	    body	string	true	"文章内容(markdown格式)"
// @Param	CategoryID			body	int64	true	"文章分类 ID"
type CreateOnePostRequest struct {
	Title           string `json:"title" xml:"title" form:"title" query:"title" validate:"required,min=1,max=225"`
	Image           string `json:"image" xml:"image" form:"image" query:"image"`
	Visibility      bool   `json:"visibility" xml:"visibility" form:"visibility" query:"visibility" validate:"omitempty,boolean"`
	ContentMarkdown string `json:"content_markdown" xml:"content_markdown" form:"content_markdown" query:"content_markdown"`
	CategoryID      int64  `json:"category_id,string" xml:"category_id,string" form:"category_id,string" query:"category_id" validate:"omitempty"`
}

// DeleteOnePostRequest    文章删除请求
// @Param ID body int true "文章 ID"
type DeleteOnePostRequest struct {
	ID int64 `json:"id,string" xml:"id,string" form:"id,string" query:"id" validate:"required"`
}

// GetOnePostRequest        获取文章的请求结构体
// @Param	ID		query	string	true	"文章 ID"
type GetOnePostRequest struct {
	ID int64 `json:"id,string" xml:"id,string" form:"id,string" query:"id" validate:"required"`
}

// UpdateOnePostRequest       更新文章请求参数结构体
// @Param   ID   			  body    int	    	true      "文章 ID"
// @Param   Title		      body    string        false	  "文章标题"
// @Param   Image 		      body 	  string        false     "文章图片(可选)"
// @Param   Visibility 	      body 	  string        false     "文章可见性(可选)"
// @Param   ContentMarkdown   body    string 		false     "文章内容(markdown格式)"
// @Param   CategoryID 	 	  body    int64         false     "文章分类ID列表(可选)"
type UpdateOnePostRequest struct {
	ID              int64  `json:"id,string" xml:"id,string" form:"id" query:"id" validate:"required"`
	Title           string `json:"title" xml:"title" form:"title" query:"title" validate:"min=0,max=255"`
	Image           string `json:"image" xml:"image" form:"image" query:"image"`
	Visibility      bool   `json:"visibility" xml:"visibility" form:"visibility" query:"visibility" validate:"omitempty,boolean"`
	ContentMarkdown string `json:"content_markdown" xml:"content_markdown" form:"content_markdown" query:"content_markdown"`
	CategoryID      int64  `json:"category_id,string" xml:"category_id,string" form:"category_id,string" query:"category_id" validate:"omitempty"`
}

// GetAllPostsRequest        获取文章列表的请求结构体
// @Param	Page		query	int	false	"页码"
// @Param	PageSize	query	int	false	"每页条数"
type GetAllPostsRequest struct {
	Page     int `json:"page" xml:"page" form:"page" query:"page" validate:"omitempty,min=1"`
	PageSize int `json:"page_size" xml:"page_size" form:"page_size" query:"page_size" validate:"omitempty,min=1,max=100"`
}
