package request

type Chat struct {
	GameId   int    `form:"gameId" json:"gameId" binding:"required"`
	Agent    string `form:"agent" json:"agent" binding:"required"`
	ServerId string `form:"serverId" json:"serverId" binding:"required"`
	RoleId   string `form:"roleId" json:"roleId" binding:"required"`
	RoleName string `form:"roleName" json:"roleName" binding:"required"`
	Ip       string `form:"ip" json:"ip" binding:"required"`
	Content  string `form:"content" json:"content" binding:"required"`
}

func (chat Chat) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"serverId.required": "区服ID不能为空",
		"gameId.required":   "游戏ID不能为空",
		"agent.required":    "渠道标识不能为空",
		"roleId.required":   "角色ID不能为空",
		"roleName.required": "角色名称不能为空",
		"ip.required":       "IP地址不能为空",
		"content.required":  "邮件内容不能为空",
	}
}