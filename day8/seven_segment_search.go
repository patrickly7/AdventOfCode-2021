package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func isEasyCode(code string) int {
	if len(code) == 2 { // 1
		return 1
	}

	if len(code) == 4 { // 4
		return 1
	}

	if len(code) == 3 { // 7
		return 1
	}

	if len(code) == 7 { // 8
		return 1
	}

	return 0
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer file.Close()

	var scanner = bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		var splitString = strings.Fields(scanner.Text())

		firstCode := splitString[len(splitString)-4]
		total += isEasyCode(firstCode)

		secondCode := splitString[len(splitString)-3]
		total += isEasyCode(secondCode)

		thirdCode := splitString[len(splitString)-2]
		total += isEasyCode(thirdCode)

		fourthCode := splitString[len(splitString)-1]
		total += isEasyCode(fourthCode)
	}

	fmt.Printf("The number of times 1, 4, 7, or 8 appear is: %d\n", total)
}
