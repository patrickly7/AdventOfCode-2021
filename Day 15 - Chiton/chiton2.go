package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

var distances = make(map[Coordinate]int)

type Coordinate struct {
	X, Y int
}

func calculateLowestTotalRisk(cave [][]int) int {
	caveRows := len(cave)
	caveColumns := len(cave[0])
	fmt.Printf("Cave Rows: %d, Cave Cols: %d\n", caveRows, caveColumns)

	var visited [][]bool
	previous := make(map[Coordinate]Coordinate)

	for x := 0; x < caveRows; x++ {
		var visitedRow []bool
		for y := 0; y < caveColumns; y++ {
			if x == 0 && y == 0 {
				distances[Coordinate{x, y}] = 0
			} else {
				distances[Coordinate{x, y}] = math.MaxInt
			}
			visitedRow = append(visitedRow, false)
		}
		visited = append(visited, visitedRow)
	}

	// Go through the Queue until it's empty
	distances[Coordinate{0, 0}] = 0

	priorityQueue := make(map[Coordinate]int)
	priorityQueue[Coordinate{0, 0}] = 1

	for len(priorityQueue) > 0 {
		current := getCoordinateWithSmallestDistanceVertex(distances, priorityQueue)
		delete(priorityQueue, current)
		//fmt.Println(current)
		visited[current.X][current.Y] = true

		// Quit Once You've Found the Goal
		if current.X == caveRows-1 && current.Y == caveColumns-1 {
			break
		}

		// Check Neighbors
		neighbors := getNeighbors(current, caveRows, caveColumns, visited)
		for _, neighbor := range neighbors {
			distance := costTaken(current, previous, cave) + cave[neighbor.X][neighbor.Y]
			if distance < distances[neighbor] {
				//fmt.Printf("Neighbor (%d, %d) @ Distance %d from Previous (%d, %d)\n", neighbor.X, neighbor.Y, distance, current.X, current.Y)
				distances[neighbor] = distance
				previous[neighbor] = current
				priorityQueue[neighbor] = 1
			}
		}
	}

	return distances[Coordinate{caveRows - 1, caveColumns - 1}]
}

func costTaken(coordinate Coordinate, previous map[Coordinate]Coordinate, cave [][]int) int {
	totalCost := 0
	current := coordinate

	for true {
		if current.X == 0 && current.Y == 0 {
			break
		}

		totalCost += cave[current.X][current.Y]
		current = previous[current]
	}

	return totalCost
}

func getCoordinateWithSmallestDistanceVertex(distances map[Coordinate]int, queue map[Coordinate]int) Coordinate {
	smallestDistance := math.MaxInt
	var nearestCoordinate Coordinate

	isFirst := true
	for coordinate := range queue {
		if isFirst {
			smallestDistance = distances[coordinate]
			nearestCoordinate = coordinate
			isFirst = false
		}

		if distances[coordinate] < smallestDistance {
			smallestDistance = distances[coordinate]
			nearestCoordinate = coordinate
		}
	}

	return nearestCoordinate
}

func getNeighbors(coordinate Coordinate, caveRows int, caveCols int, visited [][]bool) []Coordinate {
	neighbors := []Coordinate{}

	if coordinate.X != 0 {
		neighbor := Coordinate{coordinate.X - 1, coordinate.Y}
		if !visited[neighbor.X][neighbor.Y] {
			neighbors = append(neighbors, neighbor)
		}
	}

	if coordinate.X != caveRows-1 {
		neighbor := Coordinate{coordinate.X + 1, coordinate.Y}
		if !visited[neighbor.X][neighbor.Y] {
			neighbors = append(neighbors, neighbor)
		}
	}

	if coordinate.Y != 0 {
		neighbor := Coordinate{coordinate.X, coordinate.Y - 1}
		if !visited[neighbor.X][neighbor.Y] {
			neighbors = append(neighbors, neighbor)
		}
	}

	if coordinate.Y != caveCols-1 {
		neighbor := Coordinate{coordinate.X, coordinate.Y + 1}
		if !visited[neighbor.X][neighbor.Y] {
			neighbors = append(neighbors, neighbor)
		}
	}

	return neighbors
}

func buildFullCave(cavePattern [][]int) [][]int {
	cave := [][]int{}

	originalCaveRows := len(cavePattern)
	originalCaveCols := len(cavePattern[0])
	//fmt.Printf("Cave Rows: %d, Cave Cols: %d\n", originalCaveRows, originalCaveCols)

	increment := 0
	for increment < 5 {
		for x := 0; x < originalCaveRows; x++ {
			var caveRow []int

			// Expand the Map Five Times with Increments
			for y := 0; y < originalCaveCols; y++ {
				currentSpace := caveSpaceRiskLevel(cavePattern[x][y] + increment)
				caveRow = append(caveRow, currentSpace)
			}

			for y := 0; y < originalCaveCols; y++ {
				currentSpace := caveSpaceRiskLevel(cavePattern[x][y] + increment + 1)
				caveRow = append(caveRow, currentSpace)
			}

			for y := 0; y < originalCaveCols; y++ {
				currentSpace := caveSpaceRiskLevel(cavePattern[x][y] + increment + 2)
				caveRow = append(caveRow, currentSpace)
			}

			for y := 0; y < originalCaveCols; y++ {
				currentSpace := caveSpaceRiskLevel(cavePattern[x][y] + increment + 3)
				caveRow = append(caveRow, currentSpace)
			}

			for y := 0; y < originalCaveCols; y++ {
				currentSpace := caveSpaceRiskLevel(cavePattern[x][y] + increment + 4)
				caveRow = append(caveRow, currentSpace)
			}

			cave = append(cave, caveRow)
		}

		increment++
	}

	return cave
}

func caveSpaceRiskLevel(caveSpace int) int {
	if caveSpace > 9 {
		newRiskLevel := (caveSpace % 9)
		return newRiskLevel
	}

	return caveSpace
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer file.Close()

	var scanner = bufio.NewScanner(file)

	var cave [][]int

	for scanner.Scan() {
		var caveRow []int
		for _, num := range scanner.Text() {
			riskString := string(num)
			risk, _ := strconv.Atoi(riskString)
			caveRow = append(caveRow, risk)
		}

		cave = append(cave, caveRow)
	}

	fullCave := buildFullCave(cave)

	// for row, caveRow := range fullCave {
	// 	for col := range caveRow {
	// 		fmt.Printf("%d", fullCave[row][col])
	// 	}
	// 	fmt.Printf("\n")
	// }

	fmt.Printf("The lowest risk total is: %d\n", calculateLowestTotalRisk(fullCave))
}
