package day7

import (
	"advent-of-code-2023/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func newHandPart2(cardStr string, bidStr string) hand {
	h := hand{
		cards: make([]int, 5),
	}
	h.bid, _ = strconv.Atoi(bidStr)
	for i, c := range cardStr {
		h.cards[i] = getCardValuePart2(c)
	}
	h.value = h.determineValuePart2()
	return h
}

func (h *hand) determineValuePart2() int {
	cardCountMap := make(map[int]int, 5)
	for _, card := range h.cards {
		cardCountMap[card]++
	}
	var cardCounts []int
	jokerCount := 0
	for key, value := range cardCountMap {
		if key == 1 {
			jokerCount = value
		} else {
			cardCounts = append(cardCounts, value)
		}
	}
	if jokerCount == 5 {
		return FIVE_OF_A_KIND
	}
	slices.Sort(cardCounts)
	// adding jokers to the largest group is always the strongest
	cardCounts[len(cardCounts)-1] += jokerCount
	if slices.Equal(cardCounts, []int{5}) {
		return FIVE_OF_A_KIND
	}
	if slices.Equal(cardCounts, []int{1, 4}) {
		return FOUR_OF_A_KIND
	}
	if slices.Equal(cardCounts, []int{2, 3}) {
		return FULL_HOUSE
	}
	if slices.Equal(cardCounts, []int{1, 1, 3}) {
		return THREE_OF_A_KIND
	}
	if slices.Equal(cardCounts, []int{1, 2, 2}) {
		return TWO_PAIR
	}
	if slices.Equal(cardCounts, []int{1, 1, 1, 2}) {
		return ONE_PAIR
	}
	return HIGH_CARD
}

func getCardValuePart2(char rune) int {
	switch char {
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	case 'T':
		return 10
	case 'J':
		return 1
	case 'Q':
		return 12
	case 'K':
		return 13
	default:
		return 14
	}
}

func parseHandPart2(line string) hand {
	lineParts := strings.Fields(line)
	cardStr, bidStr := lineParts[0], lineParts[1]
	return newHandPart2(cardStr, bidStr)
}

func parseFilePart2(filename string) []hand {
	hands := make([]hand, 1000)
	for l := range utils.GetInputLines(filename) {
		line := l.Text
		hands[l.Index] = parseHandPart2(line)
	}
	return hands
}

func Part2() (string, error) {
	answer := 0
	hands := parseFilePart2("input.txt")
	slices.SortFunc(hands, handCmp)
	for index, hand := range hands {
		answer += hand.bid * (index + 1)
	}
	return fmt.Sprintf("%d", answer), nil
}
