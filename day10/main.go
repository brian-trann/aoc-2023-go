package main

import (
	"fmt"

	"github.com/brian-trann/aoc-2023-go/utils"
)

func main() {
	contents := utils.OpenFileTo2dArray("./input.txt")
	for _, row := range contents {
		for _, char := range row {
			fmt.Printf("%c ", char)
		}
		fmt.Println()
	}
	start := getStart(contents)
	fmt.Println("START: ", start)
	loop := findLoop(start, contents)
	output := len(loop) / 2
	fmt.Println(output)
}

type Coordinate struct {
	i, j int
}
type Direction Coordinate

// up down left right
var (
	dirNorth = Direction{-1, 0}
	dirSouth = Direction{1, 0}
	dirWest  = Direction{0, -1}
	dirEast  = Direction{0, 1}
)
var pipes = map[rune]map[Direction]Direction{
	'|': {
		dirNorth: dirNorth,
		dirSouth: dirSouth,
	},
	'-': {
		dirEast: dirEast,
		dirWest: dirWest,
	},
	'L': {
		dirSouth: dirEast,
		dirWest:  dirNorth,
	},
	'J': {
		dirEast:  dirNorth,
		dirSouth: dirWest,
	},
	'7': {
		dirEast:  dirSouth,
		dirNorth: dirWest,
	},
	'F': {
		dirNorth: dirEast,
		dirWest:  dirSouth,
	},
}

func getStart(grid [][]rune) Coordinate {
	var coord Coordinate
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 'S' {
				coord.i = i
				coord.j = j
				return coord
			}
		}
	}
	panic("did not find start")
}
func findLoop(start Coordinate, grid [][]rune) []Coordinate {
	for _, pipe := range "|-LJ7F" {
		// place each pipe character at the start position in the grid
		grid[start.i][start.j] = rune(pipe)
		// we loopin?
		loop := checkLoop(start, grid)
		if loop != nil {
			return loop
		}
	}
	panic("none found")
}

// check if placing a specific pipe char makes a loop
func checkLoop(start Coordinate, grid [][]rune) []Coordinate {
	curr := start
	direction := anyKey(pipes[grid[start.i][start.j]])
	// we need to store the actual path of the loop
	result := []Coordinate{}

	for {
		// add current position to result path
		result = append(result, curr)
		// see if new dir based on curr posiiton / direction
		newDir, ok := pipes[grid[curr.i][curr.j]][direction]

		if !ok {
			return nil
		}
		// calculate the next coordinate based on the current position and new direction
		newCoord := Coordinate{curr.i + newDir.i, curr.j + newDir.j}
		if newCoord.i < 0 || newCoord.i >= len(grid) || newCoord.j < 0 || newCoord.j >= len(grid[newCoord.i]) {
			// outside boundary ; noo loop
			return nil
		}
		//check if the new coordinate is the starting point, indicating a loop
		if newCoord == start {
			if _, ok := pipes[grid[start.i][start.j]][newDir]; !ok {
				return nil
			}
			// loop found
			break
		}
		curr = newCoord
		direction = newDir
	}
	return result
}

// get first key from map of direction
func anyKey(m map[Direction]Direction) Direction {
	for k := range m {
		return k
	}

	panic("empty")
}
