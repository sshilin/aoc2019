package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readInput() (input []int) {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, i)
	}
	return
}

func part1() {
	fuel := 0
	for _, mass := range readInput() {
		fuel += (mass / 3) - 2
	}
	fmt.Println("Part 1: ", fuel)
}

func part2() {
	fuel := 0
	for _, mass := range readInput() {
		for (mass/3)-2 > 0 {
			mass = (mass / 3) - 2
			fuel += mass
		}
	}
	fmt.Println("Part 2: ", fuel)
}

func main() {
	part1() // 3404722
	part2() // 5104215
}
