package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// main data structure including random datas
type RandomStruct struct {
	Number1            int
	Number2            int
	FigureMethod       int
	Answer             int
	FigureMethodSymbol string
}

// a method to produce random data for RandomStruct
func (r *RandomStruct) UpdateData() {

	rand.Seed(time.Now().UnixNano())
	//0:add 1:sub 2:times 3:divide
	r.FigureMethod = int(rand.Intn(4))

	switch r.FigureMethod {
	//add
	case 0:
		r.Number1 = rand.Intn(100)
		r.Number2 = rand.Intn(100)
		r.Answer = int(r.Number1 + r.Number2)
		r.FigureMethodSymbol = "+"
	//sub
	case 1:
		r.Number2 = rand.Intn(100)
		r.Number1 = rand.Intn(100) + r.Number2
		r.Answer = int(r.Number1 - r.Number2)
		r.FigureMethodSymbol = "-"
	//times
	case 2:
		r.Number1 = rand.Intn(20)
		r.Number2 = rand.Intn(10)
		r.Answer = int(r.Number1 * r.Number2)
		r.FigureMethodSymbol = "×"
	//divide
	case 3:
		r.Number2 = rand.Intn(10) + 1
		r.Number1 = rand.Intn(20) * r.Number2
		r.Answer = int(r.Number1 / r.Number2)
		r.FigureMethodSymbol = "÷"
	}
}

// main data struct including input, grade, etc.
type UserStruct struct {
	TotalTime         int
	RightTime         int
	InputNumber       string
	InputNumber_int64 int64
	Over              bool
	Timeout          bool
	TimeoutLock     bool
}

// a method of processing
func (u *UserStruct) UpdateData(data RandomStruct) {

	//show the task
	fmt.Printf("\n\n第%d题, %d %s %d = ", u.TotalTime+1, data.Number1, data.FigureMethodSymbol, data.Number2)

	go func (){
		//get data from input
		fmt.Scanln(&u.InputNumber)
		u.TimeoutLock = true
	}()

	select {
	case <- time.After(time.Second*5):
		if u.TimeoutLock == false{
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

	u.Timeout = false
	u.TimeoutLock = false
}

func main() {
	//Channel1 := make(chan string)
	for {
		//prepare data
		var RandomData RandomStruct
		var UserData UserStruct
		var Rate float32

		//show instructions
		fmt.Println("游戏开始，输入\"999\"退出并查看答题情况")
		time.Sleep(time.Second*2)

		for {
			//reset data
			RandomData.UpdateData()
			UserData.UpdateData(RandomData)

			if UserData.Over == true {
				Rate = (float32(UserData.RightTime) / float32(UserData.TotalTime)) * 100
				fmt.Printf("\n\n游戏结束，您答对了%d题，共%d题，正确率 %2.2f%%\n\n\n", UserData.RightTime, UserData.TotalTime, Rate)
				break
			}
		}
	}
}
