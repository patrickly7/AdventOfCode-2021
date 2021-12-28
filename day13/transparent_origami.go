package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func setUpTransparentPaper(length int, height int, coordinates [][2]int) [][]string {
	var transparentPaper [][]string

	for x := 0; x < length; x++ {
		var row []string
		for y := 0; y < height; y++ {
			row = append(row, ".")
		}
		transparentPaper = append(transparentPaper, row)
	}

	for _, coordinate := range coordinates {
		transparentPaper[coordinate[0]][coordinate[1]] = "#"
	}

	return transparentPaper
}

func foldTransparentPaper(transparentPaper [][]string, length int, height int, foldLines [][]string) ([][]string, int, int) {
	foldLine := foldLines[0]
	foldPoint, _ := strconv.Atoi(foldLine[1])

	fmt.Printf("Fold Line: %v\n", foldLine)

	var foldedPaper [][]string
	var foldedLength int
	var foldedHeight int

	if foldLine[0] == "x" { // Vertical Folds
		foldedPaper, foldedLength, foldedHeight = foldVertically(transparentPaper, length, height, foldPoint)
	} else if foldLine[0] == "y" { // Horizontal Folds
		foldedPaper, foldedLength, foldedHeight = foldHorizontally(transparentPaper, length, height, foldPoint)
	}

	//viewTransparentPaper(foldedPaper, foldedLength, foldedHeight)

	return foldedPaper, foldedLength, foldedHeight
}

func foldHorizontally(transparentPaper [][]string, length int, height int, foldPoint int) ([][]string, int, int) {
	var foldedPaper [][]string

	for x := 0; x < foldPoint; x++ {
		var row []string
		for y := 0; y < height; y++ {
			difference := int(math.Abs(float64(x) - float64(foldPoint)))
			respectiveX := foldPoint + difference

			if respectiveX < length && transparentPaper[respectiveX][y] == "#" {
				row = append(row, "#")
			} else {
				row = append(row, transparentPaper[x][y])
			}
		}
		foldedPaper = append(foldedPaper, row)
	}

	for x := foldPoint + foldPoint; x < length; x++ {
		var row []string
		for y := 0; y < height; y++ {
			row = append([]string{transparentPaper[x][y]}, row...)
		}
		foldedPaper = append(foldedPaper, row)
	}

	return foldedPaper, len(foldedPaper) - 1, height
}

func foldVertically(transparentPaper [][]string, length int, height int, foldPoint int) ([][]string, int, int) {
	var foldedPaper [][]string

	for x := 0; x < length; x++ {
		var row []string
		for y := 0; y < foldPoint; y++ {
			difference := int(math.Abs(float64(y) - float64(foldPoint)))
			respectiveY := foldPoint + difference

			if respectiveY < height && transparentPaper[x][foldPoint+difference] == "#" {
				row = append(row, "#")
			} else {
				row = append(row, transparentPaper[x][y])
			}
		}
		foldedPaper = append(foldedPaper, row)
	}

	for x := 0; x < length; x++ {
		for y := foldPoint + foldPoint; y < height; y++ {
			foldedPaper[x] = append([]string{transparentPaper[x][y]}, foldedPaper[x]...)
		}
	}

	return foldedPaper, length, len(foldedPaper[0])
}

func calculateDotsOnTransparentPaper(transparentPaper [][]string, length int, height int) int {
	numOfDots := 0

	for x := 0; x < length; x++ {
		for y := 0; y < height; y++ {
			if transparentPaper[x][y] == "#" {
				numOfDots++
			}
		}
	}

	return numOfDots
}

func viewTransparentPaper(transparentPaper [][]string, length int, height int) {
	for x := 0; x < length; x++ {
		for y := 0; y < height; y++ {
			fmt.Printf(transparentPaper[x][y])
		}

		fmt.Printf("\n")
	}

	fmt.Println("========")
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer file.Close()

	var scanner = bufio.NewScanner(file)
	largestX := 0
	largestY := 0

	var coordinates [][2]int
	var foldLines [][]string

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		splitText := strings.Fields(scanner.Text())

		if splitText[0] == "fold" {
			foldLine := strings.Split(splitText[2], "=")
			foldLines = append(foldLines, foldLine)

			continue
		}

		coordinateString := strings.Split(scanner.Text(), ",")

		x, _ := strconv.Atoi(coordinateString[1])
		if x > largestX {
			largestX = x
		}

		y, _ := strconv.Atoi(coordinateString[0])
		if y > largestY {
			largestY = y
		}

		coordinate := [2]int{x, y}
		coordinates = append(coordinates, coordinate)
	}

	largestX += 1
	largestY += 1

	fmt.Printf("Length: %d, Height: %d\n", largestX, largestY)

	transparentPaper := setUpTransparentPaper(largestX, largestY, coordinates)

	//fmt.Println("Starting Paper")
	//viewTransparentPaper(transparentPaper, largestX, largestY)
	foldedPaper, foldedLength, foldedHeight := foldTransparentPaper(transparentPaper, largestX, largestY, foldLines)

	fmt.Printf("The number of dots on the paper (folded once) is: %d\n", calculateDotsOnTransparentPaper(foldedPaper, foldedLength, foldedHeight))
}
