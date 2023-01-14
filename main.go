package main

import (
	"clipboardshare-server/gobalConfig"
	"clipboardshare-server/router"
	"github.com/gin-gonic/gin"

	"log"
)

func main() {

	gin.SetMode(gobalConfig.GinMode)
	r := gin.Default()
	if gobalConfig.FrontMode {
		log.Println("已开启前后端整合模式！")
		r.LoadHTMLGlob("static/index.html")
		r.Static("/static", "static")
		r.GET("/", func(context *gin.Context) {
			context.HTML(200, "index.html", "")
		})
	}
	router.RegRouter(r)
	log.Println("定时任务启动成功,服务启动成功,当前使用端口：", gobalConfig.Port)
	r.Run(":" + gobalConfig.Port)
}
