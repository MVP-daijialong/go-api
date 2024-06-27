package models

type UserTest struct {
	ID
	UserCode   string `gorm:"not null;uniqueIndex:userCode_gameId"`
	UserName   string `gorm:"not null;index:u_g_a"`
	GameId     uint   `gorm:"not null;index:u_g_a"`
	GameName   string `gorm:"not null"`
	Agent      string `gorm:"not null"`
	IP         string `gorm:"not null"`
	UDID       string `gorm:"not null"`
	Type       int8   `gorm:"not null"`
	City       string `gorm:"default:null"`
	Province   string `gorm:"default:null"`
	PlayerType int8   `gorm:"default:0"`
	PayUser    uint8  `gorm:"default:0"`
	Timestamps
}
