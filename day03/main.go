package main

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/brian-trann/aoc-2023-go/utils"
)

func main() {
	twoDArray := utils.OpenFileTo2dArray("./input.txt")

	for _, row := range twoDArray {
		for _, char := range row {
			fmt.Printf("%c ", char)
		}
		fmt.Println()
	}
	numbersWithNeighbors := ExtractNumbersWithNeighbors(twoDArray)
	for _, item := range numbersWithNeighbors {
		fmt.Printf("num: %d, neighbors: %c\n", item.Number, item.Neighbors)
	}
	filteredNumbers := FilterNumbersWithNeighbors(numbersWithNeighbors)
	total := 0
	for _, n := range filteredNumbers {
		total += n.Number
	}
	fmt.Println(total)
}

type NumberWithNeighbors struct {
	Number    int
	Neighbors []rune
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
func ExtractNumbersWithNeighbors(array [][]rune) []NumberWithNeighbors {
	var results []NumberWithNeighbors

	for i, row := range array {
		var numStr string
		var numStartCol int
		// build number, get neighbor info
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
						results = append(results, NumberWithNeighbors{Number: num, Neighbors: neighbors})
					}
				}
				//reset for next
				numStr = ""

			}
		}

		// vheck for a number at the end of the row
		if numStr != "" {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("error converting to num")
			} else {
				neighbors := GetNeighbors(array, i, numStartCol, len(numStr))
				results = append(results, NumberWithNeighbors{Number: num, Neighbors: neighbors})
			}
		}
	}

	return results
}
func FilterNumbersWithNeighbors(numbers []NumberWithNeighbors) []NumberWithNeighbors {
	var filtered []NumberWithNeighbors
	for _, number := range numbers {
		if len(number.Neighbors) > 0 {
			filtered = append(filtered, number)
		}
	}
	return filtered
}
