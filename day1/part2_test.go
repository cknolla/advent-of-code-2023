package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPart2CalibrationValue(t *testing.T) {
	testCases := []struct {
		description   string
		inputLine     string
		expectedValue int
	}{
		{
			description:   "double integer",
			inputLine:     "27six2sixjqctbbfv5",
			expectedValue: 25,
		},
		{
			description:   "double string",
			inputLine:     "fourfive5jhgpcxmqpr41two",
			expectedValue: 42,
		},
		{
			description:   "string first integer second",
			inputLine:     "six1pcknpvv",
			expectedValue: 61,
		},
		{
			description:   "integer first string second",
			inputLine:     "5tkvc5twosixspzhdfourninestpsj",
			expectedValue: 59,
		},
		{
			description:   "single integer",
			inputLine:     "ptcfnjgchx1",
			expectedValue: 11,
		},
		{
			description:   "single string",
			inputLine:     "ptcfivefnjgchx",
			expectedValue: 55,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			assert.Equal(t, testCase.expectedValue, getPart2CalibrationValue(testCase.inputLine))
		})
	}
}

func TestPart2(t *testing.T) {
	answer, err := Part2()
	assert.Nil(t, err)
	assert.Equal(t, "53855", answer)
}
