package main

import (
	"fmt"

	"unicode"

	"github.com/brian-trann/aoc-2023-go/utils"
)

func main() {
	PartOne()
	PartTwo()
}

func PartOne() {
	lines := utils.OpenFile("./input.txt")
	var result []string
	for _, line := range lines {
		first := byte('0')
		last := byte('0')
		found := false
		for i := 0; i < len(line); i++ {

			if unicode.IsDigit(rune(line[i])) {
				if !found {
					first = line[i]
					found = true
				}
				last = line[i]

			}

		}
		if found {
			result = append(result, string([]byte{first, last}))
		} else {
			// if we only found one digit; we should  make it both first and last
			result = append(result, string([]byte{first, first}))
		}

	}

	sum, err := utils.SumStringNumbers(result)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Part one Sum:", sum)
	}
}

func PartTwo() {
	panic("unimplemented cuz it is hard")
	// lines := utils.OpenFile("./input.txt")

	// var result []string
	// this edgecase has been hard to parse; maybe i can revisit

	// https://www.reddit.com/r/adventofcode/comments/1884fpl/2023_day_1for_those_who_stuck_on_part_2/
	// eighthree => 83
	// sevenine => 79

	// numWords := map[string]byte{
	// 	"one":   '1',
	// 	"two":   '2',
	// 	"three": '3',
	// 	"four":  '4',
	// 	"five":  '5',
	// 	"six":   '6',
	// 	"seven": '7',
	// 	"eight": '8',
	// 	"nine":  '9',
	// }

	// two1nine => 29
	// eightwothree => 83
	// abcone2threexyz => 13
	// xtwone3four => 24
	// 4nineeightseven2 => 42
	// zoneight234 => 14
	// 7pqrstsixteen => 76

	// sum, err := SumStringNumbers(result)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println("Part two Sum:", sum)
	// }
}
