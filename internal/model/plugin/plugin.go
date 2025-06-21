package model

import (
	"time"
)

// Plugin 插件模型
type Plugin struct {
	ID            string    `gorm:"primaryKey;unique;column:id" json:"id"`              // 唯一ID (UUID)
	Name          string    `gorm:"not null;column:name" json:"name"`                   // 插件名称（不能为空）
	Version       string    `gorm:"column:version" json:"version"`                      // 插件版本
	Description   string    `gorm:"column:description" json:"description"`              // 描述
	Author        string    `gorm:"column:author" json:"author"`                        // 作者
	Category      string    `gorm:"column:category" json:"category"`                    // 插件分类
	DownloadURL   string    `gorm:"column:download_url" json:"download_url"`            // 插件下载地址
	GitURL        string    `gorm:"column:git_url" json:"git_url"`                      // 插件仓库地址
	Address       string    `gorm:"column:address" json:"address"`                      // 插件服务监听地址
	DownloadCount int64     `gorm:"column:download_count" json:"download_count"`        // 插件下载数量
	Rating        float64   `gorm:"column:rating" json:"rating"`                        // 评分
	CreatedAt     time.Time `gorm:"autoCreateTime;column:created_at" json:"created_at"` // 创建时间（自动生成）
	UpdatedAt     time.Time `gorm:"autoUpdateTime;column:updated_at" json:"updated_at"` // 更新时间（自动更新）
	IsActive      bool      `gorm:"default:false;column:is_active" json:"is_active"`    // 是否可用（默认 false）
	Deleted       bool      `gorm:"column:deleted;default:false;" json:"deleted"`       // 删除（默认 false）
}

// TableName 指定表名
// 返回值：
//   - string: 表名
func (Plugin) TableName() string {
	return "plugins"
}
