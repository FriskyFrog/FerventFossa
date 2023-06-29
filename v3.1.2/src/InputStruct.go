package main

import (
	"fmt"
	"strconv"
	"time"
)

// data struct for recording
type InputStruct struct {
	TotalTime         int    //the total times the player done
	RightTime         int    //the sum of the correct answer the player done
	InputNumber       string //receive input from console
	InputNumber_int64 int64  //to store the converted input
	Over              bool   //a sign to quit
}

// a method for processing
func (i *InputStruct) UpdateData(data RandomStruct) {

	//show the task
	fmt.Printf("\n\n第%d题, %d %s %d = ", i.TotalTime+1, data.Number1, data.FigureMethodSymbol, data.Number2)

	go func() {
		//get data from input
		fmt.Scanln(&i.InputNumber)
	}()

	select {
	//wait for a while
	case <-time.After(time.Second * time.Duration(WaitingTime)):
		func() {
			//do nothing
		}()
	}

	//convert to int64 in order to compare later
	i.InputNumber_int64, _ = strconv.ParseInt(i.InputNumber, 10, 64)

	//compare input with right answer
	if int(i.InputNumber_int64) == data.Answer {
		i.RightTime++
		i.TotalTime++
		fmt.Printf("\n回答正确")
	} else if i.InputNumber == "999" {
		//mark to stop
		i.Over = true
	} else {
		fmt.Printf("\n回答错误")
		i.TotalTime++
	}
}
