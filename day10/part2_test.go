package day10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart2(t *testing.T) {
	answer, err := Part2()
	assert.Nil(t, err)
	assert.Equal(t, "343", answer)
}
