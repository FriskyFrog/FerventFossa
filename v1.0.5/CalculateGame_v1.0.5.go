package main

import(
  "fmt"
  "time"
  "math/rand"
)

//produce random data and the mothed to process them
func producedata()(rnm1,rnm2,rmethod int){
  rand.Seed(time.Now().UnixNano())
  rmethod = int(rand.Intn(4)+1)
  if rmethod == 4 {
    rnm2 = (rand.Intn(10)+1)
    rnm1 = rand.Intn(20)*rnm2
  } else if rmethod == 3 {
    rnm1 = rand.Intn(20)
    rnm2 = rand.Intn(10)
  } else if rmethod == 2{
    rnm2 = rand.Intn(100)
    rnm1 = rand.Intn(100)+rnm2
  } else {
    rnm1 = rand.Intn(100)
    rnm2 = rand.Intn(100)
  }
  return
}

//a func to figure out the right data
func figuredata(nm1,nm2,fgmethod int) (fgedanswer int){
  if fgmethod == 1 {
    fgedanswer = int(nm1+nm2)
  } else if fgmethod == 2 {
    fgedanswer = int(nm1-nm2)
  } else if fgmethod == 3 {
    fgedanswer = int(nm1*nm2)
  } else if fgmethod == 4 {
    fgedanswer = int(nm1/nm2)
  }
  return
}

func main(){
  for {
    var(
      num1,num2,method,answer,righttime,totaltime,unum int
      rate float32
      symbol string
    )

    for {
      totaltime++
      num1,num2,method = producedata()
      answer = figuredata(num1,num2,method)
      switch method {
      case 4: symbol = "÷"
      case 3: symbol = "×"
      case 2: symbol = "-"
      case 1: symbol = "+"
      }

      fmt.Printf("\n\n第%d题,输入999查看正确率并重新开始, %d %s %d = ",totaltime,num1,symbol,num2)
      fmt.Scanln(&unum)

      if unum == answer {
        righttime++
        fmt.Printf("\n回答正确")
      } else if unum == 999 {
        break
      } else {
        fmt.Printf("\n回答错误")
      }

      time.Sleep(time.Second)
    }
  if totaltime != 1 {
    rate = (float32(righttime)/float32(totaltime-1))*100.0
  } else {
    rate = 0.000000
  }
  fmt.Printf("\n\n游戏结束，您答对了%d题，共%d题，正确率 %2.2f%%\n\n\n",righttime,totaltime-1,rate)
  time.Sleep(time.Second*4)
  }
}
