package main

import (
	"aoc21_go/util"
	"fmt"
	"strconv"
)

var rollDirac = map[int]int{3: 1, 4: 3, 5: 6, 6: 7, 7: 6, 8: 3, 9: 1}

func rollDeterministic() func() int {
	diceValue := 0
	return func() int {
		diceValue += 1
		return diceValue
	}
}

func nextPos(currentPos, diceRoll int) int {
	result := (currentPos + diceRoll) % 10
	if result == 0 {
		return 10
	} else {
		return result
	}
}

func evalA(startingPositions []int) int {
	var scoreA, scoreB int
	dice := rollDeterministic()
	currentPosA := startingPositions[0]
	currentPosB := startingPositions[1]
	for {
		currentPosA = nextPos(currentPosA, dice()+dice()+dice())
		scoreA += currentPosA
		if scoreA >= 1000 {
			break
		}
		currentPosB = nextPos(currentPosB, dice()+dice()+dice())
		scoreB += currentPosB
		if scoreB >= 1000 {
			break
		}
	}

	return util.MinInt(scoreA, scoreB) * (dice() - 1)
}

type State struct {
	scoreA int
	posA   int
	scoreB int
	posB   int
	turnA  bool
}

func findLowestState(states map[State]int) (State, int) {
	var lowestState State
	minScore := 22
	for state, count := range states {
		if state.turnA && state.scoreA < minScore {
			lowestState = state
			minScore = count
		} else if !state.turnA && state.scoreB < minScore {
			lowestState = state
			minScore = count
		}
	}

	return lowestState, states[lowestState]
}

func evalB(startingPositions []int) int {
	var posA, scoreA, posB, scoreB int
	var winnerA, winnerB int
	states := make(map[State]int)
	states[State{0, startingPositions[0], 0, startingPositions[1], true}] += 1
	for len(states) > 0 {
		state, stateCount := findLowestState(states)
		delete(states, state)
		for diceRoll, diceRollCount := range rollDirac {
			if state.turnA {
				posA = nextPos(state.posA, diceRoll)
				scoreA = state.scoreA + posA
				posB = state.posB
				scoreB = state.scoreB
				if scoreA >= 21 {
					winnerA += diceRollCount * stateCount
					continue
				}
			} else {
				posA = state.posA
				scoreA = state.scoreA
				posB = nextPos(state.posB, diceRoll)
				scoreB = state.scoreB + posB
				if scoreB >= 21 {
					winnerB += diceRollCount * stateCount
					continue
				}
			}
			newState := State{scoreA, posA, scoreB, posB, !state.turnA}
			states[newState] *= diceRollCount * stateCount
		}
	}

	return util.MaxInt(winnerA, winnerB)
}

func getStartingPositions(lines []string) []int {
	playerA, _ := strconv.Atoi(string(lines[0][len(lines[0])-1]))
	playerB, _ := strconv.Atoi(string(lines[1][len(lines[1])-1]))

	return []int{playerA, playerB}
}

func eval(filename string, debug bool) {
	lines := util.ReadFile(filename)
	startingPositions := getStartingPositions(lines)

	resA := evalA(startingPositions)
	resB := evalB(startingPositions)
	if debug {
		fmt.Printf("A (debug): %v \n", resA)
		fmt.Printf("B (debug): %v \n", resB)
	} else {
		fmt.Printf("A: %v \n", resA)
		fmt.Printf("B: %v \n", resB)
	}

}

func main() {
	day := 21
	debugSuffix := "_debug"
	filename := fmt.Sprintf("input%02d.txt", day)
	filenameDebug := fmt.Sprintf("input%02d%v.txt", day, debugSuffix)

	fmt.Printf("Day %02d \n", day)
	eval(filenameDebug, true)
	eval(filename, false)
}
