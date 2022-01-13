package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

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

	fmt.Printf("The number of depth measurements increases is: %d\n", getDepthMeasurementIncreases(depthMeasurements))
}
