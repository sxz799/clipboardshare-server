package router

import (
	"clipboardshare-server/model"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

func read(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":      200,
		"success":   true,
		"message":   "获取成功!",
		"clipboard": model.CurrClipboard,
	})
}

func write(c *gin.Context) {
	var clipboard model.Clipboard
	c.ShouldBind(&clipboard)
	model.CurrClipboard = clipboard
	model.AddLogs(clipboard)
	c.JSON(200, gin.H{
		"status":  200,
		"success": true,
		"message": "更新成功!",
	})
}
func listLog(c *gin.Context) {
	logs, _ := json.Marshal(model.GetLogData())
	c.String(200, string(logs))
	return
}

func Clipboard(e *gin.Engine) {
	g := e.Group("/clipboard")
	{
		g.GET("/read", read)
		g.POST("/write", write)
		g.GET("/listLog", listLog)
	}
}
