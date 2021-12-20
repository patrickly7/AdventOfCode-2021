package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func calculateSyntaxScore(syntaxLine string) int {
	var syntaxStack []string

	for _, char := range syntaxLine {
		symbol := string(char)
		lastIndex := len(syntaxStack) - 1

		if symbol == ")" {
			if syntaxStack[lastIndex] != "(" {
				return -1
			} else {
				syntaxStack = syntaxStack[:lastIndex]
			}
		} else if symbol == "]" {
			if syntaxStack[lastIndex] != "[" {
				return -1
			} else {
				syntaxStack = syntaxStack[:lastIndex]
			}
		} else if symbol == "}" {
			if syntaxStack[lastIndex] != "{" {
				return -1
			} else {
				syntaxStack = syntaxStack[:lastIndex]
			}
		} else if symbol == ">" {
			if syntaxStack[lastIndex] != "<" {
				return -1
			} else {
				syntaxStack = syntaxStack[:lastIndex]
			}
		} else {
			syntaxStack = append(syntaxStack, symbol)
		}
	}

	syntaxScore := 0

	for i := len(syntaxStack) - 1; i >= 0; i-- {
		syntaxScore *= 5

		if syntaxStack[i] == "(" {
			syntaxScore += 1
		} else if syntaxStack[i] == "[" {
			syntaxScore += 2
		} else if syntaxStack[i] == "{" {
			syntaxScore += 3
		} else if syntaxStack[i] == "<" {
			syntaxScore += 4
		}
	}

	return syntaxScore
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer file.Close()

	var scanner = bufio.NewScanner(file)

	var syntaxScores []int

	for scanner.Scan() {
		syntaxScore := calculateSyntaxScore(scanner.Text())
		if syntaxScore != -1 {
			syntaxScores = append(syntaxScores, syntaxScore)
		}
	}

	sort.Ints(syntaxScores)

	fmt.Printf("The total syntax score is: %d\n", syntaxScores[len(syntaxScores)/2])
}
