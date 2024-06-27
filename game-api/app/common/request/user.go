package request

type User struct {
	UserCode string `form:"userCode" json:"userCode" binding:"required"`
	UserName string `form:"userName" json:"userName" binding:"required"`
	GameId   int    `form:"gameId" json:"gameId" binding:"required"`
	Agent    string `form:"agent" json:"agent" binding:"required"`
	Ip       string `form:"ip" json:"ip"`
	Udid     string `form:"udid" json:"udid" binding:"required"`
	Type     string `form:"type" json:"type" binding:"required"`
}

func (user User) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"userCode.required": "玩家标识不能为空",
		"userName.required": "玩家名称不能为空",
		"gameId.required":   "游戏ID不能为空",
		"agent.required":    "渠道标识不能为空",
		"udid.required":     "设备标识不能为空",
		"type.required":     "设备类型不能为空",
	}
}
