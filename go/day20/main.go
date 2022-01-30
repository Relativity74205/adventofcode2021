package main

import (
	"aoc21_go/util"
	"fmt"
	"strconv"
)

var boarderPixel = [][]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {0, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}}

func createEmptyImage(sizeX, sizeY, defaultVal int) [][]int {
	image := make([][]int, sizeY)
	for y := range image {
		image[y] = make([]int, sizeX)
		for x := range image[y] {
			image[y][x] = defaultVal
		}
	}

	return image
}

func printImage(image [][]int) {
	for _, line := range image {
		for _, val := range line {
			if val == 1 {
				print("#")
			} else {
				print(".")
			}
		}
		println()
	}
}

func createImageFromInput(lines []string) [][]int {
	image := createEmptyImage(len(lines), len(lines[0]), 0)
	for y, line := range lines {
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

func padImage(image [][]int, padSize, defaultVal int) [][]int {
	paddedImage := createEmptyImage(len(image)+2*padSize, len(image[0])+2*padSize, defaultVal)
	for y, line := range image {
		for x, val := range line {
			paddedImage[y+padSize][x+padSize] = val
		}
	}

	return paddedImage
}

func cropImage(image [][]int, cropSize int) [][]int {
	croppedImage := createEmptyImage(len(image)-2*cropSize, len(image[0])-2*cropSize, 0)
	for y, line := range croppedImage {
		for x, _ := range line {
			croppedImage[y][x] = image[y+cropSize][x+cropSize]
		}
	}

	return croppedImage
}

func bitToInt(bitArray string) int {
	i, _ := strconv.ParseInt(bitArray, 2, 64)
	return int(i)
}

func enhancePixel(x, y int, image [][]int, algorithmMap map[int]int) int {
	var bitString string
	for _, foo := range boarderPixel {
		ytest := y + foo[1]
		xtest := x + foo[0]
		bitString += strconv.Itoa(image[ytest][xtest])
	}
	decimal := bitToInt(bitString)

	return algorithmMap[decimal]
}

func countPixel(image [][]int) int {
	var count int
	for _, line := range image {
		for _, val := range line {
			count += val
		}
	}

	return count
}

func enhanceImage(image [][]int, algorithmMap map[int]int, round int) [][]int {
	var defaultVal int
	if round%2 == 0 && algorithmMap[0] == 1 {
		defaultVal = 1
	} else {
		defaultVal = 0
	}

	padSize := 2
	inputImage := padImage(image, padSize, defaultVal)
	outputImage := createEmptyImage(len(inputImage), len(inputImage[0]), defaultVal)

	for y := 1; y < len(inputImage)-1; y++ {
		for x := 1; x < len(inputImage[0])-1; x++ {
			outputImage[y][x] = enhancePixel(x, y, inputImage, algorithmMap)
		}
	}

	return cropImage(outputImage, 1)
}

func evalA(image [][]int, algorithmMap map[int]int) int {
	var outputImage [][]int
	rounds := 2

	for round := 1; round <= rounds; round++ {
		outputImage = enhanceImage(image, algorithmMap, round)
		image = outputImage
	}

	return countPixel(outputImage)
}

func evalB(image [][]int, algorithmMap map[int]int) int {
	var outputImage [][]int
	rounds := 50

	for round := 1; round <= rounds; round++ {
		outputImage = enhanceImage(image, algorithmMap, round)
		image = outputImage
	}

	return countPixel(outputImage)
}

func eval(filename string, debug bool) {
	lines := util.ReadFile(filename)
	algorithmMap := createAlgorithmMap(lines[0])
	image := createImageFromInput(lines[2:])

	resA := evalA(image, algorithmMap)
	resB := evalB(image, algorithmMap)
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
