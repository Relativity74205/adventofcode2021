package main

import (
	"aoc21_go/util"
	"fmt"
	"strings"
)

type Cave struct {
	bigCave       bool
	neighborCaves []string
}

type Path struct {
	cavesVisited          []string
	smallCaveVisitedTwice bool
}

func traverseCaveSystem(caveSystem map[string]Cave, openPaths []Path, endedPaths []Path, partB bool) ([]Path, []Path) {
	var newOpenPaths []Path
	for _, openPath := range openPaths {
		lastVisitedCave := openPath.cavesVisited[len(openPath.cavesVisited)-1]
		for _, cave := range caveSystem[lastVisitedCave].neighborCaves {
			if cave == "start" {
				continue
			} else if cave == "end" {
				endedPaths = append(endedPaths, Path{append(util.CopyStringSlice(openPath.cavesVisited), cave), openPath.smallCaveVisitedTwice})
			} else if caveSystem[cave].bigCave {
				newOpenPaths = append(newOpenPaths, Path{append(util.CopyStringSlice(openPath.cavesVisited), cave), openPath.smallCaveVisitedTwice})
			} else if !util.InSliceString(openPath.cavesVisited, cave) {
				newOpenPaths = append(newOpenPaths, Path{append(util.CopyStringSlice(openPath.cavesVisited), cave), openPath.smallCaveVisitedTwice})
			} else if util.InSliceString(openPath.cavesVisited, cave) && !openPath.smallCaveVisitedTwice && partB {
				newOpenPaths = append(newOpenPaths, Path{append(util.CopyStringSlice(openPath.cavesVisited), cave), true})
			}
		}
	}

	return newOpenPaths, endedPaths
}

func evalA(caveSystem map[string]Cave) int {
	var endedPaths []Path

	openPaths := []Path{{[]string{"start"}, false}}
	for len(openPaths) > 0 {
		openPaths, endedPaths = traverseCaveSystem(caveSystem, openPaths, endedPaths, false)
	}

	return len(endedPaths)
}

func evalB(caveSystem map[string]Cave) int {
	var endedPaths []Path

	openPaths := []Path{{[]string{"start"}, false}}
	for len(openPaths) > 0 {
		openPaths, endedPaths = traverseCaveSystem(caveSystem, openPaths, endedPaths, true)
	}

	return len(endedPaths)
}

func buildMap(lines []string) map[string]Cave {
	caveSystem := make(map[string]Cave)
	for _, line := range lines {
		path := strings.Split(line, "-")
		addPath(caveSystem, path[0], path[1])
		addPath(caveSystem, path[1], path[0])
	}

	return caveSystem
}

func addPath(caveSystem map[string]Cave, startCave, endCave string) {
	cave, knownCave := caveSystem[startCave]
	if !knownCave {
		cave = Cave{
			strings.ToUpper(startCave) == startCave,
			nil,
		}
	}
	cave.neighborCaves = append(cave.neighborCaves, endCave)
	caveSystem[startCave] = cave
}

func main() {
	day := 12
	debugSuffix := ""
	filename := fmt.Sprintf("input%02d%v.txt", day, debugSuffix)
	lines := util.ReadFile(filename)
	caveSystem := buildMap(lines)

	resA := evalA(caveSystem)
	resB := evalB(caveSystem)
	fmt.Printf("Day %02d \n", day)
	fmt.Printf("A: %v \n", resA)
	fmt.Printf("B: %v \n", resB)
}
