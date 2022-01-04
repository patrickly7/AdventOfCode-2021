package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Universe struct {
	playerOneSpace, playerOneScore, playerTwoSpace, playerTwoScore int
}

var universes = make(map[Universe][]int64)

func calculateUniverseWins(universe Universe) []int64 {
	if universe.playerOneScore >= 21 {
		return []int64{1, 0}
	}

	if universe.playerTwoScore >= 21 {
		return []int64{0, 1}
	}

	if value, isMapped := universes[universe]; isMapped {
		return value
	}

	finalAnswer := []int64{0, 0}

	for rollOne := 1; rollOne <= 3; rollOne++ {
		for rollTwo := 1; rollTwo <= 3; rollTwo++ {
			for rollThree := 1; rollThree <= 3; rollThree++ {
				movement := rollOne + rollTwo + rollThree

				newPlayerOneSpace := getNextSpace(universe.playerOneSpace, movement)
				newPlayerOneScore := universe.playerOneScore + newPlayerOneSpace

				universeWins := calculateUniverseWins(Universe{universe.playerTwoSpace, universe.playerTwoScore, newPlayerOneSpace, newPlayerOneScore})

				finalAnswer[0] += universeWins[1]
				finalAnswer[1] += universeWins[0]
			}
		}
	}

	universes[universe] = finalAnswer

	return finalAnswer
}

func getNextSpace(currSpace int, movement int) int {
	space := currSpace + movement

	if space > 10 {
		if space%10 == 0 {
			return currSpace
		}

		return space % 10
	}

	return space
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer file.Close()

	var scanner = bufio.NewScanner(file)

	playerOne := 0
	playerTwo := 0

	for scanner.Scan() {
		fmt.Println(scanner.Text())

		splitText := strings.Fields(scanner.Text())
		startingSpaceText := splitText[len(splitText)-1]
		startingSpace, _ := strconv.Atoi(startingSpaceText)

		if playerOne == 0 {
			playerOne = startingSpace
		} else {
			playerTwo = startingSpace
		}
	}

	universeWins := calculateUniverseWins(Universe{playerOne, 0, playerTwo, 0})
	fmt.Printf("The number of universes won are: %d, %d\n", universeWins[0], universeWins[1])
}
