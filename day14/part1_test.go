package day14

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFile(t *testing.T) {
	locations := parseFile("sample.txt")
	assert.Equal(t, []int{ROUND_ROCK, SPACE, SPACE, SPACE, SPACE, CUBE_ROCK, SPACE, SPACE, SPACE, SPACE}, locations[0])
	assert.Equal(t, []int{ROUND_ROCK, SPACE, CUBE_ROCK, SPACE, SPACE, ROUND_ROCK, SPACE, CUBE_ROCK, SPACE, CUBE_ROCK}, locations[5])
	assert.Equal(t, []int{CUBE_ROCK, ROUND_ROCK, ROUND_ROCK, SPACE, SPACE, CUBE_ROCK, SPACE, SPACE, SPACE, SPACE}, locations[9])
}

func TestTilt(t *testing.T) {
	locations := parseFile("sample.txt")
	locations = tilt(locations, NORTH)
	assert.Equal(t, []int{ROUND_ROCK, ROUND_ROCK, ROUND_ROCK, ROUND_ROCK, SPACE, CUBE_ROCK, SPACE, ROUND_ROCK, SPACE, SPACE}, locations[0])
	assert.Equal(t, []int{ROUND_ROCK, ROUND_ROCK, SPACE, SPACE, CUBE_ROCK, SPACE, SPACE, SPACE, SPACE, CUBE_ROCK}, locations[1])
	assert.Equal(t, []int{CUBE_ROCK, SPACE, SPACE, SPACE, SPACE, CUBE_ROCK, SPACE, SPACE, SPACE, SPACE}, locations[9])
}

func TestPart1Sample(t *testing.T) {
	answer := Part1("sample.txt")
	assert.Equal(t, 136, answer)
}

func TestPart1(t *testing.T) {
	answer := Part1("input.txt")
	assert.Equal(t, 108889, answer)
}
