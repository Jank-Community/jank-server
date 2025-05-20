# MinIO 组件

MinIO 组件用于处理应用程序与 MinIO 对象存储服务器的连接和交互。该组件提供了一个全局 MinIO 客户端实例，可以在整个应用程序中使用。

## 功能

- **连接管理**: 创建并维护与 MinIO 服务器的连接
- **连接检查**: 通过 ListBuckets 操作验证连接状态和访问权限
- **SSL 支持**: 支持安全连接配置
- **错误处理**: 提供详细的错误日志记录

## 配置项

MinIO 连接从应用配置中读取以下参数：

- 主机地址 (MinioHost)
- 端口 (MinioPort)
- 访问密钥 (MinioAccessKey)
- 密钥 (MinioSecretKey)
- SSL 配置 (MinioUseSsl)

## 使用方式

通过全局变量 `global.MinioClient` 在应用的任何位置访问 MinIO 客户端，例如：

```go
// 上传文件
global.MinioClient.FPutObject(bucketName, objectName, filePath, opts)

// 下载文件
global.MinioClient.FGetObject(bucketName, objectName, filePath, opts)

// 列出存储桶
global.MinioClient.ListBuckets()
```

## 注意事项

1. 确保在应用启动时正确初始化 MinIO 客户端
2. 所有 MinIO 操作都应该进行适当的错误处理
3. 建议使用 utils 包中的工具函数进行文件操作，而不是直接使用 MinIO 客户端
