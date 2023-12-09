package day9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPrevValue(t *testing.T) {
	histories := parseFile("sample.txt")
	nextValue := histories[0].getPrevValue()
	assert.Equal(t, -3, nextValue)
	nextValue = histories[1].getPrevValue()
	assert.Equal(t, 0, nextValue)
	nextValue = histories[2].getPrevValue()
	assert.Equal(t, 5, nextValue)
}

func TestPart2(t *testing.T) {
	answer, err := Part2()
	assert.Nil(t, err)
	assert.Equal(t, "1208", answer)
}
