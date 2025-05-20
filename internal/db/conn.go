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

// 数据库类型常量
const (
	DIALECT_POSTGRES = "postgres" // PostgreSQL 数据库
	DIALECT_SQLITE   = "sqlite"   // SQLite 数据库
	DIALECT_MYSQL    = "mysql"    // MySQL 数据库
)

// New 初始化数据库连接
// 参数：
//   - config: 应用配置
func New(config *configs.Config) {
	var err error

	switch config.DBConfig.DBDialect {
	case DIALECT_SQLITE:
		global.DB, err = connectToDB(config, config.DBConfig.DBName)
		if err != nil {
			global.SysLog.Fatalf("连接 SQLite 数据库失败: %v", err)
		}
	case DIALECT_POSTGRES, DIALECT_MYSQL:
		systemDB, err := connectToSystemDB(config)
		if err != nil {
			global.SysLog.Fatalf("连接系统数据库失败: %v", err)
		}

		if err := ensureDBExists(systemDB, config); err != nil {
			global.SysLog.Fatalf("数据库不存在且创建失败: %v", err)
		}

		sqlDB, _ := systemDB.DB()
		sqlDB.Close()

		global.DB, err = connectToDB(config, config.DBConfig.DBName)
		if err != nil {
			global.SysLog.Fatalf("连接数据库失败: %v", err)
		}
	default:
		global.SysLog.Fatalf("不支持的数据库类型: %s", config.DBConfig.DBDialect)
	}

	log.Printf("「%s」数据库连接成功...", config.DBConfig.DBName)
	global.SysLog.Infof("「%s」数据库连接成功...", config.DBConfig.DBName)

	err = autoMigrate()
	if err != nil {
		global.SysLog.Fatalf("数据库自动迁移失败: %v", err)
	}
}

// connectToSystemDB 连接到系统数据库
// 参数：
//   - config: 应用配置
//
// 返回值：
//   - *gorm.DB: 数据库连接
//   - error: 连接过程中的错误
func connectToSystemDB(config *configs.Config) (*gorm.DB, error) {
	switch config.DBConfig.DBDialect {
	case DIALECT_POSTGRES:
		return connectToDB(config, "postgres")
	case DIALECT_MYSQL:
		return connectToDB(config, "information_scshema")
	default:
		return nil, fmt.Errorf("不支持的数据库类型: %s", config.DBConfig.DBDialect)
	}
}

// connectToDB 连接到指定数据库
// 参数：
//   - config: 应用配置
//   - dbName: 数据库名称
//
// 返回值：
//   - *gorm.DB: 数据库连接
//   - error: 连接过程中的错误
func connectToDB(config *configs.Config, dbName string) (*gorm.DB, error) {
	dialector, err := getDialector(config, dbName)
	if err != nil {
		return nil, fmt.Errorf("获取数据库驱动器失败: %v", err)
	}

	return gorm.Open(dialector, &gorm.Config{})
}
