package request

type Legion struct {
	GameId       int    `form:"gameId" json:"gameId" binding:"required"`
	ServerId     int    `form:"serverId" json:"serverId" binding:"required"`
	GangsId      int    `form:"gangsId" json:"gangsId" binding:"required"`
	GangsAdminId int    `form:"gangsAdminId" json:"gangsAdminId" binding:"required"`
	GangsName    string `form:"gangsName" json:"gangsName" binding:"required"`
	Level        int    `form:"level" json:"level" binding:"required"`
	Notice       string `form:"notice" json:"notice" binding:"required"`
	Ip           string `form:"ip" json:"ip" binding:"required"`
	CreateTime   int    `form:"createTime" json:"createTime" binding:"required"`
}

func (legion Legion) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"serverId.required":     "区服ID不能为空",
		"gameId.required":       "游戏ID不能为空",
		"gangsId.required":      "军团ID不能为空",
		"gangsAdminId.required": "军团长ID不能为空",
		"gangsName.required":    "军团名称不能为空",
		"ip.required":           "IP地址不能为空",
		"level.required":        "军团等级不能为空",
		"notice.required":       "军团宣言不能为空",
		"createTime.required":   "创建时间不能为空",
	}
}
