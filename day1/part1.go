package day1

import (
	"bufio"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func getPart1CalibrationValue(inputLine string) int {
	re := regexp.MustCompile("[0-9]")
	numbers := re.FindAllString(inputLine, -1)
	calibrationValue, err := strconv.Atoi(numbers[0] + numbers[len(numbers)-1])
	if err != nil {
		log.Fatalln(err)
	}
	return calibrationValue
}

func Part1(scanner *bufio.Scanner) (string, error) {
	sum := 0
	for scanner.Scan() {
		sum += getPart1CalibrationValue(scanner.Text())
	}
	return fmt.Sprintf("%d", sum), nil
}
