package main

import (
	"aoc21_go/util"
	"fmt"
	"strings"
)

type Instruction struct {
	pair   string
	insert string
}

func evalPolymer(polymer string) int {
	counter := make(map[string]int)

	for _, char := range polymer {
		counter[string(char)] += 1
	}

	countMinChar := 9223372036854775807
	countMaxChar := 0
	for char := range counter {
		countMinChar = util.MinInt(countMinChar, counter[char])
		countMaxChar = util.MaxInt(countMaxChar, counter[char])
	}

	return countMaxChar - countMinChar
}

func calcRound(polymer string, instructions []Instruction) string {
	var newPolymer string
	for i := 0; i < len(polymer)-1; i++ {
		newPolymer += string(polymer[i])
		for _, instruction := range instructions {
			if polymer[i:i+2] == instruction.pair {
				newPolymer += instruction.insert
			}
		}
	}
	newPolymer += string(polymer[len(polymer)-1])

	return newPolymer
}

func evalA(polymer string, instructions []Instruction) int {
	for round := 0; round < 10; round++ {
		polymer = calcRound(polymer, instructions)
	}

	return evalPolymer(polymer)
}

func printCount(polymer string) {
	counter := make(map[string]int)

	for _, char := range polymer {
		counter[string(char)] += 1
	}

	for key, val := range counter {
		print(key)
		print(": ")
		print(val)
		print("  ")
	}
	print("\n")
}

func evalB(polymer string, instructions []Instruction) int {
	for round := 0; round < 10; round++ {
		printCount(polymer)
		polymer = calcRound(polymer, instructions)
	}

	return evalPolymer(polymer)
}

func getTemplate(lines []string) string {
	return lines[0]
}

func getInstructions(lines []string) []Instruction {
	instructionStrings := lines[util.GetBlankLine(lines)+1:]

	var instructions []Instruction
	for _, inst := range instructionStrings {
		instructionParts := strings.Split(inst, " -> ")
		instructions = append(instructions, Instruction{instructionParts[0], instructionParts[1]})
	}

	return instructions
}

func main() {
	day := 14
	debugSuffix := "_debug"
	filename := fmt.Sprintf("input%02d%v.txt", day, debugSuffix)
	lines := util.ReadFile(filename)
	instructions := getInstructions(lines)
	template := getTemplate(lines)

	resA := evalA(template, instructions)
	resB := evalB(template, instructions)
	fmt.Printf("Day %02d \n", day)
	fmt.Printf("A: %v \n", resA)
	fmt.Printf("B: %v \n", resB)
}
