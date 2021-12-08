package main

import (
	"aoc21_go/util"
	"fmt"
	"strconv"
	"strings"
)

type Course struct {
	direction string
	value     int
}

func linesToDirections(lines []string) []Course {
	var courses []Course
	for _, v := range lines {
		fields := strings.Fields(v)
		value, _ := strconv.Atoi(fields[1])
		courses = append(courses, Course{fields[0], value})
	}

	return courses
}

func evalA(courses []Course) int {
	var horizontal = 0
	var depth = 0
	for _, course := range courses {
		switch course.direction {
		case "forward":
			horizontal += course.value
		case "up":
			depth -= course.value
		case "down":
			depth += course.value
		}
	}

	return horizontal * depth
}

func evalB(courses []Course) int {
	var horizontal = 0
	var depth = 0
	var aim = 0
	for _, course := range courses {
		switch course.direction {
		case "forward":
			horizontal += course.value
			depth += aim * course.value
		case "up":
			aim -= course.value
		case "down":
			aim += course.value
		}
	}

	return horizontal * depth
}

func main() {
	lines := util.ReadFile("input02.txt")
	courses := linesToDirections(lines)
	resA := evalA(courses)
	resB := evalB(courses)
	fmt.Printf("A: %v \n", resA)
	fmt.Printf("B: %v \n", resB)
}
