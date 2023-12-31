package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var (
	WaitingTime      int64 = 5
	WaitingTime_temp int64
)

var RootCmd = &cobra.Command{
	Use:   "CalculateGame",
	Short: "CalculateGame is a simple game about maths",
	Long:  "CalculateGame is a simple game about maths.\n\nIt's free and open source under GNU GPLv3.\n\nVisit www.github.com\\atommation\\CalculateGame for more info",
	Run: func(cmd *cobra.Command, args []string) {
		MainLoop()
	},
}

// a command to get the version
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version of CalculateGame",
	Long:  "All software has versions, so does CalculateGame",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("-----CalculateGame v3.0.6-----\n --under GNU GPLv3 lisense\n**********2023.06  :-)")
	},
}

// set the time to wait
var settimeCmd = &cobra.Command{
	Use:   "settime",
	Short: "set time ",
	Long:  "set a suitable period of time to commit answer",
	Run: func(cmd *cobra.Command, args []string) {
		WaitingTime_temp, _ = strconv.ParseInt(args[0], 10, 64)
		if (4 < WaitingTime_temp) && (WaitingTime_temp < 10) {
			WaitingTime = WaitingTime_temp
		}
	},
	Args: cobra.ExactArgs(1),
}

// add sub-commands
func init() {
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(settimeCmd)
}
