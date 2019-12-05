package main

import (
	"fmt"
	"math"

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

func checkPasswords(start int, end int, checker func(int) bool, result chan int) {
	count := 0
	for i := start; i <= end; i++ {
		if checker(i) {
			count++
		}
	}
	result <- count
}

func partition(length int, parts int) (full int, remainder int) {
	full = int(math.Ceil(float64(length) / float64(parts)))
	remainder = length % full
	return
}

func process(chunks int) {
	part1Channel := make(chan int, chunks)
	part2Channel := make(chan int, chunks)

	full, remainder := partition(end-start+1, chunks)

	if remainder != 0 {
		go checkPasswords(end-remainder+1, end, checkP1, part1Channel)
		go checkPasswords(end-remainder+1, end, checkP2, part2Channel)
	}

	for i := 0; (remainder != 0 && i < chunks-1) || (remainder == 0 && i < chunks); i++ {
		go checkPasswords(start+full*i, start+full*(i+1)-1, checkP1, part1Channel)
		go checkPasswords(start+full*i, start+full*(i+1)-1, checkP2, part2Channel)
	}

	part1Res := 0
	part2Res := 0

	for i := 0; i < chunks*2; i++ {
		select {
		case count := <-part1Channel:
			part1Res += count
		case count := <-part2Channel:
			part2Res += count
		}
	}

	fmt.Println("Part1:", part1Res) // 2090
	fmt.Println("Part2:", part2Res) // 1419
}

func main() {
	defer utils.Duration(utils.Track("main"))
	process(4)
}
