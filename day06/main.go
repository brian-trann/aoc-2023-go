package main

import (
	"fmt"
	"strings"

	"github.com/brian-trann/aoc-2023-go/utils"
)

func main() {
	// partOne()
	partTwo()

}
func isValid(input int, time int, distance int) bool {
	tempDist := time - input
	distanceTraveled := tempDist * input
	return distanceTraveled > distance
}
func partOne() {
	contents := utils.OpenFileAsString("./input.txt")
	// fmt.Println(contents)
	lines := strings.Split(contents, "\n")
	var timeTable [][]int

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		startIndex := strings.Index(line, ":")
		timeRowStr := strings.TrimSpace(line[startIndex+1:])
		timeRow := utils.ConvertStringToIntSlice(timeRowStr, " ")
		timeTable = append(timeTable, timeRow)
	}
	// fmt.Println(timeTable)
	var allValues []int
	for i := range timeTable[0] {

		// fmt.Println(timeTable[0][i])
		// this block is each race
		raceNumber := i + 1
		time := timeTable[0][i]
		distance := timeTable[1][i]
		fmt.Printf("Race: %d | time: %d | distance: %d\n\n", raceNumber, time, distance)
		validOpts := 0
		for j := 1; j < time; j++ {
			if isValid(j, time, distance) {
				validOpts++
			}
		}
		fmt.Println("ValidOpts: ", validOpts)
		if validOpts > 0 {
			allValues = append(allValues, validOpts)
		}
	}
	product := utils.ProductOfArray(allValues)
	fmt.Println("part one answer", product)
}
func partTwo() {
	contents := utils.OpenFileAsString("./input.txt")
	// fmt.Println(contents)
	lines := strings.Split(contents, "\n")
	var timeTable [][]int

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		startIndex := strings.Index(line, ":")
		timeRowStr := strings.TrimSpace(line[startIndex+1:])
		newTimeRow := strings.ReplaceAll(timeRowStr, " ", "")
		timeRow := utils.ConvertStringToIntSlice(newTimeRow, " ")
		timeTable = append(timeTable, timeRow)
	}
	// fmt.Println(timeTable)
	var allValues []int
	for i := range timeTable[0] {

		// fmt.Println(timeTable[0][i])
		// this block is each race
		raceNumber := i + 1
		time := timeTable[0][i]
		distance := timeTable[1][i]
		fmt.Printf("Race: %d | time: %d | distance: %d\n\n", raceNumber, time, distance)
		validOpts := 0
		for j := 1; j < time; j++ {
			if isValid(j, time, distance) {
				validOpts++
			}
		}
		fmt.Println("ValidOpts: ", validOpts)
		if validOpts > 0 {
			allValues = append(allValues, validOpts)
		}
	}
	product := utils.ProductOfArray(allValues)
	fmt.Println("part two answer", product)
}
