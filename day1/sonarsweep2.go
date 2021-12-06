package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func convertToSlidingWindowSums(depthMeasurements []int) []int {
	var slidingWindowSums []int

	var firstMeasurement = -1
	var secondMeasurement = -1

	for _, depthMeasurement := range depthMeasurements {
		if firstMeasurement == -1 {
			firstMeasurement = depthMeasurement
		} else if secondMeasurement == -1 {
			secondMeasurement = depthMeasurement
		} else {
			slidingWindowSums = append(slidingWindowSums, firstMeasurement+secondMeasurement+depthMeasurement)
			firstMeasurement = secondMeasurement
			secondMeasurement = depthMeasurement
		}
	}

	return slidingWindowSums
}

func getDepthMeasurementIncreases(depthMeasurements []int) int {
	var depthMeasurementIncreases = 0
	var lastDepthMeasurement = -1

	for _, depthMeasurement := range depthMeasurements {
		if lastDepthMeasurement != -1 {
			if depthMeasurement > lastDepthMeasurement {
				depthMeasurementIncreases += 1
			}
		}

		lastDepthMeasurement = depthMeasurement
	}

	return depthMeasurementIncreases
}

func main() {
	var depthMeasurements []int

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer file.Close()

	var scanner = bufio.NewScanner(file)

	for scanner.Scan() {
		depthMeasurement, _ := strconv.Atoi(scanner.Text())
		depthMeasurements = append(depthMeasurements, depthMeasurement)
	}

	var convertedMeasurements = convertToSlidingWindowSums(depthMeasurements)

	fmt.Printf("The number of depth measurement sum increases is: %d\n", getDepthMeasurementIncreases(convertedMeasurements))
}
