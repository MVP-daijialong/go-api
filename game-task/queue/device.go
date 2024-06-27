package queue

import (
	"context"
	"encoding/json"
	"errors"
	"my-gin/app/common/utils"
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
func StartDeviceConsumer() {
	global.App.Log.Info("Starting Device Consumer")
	numWorkers := 5 // 设置工作线程数量

	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {
			for {
				job, err := global.App.Redis.RPop(context.Background(), constant.DeviceUpload).Result()
				if err == redis.Nil {
					// Queue is empty, sleep for a while and retry
					time.Sleep(1 * time.Second)
					continue
				} else if err != nil {
					global.App.Log.Error("Worker " + strconv.Itoa(workerID) + " failed to pop from deviceQueue: " + err.Error())
					continue
				}

				// 处理 job
				if err := processDeviceJob(job); err != nil {
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

func processDeviceJob(job string) error {
	var device models.Device
	if err := json.Unmarshal([]byte(job), &device); err != nil {
		return err
	}

	// 获取 IP 信息
	ipInfo, err := utils.GetIPInfo(device.Ip)
	if err != nil {
		return err
	}

	device.Province = ipInfo.Region
	device.City = ipInfo.City
	device.LastInit = time.Now().Unix()

	// 创建 DeviceLogin 对象并赋值
	deviceLogin := models.DeviceLogin{
		GameID:   device.GameID,
		Agent:    device.Agent,
		Udid:     device.Udid,
		Type:     device.Type, // 如果 Type 是必填项，这里需要赋值
		Ip:       device.Ip,
		Ver:      device.Ver,
		Province: device.Province,
		City:     device.City,
	}

	maxRetries := 3
	var dbErr error

	for i := 0; i < maxRetries; i++ {
		// 开始事务
		dbErr = global.App.DB.Transaction(func(tx *gorm.DB) error {
			var existingDevice models.Device

			// 检查 device 是否存在
			err := tx.Where("gameId = ? AND agent = ? AND udid = ?", device.GameID, device.Agent, device.Udid).First(&existingDevice).Error
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}

			// 如果记录不存在，则插入 device 表
			if errors.Is(err, gorm.ErrRecordNotFound) {
				err = tx.Create(&device).Error
				if err != nil {
					return err
				}
			}

			// 插入 DeviceLogin 表
			err = tx.Create(&deviceLogin).Error
			if err != nil {
				return err
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
