package main

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// to store the time to wait
var (
	WaitingTime      int64 = 5
	WaitingTime_temp int64
	WaitingTime_int  int
)

// main, means itself,-h or --help or no flag
var RootCmd = &cobra.Command{
	Use:   "FerventFossa",
	Short: "FerventFossa is a simple game about maths",
	Long:  "FerventFossa is a simple game about maths.\n\nIt's free and open source under GNU GPLv3.\n\nVisit www.github.com\\FriskyFrog\\FerventFossa for more info",
	Run: func(cmd *cobra.Command, args []string) {
		MainLoop()
	},
}

// a command to get the version
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version of FerventFossa",
	Long:  "All software has versions, so does FerventFossa",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("-----FerventFossa v1.0.0-----\n --under GNU GPLv3 lisense\n**********2024.05  :-)")
	},
}

// set the time to wait
var settimeCmd = &cobra.Command{
	Use:   "settime",
	Short: "set time ",
	Long:  "set a suitable period of time to commit answer",
	Run: func(cmd *cobra.Command, args []string) {
		//convert string to int64
		WaitingTime_temp, _ = strconv.ParseInt(args[0], 10, 64)
		//neither too short nor too long
		if (3 < WaitingTime_temp) && (WaitingTime_temp < 10) {
			WaitingTime = WaitingTime_temp
		}
		MainLoop()
	},
	//only accept one arg
	Args: cobra.ExactArgs(1),
}

// add sub-commands
func init() {
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(settimeCmd)
}
