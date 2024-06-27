package queue

import (
	"fmt"
	"math/rand"
	"my-gin/app/services"
	"my-gin/global"
	"sync"
	"time"

	"go.uber.org/zap"
)

// 随机字符串生成器
func randomString(r *rand.Rand, n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := make([]byte, n)
	for i := range bytes {
		bytes[i] = letters[r.Intn(len(letters))]
	}
	return string(bytes)
}

// 随机生成文档
func generateRandomDocument(r *rand.Rand) map[string]interface{} {
	return map[string]interface{}{
		"orderId":         randomString(r, 16),
		"billNo":          fmt.Sprintf("%s_%d", randomString(r, 12), r.Intn(100000)),
		"userCode":        randomString(r, 12),
		"userName":        randomString(r, 8),
		"gameId":          r.Intn(2000) + 1000,
		"agent":           fmt.Sprintf("%d_%d", r.Intn(10000), r.Intn(1000)),
		"udid":            randomString(r, 16),
		"type":            fmt.Sprintf("%d", r.Intn(5)+1),
		"roleId":          randomString(r, 12),
		"roleName":        randomString(r, 10),
		"serverId":        fmt.Sprintf("%d", r.Intn(10000)),
		"serverName":      fmt.Sprintf("%d服", r.Intn(10)),
		"channelId":       randomString(r, 4),
		"level":           r.Intn(100),
		"amount":          fmt.Sprintf("%d", r.Intn(10000)),
		"goodsCode":       randomString(r, 3),
		"giftId":          fmt.Sprintf("%d", r.Intn(100)),
		"orderType":       r.Intn(3) + 1,
		"orderStatus":     fmt.Sprintf("%d", r.Intn(3)+1),
		"gameOrderStatus": fmt.Sprintf("%d", r.Intn(3)+1),
		"createTime":      time.Now().Unix(),
		"payType":         r.Intn(5) + 1,
	}
}

// 文档插入工作函数
func addDocumentWorker(indexName string, jobs <-chan map[string]interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for doc := range jobs {
		if err := services.AddDocument(indexName, doc); err != nil {
			global.App.Log.Error("Failed to add document", zap.Error(err))
		}
	}
}

// 测试函数
func Test() {
	indexName := "orders"
	numDocuments := 10
	numWorkers := 4

	// 创建一个带缓冲的通道，用于存放要插入的文档
	jobs := make(chan map[string]interface{}, numDocuments)

	// 创建一个 WaitGroup，以便等待所有工作协程完成
	var wg sync.WaitGroup

	// 启动工作协程
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go addDocumentWorker(indexName, jobs, &wg)
	}

	// 创建一个新的随机数生成器
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 生成文档并发送到通道中
	for i := 0; i < numDocuments; i++ {
		doc := generateRandomDocument(r)
		jobs <- doc
	}

	// 关闭通道，表示没有更多的工作
	close(jobs)

	// 等待所有工作协程完成
	wg.Wait()
}
