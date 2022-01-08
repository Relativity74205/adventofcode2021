package main

import (
	"aoc21_go/util"
	"fmt"
	"strconv"
	"strings"
)

func boolToInt(val bool) int {
	if val {
		return 1
	} else {
		return 0
	}
}

func fold(paper [][]bool, instruction Instruction) [][]bool {
	if !instruction.xAxis {
		newPaper := paper[:instruction.lineNumber]
		sizeY := len(paper)
		for y, line := range newPaper {
			for x := range line {
				newPaper[y][x] = newPaper[y][x] || paper[sizeY-1-y][x]
			}
		}

		return newPaper
	} else {
		sizeX := len(paper[0])
		newPaper := make([][]bool, len(paper))
		for y, line := range paper {
			newPaper[y] = line[:instruction.lineNumber]
		}
		for y, line := range newPaper {
			for x := range line {
				newPaper[y][x] = newPaper[y][x] || paper[y][sizeX-1-x]
			}
		}

		return newPaper
	}
}

func countDots(paper [][]bool) int {
	var count int
	for _, line := range paper {
		for _, value := range line {
			count += boolToInt(value)
		}
	}

	return count
}

func evalA(paper [][]bool, instruction Instruction) int {
	foldedPaper := fold(paper, instruction)
	count := countDots(foldedPaper)

	return count
}

func evalB(paper [][]bool, instructions []Instruction) [][]bool {
	for _, instruction := range instructions {
		paper = fold(paper, instruction)
	}

	return paper
}

func getCoords(line string) (int, int) {
	coords := strings.Split(line, ",")
	x, _ := strconv.Atoi(coords[0])
	y, _ := strconv.Atoi(coords[1])

	return x, y
}

func getSize(instructions []string) (int, int) {
	var maxX, maxY int
	for _, instruction := range instructions {
		x, y := getCoords(instruction)
		maxX = util.MaxInt(maxX, x)
		maxY = util.MaxInt(maxY, y)
	}

	return maxX, maxY
}

func createPaper(lines []string) [][]bool {
	blankLineIndex := util.GetBlankLine(lines)
	coordStrings := lines[:blankLineIndex]

	sizeX, sizeY := getSize(coordStrings)

	paper := make([][]bool, sizeY+1)
	for i := range paper {
		paper[i] = make([]bool, sizeX+1)
	}

	for _, coordString := range coordStrings {
		x, y := getCoords(coordString)
		paper[y][x] = true
	}

	return paper
}

type Instruction struct {
	xAxis      bool
	lineNumber int
}

func createInstructions(lines []string) []Instruction {
	var instructions []Instruction
	blankLineIndex := util.GetBlankLine(lines)
	instructionStrings := lines[blankLineIndex+1:]
	for _, instructionString := range instructionStrings {
		instructionParts := strings.Split(instructionString, "=")
		axis := string(instructionParts[0][len(instructionParts[0])-1])
		lineNumber, _ := strconv.Atoi(instructionParts[1])
		instructions = append(instructions, Instruction{axis == "x", lineNumber})
	}

	return instructions
}

func printPaper(paper [][]bool) {
	for _, line := range paper {
		for x, val := range line {
			if (x+1)%5 == 0 {
				print(" ")
			} else {
				print(boolToInt(val))
			}
		}
		println()
	}

}

func main() {
	day := 13
	debugSuffix := ""
	filename := fmt.Sprintf("input%02d%v.txt", day, debugSuffix)
	lines := util.ReadFile(filename)
	paper := createPaper(lines)
	instructions := createInstructions(lines)

	resA := evalA(paper, instructions[0])
	resB := evalB(paper, instructions)
	fmt.Printf("Day %02d \n", day)
	fmt.Printf("A: %v \n", resA)

	println("Code for part B is:")
	printPaper(resB) // code is HEJHJRCJ
}
