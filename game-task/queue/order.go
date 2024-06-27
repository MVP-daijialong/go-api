package queue

import (
	"context"
	"encoding/json"
	"errors"
	"my-gin/app/constant"
	"my-gin/app/models"
	"my-gin/global"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

// 启动队列消费者
func StartOrderConsumer() {
	global.App.Log.Info("Starting Order Consumer")
	numWorkers := 5 // 设置工作线程数量

	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {
			for {
				job, err := global.App.Redis.RPop(context.Background(), constant.OrderUpload).Result()
				if err == redis.Nil {
					// Queue is empty, sleep for a while and retry
					time.Sleep(1 * time.Second)
					continue
				} else if err != nil {
					global.App.Log.Error("Worker " + strconv.Itoa(workerID) + " failed to pop from orderQueue: " + err.Error())
					continue
				}

				// 处理 job
				if err := processOrderJob(job); err != nil {
					global.App.Log.Error("Failed to process job: " + err.Error())
				} else {
					global.App.Log.Info("Worker " + strconv.Itoa(workerID) + " processed job: " + job)
				}
			}
		}(i)
	}

	// 保持主程序运行
	select {}
}

func processOrderJob(job string) error {
	var order models.Order
	if err := json.Unmarshal([]byte(job), &order); err != nil {
		return err
	}

	maxRetries := 3
	var dbErr error

	for i := 0; i < maxRetries; i++ {
		// 开始事务
		dbErr = global.App.DB.Transaction(func(tx *gorm.DB) error {
			var existingOrder models.Order

			// 从 Redis 中获取游戏的百分比数据
			percentsGamesData, err := global.App.Redis.Get(context.Background(), constant.PercentGameKey).Result()
			var percentsGames []int

			if err == redis.Nil {
				// 如果 Redis 中没有数据，从数据库中获取
				var gameIDs []int
				err = global.App.DB.Model(&models.Game{}).
					Where("status = ? AND percents = ?", 1, "0.1").
					Pluck("id", &gameIDs).Error
				if err != nil {
					return err
				}
				percentsGames = gameIDs
				percentsGamesBytes, _ := json.Marshal(percentsGames)
				percentsGamesData = string(percentsGamesBytes)
				// 将数据保存到 Redis
				global.App.Redis.Set(context.Background(), constant.PercentGameKey, percentsGamesData, 0)
			} else if err != nil {
				return err
			} else {
				// 解析 Redis 中的数据
				err = json.Unmarshal([]byte(percentsGamesData), &percentsGames)
				if err != nil {
					return err
				}
			}

			// 如果订单中的 gameId 在 percentsGames 中，则将 amount 除以 100
			for _, gameID := range percentsGames {
				if order.GameId == gameID {
					amount, err := strconv.Atoi(order.Amount)
					if err != nil {
						return err
					}
					order.Amount = strconv.Itoa(amount / 100)
					break
				}
			}

			// 将 createTime 处理为长度不超过 10 的字符串
			createTimeStr := strconv.Itoa(order.CreateTime)
			if len(createTimeStr) > 10 {
				order.CreateTime, err = strconv.Atoi(createTimeStr[:10])
				if err != nil {
					return err
				}
			}

			// 检查订单是否存在
			err = tx.Where("orderId = ?", order.OrderId).First(&existingOrder).Error
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}

			// 如果记录不存在，则插入 orders 表
			if errors.Is(err, gorm.ErrRecordNotFound) {
				err = tx.Create(&order).Error
				if err != nil {
					return err
				}
			}

			return nil
		})

		if dbErr == nil {
			// 事务成功，返回 nil 表示成功
			return nil
		}

		// 如果错误是因为无效连接，重新尝试连接
		if strings.Contains(dbErr.Error(), "invalid connection") {
			sqlDB, pingErr := global.App.DB.DB()
			if pingErr != nil {
				return pingErr
			}
			_ = sqlDB.Ping()
			time.Sleep(time.Second) // 等待一秒再重试
			continue
		}

		// 对于其他错误，不重试，直接返回错误
		break
	}

	return dbErr
}
