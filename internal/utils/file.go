// Package utils
// 提供文件操作相关的工具函数
// 创建者：Pey121
// 创建时间：2025-06-26
package utils

import "os"

// FileExists 检查文件是否存在
// 参数：
//   - path: 文件路径
//
// 返回值：
//   - bool: 文件是否存在
//   - error: 检查过程中发生的错误（如果有）
func FileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}
