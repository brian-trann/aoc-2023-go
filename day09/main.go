package main

import (
	"fmt"

	"github.com/brian-trann/aoc-2023-go/utils"
)

func main() {
	input := utils.OpenFileSliceStrings("./input.txt")
	// an arr to store all histories (sequences of integers).
	histories := make([][]int, 0)
	for _, row := range input {
		if len(row) == 0 {
			continue
		}
		rowAsNums := utils.ConvertStringToIntSlice(row, " ")

		histories = append(histories, rowAsNums)
	}
	// store the ending values of the upside down triangle
	listOfEndingVals := make([]int, 0)
	for _, history := range histories {
		// start with the original history as the first sublist.
		sublists := [][]int{history}

		allZeros := false
		// make sublists until a sublist of all zeroes is made
		for !allZeros {
			sublist := make([]int, 0)
			for i := 0; i < len(sublists[len(sublists)-1])-1; i++ {
				// calc the diff between curr/next elements in the last sublist
				diff := sublists[len(sublists)-1][i+1] - sublists[len(sublists)-1][i]
				sublist = append(sublist, diff)
			}
			allZeros = checkAllZerosSlice(sublist)
			// add the new sublist to the list of sublists
			sublists = append(sublists, sublist)
		}
		// append 0 to the last sublist which should be all zeros
		sublists[len(sublists)-1] = append(sublists[len(sublists)-1], 0)
		// backtrack  through the sublists and calc the final value
		for i := len(sublists) - 2; i >= 0; i-- {
			// this adds the ending values of each sublist and the sublist prior...
			extrapolated := sublists[i][len(sublists[i])-1] + sublists[i+1][len(sublists[i+1])-1]
			sublists[i] = append(sublists[i], extrapolated)
		}
		// the final extrapolated value for this history to the arr of extrapolated values
		listOfEndingVals = append(listOfEndingVals, sublists[0][len(sublists[0])-1])
	}
	partOne := sum(listOfEndingVals)
	fmt.Println(partOne)
}
func checkAllZerosSlice(nums []int) bool {
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	return true
}
func sum(nums []int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}
