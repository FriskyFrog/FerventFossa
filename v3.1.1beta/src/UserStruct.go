package main

import (
	"fmt"
	"strconv"
	"time"
)

// data struct for recording
type UserStruct struct {
	TotalTime         int    //the total times the player done
	RightTime         int    //the sum of the correct answer the player done
	InputNumber       string //receive input from console
	InputNumber_int64 int64  //to store the converted input
	Over              bool   //a sign to quit
	//Timeout           bool   //to mark if timeout
	//TimeoutLock       bool   //a lock for Timeout
}

// a method for processing
func (u *UserStruct) UpdateData(data RandomStruct) {

	//show the task
	fmt.Printf("\n\n第%d题, %d %s %d = ", u.TotalTime+1, data.Number1, data.FigureMethodSymbol, data.Number2)

	go func() {
		//get data from input
		fmt.Scanln(&u.InputNumber)
		//u.TimeoutLock = true
	}()

	select {
	//wait for a while
	case <-time.After(time.Second * time.Duration(WaitingTime)):
		func () {
			//
		} ()
	}

	//convert to int64 in order to compare later
	u.InputNumber_int64, _ = strconv.ParseInt(u.InputNumber, 10, 64)

	//compare input with right answer
	if int(u.InputNumber_int64) == data.Answer {
		u.RightTime++
		u.TotalTime++
		fmt.Printf("\n回答正确")
	} else if u.InputNumber == "999" {
		//mark to stop
		u.Over = true
	} else {
		fmt.Printf("\n回答错误")
		u.TotalTime++
	}
	//make the two variables false
	//u.Timeout = false
	//u.TimeoutLock = false
}
