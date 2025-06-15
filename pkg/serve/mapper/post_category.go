// Package mapper 提供数据模型与数据库交互的映射层，处理文章与类目关联的数据操作
// 创建者：Done-0
// 创建时间：2025-05-10
package mapper

import (
	"fmt"

	"github.com/labstack/echo/v4"

	association "jank.com/jank_blog/internal/model/association"
	"jank.com/jank_blog/internal/utils"
)

// CreateOnePostCategory 创建文章-类目关联
// 参数：
//   - c: Echo 上下文
//   - postID: 文章 ID
//   - categoryID: 类目 ID
//
// 返回值：
//   - error: 操作过程中的错误
func CreateOnePostCategory(c echo.Context, postID, categoryID int64) error {
	postCategory := &association.PostCategory{
		PostID:     postID,
		CategoryID: categoryID,
	}
	db := utils.GetDBFromContext(c)
	if err := db.Create(postCategory).Error; err != nil {
		return fmt.Errorf("创建文章-类目关联失败: %w", err)
	}
	return nil
}

// GetOnePostCategoryByPostID 获取文章-类目关联
// 参数：
//   - c: Echo 上下文
//   - postID: 文章 ID
//
// 返回值：
//   - *association.PostCategory: 文章-类目关联信息
//   - error: 操作过程中的错误
func GetOnePostCategoryByPostID(c echo.Context, postID int64) (*association.PostCategory, error) {
	var postCategory association.PostCategory
	db := utils.GetDBFromContext(c)
	err := db.Where("post_id = ? AND deleted = ?", postID, false).First(&postCategory).Error
	if err != nil {
		if err.Error() == "record not found" {
			return nil, fmt.Errorf("文章-类目关联不存在: %w", err)
		}
		return nil, fmt.Errorf("获取文章-类目关联失败: %w", err)
	}
	return &postCategory, nil
}

// UpdateOnePostCategoryByPostID 更新文章-类目关联
// 参数：
//   - c: Echo 上下文
//   - postID: 文章 ID
//   - categoryID: 类目 ID
//
// 返回值：
//   - error: 操作过程中的错误
func UpdateOnePostCategoryByPostID(c echo.Context, postID, categoryID int64) error {
	var exists int64
	db := utils.GetDBFromContext(c)
	if err := db.Where("post_id = ? AND deleted = ?", postID, false).
		Count(&exists).Error; err != nil {
		return fmt.Errorf("检查文章-类目关联失败: %w", err)
	}
	if exists > 0 {
		if err := db.Model(&association.PostCategory{}).Where("post_id = ? AND deleted = ?", postID, false).
			Update("category_id", categoryID).Error; err != nil {
			return fmt.Errorf("更新文章-类目关联失败: %w", err)
		}
	} else {
		return CreateOnePostCategory(c, postID, categoryID)
	}
	return nil
}

// DeleteOnePostCategoryByPostID 删除文章-类目关联
// 参数：
//   - c: Echo 上下文
//   - postID: 文章 ID
//
// 返回值：
//   - error: 操作过程中的错误
func DeleteOnePostCategoryByPostID(c echo.Context, postID int64) error {
	db := utils.GetDBFromContext(c)
	if err := db.Model(&association.PostCategory{}).Where("post_id = ? AND deleted = ?", postID, false).
		Update("deleted", true).Error; err != nil {
		return fmt.Errorf("删除文章-类目关联失败: %w", err)
	}
	return nil
}

// DeleteOnePostCategoryByCategoryID 根据类目ID删除文章-类目关联
// 参数：
//   - c: Echo 上下文
//   - categoryID: 类目 ID
//
// 返回值：
//   - error: 操作过程中的错误
func DeleteOnePostCategoryByCategoryID(c echo.Context, categoryID int64) error {
	db := utils.GetDBFromContext(c)
	if err := db.Model(&association.PostCategory{}).Where("category_id = ? AND deleted = ?", categoryID, false).
		Update("deleted", true).Error; err != nil {
		return fmt.Errorf("根据类目ID删除文章-类目关联失败: %w", err)
	}
	return nil
}
