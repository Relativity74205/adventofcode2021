package main

import (
	"aoc21_go/util"
	"fmt"
)

func evalA(lines []string) int {

	return 0
}

func evalB(lines []string) int {

	return 0
}

func main() {
	day := 1
	debugSuffix := ""
	filename := fmt.Sprintf("input%02d%v.txt", day, debugSuffix)
	lines := util.ReadFile(filename)

	resA := evalA(lines)
	resB := evalB(lines)
	fmt.Printf("Day %02d \n", day)
	fmt.Printf("A: %v \n", resA)
	fmt.Printf("B: %v \n", resB)
}
