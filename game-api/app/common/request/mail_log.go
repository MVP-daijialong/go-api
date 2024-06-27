package request

type MailLog struct {
	UserCode string `form:"userCode" json:"userCode" binding:"required"`
	GameId   int    `form:"gameId" json:"gameId" binding:"required"`
	Agent    string `form:"agent" json:"agent" binding:"required"`
	ServerId string `form:"serverId" json:"serverId" binding:"required"`
	RoleId   string `form:"roleId" json:"roleId" binding:"required"`
	RoleName string `form:"roleName" json:"roleName" binding:"required"`
	Title    string `form:"title" json:"title" binding:"required"`
	Content  string `form:"content" json:"content" binding:"required"`
	Item     string `form:"item" json:"item" binding:"required"`
	Status   int    `form:"status" json:"status" binding:"required"` // 1: 发送, 2: 读取, 3: 领取'
}

func (mailLog MailLog) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"userCode.required": "玩家标识不能为空",
		"serverId.required": "区服ID不能为空",
		"gameId.required":   "游戏ID不能为空",
		"agent.required":    "渠道标识不能为空",
		"roleId.required":   "角色ID不能为空",
		"roleName.required": "角色名称不能为空",
		"title.required":    "邮件标题不能为空",
		"content.required":  "邮件内容不能为空",
		"item.required":     "附件不能为空",
		"status.required":   "邮件状态不能为空",
	}
}
