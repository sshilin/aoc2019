package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
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

// ReadCSVFile reads all csv records in file
func ReadCSVFile(file string) ([][]string, error) {
	defer Duration(Track("ReadCSVFile"))
	inputFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer inputFile.Close()
	records := make([][]string, 0)
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		record := strings.Split(scanner.Text(), ",")
		records = append(records, record)
	}

	return records, nil
}

// ReadCSVInts reads a single line csv file where each value is int
func ReadCSVInts(file string) ([]int, error) {
	ints := make([]int, 0)
	records, err := ReadCSVFile(file)
	if err != nil {
		return nil, err
	}
	for _, val := range records[0] {
		i, err := strconv.Atoi(strings.TrimSpace(val))
		if err != nil {
			return nil, err
		}
		ints = append(ints, i)
	}
	return ints, nil
}

func Track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func Duration(msg string, start time.Time) {
	fmt.Printf("%v: %vms\n", msg, time.Since(start).Milliseconds())
}
