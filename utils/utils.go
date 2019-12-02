package utils

import (
	"bufio"
	"encoding/csv"
	"os"
	"strconv"
	"strings"
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

// ReadCSVInts read a single line csv file where each value is int
func ReadCSVInts(file string) ([]int, error) {
	ints := make([]int, 0)
	inputFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer inputFile.Close()
	r := csv.NewReader(inputFile)
	record, err := r.Read()
	if err != nil {
		return nil, err
	}
	for _, val := range record {
		i, err := strconv.Atoi(strings.TrimSpace(val))
		if err != nil {
			return nil, err
		}
		ints = append(ints, i)
	}
	return ints, nil
}
