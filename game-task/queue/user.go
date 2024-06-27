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
func StartUserConsumer() {
	global.App.Log.Info("Starting User Consumer")
	numWorkers := 5 // 设置工作线程数量

	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {
			for {
				job, err := global.App.Redis.RPop(context.Background(), constant.UsersUpload).Result()
				if err == redis.Nil {
					// Queue is empty, sleep for a while and retry
					time.Sleep(1 * time.Second)
					continue
				} else if err != nil {
					global.App.Log.Error("Worker " + strconv.Itoa(workerID) + " failed to pop from userQueue: " + err.Error())
					continue
				}

				// 处理 job
				if err := processUserJob(job); err != nil {
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

func processUserJob(job string) error {
	var user models.User
	if err := json.Unmarshal([]byte(job), &user); err != nil {
		return err
	}

	// 获取 IP 信息
	ipInfo, err := utils.GetIPInfo(user.Ip)
	if err != nil {
		return err
	}

	user.Province = ipInfo.Region
	user.City = ipInfo.City

	maxRetries := 3
	var dbErr error

	for i := 0; i < maxRetries; i++ {
		// 开始事务
		dbErr = global.App.DB.Transaction(func(tx *gorm.DB) error {
			var existingUser models.User
			err := tx.Where("gameId = ? AND userCode = ?", user.GameID, user.UserCode).First(&existingUser).Error

			var userLogin models.UserLogin
			now := time.Now().Unix()

			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 如果记录不存在，则插入 users 表，并设置 userLogin 的时间为当前时间戳
				userLogin = models.UserLogin{
					GameID:   user.GameID,
					Agent:    user.Agent,
					Udid:     user.Udid,
					Type:     user.Type,
					Ip:       user.Ip,
					Province: user.Province,
					City:     user.City,
					Time:     now,
					RegTime:  now,
				}
				err = tx.Create(&user).Error
				if err != nil {
					return err
				}
			} else if err != nil {
				return err
			} else {
				// 如果记录存在，设置 userLogin 的时间为 users 表的 created_at 字段的时间戳
				timestamp := existingUser.CreatedAt.Unix()
				userLogin = models.UserLogin{
					GameID:   user.GameID,
					Agent:    user.Agent,
					Udid:     user.Udid,
					UserCode: user.UserCode,
					Type:     user.Type,
					Ip:       user.Ip,
					Province: user.Province,
					City:     user.City,
					Time:     now,
					RegTime:  timestamp,
					PayUser:  existingUser.PayUser,
				}
			}

			// 写入 user_logins 表
			err = tx.Create(&userLogin).Error
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
