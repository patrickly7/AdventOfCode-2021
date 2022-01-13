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
	foldedPaper := transparentPaper[:]
	foldedLength := length
	foldedHeight := height

	for _, foldLine := range foldLines {
		foldPoint, _ := strconv.Atoi(foldLine[1])

		//fmt.Printf("Fold Line: %v\n", foldLine)

		if foldLine[0] == "x" { // Vertical Folds
			foldedPaper, foldedLength, foldedHeight = foldVertically(foldedPaper, foldedLength, foldedHeight, foldPoint)
		} else if foldLine[0] == "y" { // Horizontal Folds
			foldedPaper, foldedLength, foldedHeight = foldHorizontally(foldedPaper, foldedLength, foldedHeight, foldPoint)
		}

		// if idx == len(foldLines)-2 {
		// 	viewTransparentPaper(foldedPaper, foldedLength, foldedHeight)
		// 	fmt.Println("===========")
		// }
	}

	return foldedPaper, foldedLength, foldedHeight
}

func foldHorizontally(transparentPaper [][]string, length int, height int, foldPoint int) ([][]string, int, int) {
	//fmt.Printf("Len: %d, Hei: %d\n", length, height)
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

	return foldedPaper, length - foldPoint - 1, height
}

func foldVertically(transparentPaper [][]string, length int, height int, foldPoint int) ([][]string, int, int) {
	//fmt.Printf("Len: %d, Hei: %d\n", length, height)
	var foldedPaper [][]string

	for x := 0; x < length; x++ {
		var row []string
		for y := 0; y < foldPoint; y++ {
			difference := int(math.Abs(float64(y) - float64(foldPoint)))
			respectiveY := foldPoint + difference

			if respectiveY < height && transparentPaper[x][respectiveY] == "#" {
				row = append(row, "#")
			} else {
				row = append(row, transparentPaper[x][y])
			}
		}
		foldedPaper = append(foldedPaper, row)
	}

	return foldedPaper, length, height - foldPoint - 1
}

func viewTransparentPaper(transparentPaper [][]string, length int, height int) {
	for x := 0; x < length; x++ {
		for y := 0; y < height; y++ {
			fmt.Printf(transparentPaper[x][y])
		}

		fmt.Printf("\n")
	}
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
	largestY += 4 // "The Magic Number"

	fmt.Printf("Length: %d, Height: %d\n", largestX, largestY)

	transparentPaper := setUpTransparentPaper(largestX, largestY, coordinates)
	foldedPaper, foldedLength, foldedHeight := foldTransparentPaper(transparentPaper, largestX, largestY, foldLines)

	viewTransparentPaper(foldedPaper, foldedLength, foldedHeight)
}
