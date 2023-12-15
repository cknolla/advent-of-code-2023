package day15

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1Sample(t *testing.T) {
	answer := Part1("sample.txt")
	assert.Equal(t, 1320, answer)
}

func TestPart1(t *testing.T) {
	answer := Part1("input.txt")
	assert.Equal(t, 517315, answer)
}
