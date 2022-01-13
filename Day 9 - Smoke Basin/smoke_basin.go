package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getTotalRiskLevel(smokeBasin [][]int) int {
	totalRiskLevel := 0

	numOfRows := len(smokeBasin)
	numOfCols := len(smokeBasin[0])

	fmt.Printf("%d rows, %d cols\n", numOfRows, numOfCols)

	for row := 0; row < numOfRows; row++ {
		for col := 0; col < numOfCols; col++ {
			currHeight := smokeBasin[row][col]

			if col != 0 && smokeBasin[row][col-1] <= currHeight {
				continue
			}

			if col+1 != numOfCols && smokeBasin[row][col+1] <= currHeight {
				continue
			}

			if row != 0 && smokeBasin[row-1][col] <= currHeight {
				continue
			}

			if row+1 != numOfRows && smokeBasin[row+1][col] <= currHeight {
				continue
			}

			totalRiskLevel += (currHeight + 1)
		}
	}

	return totalRiskLevel
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer file.Close()

	var scanner = bufio.NewScanner(file)

	var smokeBasin [][]int

	for scanner.Scan() {
		var row []int

		for _, char := range scanner.Text() {
			height, _ := strconv.Atoi(string(char))
			row = append(row, height)
		}

		smokeBasin = append(smokeBasin, row)
	}

	fmt.Printf("The total risk level is: %d\n", getTotalRiskLevel(smokeBasin))
}
