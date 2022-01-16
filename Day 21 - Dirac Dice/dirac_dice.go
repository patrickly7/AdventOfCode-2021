package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func calculateScore(playerOne int, playerTwo int) int {
	playerOneScore := 0
	playerTwoScore := 0

	currDiceRoll := 1
	diceRolls := 0
	for playerOneScore < 1000 && playerTwoScore < 1000 {
		// Play Player One's Turn
		playerOneDiceTotal := calculateMovement(currDiceRoll)
		playerOne = getNextSpace(playerOne, playerOneDiceTotal)
		playerOneScore += playerOne

		currDiceRoll = getNextDiceRoll(currDiceRoll)
		diceRolls += 3

		if playerOneScore >= 1000 {
			break
		}

		// Play Player Two's Turn
		playerTwoDiceTotal := calculateMovement(currDiceRoll)
		playerTwo = getNextSpace(playerTwo, playerTwoDiceTotal)
		playerTwoScore += playerTwo

		currDiceRoll = getNextDiceRoll(currDiceRoll)
		diceRolls += 3
	}

	fmt.Printf("Dice Rolls: %d\n", diceRolls)
	fmt.Printf("Player 1 Score: %d, Player 2 Score: %d\n", playerOneScore, playerTwoScore)

	if playerOneScore < playerTwoScore {
		return playerOneScore * diceRolls
	} else {
		return playerTwoScore * diceRolls
	}
}

func calculateMovement(currDiceRoll int) int {
	return ((3 * currDiceRoll) + 3) % 10
}

func getNextDiceRoll(currDiceRoll int) int {
	currDiceRoll += 3

	if currDiceRoll > 100 {
		return currDiceRoll % 100
	}

	return currDiceRoll
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

	fmt.Printf("The number of dice rolls * the losing score is: %d\n", calculateScore(playerOne, playerTwo))
}
