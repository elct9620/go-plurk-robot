package main

import (
	"github.com/elct9620/go-plurk-robot/logger"
	"github.com/elct9620/go-plurk-robot/plurk"
	"github.com/spf13/cobra"
	"io"
	"os"
)

var (
	AppKey      string
	AppSecret   string
	Token       string
	TokenSecret string
	RobotName   string    = "Plurk Robot" // Robot Name
	LogFile     io.Writer = os.Stdout     // Robot default out message to STDOUT
	LogFileName string    = ""
)

func setupLogger() {
	if len(LogFileName) > 0 {
		var err error
		LogFile, err = os.OpenFile(LogFileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			LogFile = os.Stdout
		}
	}

	logger.Config(LogFile, RobotName)
}

func main() {
	// Try load credential from environment variable
	AppKey = os.Getenv("PLURK_APP_KEY")
	AppSecret = os.Getenv("PLURK_APP_SECRET")
	Token = os.Getenv("PLURK_OAUTH_TOKEN")
	TokenSecret = os.Getenv("PLURK_OAUTH_SECRET")

	plurk.New(AppKey, AppSecret, Token, TokenSecret)

	rootCmd := &cobra.Command{Use: "app"}
	// Add Commands
	rootCmd.AddCommand(cmdAddPlurk, cmdAddResponse, cmdServe, cmdRobot)
	// Setup Flags
	rootCmd.PersistentFlags().StringVarP(&RobotName, "name", "n", RobotName, "The robot name")
	rootCmd.PersistentFlags().StringVarP(&LogFileName, "log", "l", "", "The logfile path, default is STDOUT")
	rootCmd.Execute()

	setupLogger()
}
