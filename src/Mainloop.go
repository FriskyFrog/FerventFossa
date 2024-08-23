package main

import (
	"fmt"
	"time"
)

func MainLoop() {

	//main loop
	for {
		//prepare data
		var RandomData RandomStruct
		var InputData InputStruct
		var Rate float32
		WaitingTime_int = int(WaitingTime)

		//show instructions
		fmt.Printf("游戏开始，输入\"999\"退出并查看答题情况\n每题只有%d秒钟的时间", WaitingTime_int)
		time.Sleep(time.Second * 2)

		for {
			//reset data
			RandomData.UpdateData()
			InputData.UpdateData(RandomData)

			//in the end
			if InputData.Over == true {
				Rate = (float32(InputData.RightTime) / float32(InputData.TotalTime)) * 100
				fmt.Printf("\n\n游戏结束，您答对了%d题，共%d题，正确率 %2.2f%%\n\n\n", InputData.RightTime, InputData.TotalTime, Rate)
				time.Sleep(time.Second * 5)
				break
			}
		}
	}
}
