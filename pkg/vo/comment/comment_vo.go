// Package comment 提供评论相关的视图对象定义
// 创建者：Done-0
// 创建时间：2025-05-10
package comment

import "jank.com/jank_blog/pkg/enums"

// CommentsVO 获取评论响应
// @Description 获取单个评论的响应
// @Property id                  body string  			  true  "评论唯一标识"
// @Property content             body string  			  true  "评论内容"
// @Property account_id          body string              true  "评论所属用户ID"
// @Property audit_reason        body string              false "审核不通过原因"
// @Property audit_status        body enums.AuditStatus   true  "评论审核状态"
// @Property post_id             body string              true  "评论所属文章ID"
// @Property reply_to_comment_id body string              false "回复的目标评论ID"
// @Property replies             body []*CommentsVO true  "子评论列表"
type CommentsVO struct {
	ID               string            `json:"id"`
	Content          string            `json:"content"`
	AccountId        string            `json:"account_id"`
	AuditReason      string            `json:"audit_reason"`
	AuditStatus      enums.AuditStatus `json:"audit_status"`
	PostId           string            `json:"post_id"`
	ReplyToCommentId string            `json:"reply_to_comment_id"`
	Replies          []*CommentsVO     `json:"replies"`
}
