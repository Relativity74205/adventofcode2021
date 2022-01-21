package main

import (
	"aoc21_go/util"
	"fmt"
	"strconv"
)

var hexToBitMap = map[string]string{
	"0": "0000",
	"1": "0001",
	"2": "0010",
	"3": "0011",
	"4": "0100",
	"5": "0101",
	"6": "0110",
	"7": "0111",
	"8": "1000",
	"9": "1001",
	"A": "1010",
	"B": "1011",
	"C": "1100",
	"D": "1101",
	"E": "1110",
	"F": "1111",
}

type Packet interface {
	getCntBits() int
	getType() int
}

type Literal struct {
	version int
	value   int
	cntBits int
}

type Operator struct {
	version    int
	opType     int
	subPackets []Packet
	cntBits    int
}

func (o Literal) getCntBits() int  { return o.cntBits }
func (o Literal) getType() int     { return 4 }
func (o Operator) getCntBits() int { return o.cntBits }
func (o Operator) getType() int    { return o.opType }

func hexToBit(hex string) string {
	var bitArray string
	for _, char := range hex {
		bitArray = bitArray + hexToBitMap[string(char)]
	}

	return bitArray
}

func bitToInt(bitArray string) int {
	i, _ := strconv.ParseInt(bitArray, 2, 64)
	return int(i)
}

func decodeLiteralBinary(bitArray string) string {
	if string(bitArray[0]) == "0" {
		return bitArray[1:5]
	}

	return bitArray[1:5] + decodeLiteralBinary(bitArray[5:])
}

func decodeLiteral(bitArray string) Literal {
	version := getVersion(bitArray)
	valueBinary := decodeLiteralBinary(bitArray[6:])
	cntBits := int(len(valueBinary)/4*5) + 6
	value := bitToInt(valueBinary)

	return Literal{version, value, cntBits}
}

func getVersion(bitArray string) int {
	return bitToInt(bitArray[0:3])
}

func getType(bitArray string) int {
	return bitToInt(bitArray[3:6])
}

func decodeLengthSubPackets(bitArray string) []Packet {
	var subPackets []Packet
	length := bitToInt(bitArray[7:22])
	startPoint := 22

	for startPoint < (length + 22) {
		subPacket := decode(bitArray[startPoint:])
		startPoint += subPacket.getCntBits()
		subPackets = append(subPackets, subPacket)
	}

	return subPackets
}

func decodeNumSubPackets(bitArray string) []Packet {
	var subPackets []Packet
	cnt := bitToInt(bitArray[7:18])
	startPoint := 18

	for i := 1; i <= cnt; i++ {
		subPacket := decode(bitArray[startPoint:])
		startPoint += subPacket.getCntBits()
		subPackets = append(subPackets, subPacket)
	}

	return subPackets
}

func decodeOperator(bitArray string) Operator {
	var subPackets []Packet
	version := getVersion(bitArray)
	opType := getType(bitArray)
	cntBits := 7

	if string(bitArray[6]) == "0" {
		subPackets = decodeLengthSubPackets(bitArray)
		cntBits += 15
	} else {
		subPackets = decodeNumSubPackets(bitArray)
		cntBits += 11
	}

	for _, subPacket := range subPackets {
		cntBits += subPacket.getCntBits()
	}

	return Operator{version, opType, subPackets, cntBits}
}

func decode(bitArray string) Packet {
	if getType(bitArray) == 4 {
		return decodeLiteral(bitArray)
	} else {
		return decodeOperator(bitArray)
	}
}

func evalTypes(packet Packet) int {
	switch v := packet.(type) {
	case Literal:
		return v.version
	case Operator:
		versionNum := v.version
		for _, subPacket := range v.subPackets {
			versionNum += evalTypes(subPacket)
		}
		return versionNum
	default:
		return 0
	}
}

func evalSubPackets(packet Packet) []int {
	var values []int
	for _, subPacket := range packet.(Operator).subPackets {
		values = append(values, evalPacket(subPacket))
	}

	return values
}

func evalPacket(packet Packet) int {
	switch packet.getType() {
	case 0:
		values := evalSubPackets(packet)
		return util.SumIntegers(values)
	case 1:
		values := evalSubPackets(packet)
		prod := 1
		for _, value := range values {
			prod *= value
		}
		return prod
	case 2:
		values := evalSubPackets(packet)
		return util.MinIntegers(values)
	case 3:
		values := evalSubPackets(packet)
		return util.MaxIntegers(values)
	case 5:
		values := evalSubPackets(packet)
		if values[0] > values[1] {
			return 1
		} else {
			return 0
		}
	case 6:
		values := evalSubPackets(packet)
		if values[0] < values[1] {
			return 1
		} else {
			return 0
		}
	case 7:
		values := evalSubPackets(packet)
		if values[0] == values[1] {
			return 1
		} else {
			return 0
		}
	default:
		return packet.(Literal).value
	}
}

func evalA(transmission string) int {
	bitArray := hexToBit(transmission)
	packet := decode(bitArray)

	return evalTypes(packet)
}

func evalB(transmission string) int {
	bitArray := hexToBit(transmission)
	packet := decode(bitArray)

	return evalPacket(packet)
}

func eval(filename string) {
	lines := util.ReadFile(filename)

	resA := evalA(lines[0])
	resB := evalB(lines[0])
	fmt.Printf("A: %v \n", resA)
	fmt.Printf("B: %v \n", resB)
}

func debugA(input string, expected int) {
	actual := evalA(input)
	var outcome string
	if actual == expected {
		outcome = "passed"
	} else {
		outcome = "FAILED"
	}
	fmt.Printf("A debug: outcome: %s; expected: %d, actual: %d, input: %s \n", outcome, expected, actual, input)
}

func debugB(input string, expected int) {
	actual := evalB(input)
	var outcome string
	if actual == expected {
		outcome = "passed"
	} else {
		outcome = "FAILED"
	}
	fmt.Printf("B debug: outcome: %s; expected: %d, actual: %d, input: %s \n", outcome, expected, actual, input)
}

func main() {
	day := 16
	filename := fmt.Sprintf("input%02d.txt", day)
	fmt.Printf("Day %02d \n", day)
	debugA("8A004A801A8002F478", 16)
	debugA("620080001611562C8802118E34", 12)
	debugA("C0015000016115A2E0802F182340", 23)
	debugA("A0016C880162017C3686B18A3D4780", 31)
	debugB("C200B40A82", 3)
	debugB("04005AC33890", 54)
	debugB("880086C3E88112", 7)
	debugB("CE00C43D881120", 9)
	debugB("D8005AC2A8F0", 1)
	debugB("F600BC2D8F", 0)
	debugB("9C005AC2F8F0", 0)
	debugB("9C0141080250320F1802104A08", 1)

	eval(filename)
}
