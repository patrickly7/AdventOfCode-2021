package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func calculateSyntaxScore(syntaxLine string) int {
	var syntaxStack []string

	for _, char := range syntaxLine {
		symbol := string(char)
		lastIndex := len(syntaxStack) - 1

		if symbol == ")" {
			if syntaxStack[lastIndex] != "(" {
				return 3
			} else {
				syntaxStack = syntaxStack[:lastIndex]
			}
		} else if symbol == "]" {
			if syntaxStack[lastIndex] != "[" {
				return 57
			} else {
				syntaxStack = syntaxStack[:lastIndex]
			}
		} else if symbol == "}" {
			if syntaxStack[lastIndex] != "{" {
				return 1197
			} else {
				syntaxStack = syntaxStack[:lastIndex]
			}
		} else if symbol == ">" {
			if syntaxStack[lastIndex] != "<" {
				return 25137
			} else {
				syntaxStack = syntaxStack[:lastIndex]
			}
		} else {
			syntaxStack = append(syntaxStack, symbol)
		}
	}

	return 0
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer file.Close()

	var scanner = bufio.NewScanner(file)

	totalSyntaxScore := 0

	for scanner.Scan() {
		syntaxScore := calculateSyntaxScore(scanner.Text())
		totalSyntaxScore += syntaxScore
	}

	fmt.Printf("The total syntax score is: %d\n", totalSyntaxScore)
}
