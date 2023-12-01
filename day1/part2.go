package day1

import (
	"advent-of-code-2023/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

var strMap = map[string]string{
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func getPart2CalibrationValue(inputLine string) int {
	var numbers []string
LINE:
	for index, char := range inputLine {
		if unicode.IsNumber(char) {
			numbers = append(numbers, string(char))
			continue
		}
		for str, num := range strMap {
			if strings.HasPrefix(inputLine[index:], str) {
				numbers = append(numbers, num)
				continue LINE
			}
		}
	}
	calibrationValue, err := strconv.Atoi(numbers[0] + numbers[len(numbers)-1])
	if err != nil {
		log.Fatalln(err)
	}
	return calibrationValue
}

func Part2() (string, error) {
	lines := utils.GetInputLines()
	sum := 0
	for _, line := range lines {
		sum += getPart2CalibrationValue(line)
	}
	return fmt.Sprintf("%d", sum), nil
}
