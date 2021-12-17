package main

import (
	"aoc21_go/util"
	"fmt"
)

type Pos struct {
	x int
	y int
}

type Board struct {
	energy [][]int
	width  int
	height int
}

var neighborDeltas = []Pos{
	{-1, 0},
	{-1, -1},
	{-1, 1},
	{0, 1},
	{0, -1},
	{1, 1},
	{1, 0},
	{1, -1},
}

func checkOutOfBounds(pos Pos, width, height int) bool {
	if pos.x < 0 || pos.y < 0 || pos.x > width-1 || pos.y > height-1 {
		return true
	}

	return false
}

func getNeighbors(pos Pos, flashList map[Pos]bool, width, height int) []Pos {
	var neighbors []Pos
	for _, neighbor := range neighborDeltas {
		candidatePos := Pos{pos.x + neighbor.x, pos.y + neighbor.y}
		_, inFlashList := flashList[candidatePos]
		if !inFlashList && !checkOutOfBounds(candidatePos, width, height) {
			neighbors = append(neighbors, candidatePos)
		}
	}

	return neighbors
}

func processRound(board Board) (Board, int) {
	var increaseList []Pos
	flashList := make(map[Pos]bool)
	for x := 0; x <= board.width-1; x++ {
		for y := 0; y <= board.height-1; y++ {
			increaseList = append(increaseList, Pos{x, y})
		}
	}

	newBoard := board
	for len(increaseList) > 0 {
		pos := increaseList[0]
		increaseList = increaseList[1:]

		if newBoard.energy[pos.x][pos.y]+1 > 9 {
			flashList[pos] = true
			increaseList = append(increaseList, getNeighbors(pos, flashList, newBoard.width, newBoard.height)...)
			newBoard.energy[pos.x][pos.y] = 0
		} else {
			newBoard.energy[pos.x][pos.y] += 1
		}
	}

	return newBoard, len(flashList)
}

func evalA(board Board) int {
	var countFlashes int
	maxRounds := 100

	for i := 0; i < maxRounds; i++ {
		newBoard, newFlashes := processRound(board)
		countFlashes += newFlashes
		board = newBoard
	}

	return countFlashes
}

func evalB(board Board) int {

	return 0
}

func createBoard(startEnergy [][]int) Board {
	return Board{startEnergy, len(startEnergy[0]), len(startEnergy)}
}

func main() {
	lines := util.ReadFile("input11_debug.txt")
	startEnergy := util.GetMatrixFromLines(lines)
	energyBoard := createBoard(startEnergy)

	resA := evalA(energyBoard)
	resB := evalB(energyBoard)
	fmt.Printf("A: %v \n", resA)
	fmt.Printf("B: %v \n", resB)
}
