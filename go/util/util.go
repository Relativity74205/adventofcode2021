package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const ResourcePath = "resources"

func ReadFile(filename string) []string {
	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.
	file, err := os.Open("../" + ResourcePath + "/" + filename)

	if err != nil {
		log.Fatalf("failed to open")

	}

	// The bufio.NewScanner() function is called in which the
	// object os.File passed as its parameter and this returns a
	// object bufio.Scanner which is further used on the
	// bufio.Scanner.Split() method.
	scanner := bufio.NewScanner(file)

	// The bufio.ScanLines is used as an
	// input to the method bufio.Scanner.Split()
	// and then the scanning forwards to each
	// new line using the bufio.Scanner.Scan()
	// method.
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	// The method os.File.Close() is called
	// on the os.File object to close the file
	err = file.Close()

	// and then a loop iterates through
	// and prints each of the slice values.

	return text
}

func LinesToInt(lines []string) []int {
	var intLines []int
	for _, v := range lines {
		vint, _ := strconv.Atoi(v)
		intLines = append(intLines, vint)
	}

	return intLines
}

func SumIntegers(integers []int) int {
	sum := 0

	for _, v := range integers {
		sum += v
	}

	return sum
}

func PrintIntegers(integers []int) {
	for _, v := range integers {
		fmt.Printf("%v ", v)
	}
	fmt.Print("\n")
}

func MaxInt(int1, int2 int) int {
	if int1 > int2 {
		return int1
	} else {
		return int2
	}
}

func MinInt(int1, int2 int) int {
	if int1 < int2 {
		return int1
	} else {
		return int2
	}
}

func AbsInt(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

func MaxIntegers(integers []int) int {
	var max int
	for _, i := range integers {
		max = MaxInt(i, max)
	}
	return max
}

func MinIntegers(integers []int) int {
	var min int
	for _, i := range integers {
		min = MinInt(i, min)
	}
	return min
}
