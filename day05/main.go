package main

import (
	"fmt"
	"strings"

	"github.com/brian-trann/aoc-2023-go/utils"
)

func main() {
	contents := utils.OpenFileAsString("./input.txt")

	sections := strings.Split(contents, "\n\n")

	startIndex := strings.Index(sections[0], ":")
	seedStr := strings.TrimSpace(sections[0][startIndex+1:])
	seeds := utils.ConvertStringToIntSlice(seedStr, " ")

	fmt.Println("SEEDS: ", seeds)

	for i, section := range sections {
		for si, seed := range seeds {
			fmt.Println("CURR_SEED: ", seed)

			if i == 0 {
				continue
			}
			// fmt.Println(section)
			// i1 == seed to siol
			// i2 == seed to fertilizer
			// i3 == fertilizer to water
			// i4 == water to light
			// i5 == light to temp
			// i5 == temp to humidity
			// i6 == humidity to location

			fmt.Println("====GROUP START====")
			text := strings.Split(section, "\n")
			fmt.Println("MAP header", text[0])
			for j, group := range text {
				fmt.Println("GROUP:", group)
				if j == 0 {
					// this is the header of the group
					continue
				}
				ranges := utils.ConvertStringToIntSlice(group, " ")
				fmt.Println("==RANGES--", ranges)
				if len(ranges) == 0 {
					continue
				}

				dest := ranges[0]
				src := ranges[1]
				rangeLength := ranges[2]
				fmt.Printf("dest: %d | src: %d | range %d\n", dest, src, rangeLength)

				if (seed >= src) && (seed <= src+rangeLength-1) {
					fmt.Printf("CurrSeed: %d\n", seed)
					newSeed := seed + (dest - src)
					fmt.Printf("New seed: %d\n\n", newSeed)
					seeds[si] = newSeed
				}
				fmt.Println("    curr seeds: ", seeds)
				// sameseed equals same
			}

			fmt.Println("====GROUP END====")
		}
	}
	lowest := 9999999999999999
	for _, seed := range seeds {
		if seed < lowest {
			lowest = seed
		}
	}
	fmt.Println("seeds: ", seeds)
	fmt.Println("part 1 answer: lowest: ", lowest)
	// fmt.Println("Processed Seeds: ", processedSeeds)
}
