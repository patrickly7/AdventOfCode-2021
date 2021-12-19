package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func getBestFuelPosition(horizontalPositions []int, smallestPosition int, largestPosition int) int {
	bestFuel := -1

	for position := smallestPosition; position < largestPosition+1; position++ {
		currFuel := calculateFuelToPosition(horizontalPositions, position)

		if bestFuel == -1 || currFuel < bestFuel {
			bestFuel = currFuel
		}
	}

	return bestFuel
}

func calculateFuelToPosition(horizontalPositions []int, targetPosition int) int {
	sum := 0

	for _, horizontalPosition := range horizontalPositions {
		sum += int(math.Abs(float64(horizontalPosition) - float64(targetPosition)))
	}

	return sum
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer file.Close()

	var scanner = bufio.NewScanner(file)
	smallestPosition := -1
	largestPosition := -1

	var horizontalPositions []int

	for scanner.Scan() {
		var positions = strings.Split(scanner.Text(), ",")

		for _, position := range positions {
			convertedPosition, _ := strconv.Atoi(position)

			if convertedPosition < smallestPosition || smallestPosition == -1 {
				smallestPosition = convertedPosition
			}

			if convertedPosition > largestPosition {
				largestPosition = convertedPosition
			}

			horizontalPositions = append(horizontalPositions, convertedPosition)
		}
	}

	fmt.Printf("The fuel needed is: %d\n", getBestFuelPosition(horizontalPositions, smallestPosition, largestPosition))
}
