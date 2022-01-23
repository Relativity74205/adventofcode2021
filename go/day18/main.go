package main

import (
	"aoc21_go/util"
	"fmt"
	"strconv"
)

type sfnNode struct {
	pair   bool
	value  int
	left   *sfnNode
	right  *sfnNode
	parent *sfnNode
}

func (node sfnNode) magnitude() int {
	if node.pair {
		return calcMagnitudePair(node)
	} else {
		return node.value
	}
}

func (node sfnNode) level() int {
	if node.pair {
		if node.parent == nil {
			return 1
		}

		return node.parent.level() + 1
	} else {
		return -1
	}
}

func calcMagnitudePair(node sfnNode) int {

	return 3*node.left.magnitude() + 2*node.right.magnitude()
}

func createRegular(value int, parent *sfnNode) *sfnNode {
	return &sfnNode{false, value, nil, nil, parent}
}

func createSfnPair(line string, parent *sfnNode) *sfnNode {
	var leftNode, rightNode *sfnNode
	newPair := &sfnNode{true, -1, nil, nil, parent}

	middleIndex := findMiddleIndex(line)
	leftPart := line[1:middleIndex]
	rightPart := line[middleIndex+1 : len(line)-1]

	leftNumber, isLeftNumber := strconv.Atoi(leftPart)
	if isLeftNumber == nil {
		leftNode = createRegular(leftNumber, newPair)
	} else {
		leftNode = createSfnPair(leftPart, newPair)
	}
	newPair.left = leftNode

	rightNumber, isRightNumber := strconv.Atoi(rightPart)
	if isRightNumber == nil {
		rightNode = createRegular(rightNumber, newPair)
	} else {
		rightNode = createSfnPair(rightPart, newPair)
	}
	newPair.right = rightNode

	return newPair
}

func findMiddleIndex(line string) int {
	var cntLeft, cntRight, middleIndex int

	for i, char := range line {
		switch string(char) {
		case "[":
			cntLeft += 1
		case "]":
			cntRight += 1
		case ",":
			if cntLeft-cntRight == 1 {
				middleIndex = i
				break
			}
		default:
		}
	}

	return middleIndex
}

func debugMagnitude() {
	var outcome string
	testCases := map[string]int{
		"[[1,2],[[3,4],5]]":                                     143,
		"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]":                     1384,
		"[[[[1,1],[2,2]],[3,3]],[4,4]]":                         445,
		"[[[[3,0],[5,3]],[4,4]],[5,5]]":                         791,
		"[[[[5,0],[7,4]],[5,5]],[6,6]]":                         1137,
		"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]": 3488,
	}

	for input, expected := range testCases {
		sfnNode := createSfnPair(input, nil)
		println(sfnNode.level())
		actual := calcMagnitudePair(*sfnNode)

		if actual == expected {
			outcome = "passed"
		} else {
			outcome = "FAILED"
		}
		fmt.Printf("Magnitude debug: outcome: %s; expected: %d, actual: %d, input: %s \n", outcome, expected, actual, input)
	}
}

//
//func addition(pair1, pair2 sfnPair) sfnPair {
//	newPair := sfnPair{nil, nil, nil}
//	pair1.parent = &newPair
//	pair2.parent = &newPair
//	newPair.left = pair1
//	newPair.right = pair2
//
//	return newPair
//}
//
//func explode(pair *sfnPair) *sfnPair {
//	return pair
//}
//
//func split(regular sfnRegular) sfnPair {
//	newPair := sfnPair{nil, nil, regular.parent}
//	newPair.left = createRegular(int(math.Floor(float64(regular.magnitude())/2.0)), &newPair)
//	newPair.right = createRegular(int(math.Ceil(float64(regular.magnitude())/2.0)), &newPair)
//
//	return newPair
//}
//
//func traverse(pair *sfnPair) bool {
//	if foo(pair.left, pair) == false {
//		return false
//	}
//	if foo(pair.right, pair) == false {
//		return false
//	}
//
//	return true
//}
//
//func foo(node sfnNode, pair *sfnPair) bool {
//	switch t := node.(type) {
//	case sfnPair:
//		if t.level() >= 5 {
//			explode(pair)
//			return false
//		} else {
//			return traverse(pair)
//		}
//	case sfnRegular:
//		if t.magnitude() >= 10 {
//			node = split(node.(sfnRegular))
//			return false
//		}
//	default:
//	}
//
//	return true
//}
//
//func reduce(pair sfnPair) {
//	foo := false
//	for foo == false {
//		foo = traverse(&pair)
//	}
//}

func evalA(lines []string) int {

	return 0
}

func evalB(lines []string) int {

	return 0
}

func eval(filename string, debug bool) {
	lines := util.ReadFile(filename)

	resA := evalA(lines)
	resB := evalB(lines)
	if debug {
		fmt.Printf("A (%s): %v \n", filename, resA)
		fmt.Printf("B (%s): %v \n", filename, resB)
	} else {
		fmt.Printf("A: %v \n", resA)
		fmt.Printf("B: %v \n", resB)
	}

}

func main() {
	day := 18
	filename := fmt.Sprintf("input%02d.txt", day)
	filenameDebug1 := fmt.Sprintf("input%02d%v.txt", day, "_debug1")
	filenameDebug2 := fmt.Sprintf("input%02d%v.txt", day, "_debug2")

	fmt.Printf("Day %02d \n", day)
	debugMagnitude()
	eval(filenameDebug1, true)
	eval(filenameDebug2, true)
	eval(filename, false)
	//test := addition(createSfnPair("[1,2]", nil), createSfnPair("[[1,2],3]", nil))
	//fmt.Printf("1. main  -- i  %T: &i=%p i=%v\n", test, &test, test)
	//println(calcMagnitudePair(test))
}
