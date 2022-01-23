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

// FIXME
func GetIncreasedSumsCount(lines []string) int {
	var count int
	var nextStartWindowIndex int
	var previousWindowsSum int
	for i, line := range lines {
		if i == 0 || i > nextStartWindowIndex {
			nextStartWindowIndex = i
			if i + 2 < len(lines) {
				firstVal, _ := strconv.Atoi(line)
				secondVal, _ := strconv.Atoi(lines[i+1])
				thirdVal, _ := strconv.Atoi(lines[i+2])
				if (firstVal + secondVal + thirdVal) > previousWindowsSum {
					count++
				}
			}
		}
	}
	return count
}