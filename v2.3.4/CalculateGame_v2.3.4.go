package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// main data structure including random elements
type RandomStruct struct {
	Number1            int    //first number
	Number2            int    //second number
	FigureMethod       int    //the way to figure
	Answer             int    //correct answer
	FigureMethodSymbol string //symbol to show out
}

// a method to produce random data for RandomStruct
func (r *RandomStruct) UpdateData() {

	//use current time to make data random
	rand.Seed(time.Now().UnixNano())
	//0 for add, 1 for sub, 2 for times, 3 for divide
	r.FigureMethod = int(rand.Intn(4))

	//match method
	switch r.FigureMethod {
	case 0:
		//add
		r.Number1 = rand.Intn(100)
		r.Number2 = rand.Intn(100)
		r.Answer = int(r.Number1 + r.Number2)
		r.FigureMethodSymbol = "+"
	case 1:
		//sub
		r.Number2 = rand.Intn(100)
		r.Number1 = rand.Intn(100) + r.Number2
		r.Answer = int(r.Number1 - r.Number2)
		r.FigureMethodSymbol = "-"
	case 2:
		//times
		r.Number1 = rand.Intn(20)
		r.Number2 = rand.Intn(10)
		r.Answer = int(r.Number1 * r.Number2)
		r.FigureMethodSymbol = "×"
	case 3:
		//divide
		r.Number2 = rand.Intn(10) + 1
		r.Number1 = rand.Intn(20) * r.Number2
		r.Answer = int(r.Number1 / r.Number2)
		r.FigureMethodSymbol = "÷"
	}
}

// data struct for recording
type UserStruct struct {
	TotalTime         int    //the total times the player done
	RightTime         int    //the sum of the correct answer the player done
	InputNumber       string //receive input from console
	InputNumber_int64 int64  //to store the converted input
	Over              bool   //a sign to quit
	Timeout           bool   //to mark if timeout
	TimeoutLock       bool   //a lock for Timeout
}

// a method for processing
func (u *UserStruct) UpdateData(data RandomStruct) {

	//show the task
	fmt.Printf("\n\n第%d题, %d %s %d = ", u.TotalTime+1, data.Number1, data.FigureMethodSymbol, data.Number2)

	go func() {
		//get data from input
		fmt.Scanln(&u.InputNumber)
		u.TimeoutLock = true
	}()

	select {
	//wait for a while
	case <-time.After(time.Second * 5):
		if u.TimeoutLock == false {
			u.Timeout = true
		}
	}

	//convert to int64 in order to compare later
	u.InputNumber_int64, _ = strconv.ParseInt(u.InputNumber, 10, 64)

	//compare input with right answer
	if int(u.InputNumber_int64) == data.Answer {
		u.RightTime++
		u.TotalTime++
		fmt.Printf("\n回答正确")
		if u.Timeout == true {
			u.RightTime--
			fmt.Printf("\n回答超时")
		}
	} else if u.InputNumber == "999" {
		//mark to stop
		u.Over = true
	} else {
		fmt.Printf("\n回答错误")
		u.TotalTime++
	}
	//make the two variables false
	u.Timeout = false
	u.TimeoutLock = false
}

func main() {

	//main loop
	for {
		//prepare data
		var RandomData RandomStruct
		var UserData UserStruct
		var Rate float32

		//show instructions
		fmt.Println("游戏开始，输入\"999\"退出并查看答题情况\n每题只有5秒钟的时间")
		time.Sleep(time.Second * 2)

		for {
			//reset data
			RandomData.UpdateData()
			UserData.UpdateData(RandomData)

			//in the end
			if UserData.Over == true {
				Rate = (float32(UserData.RightTime) / float32(UserData.TotalTime)) * 100
				fmt.Printf("\n\n游戏结束，您答对了%d题，共%d题，正确率 %2.2f%%\n\n\n", UserData.RightTime, UserData.TotalTime, Rate)
				break
			}
		}
	}
}
