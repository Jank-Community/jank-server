package model

import (
	"fmt"
	"time"
)

// Plugin 插件模型
type Plugin struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Version       string    `json:"version"`
	Description   string    `json:"description"`
	Author        string    `json:"author"`
	Category      string    `json:"category"`
	Tags          []string  `json:"tags"`
	DownloadURL   string    `json:"download_url"`
	DownloadCount int64     `json:"download_count"`
	Rating        float64   `json:"rating"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	IsActive      bool      `json:"is_active"`
}

// ValidatePlugin 验证插件信息
func (p *Plugin) ValidatePlugin() error {
	if p.Name == "" {
		return fmt.Errorf("plugin name is required")
	}
	if p.Version == "" {
		return fmt.Errorf("plugin version is required")
	}
	if p.Author == "" {
		return fmt.Errorf("plugin author is required")
	}
	return nil
}
