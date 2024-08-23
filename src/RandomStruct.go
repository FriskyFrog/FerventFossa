package main

import (
	"math/rand"
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
		r.FigureMethodSymbol = "x"
	case 3:
		//divide
		r.Number2 = rand.Intn(10) + 1
		r.Number1 = rand.Intn(20) * r.Number2
		r.Answer = int(r.Number1 / r.Number2)
		r.FigureMethodSymbol = "÷"
	}
}
