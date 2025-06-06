// Package post 提供文章相关的视图对象定义
// 创建者：Done-0
// 创建时间：2025-05-10
package post

// PostsVO    获取帖子的响应结构
// @Description	获取帖子时返回的响应数据
// @Property			id			    	body	string	true	"帖子唯一标识"
// @Property			title			    body	string	true	"帖子标题"
// @Property			image			    body	string	true	"帖子封面图片 URL"
// @Property			visibility		    body	bool	true	"帖子可见性状态"
// @Property			content_html		body	string	true	"帖子 HTML 格式内容"
// @Property			category_id	    	body	string	true	"帖子所属分类 ID"
// @Property			gmt_create	    	body	string	true	"创建时间（格式化时间）"
// @Property			gmt_modified	    body	string	true	"更新时间（格式化时间）"
type PostsVO struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Image      string `json:"image"`
	Visibility bool   `json:"visibility"`
	// ContentMarkdown string `json:"content_markdown"`
	ContentHTML string `json:"content_html"`
	CategoryID  string `json:"category_id"`
	GmtCreate   string `json:"gmt_create"`
	GmtModified string `json:"gmt_modified"`
}
