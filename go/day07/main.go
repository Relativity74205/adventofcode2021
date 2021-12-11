package main

import (
	"aoc21_go/util"
	"fmt"
	"strconv"
	"strings"
)

func evalA(crabPositions []int) int {
	min := util.MinIntegers(crabPositions)
	max := util.MaxIntegers(crabPositions)
	minFuel := 9223372036854775807
	for alignPos := min; alignPos <= max; alignPos++ {
		var fuel int
		for _, crabPosition := range crabPositions {
			fuel += util.AbsInt(crabPosition - alignPos)
		}
		minFuel = util.MinInt(fuel, minFuel)
	}

	return minFuel
}

func calcFuelConsumption(distance int) int {
	var fuelConsumption int
	for i := 1; i <= distance; i++ {
		fuelConsumption += i
	}
	return fuelConsumption
}

func evalB(crabPositions []int) int {
	min := util.MinIntegers(crabPositions)
	max := util.MaxIntegers(crabPositions)
	minFuel := 9223372036854775807
	for alignPos := min; alignPos <= max; alignPos++ {
		var fuel int
		for _, crabPosition := range crabPositions {
			fuel += calcFuelConsumption(util.AbsInt(crabPosition - alignPos))
		}
		minFuel = util.MinInt(fuel, minFuel)
	}

	return minFuel
}

func createCrabPositions(line string) []int {
	var crabPositions []int
	posStrings := strings.Split(line, ",")
	for _, posString := range posStrings {
		pos, _ := strconv.Atoi(posString)
		crabPositions = append(crabPositions, pos)
	}
	return crabPositions
}

func main() {
	lines := util.ReadFile("input07.txt")
	crabPositions := createCrabPositions(lines[0])

	resA := evalA(crabPositions)
	resB := evalB(crabPositions)
	fmt.Printf("A: %v \n", resA)
	fmt.Printf("B: %v \n", resB)
}
