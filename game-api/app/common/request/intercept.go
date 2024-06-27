package request

type Intercept struct {
	GameId         int    `form:"gameId" json:"gameId" binding:"required"`
	Agent          string `form:"agent" json:"agent" binding:"required"`
	UserCode       string `form:"userCode" json:"userCode" binding:"required"`
	Udid           string `form:"udid" json:"udid" binding:"required"`
	LastLoginAgent string `form:"last_login_agent" json:"last_login_agent" binding:"required"`
	LastLoginTime  int    `form:"last_login_time" json:"last_login_time" binding:"required"`
	InterceptTime  int    `form:"intercept_time" json:"intercept_time" binding:"required"`
	Ip             string `form:"ip" json:"ip" binding:"required"`
}

func (intercept Intercept) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"gameId.required":           "游戏ID不能为空",
		"agent.required":            "渠道标识不能为空",
		"userCode.required":         "玩家标识不能为空",
		"udid.required":             "设备标识不能为空",
		"ip.required":               "IP地址不能为空",
		"last_login_agent.required": "上次登录渠道不能为空",
		"last_login_time.required":  "上次登录时间不能为空",
		"intercept_time.required":   "拦截时间不能为空",
	}
}
