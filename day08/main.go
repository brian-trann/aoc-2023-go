package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/brian-trann/aoc-2023-go/utils"
)

func main() {

	input := utils.OpenFileAsString("./input.txt")
	sections := strings.Split(input, "\n\n")
	instructions := sections[0]
	instructions = strings.TrimSpace(instructions)
	fmt.Println("INSTRUCTIONS: ", instructions)
	navigation := sections[1]
	destination := "AAA"

	// fmt.Println("\n===NAVIGATION===")
	// fmt.Println(navigation)

	nodeReg := regexp.MustCompile(`([A-Z0-9][A-Z0-9][A-Z0-9]) = \(([A-Z0-9][A-Z0-9][A-Z0-9]), ([A-Z0-9][A-Z0-9][A-Z0-9])\)`)
	count := 0
	idx := 0
	gps := make(map[string]struct{ Left, Right string })
	rows := strings.Split(navigation, "\n")
	for _, row := range rows {
		if len(row) == 0 {
			continue
		}
		nodes := nodeReg.FindStringSubmatch((row))
		if len(nodes) == 0 {
			continue
		}

		src := nodes[1]
		left := nodes[2]
		right := nodes[3]
		gps[src] = struct {
			Left  string
			Right string
		}{Left: left, Right: right}
	}

	for destination != "ZZZ" {
		turn := instructions[idx]
		switch turn {
		case 'L':
			destination = gps[destination].Left
		case 'R':
			destination = gps[destination].Right
		default:
			panic("something went wrong and ran into a tree")
		}
		count++
		idx = (idx + 1) % len(instructions)
	}
	fmt.Println("part one: ", count)
	var iterArray []int

	for _, row := range rows {
		if len(row) == 0 {
			continue
		}
		nodes := nodeReg.FindStringSubmatch((row))

		if len(nodes) == 0 {
			continue
		}
		src := nodes[1]
		copy := src
		if src[2] == 'A' {
			// how many iterations does that singular node need to get to the node ending in Z
			iterCount := 0
			instructionsIdx := 0
			for copy[2] != 'Z' {
				turn := instructions[instructionsIdx]

				switch turn {
				case 'L':
					copy = gps[copy].Left
				case 'R':
					copy = gps[copy].Right
				default:
					panic("something went wrong and ran into a tree")
				}
				iterCount++
				instructionsIdx = (instructionsIdx + 1) % len(instructions)
			}
			iterArray = append(iterArray, iterCount)

		}
	}
	partTwo := lcmArray(iterArray)
	fmt.Println("Part Two:", partTwo)
	// make the gps again; with all of the nodes

	// find all nodes that ends with A (6) done-
	// store in []

	// for node in nodes

	// call LCM with that array
}

// from https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
func lcmArray(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	result := arr[0]
	for _, num := range arr[1:] {
		result = LCM(result, num)
	}
	return result
}
