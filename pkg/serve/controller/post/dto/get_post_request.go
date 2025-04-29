package dto

// GetOnePostRequest        获取文章的请求结构体
// @Param	id		path	string	true	"文章 ID"
type GetOnePostRequest struct {
	ID int64 `json:"id" xml:"id" form:"id" query:"id" validate:"omitempty" default:"0"`
}
