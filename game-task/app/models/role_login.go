package models

type RoleLogin struct {
	ID
	UserCode      string `gorm:"column:userCode;size:100;not null"`
	Agent         string `gorm:"size:60;not null"`
	Udid          string `gorm:"size:255;not null"`
	GameID        int    `gorm:"column:gameId;not null"`
	RoleID        string `gorm:"column:roleId;size:70;not null"`
	RoleName      string `gorm:"column:roleName;size:100;not null"`
	ServerID      string `gorm:"column:serverId;size:20;not null"`
	LoginServerID int    `gorm:"column:loginServerId;not null"`
	ServerName    string `gorm:"column:serverName;size:100;not null"`
	Type          string `gorm:"not null"`
	Time          int64
	RegTime       int64  `gorm:"column:regTime"`
	LogoutTime    int64  `gorm:"column:logoutTime"`
	RoleLevel     int    `gorm:"column:roleLevel;not null"`
	Ip            string `gorm:"column:ip;size:255"`
	Timestamps
}
