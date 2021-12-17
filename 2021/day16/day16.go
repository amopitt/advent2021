package main

import (
	_ "embed"
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"strconv"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Day 16")

	fmt.Println(input)
	fmt.Println("------------------")

	// convert hex to binary
	bin, err := hexToBin(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(bin))
	// pad bin to be divisible by 4
	for len(bin)%4 != 0 {
		bin = "0" + bin
	}
	fmt.Println(bin)
	a, b, c := parsePacket(bin, 0)
	fmt.Println(a, b, c)
}

const isLiteral = 4

func parsePacket(packet string, pos int) (int, int64, int64) {
	version := binaryToDecimal(packet[pos : pos+3])
	packetType := binaryToDecimal(packet[pos+3 : pos+6])

	fmt.Println(version, packetType, pos)

	if packetType == isLiteral {
		pos += 6
		var by string
		for {
			isLastPacket := packet[pos] == '0'
			by += packet[pos+1 : pos+5]
			pos += 5
			if isLastPacket {
				break
			}
		}
		return pos, version, binaryToDecimal(by)
	}
	lengthType := packet[pos+6]
	var subPackets []int64
	if lengthType == '0' {
		l := binaryToDecimal(packet[pos+7 : pos+7+15])
		pos = pos + 7 + 15
		stopAt := pos + int(l)
		for {
			nextpos, addV, val := parsePacket(packet, pos)
			version += addV
			pos = nextpos
			subPackets = append(subPackets, val)
			if pos >= stopAt {
				break
			}
		}
	} else {
		l := binaryToDecimal(packet[pos+7 : pos+7+11])
		pos = pos + 7 + 11
		for i := int64(0); i < l; i++ {
			nextpos, addV, val := parsePacket(packet, pos)
			version += addV
			pos = nextpos
			subPackets = append(subPackets, val)
		}
	}
	var sum int64
	switch packetType {
	case 0: // sum
		for _, v := range subPackets {
			sum += v
		}
	case 1: // product
		sum = 1
		for _, v := range subPackets {
			sum *= v
		}
	case 2: // min
		sum = math.MaxInt64
		for _, v := range subPackets {
			if v < sum {
				sum = v
			}
		}
	case 3: // max
		for _, v := range subPackets {
			if v > sum {
				sum = v
			}
		}
	case 5: // greater than
		if subPackets[0] > subPackets[1] {
			sum = 1
		}
	case 6: // less than
		if subPackets[0] < subPackets[1] {
			sum = 1
		}
	case 7: // equal
		if subPackets[0] == subPackets[1] {
			sum = 1
		}
	}
	return pos, version, sum
}

func binaryToDecimal(str string) int64 {
	ui, err := strconv.ParseInt(str, 2, 64)
	if err != nil {
		return 0
	}

	return int64(ui)
}

func hexToBin(str string) (string, error) {
	decoded, err := hex.DecodeString(str)
	if err != nil {
		log.Fatal(err)
	}
	var res string
	for _, d := range decoded {
		res += fmt.Sprintf("%08b", d)
	}
	return res, nil
}
