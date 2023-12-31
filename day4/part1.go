package day4

import (
	"advent-of-code-2023/utils"
	"fmt"
	"math"
	"strings"
)

type card struct {
	winningNumbers   map[string]bool
	availableNumbers []string
	copies           int
}

func newCard() card {
	return card{
		winningNumbers: make(map[string]bool),
		copies:         1,
	}
}

func parseLine(line string) card {
	c := newCard()
	cardContent := strings.Split(line, ":")[1]
	numbers := strings.Split(cardContent, "|")
	for _, winningNumberStr := range strings.Fields(numbers[0]) {
		c.winningNumbers[winningNumberStr] = true
	}
	c.availableNumbers = strings.Fields(numbers[1])
	return c
}

func (c *card) getWinningCount() int {
	winningCount := 0
	for _, num := range c.availableNumbers {
		if c.winningNumbers[num] {
			winningCount++
		}
	}
	return winningCount
}

func Part1() (string, error) {
	sum := 0
	for l := range utils.GetInputLines("input.txt") {
		line := l.Text
		c := parseLine(line)
		winningCount := c.getWinningCount()
		if winningCount == 0 {
			continue
		}
		sum += int(math.Pow(2, float64(winningCount-1)))
	}
	return fmt.Sprintf("%d", sum), nil
}
