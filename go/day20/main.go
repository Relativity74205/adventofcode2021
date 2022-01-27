package main

import (
	"aoc21_go/util"
	"fmt"
)

func createImage(lines []string) [][]int {
	image := make([][]int, len(lines))
	for y, line := range lines {
		image[y] = make([]int, len(lines[0]))
		for x, char := range line {
			if char == '#' {
				image[y][x] = 1
			} else {
				image[y][x] = 0
			}
		}
	}

	return image
}

func createAlgorithmMap(s string) map[int]int {
	algorithm := make(map[int]int)
	for i, char := range s {
		if char == '#' {
			algorithm[i] = 1
		} else {
			algorithm[i] = 0
		}
	}
	return algorithm
}

func evalA(image [][]int, algorithmMap map[int]int) int {
	inputImage := padImage(image, 3)

	return 0
}

func evalB(lines []string) int {

	return 0
}

func eval(filename string, debug bool) {
	lines := util.ReadFile(filename)
	algorithmMap := createAlgorithmMap(lines[0])
	image := createImage(lines[2:])

	resA := evalA(image, algorithmMap)
	resB := evalB(lines)
	if debug {
		fmt.Printf("A (debug): %v \n", resA)
		fmt.Printf("B (debug): %v \n", resB)
	} else {
		fmt.Printf("A: %v \n", resA)
		fmt.Printf("B: %v \n", resB)
	}

}

func main() {
	day := 20
	debugSuffix := "_debug"
	filename := fmt.Sprintf("input%02d.txt", day)
	filenameDebug := fmt.Sprintf("input%02d%v.txt", day, debugSuffix)

	fmt.Printf("Day %02d \n", day)
	eval(filenameDebug, true)
	eval(filename, false)
}
