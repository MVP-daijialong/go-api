package models

type Order struct {
	ID
	OrderId         string `gorm:"column:orderId;size:100;not null"`
	BillNo          string `gorm:"column:billNo;size:60;not null"`
	UserCode        string `gorm:"column:userCode;not null"`
	UserName        string `gorm:"column:userName;size:70;not null"`
	Agent           string `gorm:"column:agent;size:100;not null"`
	ChannelId       string `gorm:"column:channelId;size:100;not null"`
	GameId          int    `gorm:"column:gameId;size:100;not null"`
	Amount          string `gorm:"column:amount;not null"`
	GoodsCode       string `gorm:"column:goodsCode;size:100;not null"`
	GiftId          string `gorm:"column:giftId;not null"`
	RoleId          string `gorm:"column:roleId;size:100;not null"`
	RoleName        string `gorm:"column:roleName;size:100;not null"`
	ServerID        string `gorm:"column:serverId;size:20;not null"`
	ServerName      string `gorm:"column:serverName;size:100;not null"`
	Ip              string `gorm:"not null"`
	Level           int    `gorm:"not null"`
	OrderType       int    `gorm:"column:orderType;not null"`
	Udid            string `gorm:"size:255;not null"`
	OrderStatus     string `gorm:"column:orderStatus;not null"`
	GameOrderStatus string `gorm:"column:gameOrderStatus;not null"`
	CreateTime      int    `gorm:"column:createTime;not null"`
	Type            string `gorm:"not null"`
	PayType         int    `gorm:"column:payType"`
	Timestamps
}
