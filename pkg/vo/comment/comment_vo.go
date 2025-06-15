// Package comment 提供评论相关的视图对象定义
// 创建者：Done-0
// 创建时间：2025-05-10
package comment

// CommentsVO 获取评论响应
// @Description 获取单个评论的响应
// @Property ID                  body string  			  true  "评论唯一标识"
// @Property Content             body string  		      true  "评论内容"
// @Property AccountId           body string              true  "评论所属用户 ID"
// @Property PostId              body string              true  "评论所属文章 ID"
// @Property ReplyToCommentId    body string              true  "回复的目标评论 ID"
// @Property Replies             body []*CommentsVO       true  "子评论列表"
type CommentsVO struct {
	ID               string        `json:"id"`
	Content          string        `json:"content"`
	AccountId        string        `json:"account_id"`
	PostId           string        `json:"post_id"`
	ReplyToCommentId string        `json:"reply_to_comment_id"`
	Replies          []*CommentsVO `json:"replies"`
}
