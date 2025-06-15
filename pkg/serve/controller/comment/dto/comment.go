// Package dto 提供评论相关的数据传输对象定义
// 创建者：Done-0
// 创建时间：2025-05-10
package dto

// CreateCommentRequest 创建评论请求
// @Param Content   		  body string  true  "评论内容"
// @Param post_id   		  body int64   true  "文章 ID"
// @Param reply_to_comment_id body int64   false "回复的评论 ID"
type CreateCommentRequest struct {
	Content          string `json:"content" xml:"content" form:"content" query:"content" validate:"required,min=1,max=1024"`
	PostID           int64  `json:"post_id,string" xml:"post_id,string" form:"post_id,string" query:"post_id" validate:"required"`
	ReplyToCommentId int64  `json:"reply_to_comment_id,string" xml:"reply_to_comment_id,string" form:"reply_to_comment_id,string" query:"reply_to_comment_id" validate:"omitempty"`
}

// DeleteCommentRequest 删除评论请求
// @Param ID path int64 true "评论 ID"
type DeleteCommentRequest struct {
	ID int64 `json:"id,string" xml:"id,string" form:"id,string" query:"id" validate:"required"`
}

// GetCommentGraphRequest 获取评论请求
// @Param PostID path int64 true "帖子 ID"
type GetCommentGraphRequest struct {
	PostID int64 `json:"post_id,string" xml:"post_id,string" form:"post_id,string" query:"post_id" validate:"required"`
}

// GetOneCommentRequest 获取评论请求
// @Param ID path int64 true "评论 ID"
type GetOneCommentRequest struct {
	ID int64 `json:"id,string" xml:"id,string" form:"id,string" query:"id" validate:"required"`
}
