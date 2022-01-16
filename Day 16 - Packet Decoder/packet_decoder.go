package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"
)

func calculateVersionSum(packet string, currIndex int) (int, int) {
	binaryString := hexadecimalToBinaryString(packet)
	versionSum := 0

	index := currIndex

	// Get Packet Version (First 3 Bits)
	packetVersion, _ := strconv.ParseInt(binaryString[index:index+3], 2, 64)
	fmt.Printf("Packet Version: %d\n", packetVersion)
	versionSum += int(packetVersion)

	index += 3

	// Get Packet Type ID (Next 3 Bits)
	packetTypeId, _ := strconv.ParseInt(binaryString[index:index+3], 2, 64)
	fmt.Printf("Packet Type ID: %d\n", packetTypeId)

	index += 3

	if packetTypeId == 4 { // Literal
		literalValueString := ""

		isLastGroup := false
		for !isLastGroup {
			bit := string(binaryString[index])
			if bit == "0" {
				isLastGroup = true
			}

			literalValueString += binaryString[index+1 : index+5]

			index += 5
		}

		literalValue, _ := strconv.ParseInt(literalValueString, 2, 64)
		fmt.Printf("Literal Value: %d\n", literalValue)

		return versionSum, index
	} else { // Operator
		lengthTypeId := string(binaryString[index])
		fmt.Println("Length Type ID: " + lengthTypeId)
		index++

		if lengthTypeId == "0" {
			lengthOfSubPackets, _ := strconv.ParseInt(binaryString[index:index+15], 2, 64)
			fmt.Printf("Length: %d\n", lengthOfSubPackets)

			index += 15

			currLength := 0
			for currLength < int(lengthOfSubPackets) {
				subpacketVersionSum, newIndex := calculateVersionSum(packet, index)
				versionSum += subpacketVersionSum

				currLength += newIndex - index

				index = newIndex
			}
		} else { // Length Type ID = 1
			numOfSubPackets, _ := strconv.ParseInt(binaryString[index:index+11], 2, 64)
			fmt.Printf("Number of Sub Packets: %d\n", numOfSubPackets)

			index += 11

			subpacket := 0
			for subpacket < int(numOfSubPackets) {
				subpacketVersionSum, newIndex := calculateVersionSum(packet, index)
				versionSum += subpacketVersionSum

				subpacket++

				index = newIndex
			}
		}
	}

	return versionSum, index
}

func hexadecimalToBinaryString(in string) string {
	decodedString, _ := hex.DecodeString(in)

	var binaryString string
	for _, digit := range decodedString {
		binaryString += fmt.Sprintf("%08b", digit)
	}

	return binaryString
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer file.Close()

	var scanner = bufio.NewScanner(file)

	var packet string

	for scanner.Scan() {
		packet = scanner.Text()
	}

	fmt.Println(packet)
	versionSum, _ := calculateVersionSum(packet, 0)
	fmt.Printf("The version sum is: %d\n", versionSum)
}
