package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func getCalibrationValue(inputLine string) int {
	re := regexp.MustCompile("[0-9]")
	numbers := re.FindAllString(inputLine, -1)
	calibrationValue, err := strconv.Atoi(numbers[0] + numbers[len(numbers)-1])
	if err != nil {
		log.Fatalln(err)
	}
	return calibrationValue
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		sum += getCalibrationValue(scanner.Text())
	}
	log.Println(sum)
}
