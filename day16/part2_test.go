package day16

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart2Sample(t *testing.T) {
	answer := Part2("sample.txt")
	assert.Equal(t, 51, answer)
}

func TestPart2(t *testing.T) {
	answer := Part2("input.txt")
	assert.Equal(t, 8221, answer)
}
