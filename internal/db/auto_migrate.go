// Package db 提供数据库连接和管理功能
// 创建者：Done-0
// 创建时间：2025-05-10
package db

import (
	"fmt"
	"log"

	"jank.com/jank_blog/internal/global"
	"jank.com/jank_blog/internal/model"
)

// autoMigrate 执行数据库表结构自动迁移
func autoMigrate() error {
	err := global.DB.AutoMigrate(
		model.GetAllModels()...,
	)
	if err != nil {
		return fmt.Errorf("数据库自动迁移失败 %w", err)
	}

	log.Println("数据库自动迁移成功...")
	global.SysLog.Info("数据库自动迁移成功...")

	return nil
}
