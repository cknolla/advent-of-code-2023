package day4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseLine(t *testing.T) {
	testCases := []struct {
		description  string
		inputLine    string
		expectedCard card
	}{
		{
			description: "Card 1",
			inputLine:   "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			expectedCard: card{
				copies: 1,
				winningNumbers: map[string]bool{
					"41": true,
					"48": true,
					"83": true,
					"86": true,
					"17": true,
				},
				availableNumbers: []string{
					"83", "86", "6", "31", "17", "9", "48", "53",
				},
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			c := parseLine(testCase.inputLine)
			assert.Equal(t, testCase.expectedCard, c)
		})
	}
}

func TestGetWinningCount(t *testing.T) {
	testCases := []struct {
		description        string
		card               card
		expectedPowerOfTwo int
	}{
		{
			description: "Card 1",
			card: card{
				winningNumbers: map[string]bool{
					"41": true,
					"48": true,
					"83": true,
					"86": true,
					"17": true,
				},
				availableNumbers: []string{
					"83", "86", "6", "31", "17", "9", "48", "53",
				},
			},
			expectedPowerOfTwo: 4,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			assert.Equal(t, testCase.expectedPowerOfTwo, testCase.card.getWinningCount())
		})
	}
}

func TestPart1(t *testing.T) {
	answer, err := Part1()
	assert.Nil(t, err)
	assert.Equal(t, "19855", answer)
}
