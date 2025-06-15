// Package mapper 提供数据模型与数据库交互的映射层，处理评论相关数据操作
// 创建者：Done-0
// 创建时间：2025-05-10
package mapper

import (
	"fmt"

	"github.com/labstack/echo/v4"

	comment "jank.com/jank_blog/internal/model/comment"
	"jank.com/jank_blog/internal/utils"
)

// CreateOneComment 保存评论到数据库
// 参数：
//   - c: Echo 上下文
//   - comment: 评论信息
//
// 返回值：
//   - error: 操作过程中的错误
func CreateOneComment(c echo.Context, comment *comment.Comment) error {
	db := utils.GetDBFromContext(c)
	if err := db.Create(comment).Error; err != nil {
		return fmt.Errorf("创建评论失败: %w", err)
	}
	return nil
}

// GetOneCommentByID 根据 ID 查询评论
// 参数：
//   - c: Echo 上下文
//   - id: 评论 ID
//
// 返回值：
//   - *model.Comment: 评论信息
//   - error: 操作过程中的错误
func GetOneCommentByID(c echo.Context, id int64) (*comment.Comment, error) {
	var com comment.Comment
	db := utils.GetDBFromContext(c)
	if err := db.Where("id = ? AND deleted = ?", id, false).First(&com).Error; err != nil {
		return nil, fmt.Errorf("获取评论失败: %w", err)
	}
	return &com, nil
}

// GetOneReplyByCommentID 获取评论的所有回复
// 参数：
//   - c: Echo 上下文
//   - id: 评论 ID
//
// 返回值：
//   - []*model.Comment: 回复列表
//   - error: 操作过程中的错误
func GetOneReplyByCommentID(c echo.Context, id int64) ([]*comment.Comment, error) {
	var comments []*comment.Comment
	db := utils.GetDBFromContext(c)
	if err := db.Where("reply_to_comment_id = ? AND deleted = ?", id, false).Find(&comments).Error; err != nil {
		return nil, fmt.Errorf("获取评论回复失败: %w", err)
	}
	return comments, nil
}

// GetOneCommentsByPostID 根据文章 ID 查询所有评论
// 参数：
//   - c: Echo 上下文
//   - postID: 文章 ID
//
// 返回值：
//   - []*model.Comment: 评论列表
//   - error: 操作过程中的错误
func GetOneCommentsByPostID(c echo.Context, postID int64) ([]*comment.Comment, error) {
	var comments []*comment.Comment
	db := utils.GetDBFromContext(c)
	if err := db.Where("post_id = ? AND deleted = ?", postID, false).Find(&comments).Error; err != nil {
		return nil, fmt.Errorf("获取文章评论失败: %w", err)
	}
	return comments, nil
}

// UpdateOneComment 更新评论
// 参数：
//   - c: Echo 上下文
//   - comment: 评论信息
//
// 返回值：
//   - error: 操作过程中的错误
func UpdateOneComment(c echo.Context, comment *comment.Comment) error {
	db := utils.GetDBFromContext(c)
	if err := db.Where("id = ? AND deleted = ?", comment.ID, false).Updates(comment).Error; err != nil {
		return fmt.Errorf("更新评论失败: %w", err)
	}
	return nil
}
