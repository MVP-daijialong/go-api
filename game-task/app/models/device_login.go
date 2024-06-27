package models

type DeviceLogin struct {
	ID
	GameID   int    `gorm:"column:gameId"`
	Agent    string `gorm:"column:agent"`
	Udid     string `gorm:"column:udid"`
	Type     string `gorm:"not null"`
	Ip       string `gorm:"column:ip;size:255"`
	Ver      string `gorm:"column:ver;size:255"`
	Province string `gorm:"column:province;size:255"`
	City     string `gorm:"column:city;size:255"`
	Timestamps
}
