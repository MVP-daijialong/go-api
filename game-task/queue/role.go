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

// StartRoleConsumer 启动 Role 队列消费者
func StartRoleConsumer() {
	global.App.Log.Info("Starting Role Consumer")
	numWorkers := 5 // 设置工作线程数量

	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {
			for {
				job, err := global.App.Redis.RPop(context.Background(), constant.RoleUpload).Result()
				if err == redis.Nil {
					// Queue is empty, sleep for a while and retry
					time.Sleep(1 * time.Second)
					continue
				} else if err != nil {
					global.App.Log.Error("Worker " + strconv.Itoa(workerID) + " failed to pop from roleQueue: " + err.Error())
					continue
				}

				// 处理 job
				if err := processRoleJob(job); err != nil {
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

func processRoleJob(job string) error {
	// 包含所有字段的辅助结构体
	type roleData struct {
		models.Role
		Tags string `json:"tags"`
		Ip   string `json:"ip"`
	}

	var role roleData
	if err := json.Unmarshal([]byte(job), &role); err != nil {
		return err
	}

	// 连接数据库的重试逻辑
	maxRetries := 3
	var dbErr error

	for i := 0; i < maxRetries; i++ {
		// 开始事务
		dbErr = global.App.DB.Transaction(func(tx *gorm.DB) error {
			if role.Tags == "login" {
				// 登录操作，插入 roles 表和 role_logins 表
				var existingRole models.Role
				var timestamp int64
				err := tx.Where("gameId = ? AND serverId = ? AND roleId = ?", role.GameID, role.ServerID, role.RoleID).First(&existingRole).Error
				if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
					return err
				}

				if errors.Is(err, gorm.ErrRecordNotFound) {
					// 如果角色不存在，则插入到 roles 表
					// 设置 online 为 1
					role.Role.Online = 1
					if err := tx.Create(&role.Role).Error; err != nil {
						return err
					}
					timestamp = time.Now().Unix()
				} else {
					// 如果角色存在，更新 online 为 1
					if err := tx.Model(&existingRole).Update("online", 1).Error; err != nil {
						return err
					}
					timestamp = existingRole.CreatedAt.Unix()
				}

				// 判断是否为滚服玩家
				var olderRole models.Role
				var loginServerId int
				err = tx.Where("gameId = ? AND userCode = ? AND created_at < ?", role.GameID, role.UserCode, existingRole.CreatedAt).First(&olderRole).Error
				if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
					return err
				} else {
					loginServerId = 1
				}

				// 插入 RoleLogin 表
				roleLogin := models.RoleLogin{
					UserCode:      role.UserCode,
					Agent:         role.Agent,
					Udid:          role.Udid,
					GameID:        role.GameID,
					RoleID:        role.RoleID,
					RoleName:      role.RoleName,
					ServerID:      role.ServerID,
					LoginServerID: loginServerId, // assuming login server ID is same as server ID
					ServerName:    role.ServerName,
					Type:          role.Type,
					RoleLevel:     role.Level,
					Ip:            role.Ip, // assign as needed
					Time:          time.Now().Unix(),
					LogoutTime:    0,
					RegTime:       timestamp,
				}

				if err := tx.Create(&roleLogin).Error; err != nil {
					return err
				}

			} else if role.Tags == "logout" {
				// 登出操作，更新最后一条记录的 logoutTime、updated_at 和 roleName
				var lastRoleLogin models.RoleLogin
				if err := tx.Where("gameId = ? AND serverId = ? AND roleId = ?", role.GameID, role.ServerID, role.RoleID).
					Order("id desc").Limit(1).First(&lastRoleLogin).Error; err != nil {
					if !errors.Is(err, gorm.ErrRecordNotFound) {
						return err
					}
					// 如果没有找到最后一条记录，可能是首次登出，不处理错误，直接返回
				} else {
					lastRoleLogin.LogoutTime = time.Now().Unix()
					lastRoleLogin.UpdatedAt = time.Now()
					lastRoleLogin.RoleName = role.RoleName

					// 更新 RoleLogin 表
					if err := tx.Save(&lastRoleLogin).Error; err != nil {
						return err
					}
				}

				// 更新 roles 表中的 online 为 0
				var existingRole models.Role
				if err := tx.Where("gameId = ? AND roleId = ?", role.GameID, role.RoleID).First(&existingRole).Error; err != nil {
					if !errors.Is(err, gorm.ErrRecordNotFound) {
						return err
					}
					// 如果角色不存在，不处理错误，直接返回
				} else {
					if err := tx.Model(&existingRole).Update("online", 0).Error; err != nil {
						return err
					}
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
