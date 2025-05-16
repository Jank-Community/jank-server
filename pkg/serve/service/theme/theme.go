// Package service 提供主题相关业务逻辑处理
// 创建者：Done-0
// 创建时间：2025-05-14
package service

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"

	"jank.com/jank_blog/configs"
	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/serve/controller/theme/dto"
	"jank.com/jank_blog/pkg/serve/mapper"
	"jank.com/jank_blog/pkg/vo/theme"
)

// GetActivatedTheme 获取当前激活的主题
// 参数：
//   - c: Echo上下文
//
// 返回值：
//   - *theme.ThemeVO: 主题视图对象
//   - error: 操作过程中的错误
func GetActivatedTheme(c echo.Context) (*theme.ThemeVO, error) {
	themeModel, err := mapper.GetActivatedTheme(c)
	if err != nil {
		utils.BizLogger(c).Errorf("获取已激活主题失败: %v", err)
		return nil, fmt.Errorf("获取已激活主题失败: %w", err)
	}

	if themeModel == nil {
		utils.BizLogger(c).Errorf("未找到已激活主题")
		return nil, fmt.Errorf("未找到已激活主题")
	}

	vo, err := utils.MapModelToVO(themeModel, &theme.ThemeVO{})
	if err != nil {
		utils.BizLogger(c).Errorf("映射主题视图对象失败: %v", err)
		return nil, fmt.Errorf("映射主题视图对象失败: %w", err)
	}

	return vo.(*theme.ThemeVO), nil
}

// GetThemeByID 根据ID获取主题
// 参数：
//   - c: Echo上下文
//   - req: 获取主题请求
//
// 返回值：
//   - *theme.ThemeVO: 主题视图对象
//   - error: 操作过程中的错误
func GetThemeByID(c echo.Context, req *dto.GetThemeRequest) (*theme.ThemeVO, error) {
	themeModel, err := mapper.GetThemeByID(c, req.ID)
	if err != nil {
		utils.BizLogger(c).Errorf("根据ID获取主题失败: %v", err)
		return nil, fmt.Errorf("根据ID获取主题失败: %w", err)
	}

	if themeModel == nil {
		utils.BizLogger(c).Errorf("主题不存在: %d", req.ID)
		return nil, fmt.Errorf("主题不存在: %d", req.ID)
	}

	vo, err := utils.MapModelToVO(themeModel, &theme.ThemeVO{})
	if err != nil {
		utils.BizLogger(c).Errorf("映射主题视图对象失败: %v", err)
		return nil, fmt.Errorf("映射主题视图对象失败: %w", err)
	}

	return vo.(*theme.ThemeVO), nil
}

// ListAllThemes 获取所有主题
// 参数：
//   - c: Echo上下文
//
// 返回值：
//   - []*theme.ThemeVO: 主题视图对象列表
//   - error: 操作过程中的错误
func ListAllThemes(c echo.Context) ([]*theme.ThemeVO, error) {
	themeModels, err := mapper.ListAllThemes(c)
	if err != nil {
		utils.BizLogger(c).Errorf("获取主题列表失败: %v", err)
		return nil, fmt.Errorf("获取主题列表失败: %w", err)
	}

	// 扫描主题目录更新最新主题信息
	config, err := configs.LoadConfig()
	if err == nil {
		themeDir := config.AppConfig.Theme.ThemeDir
		configFiles := config.AppConfig.Theme.ThemeConfigFiles

		_, scanErr := mapper.ScanThemeDir(c, themeDir, configFiles)
		if scanErr != nil {
			utils.BizLogger(c).Warnf("扫描主题目录失败: %v", scanErr)
		} else {
			// 重新获取更新后的主题列表
			updatedThemes, listErr := mapper.ListAllThemes(c)
			if listErr != nil {
				utils.BizLogger(c).Warnf("获取更新后的主题列表失败: %v", listErr)
			} else {
				themeModels = updatedThemes
			}
		}
	}

	// 转换为视图对象
	var themeVOs []*theme.ThemeVO
	for _, model := range themeModels {
		vo, convertErr := utils.MapModelToVO(model, &theme.ThemeVO{})
		if convertErr != nil {
			utils.BizLogger(c).Warnf("转换主题VO失败: %v", convertErr)
			continue
		}
		themeVOs = append(themeVOs, vo.(*theme.ThemeVO))
	}

	return themeVOs, nil
}

