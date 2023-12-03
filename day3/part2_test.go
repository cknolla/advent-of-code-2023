package day3

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestParsePart2Schematic(t *testing.T) {
	input := "467..114..\n...*......\n..35..633.\n.22...#$11\n617*......"
	lines := strings.Split(input, "\n")
	assert.Equal(t, schematic{
		partNumbers: []partNumber{
			{
				value:         467,
				valid:         false,
				startLocation: 0,
				endLocation:   2,
			},
			{
				value:         114,
				valid:         false,
				startLocation: 5,
				endLocation:   7,
			},
		},
	}, parsePart2Schematic(lines[0]))
	assert.Equal(t, schematic{
		symbolLocations: []int{3},
	}, parsePart2Schematic(lines[1]))
	assert.Equal(t, schematic{
		partNumbers: []partNumber{
			{
				value:         35,
				valid:         false,
				startLocation: 2,
				endLocation:   3,
			},
			{
				value:         633,
				valid:         false,
				startLocation: 6,
				endLocation:   8,
			},
		},
	}, parsePart2Schematic(lines[2]))
	assert.Equal(t, schematic{
		partNumbers: []partNumber{
			{
				value:         22,
				valid:         false,
				startLocation: 1,
				endLocation:   2,
			},
			{
				value:         11,
				valid:         false,
				startLocation: 8,
				endLocation:   9,
			},
		},
	}, parsePart2Schematic(lines[3]))
	assert.Equal(t, schematic{
		symbolLocations: []int{3},
		partNumbers: []partNumber{
			{
				value:         617,
				valid:         false,
				startLocation: 0,
				endLocation:   2,
			},
		},
	}, parsePart2Schematic(lines[4]))
}

func TestPart2(t *testing.T) {
	answer, err := Part2()
	assert.Nil(t, err)
	assert.Equal(t, "81709807", answer)
}
