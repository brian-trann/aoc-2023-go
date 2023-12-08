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

	nodeReg := regexp.MustCompile(`([A-Z][A-Z][A-Z]) = \(([A-Z][A-Z][A-Z]), ([A-Z][A-Z][A-Z])\)`)
	count := 0
	idx := 0
	gps := make(map[string]struct{ Left, Right string })
	rows := strings.Split(navigation, "\n")
	for _, row := range rows {
		if len(row) == 0 {
			continue

		}
		nodes := nodeReg.FindStringSubmatch((row))
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
}
