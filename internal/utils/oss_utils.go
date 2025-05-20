// Package utils 提供对象存储相关的工具函数
// 创建者：Done-0
// 创建时间：2025-05-10
package utils

import (
	"mime"
	"path/filepath"
	"strings"
)

const (
	MAX_FILE_SIZE = 100 * 1024 * 1024 // 100 MB
)

var mimeTypes = map[string]string{
	"jpg":  "image/jpeg",
	"jpeg": "image/jpeg",
	"png":  "image/png",
	"gif":  "image/gif",
	"webp": "image/webp",
	"svg":  "image/svg+xml",
	"bmp":  "image/bmp",
	"ico":  "image/x-icon",
	"tiff": "image/tiff",
	"pdf":  "application/pdf",
	"doc":  "application/msword",
	"docx": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	"xls":  "application/vnd.ms-excel",
	"xlsx": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	"ppt":  "application/vnd.ms-powerpoint",
	"pptx": "application/vnd.openxmlformats-officedocument.presentationml.presentation",
	"txt":  "text/plain",
	"md":   "text/markdown",
	"csv":  "text/csv",
	"rtf":  "application/rtf",
	"html": "text/html",
	"htm":  "text/html",
	"css":  "text/css",
	"js":   "application/javascript",
	"json": "application/json",
	"xml":  "application/xml",
	"yaml": "application/x-yaml",
	"yml":  "application/x-yaml",
	"zip":  "application/zip",
	"rar":  "application/x-rar-compressed",
	"7z":   "application/x-7z-compressed",
	"tar":  "application/x-tar",
	"gz":   "application/gzip",
	"bz2":  "application/x-bzip2",
	"xz":   "application/x-xz",
	"mp3":  "audio/mpeg",
	"wav":  "audio/wav",
	"ogg":  "audio/ogg",
	"flac": "audio/flac",
	"aac":  "audio/aac",
	"mp4":  "video/mp4",
	"avi":  "video/x-msvideo",
	"wmv":  "video/x-ms-wmv",
	"flv":  "video/x-flv",
	"mov":  "video/quicktime",
	"webm": "video/webm",
	"mkv":  "video/x-matroska",
	"mpeg": "video/mpeg",
	"3gp":  "video/3gpp",
}

// GetMimeType 根据文件扩展名获取MIME类型
// 参数：
//   - path: 文件路径
//
// 返回值：
//   - string: MIME类型
func GetMimeType(filename string) string {
	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(filename), "."))
	if mimeType, ok := mimeTypes[ext]; ok {
		return mimeType
	}
	return mime.TypeByExtension("." + ext)
}
