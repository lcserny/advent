package main

import (
	"advent/lib"
	"fmt"
)

func main() {
	file := lib.ParseFile()
	lines := lib.ReadLines(file)

	increasedCount := lib.GetIncreasedCount(lines)
	fmt.Println(increasedCount)

	increasedSumsCount := lib.GetIncreasedSumsCount(lines)
	fmt.Println(increasedSumsCount)
}
