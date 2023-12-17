package day17

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart2(t *testing.T) {
	answer := Part2(inputDay)
	assert.Equal(t, 1389, answer)
}
