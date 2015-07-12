package main

import (
	"github.com/elct9620/go-plurk-robot/plurk"
	"os"
)

var (
	AppKey      string
	AppSecret   string
	OauthToken  string
	TokenSecret string
)

func main() {
	AppKey = os.Getenv("PLURK_APP_KEY")
	AppSecret = os.Getenv("PLURK_APP_SECRET")
	OauthToken = os.Getenv("PLURK_OAUTH_TOKEN")
	TokenSecret = os.Getenv("PLURK_OAUTH_SECRET")

	client := plurk.New(AppKey, AppSecret, OauthToken, TokenSecret)
	client.Echo()
}
