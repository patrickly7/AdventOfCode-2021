package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getOxygenGeneratorRate(binaryNumbers []string) int64 {
	currentBinaryList := binaryNumbers

	index := 0
	for len(currentBinaryList) > 1 {
		var zeroList []string
		var oneList []string

		for _, binaryNumber := range currentBinaryList {
			if binaryNumber[index] == '0' {
				zeroList = append(zeroList, binaryNumber)
			} else {
				oneList = append(oneList, binaryNumber)
			}
		}

		if len(oneList) > len(zeroList) || len(oneList) == len(zeroList) {
			currentBinaryList = oneList
		} else {
			currentBinaryList = zeroList
		}

		index++
	}

	fmt.Printf("The oxygen binary number is: %s\n", currentBinaryList[0])
	oxygenGeneratorRate, _ := strconv.ParseInt(currentBinaryList[0], 2, 64)

	return oxygenGeneratorRate
}

func getCO2ScrubberRate(binaryNumbers []string) int64 {
	currentBinaryList := binaryNumbers

	index := 0
	for len(currentBinaryList) > 1 {
		var zeroList []string
		var oneList []string

		for _, binaryNumber := range currentBinaryList {
			if binaryNumber[index] == '0' {
				zeroList = append(zeroList, binaryNumber)
			} else {
				oneList = append(oneList, binaryNumber)
			}
		}

		if len(zeroList) < len(oneList) || len(zeroList) == len(oneList) {
			currentBinaryList = zeroList
		} else {
			currentBinaryList = oneList
		}

		index++
	}

	fmt.Printf("The CO2 binary number is: %s\n", currentBinaryList[0])
	co2ScrubberRate, _ := strconv.ParseInt(currentBinaryList[0], 2, 64)

	return co2ScrubberRate
}

func main() {
	var binaryNumbers []string

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer file.Close()

	var scanner = bufio.NewScanner(file)

	for scanner.Scan() {
		binaryNumbers = append(binaryNumbers, scanner.Text())
	}

	fmt.Printf("The life support rating is: %d\n", getOxygenGeneratorRate(binaryNumbers)*getCO2ScrubberRate(binaryNumbers))
}
