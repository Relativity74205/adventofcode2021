package main

import (
	"aoc21_go/util"
	"fmt"
	"sort"
	"strconv"
)

type Basin struct {
	x         int
	y         int
	riskLevel int
	size      int
}

type Coords struct {
	x int
	y int
}

func checkInBounds(heightMap [][]int, x, y int) bool {
	maxHeight := len(heightMap) - 1
	maxWidth := len(heightMap[0]) - 1

	if x < 0 || y < 0 || x > maxWidth || y > maxHeight {
		return false
	}

	return true
}

func checkLowPoint(heightMap [][]int, x, y int) bool {
	if checkInBounds(heightMap, x-1, y) && heightMap[y][x-1] <= heightMap[y][x] {
		return false
	}
	if checkInBounds(heightMap, x+1, y) && heightMap[y][x+1] <= heightMap[y][x] {
		return false
	}
	if checkInBounds(heightMap, x, y-1) && heightMap[y-1][x] <= heightMap[y][x] {
		return false
	}
	if checkInBounds(heightMap, x, y+1) && heightMap[y+1][x] <= heightMap[y][x] {
		return false
	}
	return true
}

func checkPartOfBasin(heightMap [][]int, x, y int) bool {
	if checkInBounds(heightMap, x, y) && heightMap[y][x] < 9 {
		return true
	}
	return false
}

func checkExistInBasinCoords(coord Coords, basinCoords map[Coords]bool) bool {
	_, ok := basinCoords[coord]
	return ok
}

func getBasinCandidates(x, y int, basinCoords map[Coords]bool) []Coords {
	var newBasinCandidates []Coords
	for _, coord := range []Coords{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}} {
		if !checkExistInBasinCoords(coord, basinCoords) {
			newBasinCandidates = append(newBasinCandidates, coord)
		}

	}
	return newBasinCandidates
}

func calcBasinSize(heightMap [][]int, basinX, basinY int) int {
	var basinCandidates, newCandidates []Coords
	basinCoords := map[Coords]bool{Coords{basinX, basinY}: true}

	newCandidates = getBasinCandidates(basinX, basinY, basinCoords)

	for len(newCandidates) > 0 {
		basinCandidates = newCandidates
		newCandidates = nil
		for _, basinCandidate := range basinCandidates {
			x := basinCandidate.x
			y := basinCandidate.y
			if checkPartOfBasin(heightMap, x, y) {
				basinCoords[basinCandidate] = true
				newCandidates = append(newCandidates, getBasinCandidates(x, y, basinCoords)...)
			}
		}
	}

	return len(basinCoords)
}

func getBasins(heightMap [][]int) []Basin {
	var basins []Basin
	for y, row := range heightMap {
		for x, value := range row {
			if checkLowPoint(heightMap, x, y) {
				basinSize := calcBasinSize(heightMap, x, y)
				basins = append(basins, Basin{x, y, value + 1, basinSize})
			}
		}
	}

	return basins
}

func evalA(heightMap [][]int) int {
	var sumRiskLevel int
	basins := getBasins(heightMap)

	for _, basin := range basins {
		sumRiskLevel += basin.riskLevel
	}

	return sumRiskLevel
}

func evalB(heightMap [][]int) int {
	basinSizeFactor := 1
	basins := getBasins(heightMap)

	sort.Slice(basins, func(i, j int) bool { return basins[i].size > basins[j].size })

	for _, basin := range basins[0:3] {
		basinSizeFactor *= basin.size
	}
	return basinSizeFactor
}

func createHeightMap(lines []string) [][]int {
	var heightMap [][]int
	for _, line := range lines {
		var heightLine []int
		for _, char := range line {
			height, _ := strconv.Atoi(string(char))
			heightLine = append(heightLine, height)
		}
		heightMap = append(heightMap, heightLine)
	}

	return heightMap
}

func main() {
	lines := util.ReadFile("input09.txt")
	heightMap := createHeightMap(lines)

	resA := evalA(heightMap)
	resB := evalB(heightMap)
	fmt.Printf("A: %v \n", resA)
	fmt.Printf("B: %v \n", resB)
}
