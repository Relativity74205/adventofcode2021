package main

import (
	"aoc21_go/util"
	"fmt"
)

func eval01A(lines []int) int {
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

func eval01B(lines []int) int {
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
	resA := eval01A(linesInt)
	resB := eval01B(linesInt)
	fmt.Printf("Day01a: %v measurements\n", resA)
	fmt.Printf("Day01b: %v measurements\n", resB)
}
