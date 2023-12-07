package day5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessSeedsPart2(t *testing.T) {
	seedRanges := processSeedsPart2("seeds: 858905075 56936593 947763189 267019426 206349064 252409474 660226451 92561087 752930744 24162055 75704321 63600948 3866217991 323477533 3356941271 54368890 1755537789 475537300 1327269841 427659734")
	assert.Equal(t, seedRange{858905075, 915841667}, seedRanges[0])
	assert.Equal(t, seedRange{947763189, 1214782614}, seedRanges[1])
	assert.Equal(t, seedRange{1327269841, 1754929574}, seedRanges[9])
	assert.Equal(t, 10, len(seedRanges))
}

func TestConvertPart2(t *testing.T) {
	testCases := []struct {
		description    string
		inputRange     *seedRange
		expectedRanges []seedRange
		expectedFound  bool
	}{
		{
			description: "not in mapping",
			inputRange: &seedRange{
				start: 5,
				end:   10,
			},
			expectedRanges: []seedRange{
				{
					5, 10,
				},
			},
			expectedFound: false,
		},
		{
			description: "fully in mapping",
			inputRange: &seedRange{
				start: 35,
				end:   40,
			},
			expectedRanges: []seedRange{
				{
					55, 60,
				},
			},
			expectedFound: true,
		},
		{
			description: "split mapping beyond end",
			inputRange: &seedRange{
				start: 45,
				end:   55,
			},
			expectedRanges: []seedRange{
				{
					65, 70,
				},
				{
					51, 55,
				},
			},
			expectedFound: true,
		},
		{
			description: "split mapping beyond start",
			inputRange: &seedRange{
				start: 25,
				end:   35,
			},
			expectedRanges: []seedRange{
				{
					25, 29,
				},
				{
					50, 55,
				},
			},
			expectedFound: true,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			m := mapping{
				destination: 50,
				source:      30,
				range_:      20,
			}
			ranges, found := m.convertPart2(testCase.inputRange)
			assert.Equal(t, testCase.expectedFound, found)
			assert.Equal(t, testCase.expectedRanges, ranges)
		})
	}
}

func TestGetMinLocation(t *testing.T) {
	seeds := []seedRange{
		{
			90, 100,
		},
		{
			50, 80,
		},
		{
			20, 30,
		},
	}

	min := getMinLocation(seeds)
	assert.Equal(t, 20, min)
}

func TestPart2(t *testing.T) {
	answer, err := Part2()
	assert.Nil(t, err)
	assert.Equal(t, "11611182", answer)
}
