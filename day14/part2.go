package day14

type cycleState struct {
	cycleNumber int
	locationStr string
	locations   [][]int
	load        int
}

func calculateBillionth(currentCycle int, previousMatchingCycle int) (billionthCycle int) {
	cyclePeriod := currentCycle - previousMatchingCycle
	totalCyclesRemaining := 1_000_000_000 - previousMatchingCycle
	billionthCycle = totalCyclesRemaining%cyclePeriod - 1
	return billionthCycle + previousMatchingCycle
}

func cycle(locations [][]int) [][]int {
	locations = tilt(locations, NORTH)
	locations = tilt(locations, WEST)
	locations = tilt(locations, SOUTH)
	locations = tilt(locations, EAST)
	return locations
}

func convertToString(locations *[][]int) string {
	locationsStr := ""
	for _, location := range *locations {
		for _, object := range location {
			char := "."
			switch object {
			case ROUND_ROCK:
				char = "O"
			case CUBE_ROCK:
				char = "#"
			}
			locationsStr += char
		}
	}
	return locationsStr
}

func Part2(filename string) int {
	answer := 0
	cache := make(map[string]cycleState)
	locations := parseFile(filename)
	for i := 0; i < 1_000_000_000; i++ {
		locations = cycle(locations)
		locationsStr := convertToString(&locations)
		previousMatchingCycleState, ok := cache[locationsStr]
		if ok {
			targetCycleNumber := calculateBillionth(i, previousMatchingCycleState.cycleNumber)
			for _, v := range cache {
				if v.cycleNumber == targetCycleNumber {
					answer = v.load
					break
				}
			}
			break
		}
		cache[locationsStr] = cycleState{i, locationsStr, locations, calculateLoad(&locations)}
	}
	return answer
}
