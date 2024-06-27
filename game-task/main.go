package main

import (
	"my-gin/bootstrap"
	"my-gin/global"
	"my-gin/queue"
)

func main() {
	// 初始化配置
	bootstrap.InitializeConfig()

	// 初始化日志
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("log init success!")

	// 初始化数据库
	global.App.DB = bootstrap.InitializeDB()
	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()

	// 初始化验证器
	bootstrap.InitializeValidator()

	// 初始化Redis
	global.App.Redis = bootstrap.InitializeRedis()

	// 初始化ElasticSearch
	global.App.ElasticSearch = bootstrap.InitializeElasticsearch()

	// 启动队列消费者
	go queue.StartRoleConsumer()
	go queue.StartDeviceConsumer()
	go queue.StartUserConsumer()
	go queue.StartOrderConsumer()
	go queue.Test()

	// 启动服务器
	bootstrap.RunServer()
}
