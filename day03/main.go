package main

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/brian-trann/aoc-2023-go/utils"
)

func main() {
	twoDArray := utils.OpenFileTo2dArray("./input.txt")

	// for _, row := range twoDArray {
	// 	for _, char := range row {
	// 		fmt.Printf("%c ", char)
	// 	}
	// 	fmt.Println()
	// }
	// numbersWithNeighbors := ExtractNumbersWithNeighbors(twoDArray)
	// for _, item := range numbersWithNeighbors {
	// 	fmt.Printf("num: %d, neighbors: %c\n", item.Number, item.Neighbors)
	// }

	numbers := ExtractNumbers(twoDArray)
	filteredNumbers := FilterNumbersWithNeighbors(numbers)
	total := 0
	for _, n := range filteredNumbers {
		total += n.Number
	}
	fmt.Printf("part 1: %d\n\n", total)
	specialCharNeighbors := FindSpecialCharNeighbors(twoDArray, numbers)
	// for _, sc := range specialCharNeighbors {
	// 	fmt.Printf("char: %c, num neighbors: %v\n", sc.Char, sc.Neighbors)
	// }
	var filterdSpecialChars []SpecialCharWithNeighbors
	for _, item := range specialCharNeighbors {
		if item.Char == '*' && len(item.Neighbors) == 2 {
			filterdSpecialChars = append(filterdSpecialChars, item)
		}
	}
	gearRatio := 0
	for _, item := range filterdSpecialChars {
		product := productOfArray(item.Neighbors)
		gearRatio += product
	}
	fmt.Printf("part 2: %d\n\n", gearRatio)
}

// 793992
type NumberWithNeighbors struct {
	Number    int
	Neighbors []rune
}
type NumberWithInfo struct {
	Number    int
	Neighbors []rune
	Row       int
	Col       int
}

func GetNeighbors(array [][]rune, row, col, length int) []rune {
	var neighbors []rune
	rows := len(array)
	cols := len(array[0])

	for offset := 0; offset < length; offset++ {
		currentCol := col + offset

		// index offset to check around a number
		checkPositions := []struct {
			dx int
			dy int
		}{
			{0, -1}, {0, 1}, // left/right
			{-1, -1}, {-1, 0}, {-1, 1}, //upper
			{1, -1}, {1, 0}, {1, 1}, // lower
		}

		// for intermediate digits, only check left and right
		if offset > 0 && offset < length-1 {
			checkPositions = []struct {
				dx int
				dy int
			}{
				{0, -1}, {0, 1}, // Left and right
			}
		}

		// check each position for neighbors
		// there seems to be a bug somewhere here where i am double counting a neighbor
		// 35, has 2 neighbors **, while 664 and 598 and others. but i think its fine; since we arent counting neighbors
		for _, pos := range checkPositions {
			newRow, newCol := row+pos.dx, currentCol+pos.dy
			// only if the offsets are valid
			if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols {
				neighbor := array[newRow][newCol]
				// only if nondigit or `.`
				if !unicode.IsDigit(neighbor) && neighbor != '.' {
					neighbors = append(neighbors, neighbor)
				}
			}
		}
	}

	return neighbors
}
func ExtractNumbers(array [][]rune) []NumberWithInfo {
	var results []NumberWithInfo

	for i, row := range array {
		var numStr string
		var numStartCol int
		for j, char := range row {
			if unicode.IsDigit(char) {
				if numStr == "" {
					numStartCol = j
				}
				numStr += string(char)
			} else {
				if numStr != "" {
					num, err := strconv.Atoi(numStr)
					if err != nil {
						fmt.Println("error converting to num")
					} else {
						neighbors := GetNeighbors(array, i, numStartCol, len(numStr))
						results = append(results, NumberWithInfo{Number: num, Row: i, Col: numStartCol, Neighbors: neighbors})
					}
				}
				//reset for next
				numStr = ""
			}
		}
		// check for a number at the end of the row
		if numStr != "" {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("error converting to num")
			} else {
				neighbors := GetNeighbors(array, i, numStartCol, len(numStr))
				results = append(results, NumberWithInfo{Number: num, Row: i, Col: numStartCol, Neighbors: neighbors})
			}
		}
	}

	return results
}
func FilterNumbersWithNeighbors(numbers []NumberWithInfo) []NumberWithInfo {
	var filtered []NumberWithInfo
	for _, number := range numbers {
		if len(number.Neighbors) > 0 {
			filtered = append(filtered, number)
		}
	}
	return filtered
}
func IsNeighbor(row, col, length, i, j int) bool {
	// check each cell of the number to see if it's adjacent to the special character
	for offset := 0; offset < length; offset++ {
		currentCol := col + offset
		if i >= row-1 && i <= row+1 && j >= currentCol-1 && j <= currentCol+1 {
			return true
		}
	}
	return false
}

type SpecialCharWithNeighbors struct {
	Char      rune
	Neighbors []int
}

// finds special characters in the array and their full number neighbors
func FindSpecialCharNeighbors(array [][]rune, numbers []NumberWithInfo) []SpecialCharWithNeighbors {
	var specialCharNeighbors []SpecialCharWithNeighbors
	for i, row := range array {
		for j, char := range row {
			if !unicode.IsDigit(char) && char != '.' {
				var neighbors []int
				for _, num := range numbers {
					if IsNeighbor(num.Row, num.Col, len(strconv.Itoa(num.Number)), i, j) {
						neighbors = append(neighbors, num.Number)
					}
				}
				specialCharNeighbors = append(specialCharNeighbors, SpecialCharWithNeighbors{Char: char, Neighbors: neighbors})
			}
		}
	}
	return specialCharNeighbors
}
func productOfArray(arr []int) int {
	product := 1
	for _, num := range arr {
		product *= num
	}
	return product
}
