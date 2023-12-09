package day9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFile(t *testing.T) {
	histories := parseFile("sample.txt")
	assert.Equal(t, history{[][]int{{0, 3, 6, 9, 12, 15}}}, histories[0])
	assert.Equal(t, history{[][]int{{1, 3, 6, 10, 15, 21}}}, histories[1])
	assert.Equal(t, history{[][]int{{10, 13, 16, 21, 30, 45}}}, histories[2])
}

func TestBuildDifferences(t *testing.T) {
	histories := parseFile("sample.txt")
	histories[0].buildDifferences(histories[0].values[0])
	assert.Equal(t, history{
		[][]int{
			{0, 3, 6, 9, 12, 15},
			{3, 3, 3, 3, 3},
			{0, 0, 0, 0},
		},
	}, histories[0])
	histories[1].buildDifferences(histories[1].values[0])
	assert.Equal(t, history{
		[][]int{
			{1, 3, 6, 10, 15, 21},
			{2, 3, 4, 5, 6},
			{1, 1, 1, 1},
			{0, 0, 0},
		},
	}, histories[1])
}

func TestGetNextValue(t *testing.T) {
	histories := parseFile("sample.txt")
	nextValue := histories[0].getNextValue()
	assert.Equal(t, 18, nextValue)
	nextValue = histories[1].getNextValue()
	assert.Equal(t, 28, nextValue)
	nextValue = histories[2].getNextValue()
	assert.Equal(t, 68, nextValue)
}

func TestPart1(t *testing.T) {
	answer, err := Part1()
	assert.Nil(t, err)
	assert.Equal(t, "2174807968", answer)
}
