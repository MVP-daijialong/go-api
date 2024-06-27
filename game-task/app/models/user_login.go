package models

type UserLogin struct {
	ID
	UserCode string `gorm:"column:userCode;index:idx_user_unique,userCode"`
	GameID   int    `gorm:"column:gameId;index:idx_user_unique,unique"`
	Time     int64
	RegTime  int64  `gorm:"column:regTime"`
	Agent    string `gorm:"index:idx_device_unique,unique"`
	Udid     string `gorm:"index:idx_device_unique,unique"`
	Type     string `gorm:"not null"`
	Ip       string `gorm:"column:ip;size:255"`
	Province string `gorm:"column:province;size:255"`
	City     string `gorm:"column:city;size:255"`
	PayUser  int    `gorm:"column:payUser"`
	Timestamps
}
