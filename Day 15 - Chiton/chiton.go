package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type Coordinate struct {
	X, Y int
}

func calculateLowestTotalRisk(cave [][]int) int {
	caveRows := len(cave)
	caveColumns := len(cave[0])
	fmt.Printf("Cave Rows: %d, Cave Cols: %d\n", caveRows, caveColumns)

	distances := make(map[Coordinate]int)
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
	for x := 0; x < caveRows; x++ {
		for y := 0; y < caveColumns; y++ {
			priorityQueue[Coordinate{x, y}] = 1
		}
	}

	for len(priorityQueue) > 0 {
		current := getCoordinateWithSmallestDistanceVertex(distances, priorityQueue)
		//fmt.Println(current)
		delete(priorityQueue, current)
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
				fmt.Printf("Neighbor (%d, %d) @ Distance %d from Previous (%d, %d)\n", neighbor.X, neighbor.Y, distance, current.X, current.Y)
				distances[neighbor] = distance
				previous[neighbor] = current
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

	//calculateLowestTotalRisk(cave)
	fmt.Printf("The lowest risk total is: %d\n", calculateLowestTotalRisk(cave))
}
