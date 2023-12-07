package day7

import (
	"advent-of-code-2023/utils"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

const (
	HIGH_CARD = iota
	ONE_PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	FULL_HOUSE
	FOUR_OF_A_KIND
	FIVE_OF_A_KIND
)

type hand struct {
	cards []int
	bid   int
	value int
}

func newHand(cardStr string, bidStr string) hand {
	h := hand{
		cards: make([]int, 5),
	}
	h.bid, _ = strconv.Atoi(bidStr)
	for i, c := range cardStr {
		h.cards[i] = getCardValue(c)
	}
	h.value = h.determineValue()
	return h
}

func (h *hand) determineValue() int {
	cardCountMap := make(map[int]int, 5)
	for _, card := range h.cards {
		cardCountMap[card]++
	}
	var cardCounts []int
	for _, value := range cardCountMap {
		cardCounts = append(cardCounts, value)
	}
	sort.Ints(cardCounts)
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

type byCards []hand

func (b byCards) Len() int      { return len(b) }
func (b byCards) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

// despite the result being Less, I prefer determining Greater, then inverting the return
func (b byCards) Less(i, j int) bool {
	if b[i].value == b[j].value {
		for index, card := range b[i].cards {
			if card != b[j].cards[index] {
				return card < b[j].cards[index]
			}
		}
	}
	return b[i].value < b[j].value
}

func getCardValue(char rune) int {
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
		return 11
	case 'Q':
		return 12
	case 'K':
		return 13
	default:
		return 14
	}
}

func parseHand(line string) hand {
	lineParts := strings.Fields(line)
	cardStr, bidStr := lineParts[0], lineParts[1]
	return newHand(cardStr, bidStr)
}

func parseFile(filename string) []hand {
	hands := make([]hand, 1000)
	for l := range utils.GetInputLines(filename) {
		line := l.Text
		hands[l.Index] = parseHand(line)
	}
	return hands
}

func Part1() (string, error) {
	answer := 0
	hands := parseFile("input.txt")
	sort.Sort(byCards(hands))
	for index, hand := range hands {
		answer += hand.bid * (index + 1)
	}
	return fmt.Sprintf("%d", answer), nil
}
