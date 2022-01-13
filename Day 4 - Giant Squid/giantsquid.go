package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func calculateScore(bingoBoard [][]string, markerBoard [5][5]int, winningNumber int) int {
	sum := 0

	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if markerBoard[row][col] == 0 {
				convertedBingoNumber, _ := strconv.Atoi(bingoBoard[row][col])
				sum += convertedBingoNumber
			}
		}
	}

	return sum * winningNumber
}

func verifyWinningBoard(markerBoard [5][5]int) bool {
	// Check the Rows
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if markerBoard[row][col] == 0 {
				break
			}

			if col == 4 {
				return true
			}
		}
	}

	// Columns
	for col := 0; col < 5; col++ {
		for row := 0; row < 5; row++ {
			if markerBoard[row][col] == 0 {
				break
			}

			if row == 4 {
				return true
			}
		}
	}

	return false
}

func getBingoTurnsToWinAndScore(bingoBoard [][]string, bingoNumbers []string, currentTurnsToWin int) (int, int) {
	fmt.Printf("Bingo Board: %v\n", bingoBoard)
	turnsToWin := 0
	boardWon := false
	var winningNumber int

	var markerBoard [5][5]int
	for _, bingoNumber := range bingoNumbers {
	out:
		for row := 0; row < 5; row++ {
			for col := 0; col < 5; col++ {
				if bingoBoard[row][col] == bingoNumber {
					markerBoard[row][col] = 1
					break out
				}
			}
		}

		turnsToWin++

		// Verify Marker Board is Won
		isVictory := verifyWinningBoard(markerBoard)
		if isVictory {
			convertedBingoNumber, _ := strconv.Atoi(bingoNumber)
			winningNumber = convertedBingoNumber

			boardWon = true
			break
		}
	}

	// Get the Score if Won
	if boardWon {
		return turnsToWin, calculateScore(bingoBoard, markerBoard, winningNumber)
	}

	return -1, -1
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	isFirstLine := true
	index := 0
	leastTurnsToWin := -1
	winningBoardScore := 0

	var bingoNumbers []string
	var bingoBoard [5][]string
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		if isFirstLine {
			bingoNumbers = strings.Split(scanner.Text(), ",")
			fmt.Printf("Bingo Numbers: %v\n", bingoNumbers)

			isFirstLine = false

			continue
		}

		bingoBoard[index] = strings.Fields(scanner.Text())
		index++

		if index == 5 {
			index = 0

			turnsToWin, score := getBingoTurnsToWinAndScore(bingoBoard[:], bingoNumbers, leastTurnsToWin)
			if turnsToWin != -1 && leastTurnsToWin == -1 || leastTurnsToWin > turnsToWin || (leastTurnsToWin == turnsToWin && score > winningBoardScore) {
				leastTurnsToWin = turnsToWin
				winningBoardScore = score
			}
		}
	}

	fmt.Printf("The score of the winning bingo board is: %d\n", winningBoardScore)
}
