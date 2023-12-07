package day7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFile(t *testing.T) {
	hands := parseFile("input.txt")
	assert.Equal(t, []int{3, 5, 2, 2, 9}, hands[0].cards)
	assert.Equal(t, 30, hands[0].bid)
	assert.Equal(t, []int{12, 3, 7, 9, 11}, hands[1].cards)
	assert.Equal(t, 837, hands[1].bid)
}

func TestPart1(t *testing.T) {
	answer, err := Part1()
	assert.Nil(t, err)
	assert.Equal(t, "248396258", answer)
}
