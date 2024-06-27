package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"my-gin/global"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"go.uber.org/zap"
)

/**
	使用示例:
	indexName := "orders"
	mapping := `
	{
		"mappings": {
			"properties": {
				"orderId": { "type": "keyword" },
				"billNo": { "type": "keyword" },
				"userCode": { "type": "keyword" },
				"userName": { "type": "keyword" },
				"gameId": { "type": "integer" },
				"agent": { "type": "keyword" },
				"udid": { "type": "keyword" },
				"type": { "type": "keyword" },
				"roleId": { "type": "keyword" },
				"roleName": { "type": "text" },
				"serverId": { "type": "keyword" },
				"serverName": { "type": "text" },
				"channelId": { "type": "keyword" },
				"level": { "type": "integer" },
				"amount": { "type": "keyword" },
				"goodsCode": { "type": "keyword" },
				"giftId": { "type": "keyword" },
				"orderType": { "type": "integer" },
				"orderStatus": { "type": "keyword" },
				"gameOrderStatus": { "type": "keyword" },
				"createTime": { "type": "integer" },
				"payType": { "type": "integer" }
			}
		}
	}
	`

	创建索引
	if err := CreateIndex(indexName, mapping); err != nil {
		global.App.Log.Error("Failed to create index", zap.Error(err))
		return
	}

	新增文档的内容
	doc := map[string]interface{}{
		"orderId":         "202412544465445",
		"billNo":          "2024125444654_45456",
		"userCode":        "adsfadsfdasf",
		"userName":        "126487987",
		"gameId":          1501,
		"agent":           "150102_102",
		"udid":            "adsfasdfasdf4646d",
		"type":            "1",
		"roleId":          "200112445646",
		"roleName":        "我是的阿道夫222",
		"serverId":        "2001",
		"serverName":      "2服",
		"channelId":       "1235",
		"level":           66,
		"amount":          "100",
		"goodsCode":       "123",
		"giftId":          "55",
		"orderType":       1,
		"orderStatus":     "1",
		"gameOrderStatus": "1",
		"createTime":      1542156454,
		"payType":         1,
	}

	// 添加文档
	if err := AddDocument(indexName, doc); err != nil {
		global.App.Log.Error("Failed to add document", zap.Error(err))
		return
	}

	// 获取文档
	docID := "QC79JJABwXHdCjzNFOj5"
	if _, err := GetDocumentInfo(indexName, docID); err != nil {
		global.App.Log.Error("Failed to get document", zap.Error(err))
		return
	}

	// 获取文档分页
	page := 1
	pageSize := 20 // 每页 20 条记录
	docs, total, err := GetDocumentList(indexName, page, pageSize)
	if err != nil {
		global.App.Log.Error("Failed to get documents", zap.Error(err))
	} else {
		global.App.Log.Info("Documents retrieved successfully", zap.String("index", indexName), zap.Int("total", total), zap.Any("documents", docs))
	}

	// 删除文档
	if err := DeleteDocument(indexName, docID); err != nil {
		global.App.Log.Error("Failed to delete document", zap.Error(err))
		return
	}

	// 删除索引
	if err := DeleteIndex(indexName); err != nil {
		global.App.Log.Error("Failed to delete index", zap.Error(err))
		return
	}
**/

// 创建索引
func CreateIndex(indexName string, mapping string) error {
	res, err := global.App.ElasticSearch.Indices.Create(
		indexName,
		global.App.ElasticSearch.Indices.Create.WithBody(bytes.NewReader([]byte(mapping))),
	)
	if err != nil {
		return fmt.Errorf("failed to create index: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("failed to create index: %s", res.Status())
	}

	global.App.Log.Info("Index created successfully", zap.String("index", indexName))
	return nil
}

// 添加文档
func AddDocument(indexName string, doc interface{}) error {
	// 将文档转换为 JSON
	docJSON, err := json.Marshal(doc)
	if err != nil {
		return fmt.Errorf("failed to marshal document: %w", err)
	}

	// 新增文档
	res, err := global.App.ElasticSearch.Index(
		indexName,
		bytes.NewReader(docJSON),
	)
	if err != nil {
		return fmt.Errorf("failed to add document: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("failed to add document: %s", res.Status())
	}

	global.App.Log.Info("Document added successfully", zap.String("index", indexName))
	return nil
}

// 获取文档
func GetDocumentInfo(indexName, docID string) (map[string]interface{}, error) {
	req := esapi.GetRequest{
		Index:      indexName,
		DocumentID: docID,
	}
	res, err := req.Do(context.Background(), global.App.ElasticSearch)
	if err != nil {
		return nil, fmt.Errorf("failed to get document: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("failed to get document: %s", res.Status())
	}

	var doc map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&doc); err != nil {
		return nil, fmt.Errorf("failed to parse document: %w", err)
	}

	global.App.Log.Info("Document retrieved successfully", zap.String("index", indexName), zap.Any("document", doc))
	return doc, nil
}

// 获取文档分页
func GetDocumentList(indexName string, page int, pageSize int) ([]map[string]interface{}, int, error) {
	// 默认每页 15 条记录
	if pageSize <= 0 {
		pageSize = 15
	}

	from := (page - 1) * pageSize

	query := map[string]interface{}{
		"from": from,
		"size": pageSize,
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, 0, fmt.Errorf("error encoding query: %w", err)
	}

	res, err := global.App.ElasticSearch.Search(
		global.App.ElasticSearch.Search.WithContext(context.Background()),
		global.App.ElasticSearch.Search.WithIndex(indexName),
		global.App.ElasticSearch.Search.WithBody(&buf),
		global.App.ElasticSearch.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get documents: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, 0, fmt.Errorf("failed to get documents: %s", res.Status())
	}

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, 0, fmt.Errorf("failed to parse response body: %w", err)
	}

	hits := result["hits"].(map[string]interface{})
	totalHits := int(hits["total"].(map[string]interface{})["value"].(float64))
	hitsArray := hits["hits"].([]interface{})

	documents := make([]map[string]interface{}, len(hitsArray))
	for i, hit := range hitsArray {
		documents[i] = hit.(map[string]interface{})["_source"].(map[string]interface{})
	}

	global.App.Log.Info("Documents retrieved successfully", zap.String("index", indexName), zap.Int("page", page), zap.Int("page_size", pageSize))
	return documents, totalHits, nil
}

// 删除文档
func DeleteDocument(indexName, docID string) error {
	req := esapi.DeleteRequest{
		Index:      indexName,
		DocumentID: docID,
	}
	res, err := req.Do(context.Background(), global.App.ElasticSearch)
	if err != nil {
		return fmt.Errorf("failed to delete document: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("failed to delete document: %s", res.Status())
	}

	global.App.Log.Info("Document deleted successfully", zap.String("index", indexName), zap.String("document_id", docID))
	return nil
}

// 删除索引
func DeleteIndex(indexName string) error {
	res, err := global.App.ElasticSearch.Indices.Delete([]string{indexName})
	if err != nil {
		return fmt.Errorf("failed to delete index: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("failed to delete index: %s", res.Status())
	}

	global.App.Log.Info("Index deleted successfully", zap.String("index", indexName))
	return nil
}
