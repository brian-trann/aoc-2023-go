package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brian-trann/aoc-2023-go/utils"
)

func main() {
	contents := utils.OpenFileSliceStrings("./input.txt")
	camelCards := contentsToCards(contents)

	bubbleSort(camelCards)
	// at this point we are sorted
	partOneSum := 0
	for i, card := range camelCards {
		multiplier := i + 1
		partOneSum += (multiplier * card.Bid)
		// fmt.Printf("%s | handTypeVal: %d | firstCardVal %d | secondCardVal %d \n", card.Hand, card.HandTypeVal, card.FirstCardVal, card.SecondCardVal)

	}
	fmt.Println("part one: ", partOneSum)

}

type CamelCard struct {
	Hand string
	Bid  int

	HandTypeVal int
	// just for convienenc
	FirstCardVal  int
	SecondCardVal int
	ThirdCardVal  int
	FourthCardVal int
	FifthCardVal  int
}

var handRanks = map[string]int{
	"fiveKind":  6,
	"fourKind":  5,
	"fullHouse": 4,
	"threekind": 3,
	"twoPair":   2,
	"onePair":   1,
	"noPair":    0,
}

var cardRank = map[string]int{
	"A": 12,
	"K": 11,
	"Q": 10,
	"J": 9,
	"T": 8,
	"9": 7,
	"8": 6,
	"7": 5,
	"6": 4,
	"5": 3,
	"4": 2,
	"3": 1,
	"2": 0,
}

func contentsToCards(input []string) []CamelCard {

	var cards []CamelCard
	for _, line := range input {
		if len(line) == 0 {
			continue
		}
		index := strings.Index(line, " ")
		hand := strings.TrimSpace(line[0:index])
		bidStr := strings.TrimSpace(line[index+1:])
		bid, err := strconv.Atoi(bidStr)
		if err != nil {
			continue
		}
		card := CamelCard{Hand: hand, Bid: bid}

		card.calculateHand()

		cards = append(cards, card)
	}
	return cards
}

func (c *CamelCard) calculateHand() {
	counts := make(map[rune]int)

	for i, card := range c.Hand {
		if i == 0 {
			c.FirstCardVal = cardRank[string(card)]
		} else if i == 1 {
			c.SecondCardVal = cardRank[string(card)]
		} else if i == 2 {
			c.ThirdCardVal = cardRank[string(card)]
		} else if i == 3 {
			c.FourthCardVal = cardRank[string(card)]
		} else {
			c.FifthCardVal = cardRank[string(card)]

		}
		counts[card]++
	}

	var pairCount, threeCount, fourCount, fiveCount int
	for _, count := range counts {
		switch count {
		case 2:
			pairCount++
		case 3:
			threeCount++
		case 4:
			fourCount++
		case 5:
			fiveCount++
		}
	}
	// assign the hand types
	switch {
	case fiveCount > 0:
		c.HandTypeVal = handRanks["fiveKind"]
	case fourCount > 0:
		c.HandTypeVal = handRanks["fourKind"]
	case threeCount > 0 && pairCount > 0:
		c.HandTypeVal = handRanks["fullHouse"]
	case threeCount > 0:
		c.HandTypeVal = handRanks["threekind"]
	case pairCount > 1:
		c.HandTypeVal = handRanks["twoPair"]
	case pairCount == 1:
		c.HandTypeVal = handRanks["onePair"]
	default:
		c.HandTypeVal = handRanks["noPair"]
	}
}
func bubbleSort(arr []CamelCard) bool {
	n := len(arr)
	neededToSort := false
	for i := 0; i < n-1; i++ {

		// Flag to check if any swap happens in this pass
		swapped := false
		// Last i elements are already in place
		for j := 0; j < n-i-1; j++ {
			curr := arr[j]
			next := arr[j+1]
			if curr.HandTypeVal > next.HandTypeVal {
				// swap; high type wins
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
				neededToSort = true
			} else if curr.HandTypeVal == next.HandTypeVal {
				// disgusting
				if curr.FirstCardVal > next.FirstCardVal {
					arr[j], arr[j+1] = arr[j+1], arr[j]
					swapped = true
				} else if curr.FirstCardVal == next.FirstCardVal && curr.SecondCardVal > next.SecondCardVal {
					arr[j], arr[j+1] = arr[j+1], arr[j]
					swapped = true
				} else if curr.FirstCardVal == next.FirstCardVal && curr.SecondCardVal == next.SecondCardVal && curr.ThirdCardVal > next.ThirdCardVal {
					arr[j], arr[j+1] = arr[j+1], arr[j]
					swapped = true
				} else if curr.FirstCardVal == next.FirstCardVal && curr.SecondCardVal == next.SecondCardVal && curr.ThirdCardVal == next.ThirdCardVal && curr.FourthCardVal > next.FourthCardVal {
					arr[j], arr[j+1] = arr[j+1], arr[j]
					swapped = true
				} else if curr.FirstCardVal == next.FirstCardVal && curr.SecondCardVal == next.SecondCardVal && curr.ThirdCardVal == next.ThirdCardVal && curr.FourthCardVal == next.FourthCardVal && curr.FifthCardVal > next.FifthCardVal {
					arr[j], arr[j+1] = arr[j+1], arr[j]
					swapped = true
				}
			}
		}

		if !swapped {
			break
		}
	}
	return neededToSort
}
