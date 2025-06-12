package es

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/olivere/elastic/v7"
	"jank.com/jank_blog/configs"
	"jank.com/jank_blog/internal/es/base"
	"jank.com/jank_blog/internal/global"
)

// SearchQuery 定义更通用的搜索查询条件
type SearchQuery struct {
	// 查询条件
	Queries []elastic.Query // 多种查询条件组合
	Query   string          // 全文搜索关键词(兼容旧版)
	Fields  []string        // 搜索字段(默认包含title,content_text)

	// 过滤条件
	Filters      map[string]interface{} // 精确匹配过滤
	RangeFilters map[string]struct {    // 范围过滤
		Gte interface{}
		Lte interface{}
	}

	// 排序选项
	Sorts    []elastic.Sorter // 多种排序条件
	SortBy   string           // 排序字段(兼容旧版)
	SortDesc bool             // 是否降序(兼容旧版)

	// 分页
	From int
	Size int

	// 聚合
	Aggregations map[string]elastic.Aggregation
}

func New(config *configs.Config) {
	host := config.EsConfig.EsHost
	port := config.EsConfig.EsPort
	user := config.EsConfig.EsUser
	password := config.EsConfig.EsPassword

	client, err := elastic.NewClient(
		elastic.SetURL("http://"+host+":"+port),
		elastic.SetSniff(false),
		elastic.SetBasicAuth(user, password),
	)
	if err != nil {
		global.SysLog.Fatalf("ElasticSearch连接失败: %v", err)
	}
	global.EsClient = client
}

// EnsureIndexExists 确保索引存在
func EnsureIndexExists(doc base.ESDocument) error {
	exists, err := global.EsClient.IndexExists(doc.EsIndexName()).Do(context.Background())
	if err != nil {
		return err
	}
	if !exists {
		_, err := global.EsClient.CreateIndex(doc.EsIndexName()).
			BodyJson(strings.NewReader(doc.EsMapping())).
			Do(context.Background())
		if err != nil {
			return err
		}
	}
	return nil
}

// IndexDocument 索引文档
func IndexDocument(ctx context.Context, doc base.ESDocument) error {
	_, err := global.EsClient.Index().
		Index(doc.EsIndexName()).
		Id(doc.EsID()).
		BodyJson(doc).
		Do(ctx)
	return err
}

// DeleteDocument 删除文档
func DeleteDocument(ctx context.Context, id string, indexName string) error {
	_, err := global.EsClient.Delete().
		Index(indexName).
		Id(id).
		Do(ctx)
	return err
}

// SearchDocuments 用的搜索文档方法
func SearchDocuments(ctx context.Context, indexName string, query *SearchQuery) ([]map[string]interface{}, *elastic.SearchResult, error) {
	searchService := global.EsClient.Search().Index(indexName)

	// 构建查询条件
	if len(query.Queries) > 0 {
		boolQuery := elastic.NewBoolQuery()
		for _, q := range query.Queries {
			boolQuery.Must(q)
		}
		searchService.Query(boolQuery)
	} else if query.Query != "" {
		fields := query.Fields
		if len(fields) == 0 {
			fields = []string{"title", "content_text"} // TODO: 默认字段配置化后续修改
		}
		searchService.Query(elastic.NewMultiMatchQuery(query.Query, fields...))
	}

	// 添加过滤条件
	for field, value := range query.Filters {
		searchService.Query(elastic.NewTermQuery(field, value))
	}
	for field, rangeFilter := range query.RangeFilters {
		searchService.Query(elastic.NewRangeQuery(field).Gte(rangeFilter.Gte).Lte(rangeFilter.Lte))
	}

	// 添加排序
	if len(query.Sorts) > 0 {
		for _, sorter := range query.Sorts {
			searchService.SortBy(sorter)
		}
	} else if query.SortBy != "" {
		searchService.Sort(query.SortBy, query.SortDesc)
	}

	// 添加聚合
	for name, agg := range query.Aggregations {
		searchService.Aggregation(name, agg)
	}

	// 执行搜索
	searchResult, err := searchService.
		From(query.From).
		Size(query.Size).
		Do(ctx)
	if err != nil {
		return nil, nil, err
	}

	// 转换结果
	var results []map[string]interface{}
	for _, hit := range searchResult.Hits.Hits {
		var result map[string]interface{}
		if err := json.Unmarshal(hit.Source, &result); err != nil {
			log.Printf("Error unmarshaling search result: %v", err)
			continue
		}
		results = append(results, result)
	}

	return results, searchResult, nil
}
