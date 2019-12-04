package main

import (
	"fmt"
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
	for i, j := 0, len(factors)-1; i < j; i, j = i+1, j-1 {
		factors[i], factors[j] = factors[j], factors[i]
	}
	return factors
}

func checkP1(password int) bool {
	p := factor(password)
	test1 := true
	test2 := false
	for i := 0; i < len(p)-1; i++ {
		if p[i] > p[i+1] {
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
		if p[i] > p[i+1] {
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
	part1()
	part2()
}
