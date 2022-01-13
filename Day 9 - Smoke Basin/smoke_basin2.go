package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Coordinate struct {
	X, Y, Height int
}

func getThreeLargestBasins(smokeBasin [][]int) int {
	numOfRows := len(smokeBasin)
	numOfCols := len(smokeBasin[0])

	fmt.Printf("%d rows, %d cols\n", numOfRows, numOfCols)

	lowPoints := []Coordinate{}

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

			lowPoints = append(lowPoints, Coordinate{row, col, currHeight})
		}
	}

	basins := [][]Coordinate{}

	for _, lowPoint := range lowPoints {
		basin := []Coordinate{}

		pointQueue := []Coordinate{lowPoint}
		visited := make(map[Coordinate]int)
		for len(pointQueue) > 0 {
			// Get the Current Point
			currentPoint := pointQueue[0]

			if _, ok := visited[currentPoint]; ok {
				pointQueue = pointQueue[1:]
				continue
			}

			basin = append(basin, currentPoint)
			visited[currentPoint] = 1

			// Add Any Valid Neighbors to the Queue
			if currentPoint.Y != 0 && smokeBasin[currentPoint.X][currentPoint.Y-1] > currentPoint.Height {
				if smokeBasin[currentPoint.X][currentPoint.Y-1] != 9 {
					pointQueue = append(pointQueue, Coordinate{currentPoint.X, currentPoint.Y - 1, smokeBasin[currentPoint.X][currentPoint.Y-1]})
				}
			}

			if currentPoint.Y+1 != numOfCols && smokeBasin[currentPoint.X][currentPoint.Y+1] > currentPoint.Height {
				if smokeBasin[currentPoint.X][currentPoint.Y+1] != 9 {
					pointQueue = append(pointQueue, Coordinate{currentPoint.X, currentPoint.Y + 1, smokeBasin[currentPoint.X][currentPoint.Y+1]})
				}
			}

			if currentPoint.X != 0 && smokeBasin[currentPoint.X-1][currentPoint.Y] > currentPoint.Height {
				if smokeBasin[currentPoint.X-1][currentPoint.Y] != 9 {
					pointQueue = append(pointQueue, Coordinate{currentPoint.X - 1, currentPoint.Y, smokeBasin[currentPoint.X-1][currentPoint.Y]})
				}
			}

			if currentPoint.X+1 != numOfRows && smokeBasin[currentPoint.X+1][currentPoint.Y] > currentPoint.Height {
				if smokeBasin[currentPoint.X+1][currentPoint.Y] != 9 {
					pointQueue = append(pointQueue, Coordinate{currentPoint.X + 1, currentPoint.Y, smokeBasin[currentPoint.X+1][currentPoint.Y]})
				}
			}

			// Pop Off the Visited Coordinate
			pointQueue = pointQueue[1:]
		}

		basins = append(basins, basin)
	}

	// for _, basin := range basins {
	// 	fmt.Printf("Basin (Size %d): %v\n", len(basin), basin)
	// }

	largestBasin := 0
	secondLargestBasin := 0
	thirdLargestBasin := 0

	for _, basin := range basins {
		basinSize := len(basin)

		if basinSize >= largestBasin {
			thirdLargestBasin = secondLargestBasin
			secondLargestBasin = largestBasin
			largestBasin = basinSize

		} else if basinSize >= secondLargestBasin {
			thirdLargestBasin = secondLargestBasin
			secondLargestBasin = basinSize

		} else if basinSize >= thirdLargestBasin {
			thirdLargestBasin = basinSize
		}
	}

	return largestBasin * secondLargestBasin * thirdLargestBasin
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

	fmt.Printf("The largest three basins multiplied is: %d\n", getThreeLargestBasins(smokeBasin))
}
