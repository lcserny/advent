package lib

import (
	"bufio"
	"fmt"
	"os"
)

func ParseFile() *os.File {
	args := os.Args[1:]
	if len(args) != 1 {
		Exit(1, "please provide input file as arg")
	}

	inputFileArg := args[0]

	file, err := os.Open(inputFileArg)
	if err != nil {
		Exit(1, fmt.Sprintf("could not open file: %q", err.Error()))
	}

	return file
}

func ReadLines(file *os.File) []string {
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		Exit(1, fmt.Sprintf("could not read lines: %q", err.Error()))
	}

	return lines
}

func Exit(code int, message string) {
	fmt.Println("ERR: " + message)
	os.Exit(code)
}