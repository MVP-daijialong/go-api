package request

type Role struct {
	UserCode   string `form:"userCode" json:"userCode" binding:"required"`
	RoleName   string `form:"roleName" json:"roleName" binding:"required"`
	RoleId     string `form:"RoleId" json:"RoleId" binding:"required"`
	GameId     int    `form:"gameId" json:"gameId" binding:"required"`
	Agent      string `form:"agent" json:"agent" binding:"required"`
	ServerId   string `form:"serverId" json:"serverId" binding:"required"`
	ServerName string `form:"serverName" json:"serverName" binding:"required"`
	Udid       string `form:"udid" json:"udid" binding:"required"`
	Type       string `form:"type" json:"type" binding:"required"`
	Currency   int    `form:"currency" json:"currency" binding:"required"`
	Level      int    `form:"level" json:"level" binding:"required"`
	Vip        int    `form:"vip" json:"vip" binding:"required"`
	Balance    int    `form:"balance" json:"balance" binding:"required"`
	Power      int    `form:"power" json:"power" binding:"required"`
	LastScene  string `form:"lastScene" json:"lastScene" binding:"required"`
	Tags       string `form:"tags" json:"tags" binding:"required"`
	Ip         string `form:"ip" json:"ip"`
}

func (role Role) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"userCode.required":   "玩家标识不能为空",
		"userName.required":   "玩家名称不能为空",
		"gameId.required":     "游戏ID不能为空",
		"agent.required":      "渠道标识不能为空",
		"udid.required":       "设备标识不能为空",
		"type.required":       "设备类型不能为空",
		"roleId.required":     "角色ID不能为空",
		"roleName.required":   "角色名称不能为空",
		"serverId.required":   "区服ID不能为空",
		"serverName.required": "区服名称不能为空",
		"currency.required":   "钻石元宝余额不能为空",
		"level.required":      "等级不能为空",
		"vip.required":        "vip等级不能为空",
		"balance.required":    "游戏币余额不能为空",
		"power.required":      "战力不能为空",
		"lastScene.required":  "最后登录场景不能为空",
		"tags.required":       "登录标识不能为空",
	}
}
