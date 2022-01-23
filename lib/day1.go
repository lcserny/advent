package lib

import "strconv"

func GetIncreasedCount(lines []string) int {
	var increasedCount int
	for i, line := range lines {
		if i == 0 {
			continue
		}
		prevVal, _ := strconv.Atoi(lines[i-1])
		currentVal, _ := strconv.Atoi(line)
		if currentVal > prevVal {
			increasedCount++
		}
	}
	return increasedCount
}