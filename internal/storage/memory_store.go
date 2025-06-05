package storage

import (
	"fmt"
	"jank.com/jank_blog/internal/model/plugin"
	"sort"
	"strings"
	"sync"
	"time"
)

//TODO 后续修改成数据库

// MemoryStore 内存存储实现
type MemoryStore struct {
	plugins map[string]*model.Plugin
	mutex   sync.RWMutex
}

// NewMemoryStore 创建新的内存存储
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		plugins: make(map[string]*model.Plugin),
	}
}

// GetPlugin 获取插件
func (ms *MemoryStore) GetPlugin(id string) (*model.Plugin, error) {
	ms.mutex.RLock()
	defer ms.mutex.RUnlock()

	plugin, exists := ms.plugins[id]
	if !exists {
		return nil, fmt.Errorf("plugin not found")
	}
	return plugin, nil
}

// CreatePlugin 创建插件
func (ms *MemoryStore) CreatePlugin(plugin *model.Plugin) error {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	if _, exists := ms.plugins[plugin.ID]; exists {
		return fmt.Errorf("plugin already exists")
	}

	plugin.CreatedAt = time.Now()
	plugin.UpdatedAt = time.Now()
	ms.plugins[plugin.ID] = plugin
	return nil
}

// UpdatePlugin 更新插件
func (ms *MemoryStore) UpdatePlugin(plugin *model.Plugin) error {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	if _, exists := ms.plugins[plugin.ID]; !exists {
		return fmt.Errorf("plugin not found")
	}

	plugin.UpdatedAt = time.Now()
	ms.plugins[plugin.ID] = plugin
	return nil
}

// DeletePlugin 删除插件
func (ms *MemoryStore) DeletePlugin(id string) error {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	if _, exists := ms.plugins[id]; !exists {
		return fmt.Errorf("plugin not found")
	}

	delete(ms.plugins, id)
	return nil
}

// ListPlugins 列出插件
func (ms *MemoryStore) ListPlugins(category string, tags []string, searchQuery string, page, pageSize int, sortBy string, ascending bool) ([]*model.Plugin, int, error) {
	ms.mutex.RLock()
	defer ms.mutex.RUnlock()

	var filteredPlugins []*model.Plugin

	// 过滤插件
	for _, plugin := range ms.plugins {
		if !plugin.IsActive {
			continue
		}

		// 分类过滤
		if category != "" && plugin.Category != category {
			continue
		}

		// 标签过滤
		if len(tags) > 0 {
			found := false
			for _, tag := range tags {
				for _, pluginTag := range plugin.Tags {
					if pluginTag == tag {
						found = true
						break
					}
				}
				if found {
					break
				}
			}
			if !found {
				continue
			}
		}

		// 搜索过滤
		if searchQuery != "" {
			query := strings.ToLower(searchQuery)
			if !strings.Contains(strings.ToLower(plugin.Name), query) &&
				!strings.Contains(strings.ToLower(plugin.Description), query) {
				continue
			}
		}

		filteredPlugins = append(filteredPlugins, plugin)
	}

	// 排序
	ms.sortPlugins(filteredPlugins, sortBy, ascending)

	total := len(filteredPlugins)

	// 分页
	start := (page - 1) * pageSize
	end := start + pageSize

	if start >= total {
		return []*model.Plugin{}, total, nil
	}

	if end > total {
		end = total
	}

	return filteredPlugins[start:end], total, nil
}

// sortPlugins 排序插件
func (ms *MemoryStore) sortPlugins(plugins []*model.Plugin, sortBy string, ascending bool) {
	sort.Slice(plugins, func(i, j int) bool {
		var less bool
		switch sortBy {
		case "name":
			less = plugins[i].Name < plugins[j].Name
		case "rating":
			less = plugins[i].Rating < plugins[j].Rating
		case "download_count":
			less = plugins[i].DownloadCount < plugins[j].DownloadCount
		case "created_at":
			less = plugins[i].CreatedAt.Before(plugins[j].CreatedAt)
		default:
			less = plugins[i].Name < plugins[j].Name
		}

		if ascending {
			return less
		}
		return !less
	})
}

// SearchPlugins 搜索插件
func (ms *MemoryStore) SearchPlugins(query string, limit int) ([]*model.Plugin, error) {
	ms.mutex.RLock()
	defer ms.mutex.RUnlock()

	var results []*model.Plugin
	query = strings.ToLower(query)

	for _, plugin := range ms.plugins {
		if !plugin.IsActive {
			continue
		}

		if strings.Contains(strings.ToLower(plugin.Name), query) ||
			strings.Contains(strings.ToLower(plugin.Description), query) ||
			strings.Contains(strings.ToLower(plugin.Author), query) {
			results = append(results, plugin)

			if len(results) >= limit {
				break
			}
		}
	}

	return results, nil
}
