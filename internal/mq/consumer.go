package mq

import (
	"context"
	"errors"
	"jank.com/jank_blog/internal/global"
	"strings"

	"github.com/redis/go-redis/v9"
)

// EnsureConsumerGroup 确保消费者组存在，如果不存在则创建
// 参数：
//   - redisClient: Redis 客户端
//   - storage: 存储结构体
//
// 返回值：
//   - error: 如果创建或检查过程中发生错误，则返回错误
func EnsureConsumerGroup(redisClient *redis.Client, storage *global.Storage) error {
	err := redisClient.XGroupCreateMkStream(context.Background(), storage.StreamKey, storage.GroupName, "0").Err()
	if err != nil && !strings.Contains(err.Error(), "BUSYGROUP") {
		return err
	}
	return nil
}

// AcknowledgeMessage 确认消息已被处理
// 参数：
//   - c: 上下文，用于取消操作
//   - redisClient: Redis 客户端
//   - storage: 存储结构体
//   - commentID: 评论的唯一标识符
//
// 返回值：
//   - error: 如果确认过程中发生错误，则返回错误
func AcknowledgeMessage(c context.Context, redisClient *redis.Client, storage *global.Storage, commentID string) error {
	messageID, err := redisClient.HGet(c, storage.MapKey, commentID).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil // 如果消息不存在，则不需要确认
		}
		return err
	}
	if err := redisClient.XAck(c, storage.StreamKey, storage.GroupName, messageID).Err(); err != nil {
		return err
	}
	if err := redisClient.HDel(c, storage.MapKey, commentID).Err(); err != nil {
		return err
	}
	return nil
}
