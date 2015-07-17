package main

import (
	"github.com/elct9620/go-plurk-robot/logger"
	"github.com/spf13/cobra"
	"strings"
)

func setupCommandFlags() {
	// Add Plurk
	cmdAddPlurk.Flags().StringP("lang", "L", "en", "Specify this plurk language")
	cmdAddPlurk.Flags().StringP("qualifier", "q", ":", "Spacify plurk qualifier, Ex. says, think")
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
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// A helpful web ui for manage the robots
var cmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Start a Web UI for robot",
	Long:  `Start a Web UI with analytics and some useful tool for robot and plurk`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// The main command, user for start a new robot
var cmdRobot = &cobra.Command{
	Use:   "robot",
	Short: "Start a robot",
	Long:  `Start a robot with task can automatic add plurk/response and other things`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
