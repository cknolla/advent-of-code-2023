package day7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart2(t *testing.T) {
	answer, err := Part2()
	assert.Nil(t, err)
	assert.Equal(t, "246436046", answer)
}
