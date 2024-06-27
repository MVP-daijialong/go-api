package models

type Role struct {
	ID
	UserCode   string `gorm:"column:userCode;size:100;not null"`
	Agent      string `gorm:"size:60;not null"`
	Udid       string `gorm:"size:255;not null"`
	GameID     int    `gorm:"column:gameId;not null"`
	RoleID     string `gorm:"column:roleId;size:70;not null"`
	RoleName   string `gorm:"column:roleName;size:100;not null"`
	ServerID   string `gorm:"column:serverId;size:20;not null"`
	ServerName string `gorm:"column:serverName;size:100;not null"`
	Type       string `gorm:"not null"`
	Level      int    `gorm:"not null"`
	Currency   int    `gorm:"not null"`
	Vip        int    `gorm:"not null"`
	Balance    int    `gorm:"not null"`
	Power      int    `gorm:"not null"`
	LastScene  string `gorm:"column:lastScene;size:255"`
	Online     int
	Timestamps
}
