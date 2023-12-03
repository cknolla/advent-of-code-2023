package day3

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestParseSchematic(t *testing.T) {
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
	}, parseSchematic(lines[0]))
	assert.Equal(t, schematic{
		symbolLocations: []int{3},
	}, parseSchematic(lines[1]))
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
	}, parseSchematic(lines[2]))
	assert.Equal(t, schematic{
		symbolLocations: []int{6, 7},
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
	}, parseSchematic(lines[3]))
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
	}, parseSchematic(lines[4]))
}

func TestValidatePartNumbers(t *testing.T) {
	//....*.....*...............
	//*89..20...3..4..54*..1...9
	//......................*...
	prevSchematic := schematic{
		symbolLocations: []int{4, 10},
	}
	schem := schematic{
		symbolLocations: []int{0, 18},
		partNumbers: []partNumber{
			{
				value:         89,
				startLocation: 1,
				endLocation:   2,
			},
			{
				value:         20,
				startLocation: 5,
				endLocation:   6,
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
				value:         1,
				startLocation: 21,
				endLocation:   21,
			},
			{
				value:         9,
				startLocation: 25,
				endLocation:   25,
			},
		},
	}
	nextSchematic := schematic{
		symbolLocations: []int{22},
	}
	schem.prev = &prevSchematic
	schem.next = &nextSchematic
	schem.validatePartNumbers()
	assert.True(t, schem.partNumbers[0].valid)
	assert.True(t, schem.partNumbers[1].valid)
	assert.True(t, schem.partNumbers[2].valid)
	assert.False(t, schem.partNumbers[3].valid)
	assert.True(t, schem.partNumbers[4].valid)
	assert.True(t, schem.partNumbers[5].valid)
	assert.False(t, schem.partNumbers[6].valid)
}

func TestPart1(t *testing.T) {
	answer, err := Part1()
	assert.Nil(t, err)
	assert.Equal(t, "538046", answer)
}
