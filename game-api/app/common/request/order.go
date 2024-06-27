package request

type Order struct {
	OrderId         string `form:"orderId" json:"orderId" binding:"required"`
	BillNo          string `form:"billNo" json:"billNo" binding:"required"`
	UserCode        string `form:"userCode" json:"userCode" binding:"required"`
	RoleName        string `form:"roleName" json:"roleName" binding:"required"`
	ChannelId       string `form:"channelId" json:"channelId" binding:"required"`
	Amount          string `form:"amount" json:"amount" binding:"required"`
	GoodsCode       string `form:"goodsCode" json:"goodsCode" binding:"required"`
	GiftId          string `form:"giftId" json:"giftId" binding:"required"`
	RoleId          string `form:"RoleId" json:"RoleId" binding:"required"`
	GameId          int    `form:"gameId" json:"gameId" binding:"required"`
	Agent           string `form:"agent" json:"agent" binding:"required"`
	ServerId        string `form:"serverId" json:"serverId" binding:"required"`
	ServerName      string `form:"serverName" json:"serverName" binding:"required"`
	Udid            string `form:"udid" json:"udid" binding:"required"`
	Type            string `form:"type" json:"type" binding:"required"`
	Ip              string `form:"ip" json:"ip"`
	Level           int    `form:"level" json:"level" binding:"required"`
	OrderType       int    `form:"orderType" json:"orderType" binding:"required"`
	OrderStatus     string `form:"orderStatus" json:"orderStatus" binding:"required"`
	GameOrderStatus string `form:"gameOrderStatus" json:"gameOrderStatus" binding:"required"`
	CreateTime      int    `form:"createTime" json:"createTime" binding:"required"`
	PayType         int    `form:"payType" json:"payType" binding:"required"`
}

func (order Order) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"orderId.required":         "公开订单号不能为空",
		"billNo.required":          "游戏订单号不能为空",
		"userCode.required":        "玩家标识不能为空",
		"userName.required":        "玩家名称不能为空",
		"gameId.required":          "游戏ID不能为空",
		"agent.required":           "渠道标识不能为空",
		"udid.required":            "设备标识不能为空",
		"type.required":            "设备类型不能为空",
		"roleId.required":          "角色ID不能为空",
		"roleName.required":        "角色名称不能为空",
		"serverId.required":        "区服ID不能为空",
		"serverName.required":      "区服名称不能为空",
		"channelId.required":       "平台渠道不能为空",
		"level.required":           "等级不能为空",
		"amount.required":          "充值金额不能为空",
		"goodsCode.required":       "商品ID不能为空",
		"giftId.required":          "礼包ID不能为空",
		"orderType.required":       "订单类型不能为空",
		"orderStatus.required":     "订单状态不能为空",
		"gameOrderStatus.required": "游戏充值状态不能为空",
		"createTime.required":      "下单时间不能为空",
		"payType.required":         "充值渠道不能为空",
	}
}
