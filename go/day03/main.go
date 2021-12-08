package main

import (
	"aoc21_go/util"
	"fmt"
	"strconv"
)

func getCommons(lines []string, pos int) (string, string) {
	var zeros, ones int
	for _, line := range lines {
		if line[pos] == '0' {
			zeros += 1
		} else {
			ones += 1
		}
	}

	if zeros > ones {
		return "0", "1"
	} else {
		return "1", "0"
	}
}

func bitstringToInt(bitstring string) int {
	intValue, _ := strconv.ParseInt(bitstring, 2, 64)
	return int(intValue)
}

func evalA(lines []string) int {
	cntBits := len(lines[0])

	var gamma, epsilon string
	for i := 0; i < cntBits; i++ {
		mostCommon, leastCommon := getCommons(lines, i)
		gamma += mostCommon
		epsilon += leastCommon
	}

	return bitstringToInt(gamma) * bitstringToInt(epsilon)
}

func getCommonLines(lines []string, pos int, flagMostCommon bool) string {
	mostCommon, leastCommon := getCommons(lines, pos)
	var commonLines []string

	for _, line := range lines {
		if string(line[pos]) == mostCommon && flagMostCommon {
			commonLines = append(commonLines, line)
		}
		if string(line[pos]) == leastCommon && !flagMostCommon {
			commonLines = append(commonLines, line)
		}
	}

	if len(commonLines) == 1 {
		return commonLines[0]
	} else {
		return getCommonLines(commonLines, pos+1, flagMostCommon)
	}
}

func evalB(lines []string) int {
	oxygen := getCommonLines(lines, 0, true)
	co2 := getCommonLines(lines, 0, false)

	return bitstringToInt(oxygen) * bitstringToInt(co2)
}

func main() {
	println("132"[2])
	lines := util.ReadFile("input03.txt")
	resA := evalA(lines)
	resB := evalB(lines)
	fmt.Printf("A: %v \n", resA)
	fmt.Printf("B: %v \n", resB)
}
