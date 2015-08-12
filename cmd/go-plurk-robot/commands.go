package main

import (
	"github.com/elct9620/go-plurk-robot/db"
	"github.com/elct9620/go-plurk-robot/logger"
	"github.com/elct9620/go-plurk-robot/robot"
	"github.com/elct9620/go-plurk-robot/server"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

func setupCommandFlags() {
	// Add Plurk
	cmdAddPlurk.Flags().StringP("lang", "L", "en", "Specify this plurk language")
	cmdAddPlurk.Flags().StringP("qualifier", "q", ":", "Spacify plurk qualifier, Ex. says, thinks")

	// Add Response
	cmdAddResponse.Flags().StringP("qualifier", "q", ":", "Specify plurk qualifier, Ex. says, thinks")

	// Add Server
	cmdServe.Flags().StringVarP(&Port, "port", "p", "5000", "Specify web server port")
}

// A shortcut for Add Plurk, can use with cronjob
var cmdAddPlurk = &cobra.Command{
	Use:   "plurk content [qualifier]",
	Short: "Add a new plruk",
	Long:  `Add a new plurk to your robot timeline`,
	Run:   addPlurk,
}

// Add Plurk Implement
func addPlurk(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		lang, _ := cmd.Flags().GetString("lang")
		qualifier, _ := cmd.Flags().GetString("qualifier")

		timeline := Client.GetTimeline()
		res, err := timeline.PlurkAdd(strings.Join(args, " "), qualifier, make([]int, 0), false, lang, true)
		if err != nil {
			logger.FError(cmd.Out(), err.Error())
			return
		}
		logger.FInfo(cmd.Out(), "Success, Plurk ID: %d", res.PlurkID)
	} else {
		logger.FError(cmd.Out(), "No plurk content specified")
	}
}

// A shortcut for Add Response, can use with cronjob
var cmdAddResponse = &cobra.Command{
	Use:   "response id content [quailfier]",
	Short: "Response to specify plurk",
	Long:  `Add a new response to specify plurk on timeline`,
	Run:   addResponse,
}

func addResponse(cmd *cobra.Command, args []string) {
	if len(args) > 1 {
		plurkID, err := strconv.Atoi(args[0:1][0])
		responseContent := args[1:]
		qualifier, _ := cmd.Flags().GetString("qualifier")

		if err != nil {
			logger.FError(cmd.Out(), "Convert plurk id error, reason: %s", err.Error())
			return
		}

		response := Client.GetResponses()
		res, err := response.ResponseAdd(plurkID, strings.Join(responseContent, " "), qualifier)

		if err != nil {
			logger.FError(cmd.Out(), err.Error())
		} else {
			logger.FInfo(cmd.Out(), "Respons success add to %d, and response id is %d", plurkID, res.Id)
		}
	} else {
		logger.FError(cmd.Out(), "No plurk id or response content specified")
	}
}

// A helpful web ui for manage the robots
var cmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Start a Web UI for robot",
	Long:  `Start a Web UI with analytics and some useful tool for robot and plurk`,
	Run: func(cmd *cobra.Command, args []string) {
		server.Serve(Port)
	},
}

// The main command, user for start a new robot
var cmdRobot = &cobra.Command{
	Use:   "robot",
	Short: "Start a robot",
	Long:  `Start a robot with task can automatic add plurk/response and other things`,
	Run: func(cmd *cobra.Command, args []string) {
		robot.SetupPlurk(AppKey, AppSecret, Token, TokenSecret)
		r := robot.New()
		if r != nil {
			r.Start()
		}
	},
}

var cmdCreateUser = &cobra.Command{
	Use:   "useradd",
	Short: "Create a user",
	Long:  `Create a user for admin plurk robot script, first argument is username, secondary is password`,
	Run: func(cmd *cobra.Command, args []string) {
		session, err := db.OpenSession("")

		if err != nil {
			logger.FError(cmd.Out(), err.Error())
			return
		}

		if len(args) < 2 {
			logger.FError(cmd.Out(), "You should specify username and password")
			return
		}

		user, err := db.CreateUser(session.DB(""), args[0], args[1])

		if err != nil {
			logger.FError(cmd.Out(), err.Error())
			return
		}

		logger.FInfo(cmd.Out(), "User \"%s\" create success!", user.Username)

	},
}
