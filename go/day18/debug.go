package main

import (
	"fmt"
	"strings"
)

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
		sfn := createSFN(input, nil)
		actual := sfn.magnitude()

		if actual == expected {
			outcome = "passed"
		} else {
			outcome = "FAILED"
		}
		fmt.Printf("Magnitude debug: outcome: %s; expected: %d, actual: %d, input: %s \n", outcome, expected, actual, input)
	}
}

func debugExplode() {
	var outcome string
	testCases := map[string]string{
		"[[[[[9,8],1],2],3],4]":                 "[[[[0,9],2],3],4]",
		"[7,[6,[5,[4,[3,2]]]]]":                 "[7,[6,[5,[7,0]]]]",
		"[[6,[5,[4,[3,2]]]],1]":                 "[[6,[5,[7,0]]],3]",
		"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]": "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
		"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]":     "[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
	}

	for input, expected := range testCases {
		sfn := createSFN(input, nil)
		traverseOnce(sfn)
		actual := serializeSFN(sfn)

		if actual == expected {
			outcome = "passed"
		} else {
			outcome = "FAILED"
		}
		fmt.Printf("Magnitude debug: outcome: %s; expected: %d, actual: %d, input: %s \n", outcome, expected, actual, input)
	}
}

func debugAdd() {
	var outcome string
	testCases := map[string]string{
		"[[[[4,3],4],4],[7,[[8,4],9]]];[1,1]":                                         "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		"[1,1];[2,2];[3,3];[4,4]":                                                     "[[[[1,1],[2,2]],[3,3]],[4,4]]",
		"[1,1];[2,2];[3,3];[4,4];[5,5]":                                               "[[[[3,0],[5,3]],[4,4]],[5,5]]",
		"[1,1];[2,2];[3,3];[4,4];[5,5];[6,6]":                                         "[[[[5,0],[7,4]],[5,5]],[6,6]]",
		"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]];[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]": "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]",
	}
	for input, expected := range testCases {
		var nodes []*Node
		for _, s := range strings.Split(input, ";") {
			nodes = append(nodes, createSFN(s, nil))
		}
		node := nodes[0]
		for _, newNode := range nodes[1:] {
			node = addition(node, newNode)
			reduce(node)
		}
		actual := serializeSFN(node)
		if actual == expected {
			outcome = "passed"
		} else {
			outcome = "FAILED"
		}
		fmt.Printf("Add debug: outcome: %s; expected: %d, actual: %d; input: %s\n", outcome, expected, actual, input)
	}
}

//[[[[4, 0], [5, 4]], [[7, 7], [6, 0]]],[[8, [7, 7]], [[7, 9], [5, 0]]]]
//[[[[4, 0], [5, 4]], [[7, 7], [6, 0]]],[[[6, 6], [5, 6]], [[6, 0], [7, 7]]]]
