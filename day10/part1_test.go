package day10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFile(t *testing.T) {
	maze := parseFile("sample.txt")
	assert.Equal(t, location{2, 0}, maze.startingLoc)
	assert.Equal(t, []rune{'.', 'F', 'J', '|', '7'}, maze.runes[1][0:5])
}

func TestPart1(t *testing.T) {
	answer, err := Part1()
	assert.Nil(t, err)
	assert.Equal(t, "6860", answer)
}
