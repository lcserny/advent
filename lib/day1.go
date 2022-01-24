package lib

import (
	"strconv"
)

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

// FIXME: reread description
func GetIncreasedSumsCount(lines []string) int {
	var count int
	var previousSum int
	for i, line := range lines {
		if i == 0 || (i+1)%3 == 0 {
			firstVal, _ := strconv.Atoi(line)
			secondVal, _ := strconv.Atoi(lines[i+1])
			thirdVal, _ := strconv.Atoi(lines[i+2])
			sum := firstVal + secondVal + thirdVal
			if sum > previousSum {
				count++
				previousSum = sum
			}
		}
	}
	return count
}
