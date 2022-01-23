package main

import (
	"aoc21_go/util"
	"fmt"
	"math"
	"strconv"
)

type Node struct {
	isPair bool
	value  int
	left   *Node
	right  *Node
	parent *Node
}

func (node Node) magnitude() int {
	if node.isPair {
		return 3*node.left.magnitude() + 2*node.right.magnitude()
	} else {
		return node.value
	}
}

func (node Node) level() int {
	if node.isPair {
		if node.parent == nil {
			return 1
		}

		return node.parent.level() + 1
	} else {
		return -1
	}
}

func createNode(nodeString string) *Node {
	var node *Node
	value, isNumber := strconv.Atoi(nodeString)
	if isNumber == nil {
		node = &Node{false, value, nil, nil, nil}
	} else {
		node = createSFN(nodeString, nil)
	}

	return node
}

func createSFN(line string, parent *Node) *Node {
	middleIndex := findMiddleIndex(line)
	leftPart := line[1:middleIndex]
	rightPart := line[middleIndex+1 : len(line)-1]

	newPair := &Node{
		true,
		-1,
		createNode(leftPart),
		createNode(rightPart),
		parent,
	}
	newPair.left.parent = newPair
	newPair.right.parent = newPair

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

func serializeSFN(node *Node) string {
	return serializeSFNRec(node, "")
}

func serializeSFNRec(node *Node, s string) string {
	if node.isPair {
		var sNew string
		sNew += "["
		sNew = serializeSFNRec(node.left, sNew)
		sNew += ","
		sNew = serializeSFNRec(node.right, sNew)
		sNew += "]"
		return s + sNew
	} else {
		return s + strconv.Itoa(node.value)
	}
}

func addition(pair1, pair2 *Node) *Node {
	newPair := &Node{true, -1, pair1, pair2, nil}
	newPair.left.parent = newPair
	newPair.right.parent = newPair

	return newPair
}

func getNodeList(node *Node, nodeList []*Node) []*Node {
	nodeList = append(nodeList, node)
	if node.isPair {
		nodeList = getNodeList(node.left, nodeList)
		nodeList = getNodeList(node.right, nodeList)
	}
	return nodeList
}

func split(nodeList []*Node, splitIndex int) {
	node := nodeList[splitIndex]
	newNode := &Node{
		true,
		-1,
		&Node{false, int(math.Floor(float64(node.magnitude()) / 2.0)), nil, nil, nil},
		&Node{false, int(math.Ceil(float64(node.magnitude()) / 2.0)), nil, nil, nil},
		node.parent,
	}
	newNode.left.parent = newNode
	newNode.right.parent = newNode

	if node.parent.left == node {
		node.parent.left = newNode
	} else {
		node.parent.right = newNode
	}
}

func explode(nodeList []*Node, explodeIndex int) {
	node := nodeList[explodeIndex]

	valLeft := node.left.value
	valRight := node.right.value
	newNode := &Node{false, 0, nil, nil, node.parent}
	if node.parent.left == node {
		node.parent.left = newNode
	} else {
		node.parent.right = newNode
	}

	for i := explodeIndex - 1; i >= 0; i-- {
		if !nodeList[i].isPair {
			nodeList[i].value += valLeft
			break
		}
	}
	for i := explodeIndex + 3; i <= len(nodeList)-1; i++ {
		if !nodeList[i].isPair {
			nodeList[i].value += valRight
			break
		}
	}
}

func traverse(nodeList []*Node) bool {
	for i, node := range nodeList {
		if node.isPair && node.level() > 4 {
			explode(nodeList, i)
			return false
		}
		if !node.isPair && node.magnitude() >= 10 {
			split(nodeList, i)
			return false
		}
	}

	return true
}

func traverseOnce(sfn *Node) bool {
	nodeList := getNodeList(sfn, nil)
	return traverse(nodeList)
}

func reduce(sfn *Node) {
	for !traverseOnce(sfn) {
	}
}

func evalA(lines []string) int {
	sfn := createSFN(lines[0], nil)
	for _, line := range lines[1:] {
		newSFN := createSFN(line, nil)
		sfn = addition(sfn, newSFN)
		reduce(sfn)
	}
	println(serializeSFN(sfn))

	return sfn.magnitude()
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
	println(filenameDebug1)
	filenameDebug2 := fmt.Sprintf("input%02d%v.txt", day, "_debug2")

	fmt.Printf("Day %02d \n", day)
	//debugMagnitude()
	//debugExplode()
	debugAdd()
	eval(filenameDebug1, true)
	eval(filenameDebug2, true)
	eval(filename, false)
}
