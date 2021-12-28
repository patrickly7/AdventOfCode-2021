package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func hasFlashingOctopus(octopuses [10][10]int) (bool, int, int) {
	for row := 0; row < 10; row++ {
		for col := 0; col < 10; col++ {
			if octopuses[row][col] > 9 {
				return true, row, col
			}
		}
	}

	return false, -1, -1
}

func getTotalFlashes(octopuses [10][10]int, steps int) int {
	totalFlashes := 0

	for steps > 0 {
		// Increase Energy Levels
		for row := 0; row < 10; row++ {
			for col := 0; col < 10; col++ {
				octopuses[row][col] += 1
			}
		}

		// Check for Flashing Octopuses
		for true {
			result, row, col := hasFlashingOctopus(octopuses)

			if !result {
				break
			}

			// Left Side
			if row != 0 {
				if col != 0 && octopuses[row-1][col-1] != -1 {
					octopuses[row-1][col-1] += 1 // top left
				}

				if col != 9 && octopuses[row-1][col+1] != -1 {
					octopuses[row-1][col+1] += 1 // top right
				}

				if octopuses[row-1][col] != -1 {
					octopuses[row-1][col] += 1 // top
				}
			}

			// Right Side
			if row != 9 {
				if col != 0 && octopuses[row+1][col-1] != -1 {
					octopuses[row+1][col-1] += 1 // bottom left
				}

				if col != 9 && octopuses[row+1][col+1] != -1 {
					octopuses[row+1][col+1] += 1 // bottom right
				}

				if octopuses[row+1][col] != -1 {
					octopuses[row+1][col] += 1 // bottom
				}
			}

			// Above and Below
			if col != 0 && octopuses[row][col-1] != -1 {
				octopuses[row][col-1] += 1 // left
			}

			if col != 9 && octopuses[row][col+1] != -1 {
				octopuses[row][col+1] += 1 // right
			}

			octopuses[row][col] = -1
			totalFlashes++
		}

		// Reset Any Flashed Octopuses
		for row := 0; row < 10; row++ {
			for col := 0; col < 10; col++ {
				if octopuses[row][col] == -1 {
					octopuses[row][col] = 0
				}
			}
		}

		steps--
	}

	return totalFlashes
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer file.Close()

	var scanner = bufio.NewScanner(file)

	var octopuses [10][10]int

	row := 0
	for scanner.Scan() {
		for col, char := range scanner.Text() {
			str := string(char)
			energyLevel, _ := strconv.Atoi(str)

			octopuses[row][col] = energyLevel
		}
		row++
	}

	fmt.Printf("The total number of flashes after 100 steps is: %d\n", getTotalFlashes(octopuses, 100))
}
