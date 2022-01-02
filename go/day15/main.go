package main

import (
	"aoc21_go/util"
	"fmt"
	"strconv"
)

var neighborDeltas = []Pos{
	{-1, 0},
	{0, 1},
	{0, -1},
	{1, 0},
}

type Pos struct {
	x int
	y int
}

type Map struct {
	riskLevels [][]int
	width      int
	height     int
}

func getNewPositions(caveMap [][]int, x, y int) []Pos {
	var newPositions []Pos
	for _, delta := range neighborDeltas {
		if util.CheckInBounds(caveMap, x+delta.x, y+delta.y) {
			newPositions = append(newPositions, Pos{x + delta.x, y + delta.y})
		}
	}

	return newPositions
}

func evalA(caveMap [][]int) int {
	posStart := Pos{0, 0}
	for _, newPosition := range getNewPositions(caveMap, posStart.x, posStart.y) {

	}
	pathsStart

	return 0
}

func evalB(lines []string) int {

	return 0
}

func buildMap(lines []string) [][]int {
	caveMap := make([][]int, len(lines))
	for y, line := range lines {
		var row []int
		for _, char := range line {
			riskLevel, _ := strconv.Atoi(string(char))
			row = append(row, riskLevel)
		}
		caveMap[y] = row
	}

	return caveMap
}

func main() {
	day := 15
	debugSuffix := ""
	filename := fmt.Sprintf("input%02d%v.txt", day, debugSuffix)
	lines := util.ReadFile(filename)
	caveMap := buildMap(lines)

	resA := evalA(caveMap)
	resB := evalB(lines)
	fmt.Printf("Day %02d \n", day)
	fmt.Printf("A: %v \n", resA)
	fmt.Printf("B: %v \n", resB)
}
