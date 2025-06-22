// Package mapper 提供数据模型与数据库交互的映射层，处理评论相关数据操作
// 创建者：Done-0
// 创建时间：2025-05-10
package mapper

import (
	"fmt"

	"github.com/labstack/echo/v4"

	model "jank.com/jank_blog/internal/model/comment"
	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/enums"
)

// CreateComment 保存评论到数据库
// 参数：
//   - c: Echo 上下文
//   - comment: 评论信息
//
// 返回值：
//   - error: 操作过程中的错误
func CreateComment(c echo.Context, comment *model.Comment) error {
	db := utils.GetDBFromContext(c)
	if err := db.Create(comment).Error; err != nil {
		return fmt.Errorf("创建评论失败: %w", err)
	}
	return nil
}

// GetCommentByID 根据 ID 查询评论
// 参数：
//   - c: Echo 上下文
//   - id: 评论 ID
//   - ignoreAudit: 是否忽略审核状态
//
// 返回值：
//   - *model.Comment: 评论信息
//   - error: 操作过程中的错误
func GetCommentByID(c echo.Context, id int64, ignoreAudit bool) (*model.Comment, error) {
	var comment model.Comment
	var query string
	var args []interface{}
	db := utils.GetDBFromContext(c)
	// 如果 all 为 true，则不考虑审核状态
	if ignoreAudit {
		query = "id = ? AND deleted = ?"
		args = append(args, id, false)
	} else {
		query = "id = ? AND deleted = ? AND audit_status = ?"
		args = append(args, id, false, enums.AuditApproved)
	}

	if err := db.Where(query, args...).First(&comment).Error; err != nil {
		return nil, fmt.Errorf("获取评论失败: %w", err)
	}
	return &comment, nil
}

// GetReplyByCommentID 获取评论的所有回复
// 参数：
//   - c: Echo 上下文
//   - id: 评论 ID
//   - ignoreAudit: 是否忽略审核状态
//
// 返回值：
//   - []*model.Comment: 回复列表
//   - error: 操作过程中的错误
func GetReplyByCommentID(c echo.Context, id int64, ignoreAudit bool) ([]*model.Comment, error) {
	var comments []*model.Comment
	var query string
	var args []interface{}
	db := utils.GetDBFromContext(c)
	// 如果 all 为 true，则不考虑审核状态
	if ignoreAudit {
		query = "reply_to_comment_id = ? AND deleted = ?"
		args = append(args, id, false)
	} else {
		query = "reply_to_comment_id = ? AND deleted = ? AND audit_status = ?"
		args = append(args, id, false, enums.AuditApproved)
	}
	if err := db.Where(query, args...).Find(&comments).Error; err != nil {
		return nil, fmt.Errorf("获取评论回复失败: %w", err)
	}
	return comments, nil
}

// GetCommentsByPostID 根据文章 ID 查询所有评论
// 参数：
//   - c: Echo 上下文
//   - postID: 文章 ID
//   - ignoreAudit: 是否忽略审核状态
//
// 返回值：
//   - []*model.Comment: 评论列表
//   - error: 操作过程中的错误
func GetCommentsByPostID(c echo.Context, postID int64, ignoreAudit bool) ([]*model.Comment, error) {
	var comments []*model.Comment
	var query string
	var args []interface{}
	db := utils.GetDBFromContext(c)
	// 如果 all 为 true，则不考虑审核状态
	if ignoreAudit {
		query = "post_id = ? AND deleted = ?"
		args = append(args, postID, false)
	} else {
		query = "post_id = ? AND deleted = ? AND audit_status = ?"
		args = append(args, postID, false, enums.AuditApproved)
	}
	if err := db.Where(query, args...).Find(&comments).Error; err != nil {
		return nil, fmt.Errorf("获取文章评论失败: %w", err)
	}
	return comments, nil
}

// GetPendingComments 获取所有待审核的评论
// 参数：
//   - c: Echo 上下文
//   - page: 页码
//   - pageSize: 每页数量
//
// 返回值：
//   - []*model.Comment: 待审核评论列表
//   - int64: 待审核评论总数
//   - error: 操作过程中的错误
func GetPendingComments(c echo.Context, page, pageSize int) ([]*model.Comment, int64, error) {
	var comments []*model.Comment
	var total int64
	db := utils.GetDBFromContext(c)

	// 查询待审核评论总数
	if err := db.Model(&model.Comment{}).Where("audit_status = ? AND deleted = ?", enums.AuditPending, false).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("获取待审核评论总数失败: %w", err)
	}

	// 分页查询待审核评论，从旧到新
	if err := db.Where("audit_status = ? AND deleted = ?", enums.AuditPending, false).
		Order("id ASC").
		Limit(pageSize).Offset((page - 1) * pageSize).
		Find(&comments).Error; err != nil {
		return nil, 0, fmt.Errorf("获取待审核评论失败: %w", err)
	}
	return comments, total, nil
}

// UpdateComment 更新评论
// 参数：
//   - c: Echo 上下文
//   - comment: 评论信息
//
// 返回值：
//   - error: 操作过程中的错误
func UpdateComment(c echo.Context, comment *model.Comment) error {
	db := utils.GetDBFromContext(c)
	if err := db.Save(comment).Error; err != nil {
		return fmt.Errorf("更新评论失败: %w", err)
	}
	return nil
}
