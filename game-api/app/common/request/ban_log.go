package request

type BanLog struct {
	GameId   int    `form:"gameId" json:"gameId" binding:"required"`
	ServerId int    `form:"serverId" json:"serverId" binding:"required"`
	Account  string `form:"account" json:"account" binding:"required"`
	DevInfo  string `form:"devInfo" json:"devInfo" binding:"required"`
	Ip       string `form:"ip" json:"ip" binding:"required"`
}

func (banLog BanLog) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"gameId.required":   "游戏ID不能为空",
		"serverId.required": "区服ID不能为空",
		"account.required":  "账号不能为空",
		"devInfo.required":  "设备标识不能为空",
		"ip.required":       "IP地址不能为空",
	}
}
