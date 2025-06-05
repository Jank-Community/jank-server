package service

import (
	"crypto/md5"
	"fmt"
	"jank.com/jank_blog/internal/model/plugin"
	"jank.com/jank_blog/internal/storage"
	"time"
)

// PluginService 插件服务
type PluginService struct {
	store *storage.MemoryStore
}

// NewPluginService 创建插件服务
func NewPluginService() *PluginService {
	return &PluginService{
		store: storage.NewMemoryStore(),
	}
}

// RegisterPlugin 注册插件
func (ps *PluginService) RegisterPlugin(plugin *model.Plugin) (string, error) {
	// 验证插件信息
	if err := plugin.ValidatePlugin(); err != nil {
		return "", err
	}

	// 生成插件ID
	if plugin.ID == "" {
		plugin.ID = ps.generatePluginID(plugin.Name, plugin.Author)
	}

	// 设置默认值
	plugin.IsActive = true
	plugin.DownloadCount = 0
	plugin.Rating = 0.0

	// 存储插件
	if err := ps.store.CreatePlugin(plugin); err != nil {
		return "", err
	}

	return plugin.ID, nil
}

// GetPlugin 获取插件
func (ps *PluginService) GetPlugin(id string) (*model.Plugin, error) {
	return ps.store.GetPlugin(id)
}

// UpdatePlugin 更新插件
func (ps *PluginService) UpdatePlugin(plugin *model.Plugin) error {
	// 验证插件信息
	if err := plugin.ValidatePlugin(); err != nil {
		return err
	}

	return ps.store.UpdatePlugin(plugin)
}

// DeletePlugin 删除插件
func (ps *PluginService) DeletePlugin(id string) error {
	return ps.store.DeletePlugin(id)
}

// ListPlugins 列出插件
func (ps *PluginService) ListPlugins(category string, tags []string, searchQuery string, page, pageSize int, sortBy string, ascending bool) ([]*model.Plugin, int, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	return ps.store.ListPlugins(category, tags, searchQuery, page, pageSize, sortBy, ascending)
}

// SearchPlugins 搜索插件
func (ps *PluginService) SearchPlugins(query string, limit int) ([]*model.Plugin, error) {
	if limit <= 0 {
		limit = 10
	}

	return ps.store.SearchPlugins(query, limit)
}

// DownloadPlugin 下载插件
func (ps *PluginService) DownloadPlugin(id string) (string, error) {
	plugin, err := ps.store.GetPlugin(id)
	if err != nil {
		return "", err
	}

	// 增加下载计数
	plugin.DownloadCount++
	ps.store.UpdatePlugin(plugin)

	return plugin.DownloadURL, nil
}

// generatePluginID 生成插件ID
func (ps *PluginService) generatePluginID(name, author string) string {
	data := fmt.Sprintf("%s-%s-%d", name, author, time.Now().UnixNano())
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}
