package day4

import (
	"advent-of-code-2023/utils"
	"fmt"
)

func getCardCount(cards []card) int {
	sum := 0
	for index, c := range cards {
		winningCount := c.getWinningCount()
		for i := winningCount; i > 0; i-- {
			cards[index+i].copies += c.copies
		}
		sum += c.copies
	}
	return sum
}

func Part2() (string, error) {
	sum := 0
	var cards []card
	for line := range utils.GetInputLines() {
		c := parseLine(line)
		cards = append(cards, c)
	}
	sum = getCardCount(cards)
	return fmt.Sprintf("%d", sum), nil
}
