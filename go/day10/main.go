package main

import (
	"aoc21_go/util"
	"fmt"
)

func findCorruption(line string) string {
	var openBrackets []string
	for _, char := range line {
		if checkOpenBracket(string(char)) {
			openBrackets = append(openBrackets, string(char))
		} else {
			switch string(char) {
			case ")":
				if string(openBrackets[len(openBrackets)-1]) == "(" {
					openBrackets = openBrackets[:len(openBrackets)-1]
				} else {
					return "("
				}
			case "]":
				if string(openBrackets[len(openBrackets)-1]) == "[" {
					openBrackets = openBrackets[:len(openBrackets)-1]
				} else {
					return "["
				}
			case "}":
				if string(openBrackets[len(openBrackets)-1]) == "{" {
					openBrackets = openBrackets[:len(openBrackets)-1]
				} else {
					return "{"
				}
			case ">":
				if string(openBrackets[len(openBrackets)-1]) == "<" {
					openBrackets = openBrackets[:len(openBrackets)-1]
				} else {
					return "<"
				}
			}
		}
	}

	return ""
}

func calcSyntaxErrorScore(corruptedChar string) int {
	switch corruptedChar {
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
		corruptedChar := findCorruption(line)
		sumSyntaxErrorScore += calcSyntaxErrorScore(corruptedChar)
	}

	return sumSyntaxErrorScore
}

func evalB(lines []string) int {
	return 0
}

func main() {
	lines := util.ReadFile("input10.txt")

	resA := evalA(lines)
	resB := evalB(lines)
	fmt.Printf("A: %v \n", resA)
	fmt.Printf("B: %v \n", resB)
}
