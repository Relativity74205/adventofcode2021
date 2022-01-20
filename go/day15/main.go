package main

import (
	"aoc21_go/util"
	"container/heap"
	"fmt"
	"strconv"
)

var neighborDeltas = []Pos{
	{-1, 0},
	{0, 1},
	{0, -1},
	{1, 0},
}

type PriorityQueue []Path

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].riskLevel < pq[j].riskLevel
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Path))
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

type Pos struct {
	x int
	y int
}

type Path struct {
	lastPos   Pos
	path      []Pos
	riskLevel int
}

func getNewPaths(caveMap [][]int, path Path) []Path {
	var newPaths []Path

	for _, delta := range neighborDeltas {
		newPos := Pos{path.lastPos.x + delta.x, path.lastPos.y + delta.y}

		if util.CheckInBounds(caveMap, newPos.x, newPos.y) {
			newPath := Path{
				newPos,
				append(path.path, newPos),
				path.riskLevel + caveMap[newPos.y][newPos.x],
			}
			newPaths = append(newPaths, newPath)
		}
	}

	return newPaths
}

func findPath(caveMap [][]int) int {
	endX := len(caveMap) - 1
	endY := len(caveMap[0]) - 1
	startPos := Pos{0, 0}

	h := &PriorityQueue{}
	heap.Init(h)
	heap.Push(h, Path{startPos, []Pos{startPos}, 0})
	visitedPos := make(map[Pos]bool)
	toVisitPos := make(map[Pos]int)
	toVisitPos[startPos] = 0

	for h.Len() > 0 {
		path := heap.Pop(h).(Path)
		if path.lastPos.x == endX && path.lastPos.y == endY {
			return path.riskLevel
		}

		visitedPos[path.lastPos] = true

		newPaths := getNewPaths(caveMap, path)
		for _, newPath := range newPaths {
			_, alreadyVisited := visitedPos[newPath.lastPos]
			oldRiskLevel, onVisitPos := toVisitPos[newPath.lastPos]
			if alreadyVisited {
				continue
			}
			if onVisitPos && newPath.riskLevel >= oldRiskLevel {
				continue
			}
			toVisitPos[newPath.lastPos] = newPath.riskLevel
			heap.Push(h, newPath)
		}
	}

	return -1
}

func evalA(caveMap [][]int) int {
	return findPath(caveMap)
}

func evalB(caveMap [][]int) int {
	unfoldedCaveMap := unfoldMap(caveMap)
	return findPath(unfoldedCaveMap)
}

func unfoldMap(caveMap [][]int) [][]int {
	orgSizeX := len(caveMap)
	orgSizeY := len(caveMap[0])
	unfoldedMap := make([][]int, orgSizeY*5)
	for y := 0; y < orgSizeY*5; y++ {
		unfoldedMap[y] = make([]int, orgSizeX*5)
	}
	foldFactor := 5

	for x := 0; x < orgSizeX; x++ {
		for y := 0; y < orgSizeY; y++ {
			for dx := 0; dx < foldFactor; dx++ {
				for dy := 0; dy < foldFactor; dy++ {
					newRiskLevel := caveMap[y][x] + dx + dy
					if newRiskLevel > 9 {
						newRiskLevel -= 9
					}
					unfoldedMap[y+dy*orgSizeY][x+dx*orgSizeX] = newRiskLevel
				}
			}
		}
	}

	return unfoldedMap
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

func eval(filename string, debug bool) {
	lines := util.ReadFile(filename)
	caveMap := buildMap(lines)

	resA := evalA(caveMap)
	resB := evalB(caveMap)
	if debug {
		fmt.Printf("A (debug): %v \n", resA)
		fmt.Printf("B (debug): %v \n", resB)
	} else {
		fmt.Printf("A: %v \n", resA)
		fmt.Printf("B: %v \n", resB)
	}

}

func main() {
	day := 15
	debugSuffix := "_debug"
	filename := fmt.Sprintf("input%02d.txt", day)
	filenameDebug := fmt.Sprintf("input%02d%v.txt", day, debugSuffix)

	fmt.Printf("Day %02d \n", day)
	eval(filenameDebug, true)
	eval(filename, false)
}
