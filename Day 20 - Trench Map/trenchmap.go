package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func calculateNumberofLitPixels(enhancedImage [][]string, imageEnhancementAlgorithm []string, steps int) int {
	currentImage := enhancedImage
	currStep := 0

	for currStep < steps {
		newImage := createEmptyImage(len(currentImage))

		for row := 1; row < len(currentImage)-1; row++ {
			for col := 1; col < len(currentImage)-1; col++ {
				binaryString := ""

				// Get Binary Representation of Current Pixel
				for imgRow := row - 1; imgRow <= row+1; imgRow++ {
					for imgCol := col - 1; imgCol <= col+1; imgCol++ {
						if currentImage[imgRow][imgCol] == "." {
							binaryString += "0"
						} else {
							binaryString += "1"
						}
					}
				}

				// Get the Corresponding Pixel Result from Algorithm
				decimalValue, _ := strconv.ParseInt(binaryString, 2, 64)
				newImage[row][col] = imageEnhancementAlgorithm[decimalValue]
			}
		}

		fmt.Println("=======")
		displayImage(newImage)

		currentImage = newImage
		currStep++
	}

	litPixelCount := 0
	notCountedPixels := 5 - steps
	for row := notCountedPixels; row < len(currentImage)-notCountedPixels; row++ {
		for col := notCountedPixels; col < len(currentImage)-notCountedPixels; col++ {
			if currentImage[row][col] == "#" {
				litPixelCount++
			}
		}
	}

	return litPixelCount
}

func createEmptyImage(imageSize int) [][]string {
	image := [][]string{}

	for i := 0; i < imageSize; i++ {
		imageRow := []string{}
		for j := 0; j < imageSize; j++ {
			imageRow = append(imageRow, ".")
		}
		image = append(image, imageRow)
	}

	return image
}

func displayImage(image [][]string) {
	for row := 0; row < len(image); row++ {
		for col := 0; col < len(image); col++ {
			fmt.Printf(image[row][col])
		}
		fmt.Printf("\n")
	}
}

func stringToList(str string) []string {
	list := []string{}

	for _, char := range str {
		list = append(list, string(char))
	}

	return list
}

func stringToListWithPadding(str string, padding int) []string {
	list := []string{}

	for i := 0; i < padding; i++ {
		list = append(list, ".")
	}

	for _, char := range str {
		list = append(list, string(char))
	}

	for i := 0; i < padding; i++ {
		list = append(list, ".")
	}

	return list
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	imageEnhancementAlgorithm := []string{}
	enhancedImage := [][]string{}
	imageSize := 0

	isFirstLine := true
	isFirstImageRow := true
	for scanner.Scan() {
		if isFirstLine {
			imageEnhancementAlgorithm = stringToList(scanner.Text())
			isFirstLine = false
			continue
		}

		if scanner.Text() == "" {
			continue
		}

		if isFirstImageRow {
			imageSize = len(scanner.Text())

			// Add Top Padding to Image
			for i := 0; i < 5; i++ {
				imagePadding := []string{}
				for pixel := 0; pixel < imageSize+10; pixel++ {
					imagePadding = append(imagePadding, ".")
				}
				enhancedImage = append(enhancedImage, imagePadding)
			}

			isFirstImageRow = false
		}

		// Add Rows of the Image with Padding
		imageRow := stringToListWithPadding(scanner.Text(), 5)
		enhancedImage = append(enhancedImage, imageRow)
	}

	// Add Bottom Padding to Image
	for i := 0; i < 5; i++ {
		imagePadding := []string{}
		for pixel := 0; pixel < imageSize+10; pixel++ {
			imagePadding = append(imagePadding, ".")
		}
		enhancedImage = append(enhancedImage, imagePadding)
	}

	//fmt.Printf("%v\n", enhancedImage)
	displayImage(enhancedImage)

	fmt.Printf("The number of lit pixels is: %d\n", calculateNumberofLitPixels(enhancedImage, imageEnhancementAlgorithm, 2))
}
