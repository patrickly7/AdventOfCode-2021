package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func displaySeaCucumbers(seaCucumbers [][]string) {
	fmt.Print("\033[H\033[2J")

	for _, seaCucumberRow := range seaCucumbers {
		for _, seaCucumber := range seaCucumberRow {
			fmt.Printf(seaCucumber)
		}
		fmt.Printf("\n")
	}

	//fmt.Println("==========")
}

func getNextPosition(position int, boundary int) int {
	if position+1 != boundary {
		return position + 1
	}

	return 0
}

func calculateNoSeaCucumberMovementStep(seaCucumbers [][]string) int {
	seaCucumberHeight := len(seaCucumbers)
	seaCucumberLength := len(seaCucumbers[0])
	stepNumber := 1

	for true {
		seaCucumberMovements := 0

		// Mark Movable Eastward Sea Cucumbers
		for row := 0; row < seaCucumberHeight; row++ {
			for col := 0; col < seaCucumberLength; col++ {
				if seaCucumbers[row][col] == ">" {
					nextCol := getNextPosition(col, seaCucumberLength)
					if seaCucumbers[row][nextCol] == "." {
						seaCucumbers[row][col] = "X"
						seaCucumberMovements++
					}
				}
			}
		}

		// Move Eastward Sea Cucumbers
		for row := 0; row < seaCucumberHeight; row++ {
			for col := 0; col < seaCucumberLength; col++ {
				if seaCucumbers[row][col] == "X" {
					seaCucumbers[row][col] = "."
					nextCol := getNextPosition(col, seaCucumberLength)
					seaCucumbers[row][nextCol] = ">"
				}
			}
		}

		// Mark Movable Southward Sea Cucumbers
		for row := 0; row < seaCucumberHeight; row++ {
			for col := 0; col < seaCucumberLength; col++ {
				if seaCucumbers[row][col] == "v" {
					nextRow := getNextPosition(row, seaCucumberHeight)
					if seaCucumbers[nextRow][col] == "." {
						seaCucumbers[row][col] = "X"
						seaCucumberMovements++
					}
				}
			}
		}

		// Move Southward Sea Cucumbers
		for row := 0; row < seaCucumberHeight; row++ {
			for col := 0; col < seaCucumberLength; col++ {
				if seaCucumbers[row][col] == "X" {
					seaCucumbers[row][col] = "."
					nextRow := getNextPosition(row, seaCucumberHeight)
					seaCucumbers[nextRow][col] = "v"
				}
			}
		}

		// displaySeaCucumbers(seaCucumbers)
		// time.Sleep(1 * time.Second)

		if seaCucumberMovements == 0 {
			break
		}

		stepNumber++
	}

	return stepNumber
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer file.Close()

	var scanner = bufio.NewScanner(file)

	var seaCucumbers [][]string

	for scanner.Scan() {
		var seaCucumberRow []string
		for _, char := range scanner.Text() {
			seaCucumber := string(char)
			seaCucumberRow = append(seaCucumberRow, seaCucumber)
		}
		seaCucumbers = append(seaCucumbers, seaCucumberRow)
	}

	displaySeaCucumbers(seaCucumbers)

	fmt.Printf("The first step no sea cucumbers move is: %d\n", calculateNoSeaCucumberMovementStep(seaCucumbers))
}
