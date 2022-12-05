package aoc2022

import (
	"fmt"
	"os"
	"strconv"
)

var solutions = []func(in string){
	DayOne,
	DayTwo,
	DayThree,
	DayFour,
	DayFive,
}

func Run(day int) {
	day -= 1
	if day < 0 || day >= len(solutions) {
		fmt.Println("Invalid day number.")
		return
	}
	solutions[day](ReadInput(day))
}

func ReadInput(day int) string {
	f, err := os.ReadFile("./input/day" + strconv.Itoa(day+1) + ".txt")
	if err != nil {
		panic(err)
	}
	return string(f)
}
