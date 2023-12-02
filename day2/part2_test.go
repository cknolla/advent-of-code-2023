package day2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPower(t *testing.T) {
	testCases := []struct {
		description   string
		inputLine     string
		expectedValue int
	}{
		{
			description:   "Game 1",
			inputLine:     "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			expectedValue: 48,
		},
		{
			description:   "Game 2",
			inputLine:     "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			expectedValue: 12,
		},
		{
			description:   "Game 3",
			inputLine:     "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			expectedValue: 1560,
		},
		{
			description:   "Game 4",
			inputLine:     "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			expectedValue: 630,
		},
		{
			description:   "Game 5",
			inputLine:     "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			expectedValue: 36,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			g := parseGame(testCase.inputLine)
			assert.Equal(t, testCase.expectedValue, g.getPower())
		})
	}
}

func TestPart2(t *testing.T) {
	answer, err := Part2()
	assert.Nil(t, err)
	assert.Equal(t, "56580", answer)
}
