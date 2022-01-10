package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func calculateNumberOfLanternFish(lanternFishes map[int]int, numOfDays int) int64 {
	currentLanternFish := lanternFishes
	dayNumber := 0

	for dayNumber < numOfDays {
		newLanternFish := make(map[int]int)

		for index := range currentLanternFish {
			if index == 0 { // Reset Timer and Make a new Lanternfish
				newLanternFish[6] += currentLanternFish[index]
				newLanternFish[8] += currentLanternFish[index]
			} else { // Decrement the Timer
				newLanternFish[index-1] += currentLanternFish[index]
			}
		}

		currentLanternFish = newLanternFish
		dayNumber++
	}

	var numberOfLanternFish int64 = 0

	for _, value := range currentLanternFish {
		numberOfLanternFish += int64(value)
	}

	return numberOfLanternFish
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer file.Close()

	var scanner = bufio.NewScanner(file)

	lanternFishes := make(map[int]int)

	for scanner.Scan() {
		lanternFishNumbers := strings.Split(scanner.Text(), ",")
		for _, number := range lanternFishNumbers {
			convertedNumber, _ := strconv.Atoi(number)
			if _, ok := lanternFishes[convertedNumber]; ok {
				lanternFishes[convertedNumber] += 1
			} else {
				lanternFishes[convertedNumber] = 1
			}
		}
	}

	fmt.Printf("Number of Lantern Fish After 256 Days: %d\n", calculateNumberOfLanternFish(lanternFishes, 256))
}
