package main

import (
	"aoc21_go/util"
	"fmt"
	"sort"
)

func analyzeLine(line string) string {
	var openBrackets []string
	for _, char := range line {
		if checkOpenBracket(string(char)) {
			openBrackets = append(openBrackets, string(char))
		} else {
			switch string(char) {
			case ")":
				if util.GetLastElementStrings(openBrackets) == "(" {
					openBrackets = util.RemoveLastElementStrings(openBrackets)
				} else {
					return "("
				}
			case "]":
				if util.GetLastElementStrings(openBrackets) == "[" {
					openBrackets = util.RemoveLastElementStrings(openBrackets)
				} else {
					return "["
				}
			case "}":
				if util.GetLastElementStrings(openBrackets) == "{" {
					openBrackets = util.RemoveLastElementStrings(openBrackets)
				} else {
					return "{"
				}
			case ">":
				if util.GetLastElementStrings(openBrackets) == "<" {
					openBrackets = util.RemoveLastElementStrings(openBrackets)
				} else {
					return "<"
				}
			}
		}
	}

	var openBracketsString string
	for _, openBracket := range openBrackets {
		openBracketsString += openBracket
	}

	return openBracketsString
}

func calcSyntaxErrorScore(lineReport string) int {
	switch lineReport {
	case "(":
		return 3
	case "[":
		return 57
	case "{":
		return 1197
	case "<":
		return 25137
	default:
		return 0
	}
}

func checkOpenBracket(char string) bool {
	if char == "(" || char == "[" || char == "{" || char == "<" {
		return true
	}
	return false
}

func evalA(lines []string) int {
	var sumSyntaxErrorScore int
	for _, line := range lines {
		lineReport := analyzeLine(line)
		sumSyntaxErrorScore += calcSyntaxErrorScore(lineReport)
	}

	return sumSyntaxErrorScore
}

func calcAutoCompleteScore(lineReport string) int {
	if checkOpenBracket(lineReport) {
		return 0
	}

	var autoCompleteScore int
	for _, char := range util.ReverseString(lineReport) {
		autoCompleteScore *= 5
		switch string(char) {
		case "(":
			autoCompleteScore += 1
		case "[":
			autoCompleteScore += 2
		case "{":
			autoCompleteScore += 3
		case "<":
			autoCompleteScore += 4
		}
	}

	return autoCompleteScore
}

func evalB(lines []string) int {
	var sumAutoCompleteScore []int
	for _, line := range lines {
		lineReport := analyzeLine(line)
		autoCompleteScore := calcAutoCompleteScore(lineReport)
		if autoCompleteScore != 0 {
			sumAutoCompleteScore = append(sumAutoCompleteScore, autoCompleteScore)
		}
	}
	sort.Ints(sumAutoCompleteScore)

	return sumAutoCompleteScore[(len(sumAutoCompleteScore)-1)/2]
}

func main() {
	day := 10
	filename := fmt.Sprintf("input%02d.txt", day)
	lines := util.ReadFile(filename)

	resA := evalA(lines)
	resB := evalB(lines)
	fmt.Printf("Day %02d \n", day)
	fmt.Printf("A: %v \n", resA)
	fmt.Printf("B: %v \n", resB)
}
