package main

import (
	"github.com/spf13/cobra"
)

// A shortcut for Add Plurk, can use with cronjob
var cmdAddPlurk = &cobra.Command{
	Use:   "plurk content [qualifier]",
	Short: "Add a new plruk",
	Long:  `Add a new plurk to your robot timeline`,
	Run: func(cmd *cobra.Command, args []string) {

	},
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
