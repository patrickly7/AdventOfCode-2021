package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func calculateExtendedPolymerization(template string, pairInsertions map[string]string, steps int) int {
	currTemplate := template
	stepNumber := 1

	for stepNumber <= steps {
		newTemplate := ""

		for i := 0; i < len(currTemplate)-1; i++ {
			pair := currTemplate[i : i+2]
			rule := pairInsertions[pair]

			// fmt.Printf("Pair: %s\n", pair)
			// fmt.Printf("Rule: %s\n", rule)

			newTemplate += (string(currTemplate[i]) + rule)
			if i == len(currTemplate)-2 {
				newTemplate += string(currTemplate[i+1])
			}

			// fmt.Println(newTemplate)
		}

		//fmt.Printf("After Step %d: %s\n", stepNumber, newTemplate)

		currTemplate = newTemplate
		stepNumber++
	}

	mostCommonElementQty := 0
	leastCommonElementQty := 0

	elementCounts := make(map[string]int)
	for _, char := range currTemplate {
		element := string(char)

		if _, ok := elementCounts[element]; ok {
			elementCounts[element] += 1
		} else {
			elementCounts[element] = 1
		}
	}

	for element, quantity := range elementCounts {
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
	fmt.Println(pairInsertions)
	fmt.Printf("Most Common Element Quantity - Least Common Element Quantity (10 Steps): %d\n",
		calculateExtendedPolymerization(template, pairInsertions, 10))
}
