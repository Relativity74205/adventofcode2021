package main

import (
	"aoc21_go/util"
	"fmt"
)

func evalA(lines []int) int {
	var cnt = 0
	for i := range lines {
		if i == 0 {
			continue
		}
		if lines[i] > lines[i-1] {
			cnt += 1
		}
	}

	return cnt
}

func evalB(lines []int) int {
	var cnt = 0
	for i := range lines {
		if i <= 2 {
			continue
		}
		if util.SumIntegers(lines[i-2:i+1]) > util.SumIntegers(lines[i-3:i]) {
			cnt += 1
		}
	}

	return cnt
}

func main() {
	lines := util.ReadFile("input01.txt")
	linesInt := util.LinesToInt(lines)
	resA := evalA(linesInt)
	resB := evalB(linesInt)
	fmt.Printf("A: %v \n", resA)
	fmt.Printf("B: %v \n", resB)
}
