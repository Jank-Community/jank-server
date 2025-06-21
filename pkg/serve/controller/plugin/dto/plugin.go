// Package dto 提供插件相关的数据传输对象定义
// 创建者：ixuemy
// 创建时间：2025-06-21
package dto

// RegisterPluginRequest         注册插件的请求结构体
// @Param	name				body	string	true	"插件名称"
// @Param	version			body	string	true	"插件版本"
// @Param	description		body	string	false	"插件描述"
// @Param	author			body	string	false	"插件作者"
// @Param	category			body	string	false	"插件分类"
// @Param	download_url	body	string	true	"插件下载地址"
// @Param	git_url			body	string	false	"插件仓库地址"
// @Param	address			body	string	false	"gRPC 服务监听地址"
type RegisterPluginRequest struct {
	Name        string `json:"name" xml:"name" form:"name" query:"name" validate:"required"`
	Version     string `json:"version" xml:"version" form:"version" query:"version" validate:"required"`
	Description string `json:"description" xml:"description" form:"description" query:"description"`
	Author      string `json:"author" xml:"author" form:"author" query:"author"`
	Category    string `json:"category" xml:"category" form:"category" query:"category"`
	DownloadURL string `json:"download_url" xml:"download_url" form:"download_url" query:"download_url" validate:"required"`
	GitURL      string `json:"git_url" xml:"git_url" form:"git_url" query:"git_url"`
	Address     string `json:"address" xml:"address" form:"address" query:"address"`
}

// DeletePluginRequest    插件删除请求
// @Param id path int true "插件 ID"
type DeletePluginRequest struct {
	ID string `json:"id" xml:"id" form:"id" query:"id" validate:"required"`
}

// DownloadPluginRequest    插件下载请求
// @Param id path int true "插件 ID"
type DownloadPluginRequest struct {
	ID string `json:"id" xml:"id" form:"id" query:"id" validate:"required"`
}

// UpdatePluginRequest       更新插件请求参数结构体
// @Param	id		path	string	true	"插件 ID"
// @Param	name				body	string	true	"插件名称"
// @Param	version			body	string	true	"插件版本"
// @Param	description		body	string	false	"插件描述"
// @Param	author			body	string	false	"插件作者"
// @Param	category			body	string	false	"插件分类"
// @Param	download_url	body	string	true	"插件下载地址"
// @Param	git_url			body	string	false	"插件仓库地址"
// @Param	address			body	string	false	"gRPC 服务监听地址"
type UpdatePluginRequest struct {
	ID          string `json:"id" xml:"id" form:"id" query:"id" validate:"required"`
	Name        string `json:"name" xml:"name" form:"name" query:"name" validate:"required"`
	Version     string `json:"version" xml:"version" form:"version" query:"version" validate:"required"`
	Description string `json:"description" xml:"description" form:"description" query:"description"`
	Author      string `json:"author" xml:"author" form:"author" query:"author"`
	Category    string `json:"category" xml:"category" form:"category" query:"category"`
	DownloadURL string `json:"download_url" xml:"download_url" form:"download_url" query:"download_url" validate:"required"`
	GitURL      string `json:"git_url" xml:"git_url" form:"git_url" query:"git_url"`
	Address     string `json:"address" xml:"address" form:"address" query:"address"`
}
