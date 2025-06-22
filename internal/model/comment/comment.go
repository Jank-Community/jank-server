// Package model 提供评论数据模型定义
// 创建者：Done-0
// 创建时间：2025-05-10
package model

import (
	"jank.com/jank_blog/internal/model/base"
	"jank.com/jank_blog/pkg/enums"
)

// Comment 评论模型
type Comment struct {
	base.Base
	Content          string            `gorm:"type:varchar(1024);not null" json:"content"`                      // 评论内容
	AccountId        int64             `gorm:"type:bigint;not null;index" json:"account_id"`                    // 所属用户ID
	AuditReason      string            `gorm:"type:varchar(255)" json:"audit_reason"`                           // 审核不通过原因
	AuditStatus      enums.AuditStatus `gorm:"type:varchar(10);not null;default:'pending'" json:"audit_status"` // 审核状态，默认值为 "pending"
	PostId           int64             `gorm:"type:bigint;not null;index" json:"post_id"`                       // 所属文章ID
	ReplyToCommentId int64             `gorm:"type:bigint;default:null" json:"reply_to_comment_id"`             // 目标评论ID
	Replies          []*Comment        `gorm:"-" json:"replies"`                                                // 子评论列表，用于构建图结构
}

// TableName 指定表名
// 返回值：
//   - string: 表名
func (Comment) TableName() string {
	return "comments"
}
