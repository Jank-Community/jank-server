// Package oss 提供对象存储相关的视图对象定义
// 创建者：Done-0
// 创建时间：2025-05-10
package oss

// UploadVO 上传文件响应
// @Description 上传文件后返回的对象存储路径
// @Property ObjectPath body string true "对象存储路径"
type UploadVO struct {
	ObjectPath string `json:"object_path"`
}

// DownloadVO 文件下载信息
// @Description 文件下载信息，包括文件名、下载链接和链接过期时间
// @Property FileName    body string true "文件名"
// @Property DownloadURL body string true "下载链接"
// @Property ExpiresAt   body string true "链接过期时间"
type DownloadVO struct {
	FileName    string `json:"file_name"`
	DownloadURL string `json:"download_url"`
	ExpiresAt   string `json:"expires_at"`
}

// ListObjectsVO 列出对象响应
// @Description 列出对象存储中的所有对象
// @Property Objects body []string true "对象列表"
// @Property Objects body string   true "对象列表"
type ListObjectsVO struct {
	Objects []string `json:"objects"`
}
