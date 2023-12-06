package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/brian-trann/aoc-2023-go/utils"
)

func main() {
	lines := utils.OpenFile("./input.txt")
	var cardData []CardData
	for _, line := range lines {
		parsedData := ParseCardString(line)
		card := NewCardData(parsedData.Id, parsedData.WinningCards, parsedData.MyCards)
		cardData = append(cardData, *card)
		// fmt.Println(card.WinningCardCount)
	}
	points := 0
	for _, card := range cardData {
		if card.WinningCardCount == 0 {
			continue
		}
		currPoint := int(math.Pow(2, float64(card.WinningCardCount-1)))

		points += currPoint
	}
	fmt.Printf("part 1 score: %d\n\n", points)
}

type CardData struct {
	Id               string
	WinningCards     []int
	WinningSet       map[int]struct{}
	MyCards          []int
	WinningCardCount int
}

func ParseCardString(input string) CardData {
	startIndex := len("Card ")
	endIndex := strings.Index(input, ":")
	if endIndex == -1 {
		return CardData{}
	}
	cardId := strings.TrimSpace(input[startIndex:endIndex])
	list := strings.TrimSpace(input[endIndex+1:])
	cardSplitStart := strings.Index(list, "|")
	winningCardsStr := strings.TrimSpace(list[0:cardSplitStart])
	myCardsStr := strings.TrimSpace(list[cardSplitStart+1:])
	winningCards := ConvertStringToIntSlice(winningCardsStr, " ")
	myCards := ConvertStringToIntSlice(myCardsStr, " ")

	return CardData{Id: cardId, WinningCards: winningCards, MyCards: myCards}

}

func NewCardData(id string, winningCards, myCards []int) *CardData {
	cardData := &CardData{
		Id:           id,
		WinningCards: winningCards,
		MyCards:      myCards,
	}
	cardData.initializeWinningSet()
	cardData.countMyWinningCards()

	return cardData
}

// change list to a set for quick lookup
func (c *CardData) initializeWinningSet() {
	c.WinningSet = make(map[int]struct{})
	for _, card := range c.WinningCards {
		c.WinningSet[card] = struct{}{}
	}
	c.WinningCardCount = len(c.WinningCards)
}

func ConvertStringToIntSlice(input string, delimiter string) []int {
	var result []int
	substrings := strings.Split(input, delimiter)
	for _, s := range substrings {
		num, err := strconv.Atoi(s)
		if err != nil {

			// fmt.Println("error parsing strign to int")
			// continue
		} else {

			result = append(result, num)
		}
	}
	return result
}
func (c *CardData) countMyWinningCards() {
	count := 0
	for _, card := range c.MyCards {
		if _, exists := c.WinningSet[card]; exists {
			count++
		}
	}
	c.WinningCardCount = count

}
