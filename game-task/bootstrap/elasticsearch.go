package bootstrap

import (
	"crypto/tls"
	"my-gin/global"
	"net/http"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8"
	"go.uber.org/zap"
)

func InitializeElasticsearch() *elasticsearch.Client {
	cfg := elasticsearch.Config{
		Addresses: []string{
			global.App.Config.Elasticsearch.Host + ":" + strconv.Itoa(global.App.Config.Elasticsearch.Port),
		},
		Username: global.App.Config.Elasticsearch.Username,
		Password: global.App.Config.Elasticsearch.Password,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // 跳过证书验证
			},
		},
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		global.App.Log.Error("Elasticsearch connection failed, err:", zap.Any("err", err))
		return nil
	}

	// Ping Elasticsearch to verify connection
	res, err := client.Info()
	if err != nil {
		global.App.Log.Error("Elasticsearch ping failed, err:", zap.Any("err", err))
		return nil
	}
	defer res.Body.Close()

	if res.IsError() {
		global.App.Log.Error("Elasticsearch ping returned error", zap.String("status", res.Status()))
		return nil
	}

	global.App.Log.Info("Elasticsearch connected successfully")
	return client
}
