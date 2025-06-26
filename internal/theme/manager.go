// Package theme 提供主题的检索和切换功能
// 主题管理器
// 创建者：Pey121
// 创建时间：2025-05-30
package theme

import (
	"errors"
	"fmt"
	"sync"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm" // 导入 GORM 数据库包

	t_model "jank.com/jank_blog/internal/model/theme" // 导入主题模型
)

// ThemeManager 结构体，负责主题的激活、切换和缓存管理
// 字段：
//   - db: 数据库连接实例
//   - activeThemeID: 当前激活主题 ID
//   - activeThemeName: 当前激活主题名称
//   - themeCache: 主题缓存
//   - configCache: 配置缓存
//   - templateEngine: 模板引擎实例
type ThemeManager struct {
	db              *gorm.DB       // 数据库连接实例
	activeThemeID   string         // 当前激活主题 ID
	activeThemeName string         // 当前激活主题名称
	themeCache      sync.Map       // 主题缓存
	configCache     sync.Map       // 配置缓存
	templateEngine  TemplateEngine // 模板引擎实例
}

// ThemeError 主题错误信息结构体，用于统一错误处理
// 字段：
//   - Code: 错误码
//   - Message: 错误信息
//   - Err: 原始错误
type ThemeError struct {
	Code    int    // 错误码
	Message string // 错误信息
	Err     error  // 原始错误
}

// NewThemeManager 创建主题管理器实例
// 参数：
//   - db: 数据库连接实例
//
// 返回值：
//   - *ThemeManager: 主题管理器实例
func NewThemeManager(db *gorm.DB) *ThemeManager {
	tm := &ThemeManager{
		db: db,
	}

	// 从数据库查询一条 is_active = true 的主题记录
	// 如果没有找到，则使用默认主题
	var activeRecord t_model.Theme
	if err := db.Where("is_active = ?", true).First(&activeRecord).Error; err == nil {
		tm.activeThemeID = activeRecord.ThemeID
		tm.activeThemeName = activeRecord.Name
	} else {
		tm.activeThemeID = "defaultID"
		tm.activeThemeName = "default"
	}
	// 初始化时扫描所有主题到缓存
	tm.ScanThemes(nil)
	return tm
}

// ScanThemes 扫描并加载所有主题信息到缓存
// 参数：
//   - c: Echo 上下文
//
// 返回值：
//   - error: 扫描过程中的错误
func (tm *ThemeManager) ScanThemes(c echo.Context) error {
	themes_slice := []t_model.Theme{}
	if err := tm.db.Find(&themes_slice).Error; err != nil {
		return err
	}

	// 将主题信息存储到缓存中
	for _, theme := range themes_slice {
		tm.themeCache.Store(theme.ThemeID, theme)
	}
	return nil
}

// ActivateTheme 激活指定的主题，并刷新模板缓存
// 参数：
//   - c: Echo 上下文
//   - themeID: 要激活的主题 ID
//
// 返回值：
//   - error: 激活过程中的错误
func (tm *ThemeManager) ActivateTheme(c echo.Context, themeID string) error {
	// 1. 验证主题是否存在
	value, ok := tm.themeCache.Load(themeID)
	if !ok {
		return errors.New("主题不存在")
	}
	theme := value.(t_model.Theme)

	// 2. 更新数据库中的激活状态
	if err := tm.updateActiveTheme(themeID); err != nil {
		return err
	}

	// 3. 清理缓存
	tm.clearCache()

	// 4. 清理模板缓存并预加载主模板
	tm.templateEngine.ClearTemplateCache()
	if _, err := tm.templateEngine.LoadThemeTemplate(theme.Name, "index.html"); err != nil {
		return fmt.Errorf("预加载主模板失败: %w", err)
	}

	// 5. 设置当前主题
	tm.activeThemeID = themeID
	tm.activeThemeName = theme.Name
	return nil
}

// GetActiveTheme 获取当前激活的主题 ID
// 返回值：
//   - string: 当前激活的主题 ID
//   - error: 获取过程中的错误
func (tm *ThemeManager) GetActiveTheme() (string, error) {
	return tm.activeThemeID, nil
}

// GetThemeMap 获取主题 ID 到名称、名称到 ID 的映射，并缓存所有主题
// 返回值：
//   - idToName: 主题 ID 到名称的映射
//   - nameToID: 主题名称到 ID 的映射
//   - error: 获取过程中的错误
func (tm *ThemeManager) GetThemeMap() (map[string]string, map[string]string, error) {
	themes := []t_model.Theme{}
	if err := tm.db.Find(&themes).Error; err != nil {
		return nil, nil, err
	}

	idToName := make(map[string]string)
	nameToID := make(map[string]string)

	for _, theme := range themes {
		idToName[theme.ThemeID] = theme.Name
		nameToID[theme.Name] = theme.ThemeID
		tm.themeCache.Store(theme.ThemeID, theme)
	}

	return idToName, nameToID, nil
}

// isThemeExists 检查指定主题是否存在
// 参数：
//   - themeID: 要检查的主题 ID
//
// 返回值：
//   - bool: 主题是否存在
func (tm *ThemeManager) isThemeExists(themeID string) bool {
	if _, ok := tm.themeCache.Load(themeID); ok {
		return true
	}
	return false
}

// updateActiveTheme 更新数据库中的激活主题状态
// 参数：
//   - themeID: 要激活的主题 ID
//
// 返回值：
//   - error: 更新过程中的错误
func (tm *ThemeManager) updateActiveTheme(themeID string) error {
	tx := tm.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Model(&t_model.Theme{}).
		Where("is_active = ?", true).
		Update("is_active", false).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&t_model.Theme{}).
		Where("theme_id = ?", themeID).
		Update("is_active", true).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// clearCache 清除主题和配置缓存
func (tm *ThemeManager) clearCache() {
	tm.themeCache = sync.Map{}
	tm.configCache = sync.Map{}
}

// Error 实现 error 接口，返回格式化的错误信息
// 返回值：
//   - string: 格式化的错误信息
func (e *ThemeError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// NewThemeError 创建新的主题错误实例
// 参数：
//   - code: 错误码
//   - message: 错误信息
//   - err: 原始错误
//
// 返回值：
//   - *ThemeError: 主题错误实例
func NewThemeError(code int, message string, err error) *ThemeError {
	return &ThemeError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
