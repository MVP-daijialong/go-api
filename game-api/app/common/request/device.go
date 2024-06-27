package request

type Device struct {
	GameId int    `form:"gameId" json:"gameId" binding:"required"`
	Agent  string `form:"agent" json:"agent" binding:"required"`
	Udid   string `form:"udid" json:"udid" binding:"required"`
	Type   string `form:"type" json:"type" binding:"required"`
	Ip     string `form:"ip" json:"ip"`
	Ver    string `form:"ver" json:"ver" binding:"required"`
}

func (device Device) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"gameId.required": "游戏ID不能为空",
		"agent.required":  "渠道标识不能为空",
		"udid.required":   "设备标识不能为空",
		"type.required":   "设备类型不能为空",
		"ver.required":    "版本不能为空",
	}
}
