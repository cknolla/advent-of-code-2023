package day13

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart2Sample(t *testing.T) {
	answer := Part2("sample.txt")
	assert.Equal(t, 400, answer)
}

func TestPart2(t *testing.T) {
	answer := Part2("input.txt")
	assert.Equal(t, 40995, answer)
}
