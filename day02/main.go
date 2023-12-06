package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brian-trann/aoc-2023-go/utils"
)

type GameData struct {
	GameID string
	List   string
}

func main() {
	PartOne()
	PartTwo()
}

const (
	redMax   = 12
	greenMax = 13
	blueMax  = 14
)

func PartOne() {
	lines := utils.OpenFileSliceStrings("./input.txt")
	var validIds []string
	for _, line := range lines {
		gameData := ParseGameString(line)
		// fmt.Println((gameData.GameID))
		validList := CheckSublists(gameData.List)
		if validList {
			validIds = append(validIds, gameData.GameID)
		}

	}
	sum, err := utils.SumStringNumbers(validIds)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Part one Sum:", sum)
	}
}

func PartTwo() {
	lines := utils.OpenFileSliceStrings("./input.txt")
	total := 0
	for _, line := range lines {
		gameData := ParseGameString(line)
		productOfList := GetMaxColorCountProduct(gameData.List)
		total += productOfList
	}
	fmt.Println("Part two Sum:", total)
}
func ParseGameString(input string) GameData {

	// Skip the "Game " part.
	startIndex := len("Game ")

	// find the end index of the Game ID, which is before the first colon.
	endIndex := strings.Index(input, ":")
	if endIndex == -1 {
		return GameData{}
	}

	gameID := strings.TrimSpace(input[startIndex:endIndex])

	// extract the list, which is anything after the colon.
	list := strings.TrimSpace(input[endIndex+1:])

	return GameData{GameID: gameID, List: list}
}
func CheckSublists(list string) bool {
	sublists := strings.Split(list, ";")

	for _, sublist := range sublists {
		redCount, greenCount, blueCount := 0, 0, 0

		// color segments.
		colorSegments := strings.Split(sublist, ",")

		for _, segment := range colorSegments {

			// trim; get color and count
			segment = strings.TrimSpace(segment)
			// fmt.Println((segment))
			parts := strings.Fields(segment)

			if len(parts) == 2 {
				count, err := strconv.Atoi(parts[0])
				if err != nil {
					fmt.Printf("error parsing count from segment '%s': %s\n", segment, err)
					return false
				}

				// incre the color count
				switch strings.ToLower(parts[1]) {
				case "red":
					redCount += count
				case "green":
					greenCount += count
				case "blue":
					blueCount += count
				}
			}
		}

		// check if any color count exceeds its maximum.
		if redCount > redMax || greenCount > greenMax || blueCount > blueMax {
			return false
		}
	}

	return true
}
func GetMaxColorCountProduct(list string) int {
	maxRed, maxGreen, maxBlue := 0, 0, 0

	sublists := strings.Split(list, ";")
	for _, sublist := range sublists {
		redCount, greenCount, blueCount := 0, 0, 0

		colorSegments := strings.Split(sublist, ",")
		for _, segment := range colorSegments {
			segment = strings.TrimSpace(segment)
			parts := strings.Fields(segment)
			if len(parts) == 2 {
				count, err := strconv.Atoi(parts[0])
				if err != nil {
					fmt.Printf("error parsing count from segment '%s': %s\n", segment, err)
					continue
				}

				switch strings.ToLower(parts[1]) {
				case "red":
					redCount = max(redCount, count)
				case "green":
					greenCount = max(greenCount, count)
				case "blue":
					blueCount = max(blueCount, count)
				}
			}
		}

		// update the overall maximum counts
		maxRed = max(maxRed, redCount)
		maxGreen = max(maxGreen, greenCount)
		maxBlue = max(maxBlue, blueCount)
	}

	return maxRed * maxGreen * maxBlue
}
