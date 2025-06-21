package service

import (
	"archive/zip"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	model "jank.com/jank_blog/internal/model/plugin"
	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/serve/controller/plugin/dto"
)

// RegisterPlugin 注册插件
func RegisterPlugin(c echo.Context, req dto.RegisterPluginRequest) (string, error) {
	plugin := model.Plugin{
		ID:          generatePluginID(req.Name, req.Author),
		Name:        req.Name,
		Version:     req.Version,
		Author:      req.Author,
		Category:    req.Category,
		DownloadURL: req.DownloadURL,
		GitURL:      req.GitURL,
		Address:     req.Address,
		Description: req.Description,
	}

	// 设置默认值
	plugin.DownloadCount = 0
	plugin.Rating = 0.0

	// 存储插件
	db := utils.GetDBFromContext(c)
	if err := db.Create(&plugin).Error; err != nil {
		return "0", fmt.Errorf("创建插件失败: %w", err)
	}

	return plugin.ID, nil
}

// GetPlugin 获取插件
func GetPlugin(c echo.Context, id string) (*model.Plugin, error) {
	db := utils.GetDBFromContext(c)
	var plugin model.Plugin
	db.Where("id = ?", id).First(&plugin)
	if plugin.ID == "" {
		return nil, errors.New("未查到相应插件")
	}
	return &plugin, nil
}

// UpdatePlugin 更新插件
func UpdatePlugin(c echo.Context, req dto.UpdatePluginRequest) error {
	plugin := model.Plugin{
		Name:        req.Name,
		Version:     req.Version,
		Author:      req.Author,
		Category:    req.Category,
		DownloadURL: req.DownloadURL,
		GitURL:      req.GitURL,
		Address:     req.Address,
		Description: req.Description,
	}

	db := utils.GetDBFromContext(c)
	err := db.Model(&model.Plugin{}).Where("id = ?", req.ID).Updates(plugin).Error
	return err
}

// DeletePlugin 删除插件
func DeletePlugin(c echo.Context, id string) error {
	db := utils.GetDBFromContext(c)
	err := db.Model(&model.Plugin{}).Where("id = ?", id).Update("deleted", true).Error
	return err
}

// ListPlugins 列出插件
func ListPlugins(c echo.Context, page, pageSize int, category, searchQuery, sortBy string, ascending bool) (map[string]interface{}, error) {
	db := utils.GetDBFromContext(c)

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	var plugins []*model.Plugin
	query := db.Model(&model.Plugin{}).Where("deleted = ?", false).Where("is_active = ?", true)

	if category != "" {
		query = query.Where("category = ?", category)
	}

	if searchQuery != "" {
		searchPattern := "%" + searchQuery + "%"
		query = query.Where("name LIKE ? OR description LIKE ?", searchPattern, searchPattern)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	if sortBy == "" {
		sortBy = "created_at"
	}

	order := sortBy
	if !ascending {
		order += " DESC"
	}

	err := query.Order(order).
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&plugins).Error
	if err != nil {
		return nil, err
	}

	data := map[string]interface{}{
		"total": total,
		"items": plugins,
	}

	return data, nil
}

// SearchPlugins 搜索插件
func SearchPlugins(query string, limit int) ([]*model.Plugin, error) {
	if limit <= 0 {
		limit = 10
	}

	return SearchPlugins(query, limit)
}

// DownloadPlugin 下载/安装插件
func DownloadPlugin(c echo.Context, id string) (string, error) {
	plugin, err := GetPlugin(c, id)
	if err != nil {
		return "", err
	}

	// 设置下载路径
	currentPath, _ := os.Getwd()
	downloadDir := currentPath + "/plugins/" + plugin.Name
	if err := os.MkdirAll(downloadDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("创建目录失败: %v", err)
	}

	// 下载插件代码
	resp, err := http.Get(plugin.DownloadURL)
	if err != nil {
		return "", fmt.Errorf("下载失败: %v", err)
	}
	defer resp.Body.Close()

	archivePath := filepath.Join(downloadDir, plugin.Name+".zip")
	out, err := os.Create(archivePath)
	if err != nil {
		return "", fmt.Errorf("创建文件失败: %v", err)
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		return "", fmt.Errorf("保存插件失败: %v", err)
	}

	// 解压插件代码
	unzipPath := currentPath + "/plugins"
	if err := unzip(archivePath, unzipPath); err != nil {
		return "", fmt.Errorf("解压失败: %v", err)
	}

	// 启动插件服务
	cmd := exec.Command("bash", "-c", fmt.Sprintf("cd %s && chmod +x ./%s && nohup ./\"%s\" -id=\"%s\" -name=\"%s\" -version=\"%s\" -author=\"%s\" -desc=\"%s\" -port=\"%s\" > /dev/null 2>&1 &",
		downloadDir, plugin.Name, plugin.Name, plugin.ID, plugin.Name, plugin.Version, plugin.Author, plugin.Description, plugin.Address))

	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("启动插件失败: %v", err)
	}

	// 更新下载计数
	db := utils.GetDBFromContext(c)
	db.Model(&model.Plugin{}).Where("id = ?", id).Update("download_count", plugin.DownloadCount+1)

	return plugin.DownloadURL, nil
}

// generatePluginID 生成插件ID
func generatePluginID(name, author string) string {
	data := fmt.Sprintf("%s-%s-%d", name, author, time.Now().UnixNano())
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}

// UploadPluginPackage 上传插件包
func UploadPluginPackage(c echo.Context) (string, error) {
	file, err := c.FormFile("file")
	if err != nil {
		return "", fmt.Errorf("无法读取上传文件: %v", err)
	}

	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("无法打开上传文件: %v", err)
	}
	defer src.Close()

	// 获取存储路径
	currentPath, _ := os.Getwd()
	uploadDir := filepath.Join(currentPath, "uploaded_plugins")
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("创建上传目录失败: %v", err)
	}

	dstPath := filepath.Join(uploadDir, file.Filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		return "", fmt.Errorf("创建目标文件失败: %v", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("保存插件文件失败: %v", err)
	}

	downUrl := "http://" + c.Request().Host + "/api/v1/plugin/download?filename=" + file.Filename
	return downUrl, nil
}

// DownloadPluginPackage 下载上传的插件包
func DownloadPluginPackage(c echo.Context) error {
	filename := c.QueryParam("filename")
	if filename == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "缺少文件名"})
	}

	currentPath, _ := os.Getwd()
	pluginPath := filepath.Join(currentPath, "uploaded_plugins", filename)

	if _, err := os.Stat(pluginPath); os.IsNotExist(err) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "文件不存在"})
	}

	return c.Attachment(pluginPath, filename)
}

func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		fPath := filepath.Join(dest, f.Name)
		if !strings.HasPrefix(fPath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("非法文件路径: %s", fPath)
		}
		if f.FileInfo().IsDir() {
			os.MkdirAll(fPath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fPath), os.ModePerm); err != nil {
			return err
		}

		dstFile, err := os.OpenFile(fPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		fileInArchive, err := f.Open()
		if err != nil {
			return err
		}
		defer fileInArchive.Close()

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			return err
		}
		dstFile.Close()
	}

	return nil
}
