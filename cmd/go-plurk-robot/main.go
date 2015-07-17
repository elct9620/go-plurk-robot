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

	polling := client.GetPolling()
	plurks, err := polling.GetPlurks(plurk.Now(), 20)

	if err != nil {
		logger.Error("Error: %s", err.Error())
	}

	logger.Info("New plurks %#v", plurks)

}
