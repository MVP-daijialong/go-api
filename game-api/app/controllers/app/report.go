package app

import (
	"my-gin/app/common/request"
	"my-gin/app/common/response"
	"my-gin/app/constant"
	"my-gin/app/controllers/common"
	"my-gin/global"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

// 账号上报
func User(c *gin.Context) {
	var form request.User
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	form.Ip = common.GetClientIP(c.Request)

	// 将请求参数序列化为 JSON 字符串
	s, err := jsoniter.MarshalToString(form)
	if err != nil {
		global.App.Log.Error("Failed to marshal request params to JSON:" + err.Error())
		response.Fail(c, 500, "Failed to marshal request params to JSON")
		return
	}

	// 异步处理将数据推入 Redis 队列
	go func() {
		global.App.Log.Info("User请求参数@" + s)

		// 将数据推入 Redis 队列
		if _, err := global.App.Redis.LPush(c, constant.UsersUpload, s).Result(); err != nil {
			global.App.Log.Error("Failed to push data to Redis:" + err.Error())
			return
		}

	}()

	// 返回成功响应
	response.Success(c, nil)
}

// 角色上报
func Role(c *gin.Context) {
	var form request.Role
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	form.Ip = common.GetClientIP(c.Request)

	// 将请求参数序列化为 JSON 字符串
	s, err := jsoniter.MarshalToString(form)
	if err != nil {
		global.App.Log.Error("Failed to marshal request params to JSON:" + err.Error())
		response.Fail(c, 500, "Failed to marshal request params to JSON")
		return
	}

	// 异步处理将数据推入 Redis 队列
	go func() {
		global.App.Log.Info("Role请求参数@" + s)

		// 将数据推入 Redis 队列
		if _, err := global.App.Redis.LPush(c, constant.RoleUpload, s).Result(); err != nil {
			global.App.Log.Error("Failed to push data to Redis:" + err.Error())
			return
		}

	}()

	// 返回成功响应
	response.Success(c, nil)
}

// 设备上报
func Device(c *gin.Context) {
	var form request.Device
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	form.Ip = common.GetClientIP(c.Request)

	// 将请求参数序列化为 JSON 字符串
	s, err := jsoniter.MarshalToString(form)
	if err != nil {
		global.App.Log.Error("Failed to marshal request params to JSON:" + err.Error())
		response.Fail(c, 500, "Failed to marshal request params to JSON")
		return
	}

	// 异步处理将数据推入 Redis 队列
	go func() {
		global.App.Log.Info("Device请求参数@" + s)

		// 将数据推入 Redis 队列
		if _, err := global.App.Redis.LPush(c, constant.DeviceUpload, s).Result(); err != nil {
			global.App.Log.Error("Failed to push data to Redis:" + err.Error())
			return
		}

	}()

	// 返回成功响应
	response.Success(c, nil)
}

// 订单上报
func Order(c *gin.Context) {
	var form request.Order
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	form.Ip = common.GetClientIP(c.Request)

	// 将请求参数序列化为 JSON 字符串
	s, err := jsoniter.MarshalToString(form)
	if err != nil {
		global.App.Log.Error("Failed to marshal request params to JSON:" + err.Error())
		response.Fail(c, 500, "Failed to marshal request params to JSON")
		return
	}

	// 异步处理将数据推入 Redis 队列
	go func() {
		global.App.Log.Info("Order请求参数@" + s)

		// 将数据推入 Redis 队列
		if _, err := global.App.Redis.LPush(c, constant.OrderUpload, s).Result(); err != nil {
			global.App.Log.Error("Failed to push data to Redis:" + err.Error())
			return
		}

	}()

	// 返回成功响应
	response.Success(c, nil)
}

// 邮件日志上报
func MailLog(c *gin.Context) {
	var form request.MailLog
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	// 将请求参数序列化为 JSON 字符串
	s, err := jsoniter.MarshalToString(form)
	if err != nil {
		global.App.Log.Error("Failed to marshal request params to JSON:" + err.Error())
		response.Fail(c, 500, "Failed to marshal request params to JSON")
		return
	}

	// 异步处理将数据推入 Redis 队列
	go func() {
		global.App.Log.Info("MailLog请求参数@" + s)

		// 将数据推入 Redis 队列
		if _, err := global.App.Redis.LPush(c, constant.MailLogUpload, s).Result(); err != nil {
			global.App.Log.Error("Failed to push data to Redis:" + err.Error())
			return
		}

	}()

	// 返回成功响应
	response.Success(c, nil)
}

// 聊天上报
func Chat(c *gin.Context) {
	var form request.Chat
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	// 将请求参数序列化为 JSON 字符串
	s, err := jsoniter.MarshalToString(form)
	if err != nil {
		global.App.Log.Error("Failed to marshal request params to JSON:" + err.Error())
		response.Fail(c, 500, "Failed to marshal request params to JSON")
		return
	}

	// 异步处理将数据推入 Redis 队列
	go func() {
		global.App.Log.Info("Chat请求参数@" + s)

		// 将数据推入 Redis 队列
		if _, err := global.App.Redis.LPush(c, constant.ChatUpload, s).Result(); err != nil {
			global.App.Log.Error("Failed to push data to Redis:" + err.Error())
			return
		}

	}()

	// 返回成功响应
	response.Success(c, nil)
}

// 私聊上报
func PrivChat(c *gin.Context) {
	var form request.PrivChat
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	// 将请求参数序列化为 JSON 字符串
	s, err := jsoniter.MarshalToString(form)
	if err != nil {
		global.App.Log.Error("Failed to marshal request params to JSON:" + err.Error())
		response.Fail(c, 500, "Failed to marshal request params to JSON")
		return
	}

	// 异步处理将数据推入 Redis 队列
	go func() {
		global.App.Log.Info("PrivChat请求参数@" + s)

		// 将数据推入 Redis 队列
		if _, err := global.App.Redis.LPush(c, constant.PrivChatUpload, s).Result(); err != nil {
			global.App.Log.Error("Failed to push data to Redis:" + err.Error())
			return
		}

	}()

	// 返回成功响应
	response.Success(c, nil)
}

// 军团日志上报
func Legion(c *gin.Context) {
	var form request.Legion
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	// 将请求参数序列化为 JSON 字符串
	s, err := jsoniter.MarshalToString(form)
	if err != nil {
		global.App.Log.Error("Failed to marshal request params to JSON:" + err.Error())
		response.Fail(c, 500, "Failed to marshal request params to JSON")
		return
	}

	// 异步处理将数据推入 Redis 队列
	go func() {
		global.App.Log.Info("Legion请求参数@" + s)

		// 将数据推入 Redis 队列
		if _, err := global.App.Redis.LPush(c, constant.LegionUpload, s).Result(); err != nil {
			global.App.Log.Error("Failed to push data to Redis:" + err.Error())
			return
		}

	}()

	// 返回成功响应
	response.Success(c, nil)
}

// 礼包推送上报
func SendGifts(c *gin.Context) {
	var form request.SendGifts
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	// 将请求参数序列化为 JSON 字符串
	s, err := jsoniter.MarshalToString(form)
	if err != nil {
		global.App.Log.Error("Failed to marshal request params to JSON:" + err.Error())
		response.Fail(c, 500, "Failed to marshal request params to JSON")
		return
	}

	// 异步处理将数据推入 Redis 队列
	go func() {
		global.App.Log.Info("SendGifts请求参数@" + s)

		// 将数据推入 Redis 队列
		if _, err := global.App.Redis.LPush(c, constant.SendGiftsUpload, s).Result(); err != nil {
			global.App.Log.Error("Failed to push data to Redis:" + err.Error())
			return
		}

	}()

	// 返回成功响应
	response.Success(c, nil)
}

// 问卷调查上报
func Answer(c *gin.Context) {
	var form request.Answer
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	// 将请求参数序列化为 JSON 字符串
	s, err := jsoniter.MarshalToString(form)
	if err != nil {
		global.App.Log.Error("Failed to marshal request params to JSON:" + err.Error())
		response.Fail(c, 500, "Failed to marshal request params to JSON")
		return
	}

	// 异步处理将数据推入 Redis 队列
	go func() {
		global.App.Log.Info("Answer请求参数@" + s)

		// 将数据推入 Redis 队列
		if _, err := global.App.Redis.LPush(c, constant.AnswerUpload, s).Result(); err != nil {
			global.App.Log.Error("Failed to push data to Redis:" + err.Error())
			return
		}

	}()

	// 返回成功响应
	response.Success(c, nil)
}

// 拦截日志上报
func Intercept(c *gin.Context) {
	var form request.Intercept
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	// 将请求参数序列化为 JSON 字符串
	s, err := jsoniter.MarshalToString(form)
	if err != nil {
		global.App.Log.Error("Failed to marshal request params to JSON:" + err.Error())
		response.Fail(c, 500, "Failed to marshal request params to JSON")
		return
	}

	// 异步处理将数据推入 Redis 队列
	go func() {
		global.App.Log.Info("Intercept请求参数@" + s)

		// 将数据推入 Redis 队列
		if _, err := global.App.Redis.LPush(c, constant.InterceptUpload, s).Result(); err != nil {
			global.App.Log.Error("Failed to push data to Redis:" + err.Error())
			return
		}

	}()

	// 返回成功响应
	response.Success(c, nil)
}

// 封禁上报
func BanLog(c *gin.Context) {
	var form request.BanLog
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	// 将请求参数序列化为 JSON 字符串
	s, err := jsoniter.MarshalToString(form)
	if err != nil {
		global.App.Log.Error("Failed to marshal request params to JSON:" + err.Error())
		response.Fail(c, 500, "Failed to marshal request params to JSON")
		return
	}

	// 异步处理将数据推入 Redis 队列
	go func() {
		global.App.Log.Info("BanLog请求参数@" + s)

		// 将数据推入 Redis 队列
		if _, err := global.App.Redis.LPush(c, constant.BanLogUpload, s).Result(); err != nil {
			global.App.Log.Error("Failed to push data to Redis:" + err.Error())
			return
		}

	}()

	// 返回成功响应
	response.Success(c, nil)
}
