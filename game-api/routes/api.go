package routes

import (
	"my-gin/app/controllers/app"

	"github.com/gin-gonic/gin"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.POST("/users/upload", app.User)
	router.POST("/role/upload", app.Role)
	router.POST("/device/upload", app.Device)
	router.POST("/order/upload", app.Order)
	router.POST("/mail_log/upload", app.MailLog)
	router.POST("/chat/upload", app.Chat)
	router.POST("/priv_chat/upload", app.PrivChat)
	router.POST("/legion/upload", app.Legion)
	router.POST("/send_gifts/upload", app.SendGifts)
	router.POST("/answer/upload", app.Answer)
	router.POST("/intercept/upload", app.Intercept)
	router.POST("/ban_log/upload", app.BanLog)
}
