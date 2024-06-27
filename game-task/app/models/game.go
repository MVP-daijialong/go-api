package models

import (
	"time"
)

// Game 代表游戏子包表的结构体
type Game struct {
	ID
	UserID        uint      `gorm:"default:0;comment:'操作人id'"`
	MotherID      uint8     `gorm:"not null;comment:'游戏母包id'"`
	GameName      string    `gorm:"size:50;not null;comment:'游戏名称'"`
	GameAlias     string    `gorm:"size:50;not null;comment:'游戏别名缩写'"`
	Percents      string    `gorm:"size:50;default:null;comment:'订单金额折扣例0.1'"`
	GmURL         string    `gorm:"size:255;default:null;comment:'游戏GM地址'"`
	StandbyGmURL  string    `gorm:"size:255;default:null;comment:'游戏备用GM地址'"`
	StandbyStatus uint8     `gorm:"not null;default:0;comment:'是否启用游戏备用GM地址 1: 启用 0: 禁用'"`
	Status        uint8     `gorm:"default:1;comment:'运营状态 1: 运营 0: 停运'"`
	OnTime        time.Time `gorm:"default:null;comment:'上线时间'"`
	DownTime      time.Time `gorm:"default:null;comment:'下线时间'"`
	ReadLogType   uint8     `gorm:"default:1;comment:'读取日志类型 1本地库 2游戏库'"`
	IsBt          uint8     `gorm:"default:0;comment:'是否BT服 0否 1是'"`
	Timestamps
}
