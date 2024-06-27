package models

type Device struct {
	ID
	GameID   int    `gorm:"column:gameId;index:idx_device_unique,unique"`
	Agent    string `gorm:"index:idx_device_unique,unique"`
	Udid     string `gorm:"index:idx_device_unique,unique"`
	Type     string `gorm:"not null"`
	Ip       string `gorm:"column:ip;size:255"`
	Ver      string `gorm:"column:ver;size:255"`
	Province string `gorm:"column:province;size:255"`
	City     string `gorm:"column:city;size:255"`
	LastInit int64  `gorm:"column:lastInit;size:255"`
	Timestamps
}
