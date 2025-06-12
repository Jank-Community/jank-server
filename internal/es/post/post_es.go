package es

import (
	"fmt"

	"jank.com/jank_blog/internal/es/base"
	model "jank.com/jank_blog/internal/model/post"
	"jank.com/jank_blog/internal/utils"
)

// PostESDocument Post模型的ElasticSearch文档适配器
type PostESDocument struct {
	base.ESModelBase
	Title       string `json:"title"`
	ContentText string `json:"content_text"`
	Visibility  bool   `json:"visibility"`
}

// NewPostESDocument 从Post模型创建ES文档
func NewPostESDocument(p *model.Post) (*PostESDocument, error) {
	contentText := utils.HTMLToPlainText(p.ContentHTML)
	return &PostESDocument{
		ESModelBase: base.ESModelBase{
			ID: fmt.Sprintf("%d", p.ID),
		},
		Title:       p.Title,
		ContentText: contentText,
		Visibility:  p.Visibility,
	}, nil
}

// EsIndexName 实现ESDocument接口
func (d *PostESDocument) EsIndexName() string {
	return "posts_es_index"
}

// EsMapping 实现ESDocument接口
func (d *PostESDocument) EsMapping() string {
	return `
	{
		"mappings": {
			"properties": {
				"id": {
					"type": "keyword"
				},
				"title": {
					"type": "text",
					"analyzer": "ik_max_word",
					"search_analyzer": "ik_smart"
				},
				"content_text": {
					"type": "text",
					"analyzer": "ik_max_word",
					"search_analyzer": "ik_smart"
				},
				"visibility": {
					"type": "boolean"
				}
			}
		}
	}`
}

// EsID 实现ESDocument接口
func (d *PostESDocument) EsID() string {
	return d.ID
}
