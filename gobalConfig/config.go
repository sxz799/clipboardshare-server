package gobalConfig

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var (
	Port       string
	FrontMode  bool
	HistoryNum int
	GinMode    string
)

func init() {
	log.Println("正在检查files文件夹")
	_, err := os.Stat("files")
	if err != nil && os.IsNotExist(err) {
		os.Mkdir("files", os.ModePerm)
	}
}

func init() {
	log.Println("正在应用配置文件...")
	viper.SetConfigName("conf")
	viper.SetConfigType("yml")
	viper.AddConfigPath("conf")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicln("viper load fail ...")
		return
	}

	Port = viper.GetString("server.port")
	GinMode = viper.GetString("server.ginMode")

	FrontMode = viper.GetBool("config.frontMode")
	HistoryNum = viper.GetInt("config.historyNum")
}
