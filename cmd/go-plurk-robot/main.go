package main

import (
	"github.com/elct9620/go-plurk-robot/logger"
	"github.com/elct9620/go-plurk-robot/plurk"
	"os"
)

var (
	AppKey      string
	AppSecret   string
	Token       string
	TokenSecret string
)

func main() {
	AppKey = os.Getenv("PLURK_APP_KEY")
	AppSecret = os.Getenv("PLURK_APP_SECRET")
	Token = os.Getenv("PLURK_OAUTH_TOKEN")
	TokenSecret = os.Getenv("PLURK_OAUTH_SECRET")

	logger.Config(os.Stdout, "Plurk Robot")

	client := plurk.New(AppKey, AppSecret, Token, TokenSecret)
	echo, _ := client.Echo("Hello World?")
	logger.Info("Echo: %s", echo.Data)

	timeline := client.GetTimeline()

	_, err := timeline.PlurkAdd("發噗測試 #蒼時機器人", "says", make([]int, 0), false, "tr_ch")
	if err != nil {
		logger.Error("Error plurk add %s", err.Error())
	}
}
