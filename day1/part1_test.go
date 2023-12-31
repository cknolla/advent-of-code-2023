package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPart1CalibrationValue(t *testing.T) {
	testCases := []struct {
		description   string
		inputLine     string
		expectedValue int
	}{
		{
			description:   "simple two-number",
			inputLine:     "ninevct4cpdvqfxmspbz9xrvxfvbpzthreesfnncrqn",
			expectedValue: 49,
		},
		{
			description:   "adjacent two-number",
			inputLine:     "78four",
			expectedValue: 78,
		},
		{
			description:   "single number",
			inputLine:     "3tgsppcfpk",
			expectedValue: 33,
		},
		{
			description:   "multiple numbers",
			inputLine:     "9onert32m5",
			expectedValue: 95,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			assert.Equal(t, testCase.expectedValue, getPart1CalibrationValue(testCase.inputLine))
		})
	}
}

func TestPart1(t *testing.T) {
	answer, err := Part1()
	assert.Nil(t, err)
	assert.Equal(t, "54634", answer)
}
