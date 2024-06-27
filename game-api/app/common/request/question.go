package request

type Answer struct {
	GameId     int    `form:"gameId" json:"gameId" binding:"required"`
	ServerId   int    `form:"serverId" json:"serverId" binding:"required"`
	QuestionId int    `form:"questionId" json:"questionId" binding:"required"`
	Agent      string `form:"agent" json:"agent" binding:"required"`
	RoleId     string `form:"roleId" json:"roleId" binding:"required"`
	RoleName   string `form:"roleName" json:"roleName" binding:"required"`
	Type       int    `form:"type" json:"type" binding:"required"`
	Answer     string `form:"answer" json:"answer" binding:"required"`
}

func (answer Answer) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"serverId.required":   "区服ID不能为空",
		"gameId.required":     "游戏ID不能为空",
		"questionId.required": "问题ID不能为空",
		"agent.required":      "渠道标识不能为空",
		"roleId.required":     "角色ID不能为空",
		"roleName.required":   "角色名称不能为空",
		"type.required":       "问题类型不能为空",
		"answer.required":     "答案不能为空",
	}
}
