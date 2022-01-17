package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func calculateExtendedPolymerization(template string, pairInsertions map[string]string, steps int) int64 {
	stepNumber := 1

	// Count the Pairs in the Initial Template
	pairs := make(map[string]int64)
	for i := 0; i < len(template)-1; i++ {
		pair := template[i : i+2]
		pairs[pair] += 1
	}

	// Keep Constructing and Counting New Pairs
	for stepNumber <= steps {
		newPairs := make(map[string]int64)
		for pair, quantity := range pairs {
			newPairs[string(pair[0])+pairInsertions[pair]] += quantity
			newPairs[pairInsertions[pair]+string(pair[1])] += quantity
		}
		pairs = newPairs

		stepNumber++
	}

	// Count Up the Quantities of the Elements
	elements := make(map[string]int64)
	for pair, quantity := range pairs {
		elements[string(pair[0])] += quantity
	}

	// Add the Last Element of the Template that's left over
	elements[string(template[len(template)-1])] += 1

	var mostCommonElementQty int64 = 0
	var leastCommonElementQty int64 = 0

	for element, quantity := range elements {
		fmt.Printf("Element %s: %d\n", element, quantity)

		if quantity > mostCommonElementQty {
			mostCommonElementQty = quantity
		}

		if leastCommonElementQty == 0 || quantity < leastCommonElementQty {
			leastCommonElementQty = quantity
		}
	}

	return mostCommonElementQty - leastCommonElementQty
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer file.Close()

	var scanner = bufio.NewScanner(file)

	template := ""
	pairInsertions := make(map[string]string)

	isTemplate := true
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		if isTemplate {
			template = scanner.Text()
			isTemplate = false
			continue
		}

		splitText := strings.Fields(scanner.Text())
		pairInsertions[splitText[0]] = splitText[2]
	}

	fmt.Printf("Template: %s\n", template)
	fmt.Printf("Most Common Element Quantity - Least Common Element Quantity (40 Steps): %d\n",
		calculateExtendedPolymerization(template, pairInsertions, 40))
}
