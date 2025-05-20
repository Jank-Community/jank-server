// Package oss 提供对象存储服务功能
// 创建者：Done-0
// 创建时间：2025-05-10
package oss

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"jank.com/jank_blog/configs"
	"jank.com/jank_blog/internal/global"
)

// New 初始化 MinIO 客户端
// 参数：
//   - config: 应用配置
func New(config *configs.Config) {
	client := newMinioClient(config)
	if _, err := client.ListBuckets(context.Background()); err != nil {
		log.Printf("MinIO 连接失败: %v", err)
		global.SysLog.Errorf("MinIO 连接失败: %v", err)
		return
	}
	global.MinioClient = client

	log.Println("MinIO 连接成功...")
	global.SysLog.Infoln("MinIO 连接成功...")
}

// newMinioClient 创建新的 MinIO 客户端
// 参数：
//   - config: 应用配置（包含 MinIO 连接参数）
//
// 返回值：
//   - *minio.Client: MinIO 客户端实例
func newMinioClient(config *configs.Config) *minio.Client {
	minioConfig := config.MinioConfig
	client, err := minio.New(
		fmt.Sprintf("%s:%s", minioConfig.MinioHost, minioConfig.MinioPort), // MinIO 服务地址
		&minio.Options{
			Creds: credentials.NewStaticV4(
				minioConfig.MinioAccessKey,    // MinIO 访问密钥
				minioConfig.MinioSecretKey,    // MinIO 密钥
				minioConfig.MinioSessionToken, // MinIO 临时安全凭证（可选，通常为空字符串）
			),
			Secure: minioConfig.MinioUseSsl, // 是否启用 SSL
		},
	)
	if err != nil {
		global.SysLog.Errorf("创建 MinIO 客户端失败: %v", err)
		return nil
	}
	return client
}
