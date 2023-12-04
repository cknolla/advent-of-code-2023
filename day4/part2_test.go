package day4

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGetCardCount(t *testing.T) {
	inputLine := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53\nCard 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19\nCard 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1\nCard 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83\nCard 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36\nCard 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"
	cards := make([]card, 6)
	for index, line := range strings.Split(inputLine, "\n") {
		cards[index] = parseLine(line)
	}
	assert.Equal(t, 30, getCardCount(cards))
}

func TestPart2(t *testing.T) {
	answer, err := Part2()
	assert.Nil(t, err)
	assert.Equal(t, "10378710", answer)
}
