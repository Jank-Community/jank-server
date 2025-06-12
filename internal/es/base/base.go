package base

import (
	"context"

	"github.com/olivere/elastic/v7"
)

type ESModelBase struct {
	ID string `json:"id"`
}

// ESDocument 定义ElasticSearch文档接口
type ESDocument interface {
	EsIndexName() string
	EsMapping() string
	EsID() string
}

// ESIndexer 定义索引操作接口
type ESIndexer interface {
	IndexDocument(ctx context.Context, doc ESDocument) error
	DeleteDocument(ctx context.Context, id string) error
	SearchDocuments(ctx context.Context, query string, params map[string]interface{}) ([]map[string]interface{}, *elastic.SearchResult, error)
}

//
