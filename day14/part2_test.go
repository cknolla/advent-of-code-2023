package day14

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTiltPart2(t *testing.T) {
	testCases := []struct {
		description     string
		direction       int
		expectedValues  [][]int
		locationIndexes []int
	}{
		{
			description: "South",
			direction:   SOUTH,
			expectedValues: [][]int{
				{SPACE, SPACE, SPACE, SPACE, SPACE, CUBE_ROCK, SPACE, SPACE, SPACE, SPACE},
				{SPACE, SPACE, SPACE, ROUND_ROCK, SPACE, CUBE_ROCK, CUBE_ROCK, SPACE, SPACE, SPACE},
				{ROUND_ROCK, SPACE, ROUND_ROCK, SPACE, SPACE, SPACE, SPACE, ROUND_ROCK, CUBE_ROCK, ROUND_ROCK},
				{CUBE_ROCK, ROUND_ROCK, ROUND_ROCK, SPACE, ROUND_ROCK, CUBE_ROCK, SPACE, SPACE, SPACE, ROUND_ROCK},
			},
			locationIndexes: []int{0, 2, 4, 9},
		},
		{
			description: "East",
			direction:   EAST,
			expectedValues: [][]int{
				{SPACE, SPACE, SPACE, SPACE, ROUND_ROCK, CUBE_ROCK, SPACE, SPACE, SPACE, SPACE},
				{SPACE, ROUND_ROCK, ROUND_ROCK, ROUND_ROCK, CUBE_ROCK, SPACE, SPACE, SPACE, SPACE, CUBE_ROCK},
				{SPACE, ROUND_ROCK, ROUND_ROCK, CUBE_ROCK, SPACE, SPACE, SPACE, SPACE, ROUND_ROCK, ROUND_ROCK},
				{SPACE, SPACE, SPACE, SPACE, SPACE, SPACE, SPACE, SPACE, SPACE, ROUND_ROCK},
				{CUBE_ROCK, SPACE, SPACE, ROUND_ROCK, ROUND_ROCK, CUBE_ROCK, SPACE, SPACE, SPACE, SPACE},
			},
			locationIndexes: []int{0, 1, 3, 7, 9},
		},
		{
			description: "West",
			direction:   WEST,
			expectedValues: [][]int{
				{ROUND_ROCK, SPACE, SPACE, SPACE, SPACE, CUBE_ROCK, SPACE, SPACE, SPACE, SPACE},
				{ROUND_ROCK, ROUND_ROCK, ROUND_ROCK, SPACE, CUBE_ROCK, SPACE, SPACE, SPACE, SPACE, CUBE_ROCK},
				{ROUND_ROCK, ROUND_ROCK, SPACE, CUBE_ROCK, ROUND_ROCK, ROUND_ROCK, SPACE, SPACE, SPACE, SPACE},
				{ROUND_ROCK, SPACE, CUBE_ROCK, ROUND_ROCK, SPACE, SPACE, SPACE, CUBE_ROCK, SPACE, CUBE_ROCK},
				{ROUND_ROCK, SPACE, SPACE, SPACE, SPACE, SPACE, SPACE, SPACE, SPACE, SPACE},
				{CUBE_ROCK, ROUND_ROCK, ROUND_ROCK, SPACE, SPACE, CUBE_ROCK, SPACE, SPACE, SPACE, SPACE},
			},
			locationIndexes: []int{0, 1, 3, 5, 7, 9},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			locations := parseFile("sample.txt")
			locations = tilt(locations, testCase.direction)
			printDish(&locations)
			for index, locationIndex := range testCase.locationIndexes {
				assert.Equal(t, testCase.expectedValues[index], locations[locationIndex])
			}
		})
	}
}

func TestCycle(t *testing.T) {
	testCases := []struct {
		description     string
		cycles          int
		expectedValues  [][]int
		locationIndexes []int
	}{
		{
			description: "1 cycle",
			cycles:      1,
			expectedValues: [][]int{
				{SPACE, SPACE, SPACE, SPACE, SPACE, CUBE_ROCK, SPACE, SPACE, SPACE, SPACE},
				{SPACE, SPACE, SPACE, SPACE, CUBE_ROCK, SPACE, SPACE, SPACE, ROUND_ROCK, CUBE_ROCK},
				{SPACE, SPACE, SPACE, ROUND_ROCK, ROUND_ROCK, CUBE_ROCK, CUBE_ROCK, SPACE, SPACE, SPACE},
				{CUBE_ROCK, SPACE, SPACE, ROUND_ROCK, ROUND_ROCK, CUBE_ROCK, SPACE, SPACE, SPACE, SPACE},
			},
			locationIndexes: []int{0, 1, 2, 9},
		},
		{
			description: "2 cycles",
			cycles:      2,
			expectedValues: [][]int{
				{SPACE, SPACE, SPACE, SPACE, SPACE, CUBE_ROCK, SPACE, SPACE, SPACE, SPACE},
				{SPACE, SPACE, SPACE, SPACE, CUBE_ROCK, SPACE, SPACE, SPACE, ROUND_ROCK, CUBE_ROCK},
				{SPACE, SPACE, SPACE, SPACE, SPACE, CUBE_ROCK, CUBE_ROCK, SPACE, SPACE, SPACE},
				{CUBE_ROCK, SPACE, ROUND_ROCK, ROUND_ROCK, ROUND_ROCK, CUBE_ROCK, SPACE, SPACE, SPACE, ROUND_ROCK},
			},
			locationIndexes: []int{0, 1, 2, 9},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			locations := parseFile("sample.txt")
			for i := testCase.cycles; i > 0; i-- {
				locations = cycle(locations)
			}
			printDish(&locations)
			for index, locationIndex := range testCase.locationIndexes {
				assert.Equal(t, testCase.expectedValues[index], locations[locationIndex])
			}
		})
	}
}

func TestPart2Sample(t *testing.T) {
	answer := Part2("sample.txt")
	assert.Equal(t, 64, answer)
}

func TestPart2(t *testing.T) {
	answer := Part2("input.txt")
	assert.Equal(t, 104671, answer)
}