// ActivateOneTheme 激活主题
// 参数：
//   - c: Echo上下文
//   - req: 激活主题请求
//
// 返回值：
//   - *theme.ThemeVO: 主题视图对象
//   - error: 操作过程中的错误
func ActivateOneTheme(c echo.Context, req *dto.ActivateThemeRequest) (*theme.ThemeVO, error) {
	var themeVO *theme.ThemeVO

	err := utils.RunDBTransaction(c, func(tx error) error {
		// 确认主题存在
		themeModel, err := mapper.GetThemeByID(c, req.ID)
		if err != nil {
			utils.BizLogger(c).Errorf("获取主题失败: %v", err)
			return fmt.Errorf("获取主题失败: %w", err)
		}

		if themeModel == nil {
			utils.BizLogger(c).Errorf("主题ID「%d」不存在", req.ID)
			return fmt.Errorf("主题ID「%d」不存在", req.ID)
		}

		// 激活主题
		if err := mapper.ActivateOneThemeByID(c, req.ID); err != nil {
			utils.BizLogger(c).Errorf("激活主题失败: %v", err)
			return fmt.Errorf("激活主题失败: %w", err)
		}

		// 获取激活后的主题并转换为VO
		vo, err := utils.MapModelToVO(themeModel, &theme.ThemeVO{})
		if err != nil {
			utils.BizLogger(c).Errorf("激活主题时映射VO失败: %v", err)
			return fmt.Errorf("激活主题时映射VO失败: %w", err)
		}

		themeVO = vo.(*theme.ThemeVO)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return themeVO, nil
}

// DeleteThemeByID 删除主题
// 参数：
//   - c: Echo上下文
//   - req: 删除主题请求
//
// 返回值：
//   - error: 操作过程中的错误
func DeleteThemeByID(c echo.Context, req *dto.DeleteThemeRequest) error {
	return utils.RunDBTransaction(c, func(tx error) error {
		// 确认主题存在
		themeModel, err := mapper.GetThemeByID(c, req.ID)
		if err != nil {
			utils.BizLogger(c).Errorf("获取主题失败: %v", err)
			return fmt.Errorf("获取主题失败: %w", err)
		}

		if themeModel == nil {
			utils.BizLogger(c).Errorf("主题ID「%d」不存在", req.ID)
			return fmt.Errorf("主题ID「%d」不存在", req.ID)
		}

		// 不能删除激活的主题
		if themeModel.Activated {
			utils.BizLogger(c).Errorf("无法删除当前激活的主题: %d", req.ID)
			return fmt.Errorf("无法删除当前激活的主题: %d", req.ID)
		}

		// 删除主题
		if err := mapper.DeleteOneThemeByID(c, req.ID); err != nil {
			utils.BizLogger(c).Errorf("删除主题失败: %v", err)
			return fmt.Errorf("删除主题失败: %w", err)
		}

		// 删除主题设置
		if req.DeleteSetting {
			if err := mapper.DeleteThemeSettings(c, req.ID); err != nil {
				utils.BizLogger(c).Errorf("删除主题设置失败: %v", err)
				return fmt.Errorf("删除主题设置失败: %w", err)
			}
		}

		// 删除主题文件
		if themeModel.ThemePath != "" {
			if err := os.RemoveAll(themeModel.ThemePath); err != nil {
				utils.BizLogger(c).Errorf("删除主题文件失败: %v", err)
				return fmt.Errorf("删除主题文件失败: %w", err)
			}
		}

		return nil
	})
}
