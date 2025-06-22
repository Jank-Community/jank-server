package mq

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"jank.com/jank_blog/internal/global"
	model "jank.com/jank_blog/internal/model/comment"
	"jank.com/jank_blog/pkg/enums"
)

// PushCommentToStream 将评论消息推送到 Redis Stream 中
// 参数：
//   - c: 上下文，用于取消操作
//   - redisClient: Redis 客户端
//   - storage: 存储结构体
//   - commentID: 评论的唯一标识符
//   - content: 评论内容
//
// 返回值：
//   - string: 推送的消息 ID
//   - error: 如果推送过程中发生错误，则返回错误
func PushCommentToStream(c context.Context, redisClient *redis.Client, storage *global.Storage, commentID, content string) (string, error) {
	exists, err := redisClient.HExists(c, storage.MapKey, commentID).Result()
	if err != nil {
		return "", err
	}
	if exists {
		return redisClient.HGet(c, storage.MapKey, commentID).Result()
	}
	message := map[string]interface{}{
		"comment_id": commentID,
		"content":    content,
	}

	// 使用 Redis XAdd 将评论消息推送到流中
	msgID, err := redisClient.XAdd(context.Background(), &redis.XAddArgs{
		Stream: storage.StreamKey,
		MaxLen: int64(storage.MaxLength),
		Values: message,
	}).Result()
	if err != nil {
		return "", err
	}
	err = redisClient.HSet(c, storage.MapKey, commentID, msgID).Err()
	if err != nil {
		return "", err
	}
	return msgID, nil
}

// RestoreMessagesFromDB 从数据库中恢复评论消息到 Redis Stream
// 参数：
//   - db: 数据库连接
//   - redisClient: Redis 客户端
//   - storage: 存储结构体
//
// 返回值：
//   - error: 如果恢复过程中发生错误，则返回错误
func RestoreMessagesFromDB(db *gorm.DB, redisClient *redis.Client, storage *global.Storage) error {
	var comments []*model.Comment
	// 从数据库中查询所有未删除的评论
	if err := db.Where("audit_status = ? AND deleted = ?", enums.AuditPending, false).
		Order("id ASC").Limit(storage.MaxLength).Find(&comments).Error; err != nil {
		return err
	}

	for _, comment := range comments {
		_, err := PushCommentToStream(context.Background(), redisClient, storage, fmt.Sprintf("%d", comment.ID), comment.Content)
		if err != nil {
			return err
		}
	}

	return nil
}
