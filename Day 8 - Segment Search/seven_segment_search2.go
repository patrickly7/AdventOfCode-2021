package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Segment struct {
	A, B, C, D, E, F, G int
}

func convertStringToSegment(segmentCode string) Segment {
	var segment Segment

	for _, char := range segmentCode {
		str := string(char)

		if str == "a" {
			segment.A = 1
		} else if str == "b" {
			segment.B = 1
		} else if str == "c" {
			segment.C = 1
		} else if str == "d" {
			segment.D = 1
		} else if str == "e" {
			segment.E = 1
		} else if str == "f" {
			segment.F = 1
		} else if str == "g" {
			segment.G = 1
		}
	}

	return segment
}

func isSegmentInSegment(segment1 Segment, segment2 Segment) bool {
	if segment1.A == 1 {
		if segment2.A != 1 {
			return false
		}
	}

	if segment1.B == 1 {
		if segment2.B != 1 {
			return false
		}
	}

	if segment1.C != segment2.C {
		if segment2.C != 1 {
			return false
		}
	}

	if segment1.D != segment2.D {
		if segment2.D != 1 {
			return false
		}
	}

	if segment1.E != segment2.E {
		if segment2.E != 1 {
			return false
		}
	}

	if segment1.F != segment2.F {
		if segment2.F != 1 {
			return false
		}
	}

	if segment1.G == 1 {
		if segment2.G != 1 {
			return false
		}
	}

	return true
}

// func removeSegments(segment1 Segment, segment2 Segment) Segment {
// 	newSegment := segment2

// 	if segment1.A == 1 {
// 		if segment2.A == 1 {
// 			newSegment.A = 0
// 		}
// 	}

// 	if segment1.B == 1 {
// 		if segment2.B == 1 {
// 			newSegment.B = 0
// 		}
// 	}

// 	if segment1.C == 1 {
// 		if segment2.B == 1 {
// 			newSegment.B = 0
// 		}
// 	}

// 	if segment1.D != segment2.D {
// 		if segment2.A == 1 {
// 			newSegment.A = 0
// 		}
// 	}

// 	if segment1.E != segment2.E {
// 		if segment2.A == 1 {
// 			newSegment.A = 0
// 		}
// 	}

// 	if segment1.F != segment2.F {
// 		if segment2.A == 1 {
// 			newSegment.A = 0
// 		}
// 	}

// 	if segment1.G == 1 {
// 		if segment2.A == 1 {
// 			newSegment.A = 0
// 		}
// 	}

// 	return newSegment
// }

func intersectsThreeTimes(segment1 Segment, segment2 Segment) bool {
	intersections := 0

	if segment1.A == 1 {
		if segment2.A == 1 {
			intersections++
		}
	}

	if segment1.B == 1 {
		if segment2.B == 1 {
			intersections++
		}
	}

	if segment1.C == 1 {
		if segment2.C == 1 {
			intersections++
		}
	}

	if segment1.D == 1 {
		if segment2.D == 1 {
			intersections++
		}
	}

	if segment1.E == 1 {
		if segment2.E == 1 {
			intersections++
		}
	}

	if segment1.F == 1 {
		if segment2.F == 1 {
			intersections++
		}
	}

	if segment1.G == 1 {
		if segment2.G == 1 {
			intersections++
		}
	}

	return intersections == 3
}

func calculateOutputValueSum(segmentCodesList [][]string, outputValuesList [][]string) int {
	outputValueSum := 0

	for index, outputValue := range outputValuesList {
		currSegmentCodes := segmentCodesList[index]

		patternToSegment := make(map[string]Segment)

		// Get the Simple Segments (1, 4, 7, 8)
		for _, segmentCode := range currSegmentCodes {
			if len(segmentCode) == 2 {
				convertedSegment := convertStringToSegment(segmentCode)
				patternToSegment["1"] = convertedSegment
			} else if len(segmentCode) == 3 {
				convertedSegment := convertStringToSegment(segmentCode)
				patternToSegment["7"] = convertedSegment
			} else if len(segmentCode) == 4 {
				convertedSegment := convertStringToSegment(segmentCode)
				patternToSegment["4"] = convertedSegment
			} else if len(segmentCode) == 7 {
				convertedSegment := convertStringToSegment(segmentCode)
				patternToSegment["8"] = convertedSegment
			}
		}

		// Get the Complex Segments (2, 3, 5 and 0, 6, 9)
		for _, segmentCode := range currSegmentCodes {
			if len(segmentCode) == 5 {
				convertedSegment := convertStringToSegment(segmentCode)
				if isSegmentInSegment(patternToSegment["7"], convertedSegment) {
					patternToSegment["3"] = convertedSegment
				} else {
					if intersectsThreeTimes(patternToSegment["4"], convertedSegment) {
						patternToSegment["5"] = convertedSegment
					} else {
						patternToSegment["2"] = convertedSegment
					}
				}
			} else if len(segmentCode) == 6 {
				convertedSegment := convertStringToSegment(segmentCode)

				if isSegmentInSegment(patternToSegment["4"], convertedSegment) {
					patternToSegment["9"] = convertedSegment
				} else if isSegmentInSegment(patternToSegment["7"], convertedSegment) {
					patternToSegment["0"] = convertedSegment
				} else {
					patternToSegment["6"] = convertedSegment
				}
			}
		}

		decodedOutput := ""
		for _, outputValueDigit := range outputValue {
			convertedSegment := convertStringToSegment(outputValueDigit)
			for number, segment := range patternToSegment {
				if segment == convertedSegment {
					decodedOutput += number
				}
			}
		}
		fmt.Println(decodedOutput)

		finalOutputValue, _ := strconv.Atoi(decodedOutput)
		outputValueSum += finalOutputValue
	}

	return outputValueSum
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer file.Close()

	var scanner = bufio.NewScanner(file)

	var segmentCodesList [][]string
	var outputValuesList [][]string

	for scanner.Scan() {
		var splitString = strings.Fields(scanner.Text())

		var segmentCodes []string
		for i := 0; i < 10; i++ {
			segmentCodes = append(segmentCodes, splitString[i])
		}

		var outputValues []string
		for i := 11; i < 15; i++ {
			outputValues = append(outputValues, splitString[i])
		}

		segmentCodesList = append(segmentCodesList, segmentCodes)
		outputValuesList = append(outputValuesList, outputValues)
	}

	fmt.Println(segmentCodesList)
	fmt.Println(outputValuesList)

	fmt.Printf("The total output value sum is: %d\n", calculateOutputValueSum(segmentCodesList, outputValuesList))
}
