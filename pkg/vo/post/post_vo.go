// Package post 提供文章相关的视图对象定义
// 创建者：Done-0
// 创建时间：2025-05-10
package post

// PostsVO    获取帖子的响应结构
// @Description	获取帖子时返回的响应数据
// @Property			ID			    	body	string	true	"帖子唯一标识"
// @Property			Title			    body	string	true	"帖子标题"
// @Property			Image			    body	string	true	"帖子封面图片 URL"
// @Property			Visibility		    body	bool	true	"帖子可见性状态"
// @Property			ContentHTML		    body	string	true	"帖子 HTML 格式内容"
// @Property			CategoryID	    	body	string	true	"帖子所属分类 ID"
// @Property			GmtCreate	    	body	string	true	"创建时间（格式化时间）"
// @Property			GmtModified	        body	string	true	"更新时间（格式化时间）"
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
