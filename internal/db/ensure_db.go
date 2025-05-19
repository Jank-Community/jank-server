// Package db 提供数据库连接和管理功能
// 创建者：Done-0
// 创建时间：2025-05-10
package db

import (
	"fmt"
	"log"

	"gorm.io/gorm"

	"jank.com/jank_blog/configs"
	"jank.com/jank_blog/internal/global"
)

// ensureDBExists 确保数据库存在，不存在则创建
// 参数：
//   - db: 数据库连接
//   - config: 应用配置
//
// 返回值：
//   - error: 创建过程中的错误
func ensureDBExists(db *gorm.DB, config *configs.Config) error {
	switch config.DBConfig.DBDialect {
	case DIALECT_POSTGRES:
		return ensurePostgresDBExists(db, config.DBConfig.DBName, config.DBConfig.DBUser)
	case DIALECT_MYSQL:
		return ensureMySQLDBExists(db, config.DBConfig.DBName)
	default:
		return nil
	}
}

// ensurePostgresDBExists 确保 PostgreSQL 数据库存在，不存在则创建
// 参数：
//   - db: 数据库连接
//   - dbName: 数据库名称
//   - dbUser: 数据库用户
//
// 返回值：
//   - error: 创建过程中的错误
func ensurePostgresDBExists(db *gorm.DB, dbName, dbUser string) error {
	var exists bool
	query := "SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = ?)"
	if err := db.Raw(query, dbName).Scan(&exists).Error; err != nil {
		log.Printf("查询「%s」数据库是否存在时失败: %v", dbName, err)
		return fmt.Errorf("查询「%s」数据库是否存在时失败: %v", dbName, err)
	}

	if !exists {
		log.Printf("「%s」数据库不存在，正在创建...", dbName)
		global.SysLog.Infof("「%s」数据库不存在，正在创建...", dbName)

		createSQL := fmt.Sprintf("CREATE DATABASE %s ENCODING 'UTF8' OWNER %s", dbName, dbUser)
		if err := db.Exec(createSQL).Error; err != nil {
			return fmt.Errorf("创建「%s」数据库失败: %v", dbName, err)
		}
	}
	return nil
}

// ensureMySQLDBExists 确保 MySQL 数据库存在，不存在则创建
// 参数：
//   - db: 数据库连接
//   - dbName: 数据库名称
//
// 返回值：
//   - error: 创建过程中的错误
func ensureMySQLDBExists(db *gorm.DB, dbName string) error {
	var count int64
	query := "SELECT COUNT(*) FROM information_schema.schemata WHERE schema_name = ?"
	if err := db.Raw(query, dbName).Scan(&count).Error; err != nil {
		log.Printf("查询「%s」数据库是否存在时失败: %v", dbName, err)
		return fmt.Errorf("查询「%s」数据库是否存在时失败: %v", dbName, err)
	}

	if count == 0 {
		log.Printf("「%s」数据库不存在，正在创建...", dbName)
		global.SysLog.Infof("「%s」数据库不存在，正在创建...", dbName)
		createSQL := fmt.Sprintf("CREATE DATABASE `%s` CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci", dbName)
		if err := db.Exec(createSQL).Error; err != nil {
			return fmt.Errorf("创建「%s」数据库失败: %v", dbName, err)
		}
	}
	return nil
}
