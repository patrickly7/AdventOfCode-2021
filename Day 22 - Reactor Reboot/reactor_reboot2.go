package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	on     bool
	cuboid Cuboid
}

type Cuboid struct {
	fromXYZ []int
	toXYZ   []int
}

func (cuboid Cuboid) calculateVolume() int {
	return (cuboid.toXYZ[0] - cuboid.fromXYZ[0] + 1) *
		(cuboid.toXYZ[1] - cuboid.fromXYZ[1] + 1) *
		(cuboid.toXYZ[2] - cuboid.fromXYZ[2] + 1)
}

type BoundingCube struct {
	on     bool
	cuboid Cuboid
}

func (boundingCube BoundingCube) getIntersectionOfCuboids(otherCuboid Cuboid) Cuboid {
	fromX := getMax(boundingCube.cuboid.fromXYZ[0], otherCuboid.fromXYZ[0])
	fromY := getMax(boundingCube.cuboid.fromXYZ[1], otherCuboid.fromXYZ[1])
	fromZ := getMax(boundingCube.cuboid.fromXYZ[2], otherCuboid.fromXYZ[2])

	newFromXYZ := []int{fromX, fromY, fromZ}

	toX := getMin(boundingCube.cuboid.toXYZ[0], otherCuboid.toXYZ[0])
	toY := getMin(boundingCube.cuboid.toXYZ[1], otherCuboid.toXYZ[1])
	toZ := getMin(boundingCube.cuboid.toXYZ[2], otherCuboid.toXYZ[2])

	newToXYZ := []int{toX, toY, toZ}

	if fromX > toX || fromY > toY || fromZ > toZ {
		emptyCuboid := Cuboid{fromXYZ: nil, toXYZ: nil}
		return emptyCuboid
	}

	return Cuboid{newFromXYZ, newToXYZ}
}

func calculateNumberOfLitBoxes(commands []Command) int {
	boundingCubes := []BoundingCube{}

	isFirstCommand := true
	for _, command := range commands {
		if isFirstCommand {
			boundingCubes = append(boundingCubes, BoundingCube{true, command.cuboid})
			isFirstCommand = false
			continue
		}

		// Check Intersections with Existing Cubes
		for _, cube := range boundingCubes {
			intersection := BoundingCube{!cube.on, cube.getIntersectionOfCuboids(command.cuboid)}

			if intersection.cuboid.fromXYZ != nil {
				boundingCubes = append(boundingCubes, intersection)
			}

		}

		if command.on {
			boundingCubes = append(boundingCubes, BoundingCube{true, command.cuboid})
		}
	}

	numberOfLitBoxes := 0

	for _, cube := range boundingCubes {
		if cube.on {
			numberOfLitBoxes += cube.cuboid.calculateVolume()
		} else {
			numberOfLitBoxes -= cube.cuboid.calculateVolume()
		}
	}

	return numberOfLitBoxes
}

func getMin(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func getMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer file.Close()

	var scanner = bufio.NewScanner(file)

	var commands []Command

	for scanner.Scan() {
		splitString := strings.Fields(scanner.Text())

		command := Command{}
		if splitString[0] == "on" {
			command.on = true
		}

		var newCuboid Cuboid
		xyz := strings.Split(splitString[1], ",")

		// Add X Range
		xBounds := xyz[0][2:]
		xBoundList := strings.Split(xBounds, "..")

		xLow, _ := strconv.Atoi(xBoundList[0])
		xHigh, _ := strconv.Atoi(xBoundList[1])

		newCuboid.fromXYZ = append(newCuboid.fromXYZ, xLow)
		newCuboid.toXYZ = append(newCuboid.toXYZ, xHigh)

		// Add Y Range
		yBounds := xyz[1][2:]
		yBoundList := strings.Split(yBounds, "..")

		yLow, _ := strconv.Atoi(yBoundList[0])
		yHigh, _ := strconv.Atoi(yBoundList[1])

		newCuboid.fromXYZ = append(newCuboid.fromXYZ, yLow)
		newCuboid.toXYZ = append(newCuboid.toXYZ, yHigh)

		// Add Z Range
		zBounds := xyz[2][2:]
		zBoundList := strings.Split(zBounds, "..")

		zLow, _ := strconv.Atoi(zBoundList[0])
		zHigh, _ := strconv.Atoi(zBoundList[1])

		newCuboid.fromXYZ = append(newCuboid.fromXYZ, zLow)
		newCuboid.toXYZ = append(newCuboid.toXYZ, zHigh)

		command.cuboid = newCuboid

		commands = append(commands, command)
	}

	//fmt.Println(commands)

	fmt.Println("Number of Cubes that are ON:", calculateNumberOfLitBoxes(commands))
}
