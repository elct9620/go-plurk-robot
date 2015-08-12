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
	Client      *plurk.PlurkClient
	Port        string
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

func setupClient(cmd *cobra.Command, args []string) {
	setupLogger()
	if Client == nil {
		Client = plurk.New(AppKey, AppSecret, Token, TokenSecret)
	}
}

func main() {
	// Try load credential from environment variable
	AppKey = os.Getenv("PLURK_APP_KEY")
	AppSecret = os.Getenv("PLURK_APP_SECRET")
	Token = os.Getenv("PLURK_OAUTH_TOKEN")
	TokenSecret = os.Getenv("PLURK_OAUTH_SECRET")
	Port = os.Getenv("PORT")

	// Try get robot name from environment
	RobotName = os.Getenv("PLURK_ROBOT_NAME")
	if len(RobotName) <= 0 {
		RobotName = "Plurk Robot"
	}

	rootCmd := &cobra.Command{
		Use:              "app",
		PersistentPreRun: setupClient,
	}
	// Add Commands
	rootCmd.AddCommand(cmdAddPlurk, cmdAddResponse, cmdServe, cmdRobot, cmdCreateUser)
	// Setup Flags
	rootCmd.PersistentFlags().StringVarP(&RobotName, "name", "n", RobotName, "The robot name")
	rootCmd.PersistentFlags().StringVarP(&LogFileName, "log", "l", "", "The logfile path, default is STDOUT")
	rootCmd.PersistentFlags().StringVar(&AppKey, "app_key", AppKey, "The plurk app key, suggest use environment variable")
	rootCmd.PersistentFlags().StringVar(&AppSecret, "app_secret", AppSecret, "The plurk app secret, suggest use environment variable")
	rootCmd.PersistentFlags().StringVar(&Token, "oauth_token", Token, "The plurk user oauth token, suggest use environment variable")
	rootCmd.PersistentFlags().StringVar(&TokenSecret, "oauth_secret", TokenSecret, "The plurk oauth secret token, suggest use environment variable")
	// Prepare Sub-Command Flags
	setupCommandFlags()
	// Ready!
	rootCmd.Execute()
}
