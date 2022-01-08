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

type PriorityQueue []*Path

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest based on expiration number as the priority
	// The lower the expiry, the higher the priority
	return pq[i].riskLevel < pq[j].riskLevel
}

// We just implement the pre-defined function in interface of heap.

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Path)
	*pq = append(*pq, item)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
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

type Path struct {
	visitedPos []Pos
	riskLevel  int
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
	cumRisk := make([][]int, len(caveMap))
	for y := range cumRisk {
		cumRisk[y] = make([]int, len(caveMap[0]))
	}

	startPath := Path{[]Pos{{0, 0}}, 0}
	fooQueue := []Path{startPath}
	for len(fooQueue) > 0 {

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
