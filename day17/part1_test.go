package day17

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1Sample(t *testing.T) {
	answer := Part1(sampleDay)
	assert.Equal(t, 102, answer)
}

func TestPart1(t *testing.T) {
	answer := Part1(inputDay)
	assert.Equal(t, 1246, answer)
}
