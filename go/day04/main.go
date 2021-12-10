package main

import (
	"aoc21_go/util"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

type Board struct {
	numbers map[int]Position
	tiles   [5][5]int
}

func allOnes(slice []int) bool {
	for _, number := range slice {
		if number == 0 {
			return false
		}
	}

	return true
}

func checkBoard(board Board) bool {
	for x := 0; x <= 4; x++ {
		if allOnes(board.tiles[x][:]) {
			return true
		}
		if allOnes(board.tiles[:][x][:]) {
			return true
		}
	}
	return false
}

func evalA(boards []Board, numbers []int) int {
	for _, number := range numbers {
		for _, board := range boards {
			pos, ok := board.numbers[number]

			if ok {
				board.tiles[pos.x][pos.y] = 1
			}
		}
	}

	for _, board := range boards {
		if checkBoard(board) {
			return 1
		}
	}

	return 0
}

func evalB(boards []Board, numbers []int) int {
	return 0
}

func makeBoard(lines []string) *Board {
	var numbers []int
	for _, line := range lines {
		numbersString := strings.Fields(line)
		for _, numberString := range numbersString {
			number, _ := strconv.Atoi(numberString)
			numbers = append(numbers, number)
		}
	}

	var tiles [5][5]int
	numbersMap := make(map[int]Position)
	for i, number := range numbers {
		x := i - int(i/5.0)*5
		y := i / 5.0
		pos := Position{x, y}
		numbersMap[number] = pos

	}
	return &Board{numbersMap, tiles}
}

func readBoards(lines []string) []*Board {
	var boards []*Board

	var boardLines []string
	for _, line := range lines {
		if line == "" {
			boards = append(boards, makeBoard(boardLines))
			boardLines = nil
			continue
		}

		boardLines = append(boardLines, line)
	}

	return boards
}

func readNumbers(line string) []int {
	numbersString := strings.Split(line, ",")
	var numbers []int
	for _, numberString := range numbersString {
		number, _ := strconv.Atoi(numberString)
		numbers = append(numbers, number)
	}

	return numbers
}

func main() {
	lines := util.ReadFile("input04.txt")
	//numbers := readNumbers(lines[0])
	boards := readBoards(lines[2:])
	for _, board := range boards {
		board.tiles[0][0] = 2
	}

	println(boards[0].tiles[0][0])
	//resA := evalA(boards, numbers)
	//resB := evalB(boards, numbers)
	//fmt.Printf("A: %v \n", resA)
	//fmt.Printf("B: %v \n", resB)
}
