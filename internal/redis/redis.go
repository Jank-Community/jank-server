// Package redis 提供Redis连接和管理功能
// 创建者：Done-0
// 创建时间：2025-05-10
package redis

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"

	"jank.com/jank_blog/configs"
	"jank.com/jank_blog/internal/global"
	"jank.com/jank_blog/internal/mq"
)

// New 初始化Redis连接
// 参数：
//   - config: 应用配置
func New(config *configs.Config) {
	client := newRedisClient(config)
	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Printf("Redis 连接失败: %v", err)
		global.SysLog.Errorf("Redis 连接失败: %v", err)
		return
	}
	storage := &global.Storage{
		MapKey:       config.CommentMQConfig.MapKey,
		StreamKey:    config.CommentMQConfig.StreamKey,
		GroupName:    config.CommentMQConfig.GroupName,
		ConsumerName: config.CommentMQConfig.ConsumerName,
		MaxLength:    config.CommentMQConfig.MaxLength,
	}

	global.MQStorage = storage
	global.RedisClient = client

	// 确保消费者组存在
	if err := mq.EnsureConsumerGroup(client, storage); err != nil {
		log.Printf("创建消费者组失败: %v", err)
		global.SysLog.Errorf("创建消费者组失败: %v", err)
		return
	}
	if err := mq.RestoreMessagesFromDB(global.DB, client, storage); err != nil {
		log.Printf("从数据库恢复消息失败: %v", err)
		global.SysLog.Errorf("从数据库恢复消息失败: %v", err)
		return
	}
	log.Println("Redis 连接成功...")
	global.SysLog.Infof("Redis 连接成功...")
}

// newRedisClient 创建新的Redis客户端
// 参数：
//   - config: 应用配置
//
// 返回值：
//   - *redis.Client: Redis客户端实例
func newRedisClient(config *configs.Config) *redis.Client {
	db, _ := strconv.Atoi(config.RedisConfig.RedisDB)
	return redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", config.RedisConfig.RedisHost, config.RedisConfig.RedisPort),
		Password:     config.RedisConfig.RedisPassword, // 数据库密码，默认为空字符串
		DB:           db,                               // 数据库索引
		DialTimeout:  10 * time.Second,                 // 连接超时时间
		ReadTimeout:  1 * time.Second,                  // 读超时时间
		WriteTimeout: 2 * time.Second,                  // 写超时时间
		PoolSize:     runtime.GOMAXPROCS(10),           // 最大连接池大小
		MinIdleConns: 50,                               // 最小空闲连接数
	})
}
