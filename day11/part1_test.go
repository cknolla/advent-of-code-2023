package day11

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSample(t *testing.T) {
	universe := parseFile("sample.txt", 2)
	universe.expandColumns()
	answer := universe.getShortestDistances()
	assert.Equal(t, 374, answer)
}

func TestPart1(t *testing.T) {
	answer, err := Part1()
	assert.Nil(t, err)
	assert.Equal(t, "9918828", answer)
}
