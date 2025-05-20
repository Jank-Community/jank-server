// Package oss 提供对象存储相关的视图对象定义
// 创建者：Done-0
// 创建时间：2025-05-10
package oss

// UploadVO 上传文件响应
type UploadVO struct {
	ObjectPath string `json:"object_path"` // 对象存储路径
}

// DownloadVO 文件下载信息
type DownloadVO struct {
	FileName    string `json:"file_name"`    // 文件名
	DownloadURL string `json:"download_url"` // 下载链接
	ExpiresAt   string `json:"expires_at"`   // 链接过期时间
}

// ListObjectsVO 列出对象响应
type ListObjectsVO struct {
	Objects []string `json:"objects"` // 对象列表
}
