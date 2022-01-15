package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func calculateHighestYPosition(xBounds []int, yBounds []int) int {
	highestYPosition := 0

	for x := 0; x < xBounds[1]; x++ {
		for y := yBounds[0]; y < 1000; y++ {
			currPosition := []int{0, 0}
			currVelocity := []int{x, y}
			highestY := 0

			loops := 0
			for loops < 1000 {
				// Position Increases Based on X & Y Velocities
				currPosition[0] += currVelocity[0]
				currPosition[1] += currVelocity[1]

				if highestY < currPosition[1] {
					highestY = currPosition[1]
				}

				// Decrease X Velocity Towards 0
				if currVelocity[0] < 0 {
					currVelocity[0] += 1
				} else if currVelocity[0] > 0 {
					currVelocity[0] -= 1
				}

				// Decrease Y Velocity By 1 (Gravity)
				currVelocity[1] -= 1

				// If it hit the Target Area, Update the Highest Y Achieved
				if currPosition[0] >= xBounds[0] && currPosition[1] <= xBounds[1] && currPosition[1] >= yBounds[0] && currPosition[1] <= yBounds[1] {
					if highestY > highestYPosition {
						highestYPosition = highestY
						fmt.Printf("Reached Y Position: %d @ Velocity (%d, %d)\n", highestYPosition, x, y)
					}
					break
				}

				if currPosition[0] > xBounds[1] && currPosition[1] < yBounds[0] {
					break
				}

				loops++
			}
		}
	}

	return highestYPosition
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer file.Close()

	var scanner = bufio.NewScanner(file)

	var xBoundaries []int
	var yBoundaries []int

	for scanner.Scan() {
		splitString := strings.Fields(scanner.Text())

		xBounds := splitString[2][2 : len(splitString[2])-1]
		xBoundList := strings.Split(xBounds, "..")
		for _, bound := range xBoundList {
			value, _ := strconv.Atoi(bound)
			xBoundaries = append(xBoundaries, value)
		}

		yBounds := splitString[3][2:]
		yBoundList := strings.Split(yBounds, "..")
		for _, bound := range yBoundList {
			value, _ := strconv.Atoi(bound)
			yBoundaries = append(yBoundaries, value)
		}
	}

	fmt.Println(xBoundaries)
	fmt.Println(yBoundaries)

	fmt.Printf("Highest Y Position is: %d\n", calculateHighestYPosition(xBoundaries, yBoundaries))
}
