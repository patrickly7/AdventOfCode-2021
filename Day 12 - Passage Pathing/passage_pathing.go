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

func calculateNumberOfPaths(currentPoint string, pathTaken []string) int {
	if currentPoint == "end" {
		finalPath := append(pathTaken, "end")
		fmt.Println(finalPath)

		return 1
	}

	totalPaths := 0
	// fmt.Printf("Currently at %s, can go to %v\n", currentPoint, caves[currentPoint])
	for _, path := range caves[currentPoint] {
		if stringInList(path, pathTaken) && isLower(path) {
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
