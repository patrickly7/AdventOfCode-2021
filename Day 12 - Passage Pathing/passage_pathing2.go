package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

var caves map[string][]string
var visitedPaths map[string]int

func calculateNumberOfPaths(currentPoint string, pathTaken []string) int {
	if currentPoint == "start" && len(pathTaken) > 0 {
		return 0
	}

	if currentPoint == "end" {
		finalPath := append(pathTaken, "end")

		if _, ok := visitedPaths[strings.Join(finalPath[:], ",")]; ok {
			return 0
		}

		if hasVisitedMultipleSmallCavesTwice(finalPath) {
			return 0
		}

		fmt.Println(finalPath)

		visitedPaths[strings.Join(finalPath[:], ",")] = 1

		return 1
	}

	totalPaths := 0
	// fmt.Printf("Currently at %s, can go to %v\n", currentPoint, caves[currentPoint])
	for _, path := range caves[currentPoint] {
		if stringInList(path, pathTaken) && isLower(path) && hasVisitedSmallCaveTwice(pathTaken) {
			continue
		}

		updatedPathTaken := append(pathTaken, currentPoint)
		totalPaths += calculateNumberOfPaths(path, updatedPathTaken)
	}

	return totalPaths
}

func stringInList(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}

	return false
}

func hasVisitedMultipleSmallCavesTwice(list []string) bool {
	smallCaveMap := make(map[string]int)

	for _, b := range list {
		if _, ok := smallCaveMap[b]; ok {
			smallCaveMap[b] += 1
		} else if isLower(b) {
			smallCaveMap[b] = 1
		}
	}

	visitedTwiceCount := 0
	for _, caveVisits := range smallCaveMap {
		if caveVisits > 1 {
			visitedTwiceCount++
		}
	}

	return visitedTwiceCount > 1
}

func hasVisitedSmallCaveTwice(list []string) bool {
	smallCaveMap := make(map[string]int)

	for _, b := range list {
		if _, ok := smallCaveMap[b]; ok {
			return true
		}

		if isLower(b) {
			smallCaveMap[b] = 1
		}
	}

	return false
}

func isLower(str string) bool {
	for _, r := range str {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}

	return true
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer file.Close()

	var scanner = bufio.NewScanner(file)

	caves = make(map[string][]string)
	visitedPaths = make(map[string]int)

	for scanner.Scan() {
		var splitString = strings.Split(scanner.Text(), "-")

		if _, ok := caves[splitString[0]]; ok {
			caves[splitString[0]] = append(caves[splitString[0]], splitString[1])
		} else {
			caves[splitString[0]] = []string{splitString[1]}
		}

		if _, ok := caves[splitString[1]]; ok {
			caves[splitString[1]] = append(caves[splitString[1]], splitString[0])
		} else {
			caves[splitString[1]] = []string{splitString[0]}
		}
	}

	for start, paths := range caves {
		fmt.Printf("%s goes to %v\n", start, paths)
	}

	fmt.Printf("Num of Paths: %d\n", calculateNumberOfPaths("start", []string{}))
}
