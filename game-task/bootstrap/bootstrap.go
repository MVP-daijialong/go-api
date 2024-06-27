package bootstrap

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// RunServer 启动服务器
func RunServer() {
	// 打印成功启动消息的 ASCII 艺术图案
	printSuccessMessage()

	// 等待中断信号以优雅地关闭调度器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Scheduler ...")

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Optionally, add any cleanup logic here

	log.Println("Scheduler exiting")
}

func printSuccessMessage() {
	message := `
	____   ____  ___ ___    ___      ______   ____  _____ __  _      ____  __ __  ____        _____ __ __    __    __    ___  _____ _____
	/    | /    ||   |   |  /  _]    |      | /    |/ ___/|  |/ ]    |    \|  |  ||    \      / ___/|  |  |  /  ]  /  ]  /  _]/ ___// ___/
   |   __||  o  || _   _ | /  [_     |      ||  o  (   \_ |  ' /     |  D  )  |  ||  _  |    (   \_ |  |  | /  /  /  /  /  [_(   \_(   \_ 
   |  |  ||     ||  \_/  ||    _]    |_|  |_||     |\__  ||    \     |    /|  |  ||  |  |     \__  ||  |  |/  /  /  /  |    _]\__  |\__  |
   |  |_ ||  _  ||   |   ||   [_       |  |  |  _  |/  \ ||     \    |    \|  :  ||  |  |     /  \ ||  :  /   \_/   \_ |   [_ /  \ |/  \ |
   |     ||  |  ||   |   ||     |      |  |  |  |  |\    ||  .  |    |  .  \     ||  |  |     \    ||     \     \     ||     |\    |\    |
   |___,_||__|__||___|___||_____|      |__|  |__|__| \___||__|\_|    |__|\_|\__,_||__|__|      \___| \__,_|\____|\____||_____| \___| \___|
																																		                                                            
`
	fmt.Println(message)
}
