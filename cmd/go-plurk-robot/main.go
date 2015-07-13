package main

import (
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

	client := plurk.New(AppKey, AppSecret, Token, TokenSecret)
	client.Echo("Hello World++ WTF 囧 __**//\\QQ AA")
	client.Echo("")
}
