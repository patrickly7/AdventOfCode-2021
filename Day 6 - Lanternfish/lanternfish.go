package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func calculateNumberOfLanternFish(lanternFishes []int, numOfDays int) int {
	dayNumber := 0

	for dayNumber < numOfDays {
		newLanternFishNumber := 0

		for index := range lanternFishes {
			lanternFishes[index] -= 1
			if lanternFishes[index] < 0 {
				lanternFishes[index] = 6
				newLanternFishNumber++
			}
		}

		for newLanternFishNumber > 0 {
			lanternFishes = append(lanternFishes, 8)
			newLanternFishNumber--
		}

		dayNumber++
	}

	return len(lanternFishes)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer file.Close()

	var scanner = bufio.NewScanner(file)

	var lanternFishes []int

	for scanner.Scan() {
		lanternFishNumbers := strings.Split(scanner.Text(), ",")
		for _, number := range lanternFishNumbers {
			convertedNumber, _ := strconv.Atoi(number)
			lanternFishes = append(lanternFishes, convertedNumber)
		}
	}

	fmt.Printf("Number of Lantern Fish After 80 Days: %d\n", calculateNumberOfLanternFish(lanternFishes, 80))
}
