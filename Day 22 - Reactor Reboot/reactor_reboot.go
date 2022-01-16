package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Cube struct {
	X, Y, Z int
}

type Cuboid struct {
	X, Y, Z []int
}

func calculateNumberOfCubes(commands []string, cuboids []Cuboid) int {
	turnedOnCuboids := make(map[Cube]int)

	for index, cuboid := range cuboids {
		//fmt.Printf("Step %d (%s):\n", index, commands[index])

		// Skip Out of Range Cubes
		if (cuboid.X[0] < -50 && cuboid.X[1] < -50) || (cuboid.X[0] > 50 && cuboid.X[1] > 50) {
			//fmt.Println("Skipped")
			continue
		}

		if (cuboid.Y[0] < -50 && cuboid.Y[1] < -50) || (cuboid.Y[0] > 50 && cuboid.Y[1] > 50) {
			//fmt.Println("Skipped")
			continue
		}

		if (cuboid.Z[0] < -50 && cuboid.Z[1] < -50) || (cuboid.Z[0] > 50 && cuboid.Z[1] > 50) {
			//fmt.Println("Skipped")
			continue
		}

		// Turn ON or OFF Specified Cubes
		if commands[index] == "on" {
			for x := cuboid.X[0]; x <= cuboid.X[1]; x++ {
				for y := cuboid.Y[0]; y <= cuboid.Y[1]; y++ {
					for z := cuboid.Z[0]; z <= cuboid.Z[1]; z++ {
						turnedOnCuboids[Cube{x, y, z}] = 1
					}
				}
			}
		} else {
			for x := cuboid.X[0]; x <= cuboid.X[1]; x++ {
				for y := cuboid.Y[0]; y <= cuboid.Y[1]; y++ {
					for z := cuboid.Z[0]; z <= cuboid.Z[1]; z++ {
						delete(turnedOnCuboids, Cube{x, y, z})
					}
				}
			}
		}
	}

	return len(turnedOnCuboids)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer file.Close()

	var scanner = bufio.NewScanner(file)

	var commands []string
	var cuboids []Cuboid

	for scanner.Scan() {
		splitString := strings.Fields(scanner.Text())

		commands = append(commands, splitString[0])

		var newCuboid Cuboid
		xyz := strings.Split(splitString[1], ",")

		// Add X Range
		xBounds := xyz[0][2:]
		xBoundList := strings.Split(xBounds, "..")

		xLow, _ := strconv.Atoi(xBoundList[0])
		xHigh, _ := strconv.Atoi(xBoundList[1])

		newCuboid.X = []int{xLow, xHigh}

		// Add Y Range
		yBounds := xyz[1][2:]
		yBoundList := strings.Split(yBounds, "..")

		yLow, _ := strconv.Atoi(yBoundList[0])
		yHigh, _ := strconv.Atoi(yBoundList[1])

		newCuboid.Y = []int{yLow, yHigh}

		// Add Z Range
		zBounds := xyz[0][2:]
		zBoundList := strings.Split(zBounds, "..")

		zLow, _ := strconv.Atoi(zBoundList[0])
		zHigh, _ := strconv.Atoi(zBoundList[1])

		newCuboid.Z = []int{zLow, zHigh}

		cuboids = append(cuboids, newCuboid)
	}

	fmt.Println(commands)
	fmt.Println(cuboids)

	fmt.Printf("Number of Cubes ON is: %d\n", calculateNumberOfCubes(commands, cuboids))
}
