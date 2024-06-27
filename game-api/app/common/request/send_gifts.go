package request

type SendGifts struct {
	GameId       int     `form:"gameId" json:"gameId" binding:"required"`
	ServerId     int     `form:"serverId" json:"serverId" binding:"required"`
	Agent        string  `form:"agent" json:"agent" binding:"required"`
	RoleId       string  `form:"roleId" json:"roleId" binding:"required"`
	UserCode     string  `form:"userCode" json:"userCode" binding:"required"`
	GiftId       int     `form:"giftId" json:"giftId" binding:"required"`
	GiftType     int     `form:"giftType" json:"giftType" binding:"required"`
	TriggerTimes int     `form:"triggerTimes" json:"triggerTimes" binding:"required"`
	GiftMoney    float32 `form:"giftMoney" json:"giftMoney" binding:"required"`
	BuyCount     int     `form:"buyCount" json:"buyCount" binding:"required"`
}

func (sendGifts SendGifts) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"serverId.required":     "区服ID不能为空",
		"gameId.required":       "游戏ID不能为空",
		"agent.required":        "渠道标识不能为空",
		"roleId.required":       "角色ID不能为空",
		"userCode.required":     "玩家标识不能为空",
		"giftId.required":       "礼包ID不能为空",
		"giftType.required":     "礼包类型不能为空",
		"triggerTimes.required": "总触发次数不能为空",
		"giftMoney.required":    "总购买金额不能为空",
		"buyCount.required":     "总购买次数不能为空",
	}
}
