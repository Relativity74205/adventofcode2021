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
	cavesVisited []string
	valid        bool
	end          bool
}

func foo(caveSystem map[string]Cave, paths []Path) []Path {
	var newPaths []Path
	for _, path := range paths {
		if !path.end {
			lastVisitedCave := path.cavesVisited[len(path.cavesVisited)-1]
			for _, cave := range caveSystem[lastVisitedCave].neighborCaves {
				newCavesVisited := append(path.cavesVisited, cave)
				if cave
					newValid := true
				newEnd := true
				newPath := Path{newCavesVisited, newValid, newEnd}
				newPaths = append(newPaths, newPath)
			}
		}
	}
}

func evalA(caveSystem map[string]Cave) int {
	paths := []Path{Path{[]string{"start"}, true, false}}
	foo(caveSystem, paths)

	return 0
}

func evalB(caveSystem map[string]Cave) int {

	return 0
}

func buildMap(lines []string) map[string]Cave {
	caveSystem := make(map[string]Cave)
	for _, line := range lines {
		path := strings.Split(line, "-")
		cave, knownCave := caveSystem[path[0]]
		if !knownCave {
			cave = Cave{
				strings.ToUpper(path[0]) == path[0],
				nil,
			}
		}
		cave.neighborCaves = append(cave.neighborCaves, path[1])
		caveSystem[path[0]] = cave
	}

	return caveSystem
}

func main() {
	day := 12
	debugSuffix := "_debug2"
	filename := fmt.Sprintf("input%02d%v.txt", day, debugSuffix)
	lines := util.ReadFile(filename)
	caveSystem := buildMap(lines)

	resA := evalA(caveSystem)
	resB := evalB(caveSystem)
	fmt.Printf("Day %02d \n", day)
	fmt.Printf("A: %v \n", resA)
	fmt.Printf("B: %v \n", resB)
}
