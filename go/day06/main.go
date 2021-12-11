package main

import (
	"aoc21_go/util"
	"fmt"
	"strconv"
	"strings"
)

func nextDay(cohorts map[int]int) map[int]int {
	nextCohorts := make(map[int]int)

	for timeToGenerate, cohortCount := range cohorts {
		if timeToGenerate == 0 {
			nextCohorts[6] += cohortCount
			nextCohorts[8] += cohortCount
		} else {
			nextCohorts[timeToGenerate-1] += cohortCount
		}
	}

	return nextCohorts
}

func cohortCount(cohorts map[int]int) int {
	var totalCount int
	for _, cohortCount := range cohorts {
		totalCount += cohortCount
	}

	return totalCount
}

func evalA(cohorts map[int]int) int {
	for day := 1; day <= 80; day++ {
		cohorts = nextDay(cohorts)
		fmt.Printf("day: %v pop: %v\n", day, cohortCount(cohorts))
	}

	return cohortCount(cohorts)
}

func evalB(cohorts map[int]int) int {
	for day := 1; day <= 256; day++ {
		cohorts = nextDay(cohorts)
	}

	return cohortCount(cohorts)
}

func createCohorts(line string) map[int]int {
	cohorts := make(map[int]int)
	ageStrings := strings.Split(line, ",")
	for _, ageString := range ageStrings {
		age, _ := strconv.Atoi(ageString)
		cohorts[age] += 1
	}
	return cohorts
}

func main() {
	lines := util.ReadFile("input06.txt")
	cohorts := createCohorts(lines[0])

	resA := evalA(cohorts)
	resB := evalB(cohorts)
	fmt.Printf("A: %v \n", resA)
	fmt.Printf("B: %v \n", resB)
}
