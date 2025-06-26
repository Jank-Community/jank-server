// Package theme 提供检索、切换主题的功能
// 模板引擎实现：html/template + Sprig
// 创建者：Pey121
// 创建时间：2025-05-30
package theme

import (
	"errors"
	"html/template"
	"io"
	"path/filepath"
	"sync"

	"github.com/Masterminds/sprig/v3"

	"jank.com/jank_blog/internal/utils"
)

// TemplateEngine 模板引擎接口，定义模板渲染与缓存操作
// 方法：
//   - RenderTemplate: 渲染模板
//   - LoadThemeTemplate: 加载模板
//   - ClearTemplateCache: 清除模板缓存
type TemplateEngine interface {
	RenderTemplate(themeID, templateName string, data any, w io.Writer) error
	LoadThemeTemplate(themeID string, templateName string) (*template.Template, error)
	ClearTemplateCache()
}

// templateManager 模板引擎具体实现
// 字段：
//   - templateCache: 模板缓存（map[string]*template.Template）
type templateManager struct {
	templateCache sync.Map // 模板缓存，按路径缓存
}

// NewTemplateManager 创建模板引擎实例
// 返回值：
//   - TemplateEngine: 模板引擎实例
func NewTemplateManager() TemplateEngine {
	return &templateManager{}
}

// RenderTemplate 渲染模板（从缓存或磁盘加载）
// 参数：
//   - themeID: 主题 ID
//   - templateName: 模板文件名
//   - data: 模板数据
//   - w: 输出 Writer
//
// 返回值：
//   - error: 渲染过程中的错误
func (tm *templateManager) RenderTemplate(themeID, templateName string, data any, w io.Writer) error {
	tmpl, err := tm.LoadThemeTemplate(themeID, templateName)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, data)
}

// LoadThemeTemplate 尝试按顺序加载模板并缓存（当前主题 → 默认主题）
// 参数：
//   - themeID: 主题 ID
//   - templateName: 模板文件名
//
// 返回值：
//   - *template.Template: 模板对象
//   - error: 加载过程中的错误
func (tm *templateManager) LoadThemeTemplate(themeID string, templateName string) (*template.Template, error) {
	paths := []string{
		filepath.Join("themes", themeID, "templates", templateName),
		filepath.Join("themes", "default", "templates", templateName),
	}

	for _, path := range paths {
		if tmpl, ok := tm.templateCache.Load(path); ok {
			return tmpl.(*template.Template), nil
		}

		if utils.FileExists(path) {
			parsedTmpl, err := template.New(filepath.Base(path)).Funcs(sprig.FuncMap()).ParseFiles(path)
			if err != nil {
				return nil, err
			}

			tm.templateCache.Store(path, parsedTmpl)
			return parsedTmpl, nil
		}
	}

	return nil, errors.New("模板文件不存在: " + templateName)
}

// ClearTemplateCache 清除所有模板缓存
// 无参数
// 无返回值
func (tm *templateManager) ClearTemplateCache() {
	tm.templateCache.Range(func(key, _ interface{}) bool {
		tm.templateCache.Delete(key)
		return true
	})
}
