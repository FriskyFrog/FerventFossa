package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//main data struct
type struct1 struct {
	struct1_number1            int
	struct1_number2            int
	struct1_FigureMethod       int
	struct1_answer             int
	struct1_FigureMethodSymbol string
}

//produce random data for struct1
func (s *struct1) UpdateData() {

	rand.Seed(time.Now().UnixNano())
	//0add 1sub 2times 3divide
	s.struct1_FigureMethod = int(rand.Intn(4))

	switch s.struct1_FigureMethod {
	//add
	case 0:
		s.struct1_number1 = rand.Intn(100)
		s.struct1_number2 = rand.Intn(100)
		s.struct1_answer = int(s.struct1_number1 + s.struct1_number2)
		s.struct1_FigureMethodSymbol = "+"
	//sub
	case 1:
		s.struct1_number2 = rand.Intn(100)
		s.struct1_number1 = rand.Intn(100) + s.struct1_number2
		s.struct1_answer = int(s.struct1_number1 - s.struct1_number2)
		s.struct1_FigureMethodSymbol = "-"
	//times
	case 2:
		s.struct1_number1 = rand.Intn(20)
		s.struct1_number2 = rand.Intn(10)
		s.struct1_answer = int(s.struct1_number1 * s.struct1_number2)
		s.struct1_FigureMethodSymbol = "×"
	//divide
	case 3:
		s.struct1_number2 = rand.Intn(10) + 1
		s.struct1_number1 = rand.Intn(20) * s.struct1_number2
		s.struct1_answer = int(s.struct1_number1 / s.struct1_number2)
		s.struct1_FigureMethodSymbol = "÷"
	}

}

func Process(data struct1, totaltime_old int, righttime_old int) (totaltime_new int, righttime_new int, over bool) {
	fmt.Printf("\n\n第%d题, %d %s %d = ", totaltime_old+1, data.struct1_number1, data.struct1_FigureMethodSymbol, data.struct1_number2)

	var int64_unum int64
	var unum string
	fmt.Scanln(&unum)
	int64_unum, _ = strconv.ParseInt(unum, 10, 64)

	if int(int64_unum) == data.struct1_answer {
		righttime_new = righttime_old + 1
		totaltime_new = totaltime_old + 1
		fmt.Printf("\n回答正确")
	} else if unum == "999" {
		righttime_new = righttime_old
		totaltime_new = totaltime_old
		//to stop
		over = true
	} else {
		fmt.Printf("\n回答错误")
		righttime_new = righttime_old
		totaltime_new = totaltime_old + 1
	}

	return
}

func main() {
	//prepare data
	var Data struct1
	var (
		RightTime, TotalTime int
		Over                 bool
		Rate                 float32
	)

	for {
		//reset data
		Data.UpdateData()

		TotalTime, RightTime, Over = Process(Data, TotalTime, RightTime)
		if Over == true {
			Rate = (float32(RightTime) / float32(TotalTime)) * 100
			fmt.Printf("\n\n游戏结束，您答对了%d题，共%d题，正确率 %2.2f%%\n\n\n", RightTime, TotalTime, Rate)
			break
		}
	}

}
