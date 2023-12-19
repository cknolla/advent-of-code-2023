package day18

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFile(t *testing.T) {
	digPlan := parseFile("sample.txt")
	assert.Equal(t, Dig{"R", 6, "#70c710"}, digPlan[0])
	assert.Equal(t, Dig{"U", 2, "#7a21e3"}, digPlan[13])
	assert.Equal(t, 7, grid.Width)
	assert.Equal(t, 10, grid.Height)
}

func TestPart1Sample(t *testing.T) {
	answer := Part1("sample.txt")
	assert.Equal(t, 62, answer)
}

func TestPart1(t *testing.T) {
	answer := Part1("input.txt")
	assert.Equal(t, 50603, answer)
}
