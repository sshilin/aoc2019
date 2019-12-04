package main

import (
	"fmt"

	"github.com/sshilin/aoc2019/utils"
)

const (
	start = 130254
	end   = 678275
)

func factor(number int) []int {
	var factors []int
	for number > 0 {
		factors = append(factors, number%10)
		number /= 10
	}
	return factors
}

func checkP1(password int) bool {
	p := factor(password)
	test1 := true
	test2 := false
	for i := 0; i < len(p)-1; i++ {
		if p[i] < p[i+1] {
			test1 = false
		}
		if p[i] == p[i+1] {
			test2 = true
		}
	}
	return test1 && test2
}

func checkP2(password int) bool {
	p := factor(password)
	test1 := true
	test2 := false
	for i := 0; i < len(p)-1; i++ {
		if p[i] < p[i+1] {
			test1 = false
		}
		if (p[i] == p[i+1]) && (i == 0 || p[i-1] != p[i]) && (i == len(p)-2 || p[i+1] != p[i+2]) {
			test2 = true
		}
	}
	return test1 && test2
}

func part1() {
	count := 0
	for i := start; i <= end; i++ {
		if checkP1(i) {
			count++
		}
	}
	fmt.Println("Part1:", count) // 2090
}

func part2() {
	count := 0
	for i := start; i <= end; i++ {
		if checkP2(i) {
			count++
		}
	}
	fmt.Println("Part2:", count) // 1419
}

func main() {
	defer utils.Duration(utils.Track("main"))
	part1()
	part2()
}
