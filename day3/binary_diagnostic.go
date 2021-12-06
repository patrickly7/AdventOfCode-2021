package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getPowerConsumption(binaryNumbers []string) int64 {
	var gammaRate = ""
	var epsilonRate = ""

	bitLength := len(binaryNumbers[0])

	for i := 0; i < bitLength; i++ {
		oneCount := 0
		zeroCount := 0

		for _, binaryNumber := range binaryNumbers {
			if binaryNumber[i] == '0' {
				zeroCount++
			} else {
				oneCount++
			}
		}

		if oneCount > zeroCount {
			gammaRate += "1"
			epsilonRate += "0"
		} else {
			gammaRate += "0"
			epsilonRate += "1"
		}
	}

	finalGammaRate, _ := strconv.ParseInt(gammaRate, 2, 64)
	finalEpsilonRate, _ := strconv.ParseInt(epsilonRate, 2, 64)

	return finalGammaRate * finalEpsilonRate
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

	fmt.Printf("The power consumption rate is: %d\n", getPowerConsumption(binaryNumbers))
}
