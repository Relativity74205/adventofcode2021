package main

import (
	"aoc21_go/util"
	"fmt"
	"strconv"
	"strings"
)

type pos struct {
	x int
	y int
}

type pipe struct {
	start pos
	end   pos
}

func createBoard(pipes []pipe) [][]int {
	var size int
	for _, pipe := range pipes {
		integers := []int{size, pipe.start.x, pipe.start.y, pipe.end.x, pipe.end.y}
		size = util.MaxIntegers(integers)
	}

	board := make([][]int, size+1)
	for i := range board {
		board[i] = make([]int, size+1)
	}

	return board
}

func evalBoard(board [][]int) int {
	var count int
	for i := range board {
		for j := range board[i] {
			if board[i][j] >= 2 {
				count += 1
			}
		}
	}
	return count
}

func evalA(pipes []pipe) int {
	board := createBoard(pipes)
	for _, pipe := range pipes {
		if pipe.start.x != pipe.end.x && pipe.start.y != pipe.end.y {
			continue
		}

		if pipe.start.x != pipe.end.x {
			start := util.MinInt(pipe.start.x, pipe.end.x)
			end := util.MaxInt(pipe.start.x, pipe.end.x)
			for i := start; i <= end; i++ {
				board[i][pipe.start.y] += 1
			}
		}

		if pipe.start.y != pipe.end.y {
			start := util.MinInt(pipe.start.y, pipe.end.y)
			end := util.MaxInt(pipe.start.y, pipe.end.y)
			for i := start; i <= end; i++ {
				board[pipe.start.x][i] += 1
			}
		}

	}

	return evalBoard(board)
}

func evalB(pipes []pipe) int {
	board := createBoard(pipes)
	for _, pipe := range pipes {
		var dx, dy int

		if pipe.start.x > pipe.end.x {
			dx = -1
		}
		if pipe.start.x < pipe.end.x {
			dx = 1
		}
		if pipe.start.y > pipe.end.y {
			dy = -1
		}
		if pipe.start.y < pipe.end.y {
			dy = 1
		}
		steps := util.MaxInt(util.AbsInt(pipe.end.x-pipe.start.x), util.AbsInt(pipe.end.y-pipe.start.y))
		for i := 0; i <= steps; i++ {
			x := pipe.start.x + dx*i
			y := pipe.start.y + dy*i
			board[x][y] += 1
		}

	}

	return evalBoard(board)
}

func createPos(posString string) pos {
	coords := strings.Split(posString, ",")
	x, _ := strconv.Atoi(coords[0])
	y, _ := strconv.Atoi(coords[1])
	return pos{x, y}
}

func createPipes(lines []string) []pipe {
	var pipes []pipe
	for _, line := range lines {
		positions := strings.Split(line, " -> ")
		start := createPos(positions[0])
		end := createPos(positions[1])
		pipes = append(pipes, pipe{start, end})
	}

	return pipes
}

func main() {
	lines := util.ReadFile("input05.txt")
	pipes := createPipes(lines)

	resA := evalA(pipes)
	resB := evalB(pipes)
	fmt.Printf("A: %v \n", resA)
	fmt.Printf("B: %v \n", resB)
}
