package model

import (
	"clipboardshare-server/gobalConfig"
	"time"
)

var logs []ClipboardLog
var CurrClipboard Clipboard

type Clipboard struct {
	Content string `json:"content"`
}

type Result struct {
	Code      int       `json:"code"`
	Success   bool      `json:"success"`
	Message   string    `json:"message"`
	Clipboard Clipboard `json:"clipboard"`
}

type ClipboardLog struct {
	Id   int    `json:"id"`
	Msg  string `json:"msg"`
	Date string `json:"date"`
}

func AddLogs(c Clipboard) {
	if len(logs) < gobalConfig.HistoryNum {
		logs = append(logs, ClipboardLog{
			Id:   0,
			Msg:  c.Content,
			Date: time.Now().String(),
		})
	} else {
		logs = logs[:gobalConfig.HistoryNum-1]
		logs = append(logs, ClipboardLog{
			Id:   0,
			Msg:  c.Content,
			Date: time.Now().String(),
		})
	}
}

func GetLogData() []ClipboardLog {
	return logs
}
