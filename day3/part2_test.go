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

func TestGetGearRatios(t *testing.T) {
	//21.2*2...1*...5........9*9..8
	//*89.......3..4..54*......9..*
	//.....5..2....7*......3*3....8
	prevSchematic := schematic{
		symbolLocations: []int{4, 10, 24},
		partNumbers: []partNumber{
			{
				value:         21,
				startLocation: 0,
				endLocation:   1,
			},
			{
				2,
				false,
				3,
				3,
			},
			{
				2,
				false,
				5,
				5,
			},
			{
				1,
				false,
				9,
				9,
			},
			{
				5,
				false,
				14,
				14,
			},
			{
				9,
				false,
				23,
				25,
			},
			{
				9,
				false,
				23,
				23,
			},
			{
				8,
				false,
				28,
				28,
			},
		},
	}
	schem := schematic{
		symbolLocations: []int{0, 18, 28},
		partNumbers: []partNumber{
			{
				value:         89,
				startLocation: 1,
				endLocation:   2,
			},
			{
				value:         3,
				startLocation: 10,
				endLocation:   10,
			},
			{
				value:         4,
				startLocation: 13,
				endLocation:   13,
			},
			{
				value:         54,
				startLocation: 16,
				endLocation:   17,
			},
			{
				value:         9,
				startLocation: 25,
				endLocation:   25,
			},
		},
	}
	nextSchematic := schematic{
		symbolLocations: []int{14, 22},
		partNumbers: []partNumber{
			{
				value:         5,
				startLocation: 5,
				endLocation:   5,
			},
			{
				value:         2,
				startLocation: 8,
				endLocation:   8,
			},
			{
				value:         7,
				startLocation: 13,
				endLocation:   13,
			},
			{
				value:         3,
				startLocation: 21,
				endLocation:   21,
			},
			{
				value:         3,
				startLocation: 23,
				endLocation:   23,
			},
			{
				value:         8,
				startLocation: 28,
				endLocation:   28,
			},
		},
	}
	schem.prev = &prevSchematic
	schem.next = &nextSchematic
	prevSchematic.next = &schem
	nextSchematic.prev = &schem
	ratioSum := prevSchematic.getGearRatios()
	assert.Equal(t, 7, ratioSum)
	ratioSum = schem.getGearRatios()
	assert.Equal(t, 1933, ratioSum)
	ratioSum = nextSchematic.getGearRatios()
	assert.Equal(t, 37, ratioSum)

}

func TestPart2(t *testing.T) {
	answer, err := Part2()
	assert.Nil(t, err)
	assert.Equal(t, "81709807", answer)
}
