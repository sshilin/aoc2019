package utils

import (
	"bufio"
	"fmt"
	"io"
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

// ReadStrings reads all lines
func ReadStrings(file string) ([]string, error) {
	strings := make([]string, 0)
	inputFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		strings = append(strings, scanner.Text())
	}
	return strings, nil
}

// ReadCSVFile reads all csv records in file
func ReadCSVFile(file string) ([][]string, error) {
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

// ReadDigitsLine reads a single line of digits
func ReadDigitsLine(file string) ([]int, error) {
	digits := make([]int, 0)
	inputFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer inputFile.Close()
	reader := bufio.NewReader(inputFile)
	for {
		s, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if digit, err := strconv.Atoi(string(s)); err == nil {
			digits = append(digits, digit)
		} else {
			return nil, err
		}
	}
	return digits, nil
}

func Track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func Duration(msg string, start time.Time) {
	fmt.Printf("%v: %vms\n", msg, time.Since(start).Milliseconds())
}
