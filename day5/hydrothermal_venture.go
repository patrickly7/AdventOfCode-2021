package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getNumberOfIntersectingPoints(hydroThermalVents [][2]int, size int) int {
	var oceanFloor [][]int

	for x := 0; x < size; x++ {
		var oceanRow []int
		for index := 0; index < size; index++ {
			oceanRow = append(oceanRow, 0)
		}
		oceanFloor = append(oceanFloor, oceanRow)
	}

	for index := 0; index < len(hydroThermalVents); index += 2 {
		firstCoord := hydroThermalVents[index]
		secondCoord := hydroThermalVents[index+1]

		if firstCoord[1] == secondCoord[1] { // y1 = y2
			y := firstCoord[1]
			if firstCoord[0] > secondCoord[0] {
				for x := secondCoord[0]; x < firstCoord[0]+1; x++ {
					oceanFloor[x][y] += 1
				}
			} else if firstCoord[0] < secondCoord[0] {
				for x := firstCoord[0]; x < secondCoord[0]+1; x++ {
					oceanFloor[x][y] += 1
				}
			} else {
				oceanFloor[firstCoord[0]][y] += 1
			}

		} else if firstCoord[0] == secondCoord[0] { // x1 = x2
			x := firstCoord[0]
			if firstCoord[1] > secondCoord[1] {
				for y := secondCoord[1]; y < firstCoord[1]+1; y++ {
					oceanFloor[x][y] += 1
				}
			} else if firstCoord[1] < secondCoord[1] {
				for y := firstCoord[1]; y < secondCoord[1]+1; y++ {
					oceanFloor[x][y] += 1
				}
			} else {
				oceanFloor[x][firstCoord[1]] += 1
			}
		}
	}

	for _, row := range oceanFloor {
		fmt.Printf("%v\n", row)
	}

	intersectingPoints := 0

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			if oceanFloor[x][y] > 1 {
				intersectingPoints++
			}
		}
	}

	return intersectingPoints
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer file.Close()

	var scanner = bufio.NewScanner(file)
	largestCoordinate := 0

	var hydroThermalVents [][2]int

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

			coord1 := [2]int{firstCoordX, firstCoordY}
			hydroThermalVents = append(hydroThermalVents, coord1)

			secondCoordX, _ := strconv.Atoi(secondCoord[0])
			if secondCoordX > largestCoordinate {
				largestCoordinate = secondCoordX
			}

			secondCoordY, _ := strconv.Atoi(secondCoord[1])
			if secondCoordY > largestCoordinate {
				largestCoordinate = secondCoordY
			}

			coord2 := [2]int{secondCoordX, secondCoordY}
			hydroThermalVents = append(hydroThermalVents, coord2)
		}
	}

	fmt.Printf("The number of intersected points is: %d\n", getNumberOfIntersectingPoints(hydroThermalVents, largestCoordinate+1))
}
