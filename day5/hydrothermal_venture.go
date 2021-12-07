package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer file.Close()

	var scanner = bufio.NewScanner(file)
	largestCoordinate := 0

	for scanner.Scan() {
		var splitLine = strings.Fields(scanner.Text())
		firstCoord := strings.Split(splitLine[0], ",")
		secondCoord := strings.Split(splitLine[2], ",")

		// x1 = x2 or y1 = y2 (Horizontal Lines)
		if firstCoord[0] == secondCoord[0] || firstCoord[1] == secondCoord[1] {
			firstCoordX, _ := strconv.Atoi(firstCoord[0])
			if firstCoordX > largestCoordinate {
				largestCoordinate = firstCoordX
			}

			firstCoordY, _ := strconv.Atoi(firstCoord[1])
			if firstCoordY > largestCoordinate {
				largestCoordinate = firstCoordY
			}

			secondCoordX, _ := strconv.Atoi(secondCoord[0])
			if secondCoordX > largestCoordinate {
				largestCoordinate = secondCoordX
			}

			secondCoordY, _ := strconv.Atoi(secondCoord[1])
			if secondCoordY > largestCoordinate {
				largestCoordinate = secondCoordY
			}

			fmt.Println(scanner.Text())
		}
	}
}
