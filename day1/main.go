package main

import (
	"fmt"

	"github.com/sshilin/aoc2019/utils"
)

func part1(input []int) int {
	fuel := 0
	for _, mass := range input {
		fuel += (mass / 3) - 2
	}
	return fuel
}

func part2(input []int) int {
	fuel := 0
	for _, mass := range input {
		for (mass/3)-2 > 0 {
			mass = (mass / 3) - 2
			fuel += mass
		}
	}
	return fuel
}

func main() {
	if input, err := utils.ReadInts("input.txt"); err == nil {
		fmt.Println("Part 1:", part1(input)) // 3404722
		fmt.Println("Part 2:", part2(input)) // 5104215
	} else {
		fmt.Println(err)
	}
}
