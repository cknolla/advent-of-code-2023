package day11

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSampleRatio100(t *testing.T) {
	universe := parseFile("sample.txt", 10)
	universe.expandColumns()
	answer := universe.getShortestDistances()
	assert.Equal(t, 1030, answer)
}

func TestPart2(t *testing.T) {
	answer, err := Part2()
	assert.Nil(t, err)
	assert.Equal(t, "692506533832", answer)
}
