package day2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseGame(t *testing.T) {
	g := parseGame("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red")
	assert.Equal(t, 3, g.id)
	assert.Equal(t, []round{
		{
			red:   20,
			green: 8,
			blue:  6,
		},
		{
			red:   4,
			green: 13,
			blue:  5,
		},
		{
			red:   1,
			green: 5,
			blue:  0,
		},
	}, g.rounds)
}

func TestRoundMaximum(t *testing.T) {
	testCases := []struct {
		description   string
		inputLine     string
		expectedValue round
	}{
		{
			description: "Game 1",
			inputLine:   "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			expectedValue: round{
				red:   4,
				green: 2,
				blue:  6,
			},
		},
		{
			description: "Game 2",
			inputLine:   "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			expectedValue: round{
				red:   1,
				green: 3,
				blue:  4,
			},
		},
		{
			description: "Game 3",
			inputLine:   "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			expectedValue: round{
				red:   20,
				green: 13,
				blue:  6,
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			g := parseGame(testCase.inputLine)
			assert.Equal(t, testCase.expectedValue, g.roundMaximum())
		})
	}
}

func TestIsPossible(t *testing.T) {
	testCases := []struct {
		description   string
		inputLine     string
		expectedValue bool
	}{
		{
			description:   "Game 1",
			inputLine:     "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			expectedValue: true,
		},
		{
			description:   "Game 2",
			inputLine:     "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			expectedValue: true,
		},
		{
			description:   "Game 3",
			inputLine:     "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			expectedValue: false,
		},
		{
			description:   "Game 4",
			inputLine:     "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			expectedValue: false,
		},
		{
			description:   "Game 5",
			inputLine:     "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			expectedValue: true,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			g := parseGame(testCase.inputLine)
			assert.Equal(t, testCase.expectedValue, g.isPossible(&round{
				red:   12,
				green: 13,
				blue:  14,
			}))
		})
	}
}

func TestPart1(t *testing.T) {
	answer, err := Part1()
	assert.Nil(t, err)
	assert.Equal(t, "2727", answer)
}
