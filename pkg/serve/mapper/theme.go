// Package mapper 提供数据访问层功能
// 创建者：Done-0
// 创建时间：2025-05-14
package mapper

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	model "jank.com/jank_blog/internal/model/theme"
	"jank.com/jank_blog/internal/utils"
)

// GetActivatedTheme 获取当前激活的主题
// 参数：
//   - c: Echo上下文
//
// 返回值：
//   - *model.Theme: 激活的主题对象
//   - error: 操作过程中的错误
func GetActivatedTheme(c echo.Context) (*model.Theme, error) {
	var themeModel model.Theme
	db := utils.GetDBFromContext(c)
	if err := db.Model(&model.Theme{}).Where("activated = ? AND deleted = ?", true, false).First(&themeModel).Error; err != nil {
		return nil, fmt.Errorf("获取激活的主题失败: %w", err)
	}
	return &themeModel, nil
}

// GetThemeByID 根据ID获取主题
// 参数：
//   - c: Echo上下文
//   - themeID: 主题ID
//
// 返回值：
//   - *model.Theme: 主题对象
//   - error: 操作过程中的错误
func GetThemeByID(c echo.Context, themeID int64) (*model.Theme, error) {
	var theme model.Theme
	db := utils.GetDBFromContext(c)
	if err := db.Model(&model.Theme{}).Where("id = ? AND deleted = ?", themeID, false).First(&theme).Error; err != nil {
		return nil, fmt.Errorf("获取主题失败: %w", err)
	}
	return &theme, nil
}

// ListAllThemes 获取所有主题
// 参数：
//   - c: Echo上下文
//
// 返回值：
//   - []*model.Theme: 主题列表
//   - error: 操作过程中的错误
func ListAllThemes(c echo.Context) ([]*model.Theme, error) {
	var themes []*model.Theme
	db := utils.GetDBFromContext(c)
	if err := db.Model(&model.Theme{}).Where("deleted = ?", false).Find(&themes).Error; err != nil {
		return nil, fmt.Errorf("获取主题列表失败: %w", err)
	}
	return themes, nil
}

// ScanThemeDir 扫描主题目录，提取主题信息
// 参数：
//   - c: Echo上下文
//   - themeDir: 主题目录路径
//
// 返回值：
//   - []*model.Theme: 主题对象列表
//   - error: 操作过程中的错误
func ScanThemeDir(c echo.Context, themeDir string, configFiles []string) ([]*model.Theme, error) {
	if err := os.MkdirAll(themeDir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("创建主题目录失败: %w", err)
	}

	// 读取目录中的所有文件和文件夹
	folders, err := os.ReadDir(themeDir)
	if err != nil {
		return nil, fmt.Errorf("读取主题目录失败: %w", err)
	}

	db := utils.GetDBFromContext(c)
	// 获取当前激活的主题，错误不影响继续执行
	activatedTheme, err := GetActivatedTheme(c)
	if err != nil {
		utils.BizLogger(c).Warnf("获取激活主题失败: %v", err)
	}

	var themes []*model.Theme
	// 遍历每个文件夹检查是否为有效主题
	for _, folder := range folders {
		if !folder.IsDir() {
			continue
		}

		folderName := folder.Name()
		themePath := filepath.Join(themeDir, folderName)

		// 检查是否存在任一有效配置文件
		configFileExists := false
		for _, configFile := range configFiles {
			_, statErr := os.Stat(filepath.Join(themePath, configFile))
			if statErr == nil {
				configFileExists = true
				break
			}
		}

		if !configFileExists {
			continue
		}

		// 检查数据库中是否已存在该主题
		var theme model.Theme
		isNew := db.Where("folder_name = ? AND deleted = ?", folderName, false).First(&theme).Error != nil

		if isNew {
			// 创建新主题记录
			theme = model.Theme{
				Name:       cases.Title(language.English).String(folderName),
				FolderName: folderName,
				ThemePath:  themePath,
			}
			if err := db.Create(&theme).Error; err != nil {
				utils.BizLogger(c).Errorf("保存主题失败: %v", err)
				continue
			}
		} else {
			// 更新已存在的主题记录
			theme.Name = cases.Title(language.English).String(folderName)
			theme.ThemePath = themePath
			// 根据当前激活主题设置激活状态
			theme.Activated = activatedTheme != nil && theme.ID == activatedTheme.ID
			if err := db.Save(&theme).Error; err != nil {
				utils.BizLogger(c).Errorf("更新主题失败: %v", err)
				continue
			}
		}

		themes = append(themes, &theme)
	}

	return themes, nil
}

// ActivateOneThemeByID 激活主题
// 参数：
//   - c: Echo上下文
//   - themeID: 主题ID
//
// 返回值：
//   - error: 操作过程中的错误
func ActivateOneThemeByID(c echo.Context, themeID int64) error {
	db := utils.GetDBFromContext(c)

	// 取消所有主题的激活状态，保持只有一个激活的主题
	if err := db.Model(&model.Theme{}).Where("activated = ? AND deleted = ?", true, false).Update("activated", false).Error; err != nil {
		return fmt.Errorf("取消主题激活状态失败: %w", err)
	}

	// 激活指定的主题
	if err := db.Model(&model.Theme{}).Where("id = ? AND deleted = ?", themeID, false).Update("activated", true).Error; err != nil {
		return fmt.Errorf("激活主题失败: %w", err)
	}

	return nil
}

// DeleteOneThemeByID 删除主题
// 参数：
//   - c: Echo上下文
//   - themeID: 主题ID
//
// 返回值：
//   - error: 操作过程中的错误
func DeleteOneThemeByID(c echo.Context, themeID int64) error {
	db := utils.GetDBFromContext(c)
	if err := db.Model(&model.Theme{}).Where("id = ? AND deleted = ?", themeID, false).Update("deleted", true).Error; err != nil {
		return fmt.Errorf("删除主题失败: %w", err)
	}
	return nil
}

// DeleteThemeSettings 删除主题设置
// 参数：
//   - c: Echo上下文
//   - themeID: 主题ID
//
// 返回值：
//   - error: 操作过程中的错误
func DeleteThemeSettings(c echo.Context, themeID int64) error {
	db := utils.GetDBFromContext(c)
	if err := db.Model(&model.Theme{}).Where("theme_id = ?", themeID).Delete(&model.ThemeSetting{}).Error; err != nil {
		return fmt.Errorf("删除主题设置失败: %w", err)
	}
	return nil
}
