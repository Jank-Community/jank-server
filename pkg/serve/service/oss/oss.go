// Package service 提供对象存储相关的业务逻辑处理
// 创建者：Done-0
// 创建时间：2025-05-10
package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/url"
	"path/filepath"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"

	"jank.com/jank_blog/internal/global"
	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/serve/controller/oss/dto"
	"jank.com/jank_blog/pkg/vo/oss"
)

// UploadOneFile 上传文件到 MinIO（支持大文件流式传输）
// 参数：
//   - c: echo.Context 上下文
//   - req: 上传文件请求
//
// 返回值：
//   - *oss.UploadVO: 上传结果
//   - error: 上传过程中的错误
func UploadOneFile(c echo.Context, req *dto.UploadOneFileRequest) (*oss.UploadVO, error) {
	file, err := c.FormFile("upload_file")
	if err != nil {
		utils.BizLogger(c).Errorf("获取上传文件失败: %v", err)
		return nil, fmt.Errorf("获取上传文件失败: %w", err)
	}
	if file.Size > utils.MAX_FILE_SIZE {
		maxMB := float64(utils.MAX_FILE_SIZE) / (1024 * 1024)
		return nil, fmt.Errorf("文件大小超过限制（最大%.2fMB）", maxMB)
	}

	src, err := file.Open()
	if err != nil {
		utils.BizLogger(c).Errorf("打开上传文件失败: %v", err)
		return nil, fmt.Errorf("打开上传文件失败: %w", err)
	}
	defer func(src multipart.File) {
		if err := src.Close(); err != nil {
			utils.BizLogger(c).Errorf("关闭上传文件失败: %v", err)
		}
	}(src)

	// 生成唯一文件名：原文件名_雪花ID.后缀
	snowID, err := utils.GenerateID()
	if err != nil {
		utils.BizLogger(c).Errorf("生成雪花ID失败: %v", err)
		return nil, fmt.Errorf("生成雪花ID失败: %w", err)
	}
	ext := filepath.Ext(file.Filename)
	base := file.Filename[:len(file.Filename)-len(ext)]
	objectName := fmt.Sprintf("%s_%d%s", base, snowID, ext)

	ctx := context.Background()
	exists, err := global.MinioClient.BucketExists(ctx, req.BucketName)
	if err != nil {
		utils.BizLogger(c).Errorf("检查桶存在时出错: %v", err)
		return nil, fmt.Errorf("检查桶存在时出错: %w", err)
	}
	if !exists {
		if err := global.MinioClient.MakeBucket(ctx, req.BucketName, minio.MakeBucketOptions{}); err != nil {
			utils.BizLogger(c).Errorf("创建桶失败: %v", err)
			return nil, fmt.Errorf("创建桶失败: %w", err)
		}
	}

	_, err = global.MinioClient.PutObject(ctx, req.BucketName, objectName, src, file.Size,
		minio.PutObjectOptions{ContentType: utils.GetMimeType(file.Filename)})
	if err != nil {
		utils.BizLogger(c).Errorf("上传到 MinIO 失败: %v", err)
		return nil, fmt.Errorf("上传到 MinIO 失败: %w", err)
	}

	return &oss.UploadVO{
		ObjectPath: fmt.Sprintf("/%s/%s", req.BucketName, objectName),
	}, nil
}

// DownloadOneFile 获取文件下载链接
// 参数：
//   - c: echo.Context 上下文
//   - req: 下载文件请求
//
// 返回值：
//   - *oss.DownloadVO: 下载信息
//   - error: 下载过程中的错误
func DownloadOneFile(c echo.Context, req *dto.DownloadOneFileRequest) (*oss.DownloadVO, error) {
	stat, err := global.MinioClient.StatObject(context.Background(), req.BucketName, req.ObjectName, minio.StatObjectOptions{})
	if err != nil {
		utils.BizLogger(c).Errorf("文件不存在: %v", err)
		return nil, fmt.Errorf("文件不存在: %w", err)
	}

	// 设置请求参数
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", fmt.Sprintf("attachment; filename*=UTF-8''%s", url.QueryEscape(req.ObjectName)))
	reqParams.Set("response-content-type", stat.ContentType)
	reqParams.Set("response-content-encoding", "UTF-8")

	// 设置过期时间（48 hours）
	expiration := 48 * time.Hour
	expiresAt := time.Now().Add(expiration)

	// 生成预签名 URL
	presignedURL, err := global.MinioClient.PresignedGetObject(context.Background(), req.BucketName, req.ObjectName, expiration, reqParams)
	if err != nil {
		utils.BizLogger(c).Errorf("生成预签名URL失败: %v", err)
		return nil, fmt.Errorf("生成预签名URL失败: %w", err)
	}

	// 设置响应头
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	return &oss.DownloadVO{
		FileName:    req.ObjectName,
		DownloadURL: presignedURL.String(),
		ExpiresAt:   expiresAt.Format("2006-01-02 15:04:05"),
	}, nil
}

// DeleteOneFile 从 MinIO 删除文件
// 参数：
//   - c: echo.Context 上下文
//   - req: 删除文件请求
//
// 返回值：
//   - error: 删除过程中的错误
func DeleteOneFile(c echo.Context, req *dto.DeleteOneFileRequest) error {
	err := global.MinioClient.RemoveObject(context.Background(), req.BucketName, req.ObjectName, minio.RemoveObjectOptions{})
	if err != nil {
		utils.BizLogger(c).Errorf("从 MinIO 删除文件失败: %v", err)
		return fmt.Errorf("从 MinIO 删除文件失败: %w", err)
	}
	return nil
}

// ListAllObjects 列出 MinIO 桶中的对象
// 参数：
//   - c: echo.Context 上下文
//   - req: 列出对象请求
//
// 返回值：
//   - *oss.ListObjectsVO: 对象列表
//   - error: 操作过程中的错误
func ListAllObjects(c echo.Context, req *dto.ListAllObjectsRequest) (*oss.ListObjectsVO, error) {
	objectCh := global.MinioClient.ListObjects(context.Background(), req.BucketName, minio.ListObjectsOptions{
		Prefix:    req.Prefix,
		Recursive: true,
	})

	var objects []string
	for object := range objectCh {
		if object.Err != nil {
			utils.BizLogger(c).Errorf("列出对象时失败: %v", object.Err)
			return nil, fmt.Errorf("列出对象时失败: %w", object.Err)
		}
		objects = append(objects, object.Key)
	}

	return &oss.ListObjectsVO{
		Objects: objects,
	}, nil
}
