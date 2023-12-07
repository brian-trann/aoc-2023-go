package main

import (
	"fmt"
	"strings"

	"github.com/brian-trann/aoc-2023-go/utils"
)

func main() {
	contents := utils.OpenFileAsString("./input.txt")
	// fmt.Println(contents)
	lines := strings.Split(contents, "\n")
	var timeTable [][]int
	var timeTablePartTwo [][]int

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		startIndex := strings.Index(line, ":")
		timeRowStr := strings.TrimSpace(line[startIndex+1:])
		timeRowStr2 := strings.ReplaceAll(timeRowStr, " ", "")
		timeRow := utils.ConvertStringToIntSlice(timeRowStr, " ")
		timeRow2 := utils.ConvertStringToIntSlice(timeRowStr2, " ")
		timeTable = append(timeTable, timeRow)
		timeTablePartTwo = append(timeTablePartTwo, timeRow2)
	}
	partOneAnswer := processTimeTable(timeTable)
	partTwoAnswer := processTimeTable(timeTablePartTwo)

	fmt.Println("p1:", partOneAnswer)
	fmt.Println("p2:", partTwoAnswer)

}
func isValid(input int, time int, distance int) bool {
	tempDist := time - input
	distanceTraveled := tempDist * input
	return distanceTraveled > distance
}
func processTimeTable(table [][]int) int {

	var allValues []int
	for i := range table[0] {

		// this block is each race
		raceNumber := i + 1
		time := table[0][i]
		distance := table[1][i]
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
	return product
}
