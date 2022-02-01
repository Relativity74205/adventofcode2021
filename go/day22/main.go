package main

import (
	"aoc21_go/util"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func countOnCubes(reactorMap [101][101][101]int) int {
	size := 50
	var count int
	for x := 0; x <= 2*size; x++ {
		for y := 0; y <= 2*size; y++ {
			for z := 0; z <= 2*size; z++ {
				count += reactorMap[z][y][x]
			}
		}
	}

	return count
}

func evalA(operations []Operation) int {
	size := 50
	var reactorMap [101][101][101]int
	for _, operation := range operations {
		var value int
		if operation.opType == "on" {
			value = 1
		} else {
			value = 0
		}

		for x := util.MaxInt(operation.x.min, -size); x <= util.MinInt(operation.x.max, size); x++ {
			for y := util.MaxInt(operation.y.min, -size); y <= util.MinInt(operation.y.max, size); y++ {
				for z := util.MaxInt(operation.z.min, -size); z <= util.MinInt(operation.z.max, size); z++ {
					reactorMap[z+50][y+50][x+50] = value
				}
			}
		}
	}

	return countOnCubes(reactorMap)
}

func intersectDimension(bounds1 Bounds, bounds2 Bounds) (Bounds, error) {
	var intersectBounds Bounds
	if util.MaxInt(bounds1.min, bounds2.min) > util.MinInt(bounds1.max, bounds2.max) {
		return intersectBounds, errors.New("NoIntersection")
	} else {
		return Bounds{util.MaxInt(bounds1.min, bounds2.min), util.MinInt(bounds1.max, bounds2.max)}, nil
	}
}

func intersectCuboid(oldCuboid Cuboid, newCuboid Cuboid, sign int) (Cuboid, error) {
	xBounds, xIntersect := intersectDimension(oldCuboid.x, newCuboid.x)
	yBounds, yIntersect := intersectDimension(oldCuboid.y, newCuboid.y)
	zBounds, zIntersect := intersectDimension(oldCuboid.z, newCuboid.z)
	if xIntersect == nil && yIntersect == nil && zIntersect == nil {
		return Cuboid{-1 * sign, xBounds, yBounds, zBounds}, nil
	} else {
		return Cuboid{}, errors.New("NoIntersection")
	}

}

func addOperation(newOperation Operation, cuboids []Cuboid) []Cuboid {
	var newCuboids []Cuboid
	var sign int
	if newOperation.opType == "on" {
		sign = 1
	} else {
		sign = -1
	}
	newCuboid := Cuboid{sign, newOperation.x, newOperation.y, newOperation.z}

	for _, cuboid := range cuboids {
		extraCuboid, hasIntersection := intersectCuboid(cuboid, newCuboid, sign)
		if hasIntersection == nil {
			newCuboids = append(newCuboids, extraCuboid)
		}
	}
	newCuboids = append(newCuboids, newCuboid)

	return newCuboids
}

func evalCuboids(cuboids []Cuboid) int {
	var count int
	for _, cuboid := range cuboids {
		count += (cuboid.x.max - cuboid.x.min + 1) * (cuboid.y.max - cuboid.y.min + 1) * (cuboid.z.max - cuboid.z.min + 1) * cuboid.sign
	}

	return count
}

func evalB(operations []Operation) int {
	var cuboids []Cuboid
	for _, operation := range operations {
		cuboids = append(cuboids, addOperation(operation, cuboids)...)
	}

	return evalCuboids(cuboids)
}

type Operation struct {
	opType  string
	x, y, z Bounds
}

type Cuboid struct {
	sign    int
	x, y, z Bounds
}

type Bounds struct {
	min, max int
}

func createBounds(s string) Bounds {
	parts := strings.Split(s[2:], "..")
	min, _ := strconv.Atoi(parts[0])
	max, _ := strconv.Atoi(parts[1])

	return Bounds{min, max}
}

func createOperations(lines []string) []Operation {
	var operations []Operation

	for _, line := range lines {
		mainParts := strings.Split(line, " ")
		boundsParts := strings.Split(mainParts[1], ",")
		x := createBounds(boundsParts[0])
		y := createBounds(boundsParts[1])
		z := createBounds(boundsParts[2])
		operations = append(operations, Operation{mainParts[0], x, y, z})
	}

	return operations
}

func eval(filename string, debug bool) {
	lines := util.ReadFile(filename)
	operations := createOperations(lines)

	resA := evalA(operations)
	resB := evalB(operations)
	if debug {
		fmt.Printf("A (%s): %v \n", filename, resA)
		fmt.Printf("B (%s): %v \n", filename, resB)
	} else {
		fmt.Printf("A: %v \n", resA)
		fmt.Printf("B: %v \n", resB)
	}

}

func main() {
	day := 22
	debugSuffix1 := "_debug1"
	debugSuffix2 := "_debug2"
	filename := fmt.Sprintf("input%02d.txt", day)
	filenameDebug1 := fmt.Sprintf("input%02d%v.txt", day, debugSuffix1)
	filenameDebug2 := fmt.Sprintf("input%02d%v.txt", day, debugSuffix2)

	fmt.Printf("Day %02d \n", day)
	eval(filenameDebug1, true)
	eval(filenameDebug2, true)
	eval(filename, false)
}
