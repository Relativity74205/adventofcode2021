package main

import (
	"aoc21_go/util"
	"fmt"
	"strconv"
	"strings"
)

type targetArea struct {
	xMin int
	xMax int
	yMin int
	yMax int
}

type Probe struct {
	x    int
	y    int
	vx   int
	vy   int
	yMax int
	hit  bool
}

func inTargetArea(probe Probe, area targetArea) bool {
	if probe.x >= area.xMin && probe.x <= area.xMax && probe.y >= area.yMin && probe.y <= area.yMax {
		return true
	} else {
		return false
	}
}

func getStartVelocity(l int) int {
	var dx int
	i := 1
	for dx < l {
		dx += i
		i++
	}

	if dx == l {
		return i - 1
	} else {
		return -1
	}
}

func evalA(area targetArea) int {

	probe := Probe{0, 0, 6, 9, 0, false}
	for probe.y >= area.yMin {
		probe.x += probe.vx
		probe.y += probe.vy
		probe.vx -= 1
		if probe.vx < 0 {
			probe.vx = 0
		}
		probe.vy -= 1
		probe.yMax = util.MaxInt(probe.yMax, probe.y)
		if inTargetArea(probe, area) {
			probe.hit = true
		}
	}
	if probe.hit {
		return probe.yMax
	} else {
		return -1
	}

}

func evalB(area targetArea) int {

	return 0
}

func getTargetArea(line string) targetArea {
	parts := strings.Split(line[12:], ",")
	xParts := strings.Split(parts[0][3:], "..")
	yParts := strings.Split(parts[1][3:], "..")
	xMin, _ := strconv.Atoi(xParts[0])
	xMax, _ := strconv.Atoi(xParts[1])
	yMin, _ := strconv.Atoi(yParts[0])
	yMax, _ := strconv.Atoi(yParts[1])
	return targetArea{xMin, xMax, yMin, yMax}
}

func eval(filename string, debug bool) {
	lines := util.ReadFile(filename)
	targetArea := getTargetArea(lines[0])

	resA := evalA(targetArea)
	resB := evalB(targetArea)
	if debug {
		fmt.Printf("A (debug): %v \n", resA)
		fmt.Printf("B (debug): %v \n", resB)
	} else {
		fmt.Printf("A: %v \n", resA)
		fmt.Printf("B: %v \n", resB)
	}

}

func main() {
	day := 17
	debugSuffix := "_debug"
	filename := fmt.Sprintf("input%02d.txt", day)
	filenameDebug := fmt.Sprintf("input%02d%v.txt", day, debugSuffix)

	fmt.Printf("Day %02d \n", day)
	eval(filenameDebug, true)
	eval(filename, false)
}
