package utils

import (
	"bufio"
	"os"
	"strconv"
)

// ReadInts reads a file where each line is an integer
func ReadInts(file string) ([]int, error) {
	ints := make([]int, 0)
	inputFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		ints = append(ints, i)
	}
	return ints, nil
}
