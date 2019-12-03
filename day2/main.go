package main

import (
	"fmt"

	"github.com/sshilin/aoc2019/utils"
)

type program []int

func (p program) add(arg1Addr int, arg2Addr int, resAddr int) {
	p[resAddr] = p[arg1Addr] + p[arg2Addr]
}

func (p program) mul(arg1Addr int, arg2Addr int, resAddr int) {
	p[resAddr] = p[arg1Addr] * p[arg2Addr]
}

func (p program) run() {
	ip := 0
	halt := false
	for !halt {
		switch p[ip] {
		case 1:
			p.add(p[ip+1], p[ip+2], p[ip+3])
			ip += 4
		case 2:
			p.mul(p[ip+1], p[ip+2], p[ip+3])
			ip += 4
		case 99:
			halt = true
		}
	}
}

func part1(p program) {
	tmp := make(program, len(p))
	copy([]int(tmp), []int(p))
	tmp[1] = 12
	tmp[2] = 2
	tmp.run()
	fmt.Println("Part 1:", tmp[0]) // 2894520
}

func part2(p program) {
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			tmp := make(program, len(p))
			copy([]int(tmp), []int(p))
			tmp[1] = noun
			tmp[2] = verb
			tmp.run()
			if tmp[0] == 19690720 {
				fmt.Println("Part 2:", 100*noun+verb) // 9342
				return
			}
		}
	}
	fmt.Println("not found")
}

func main() {
	if input, err := utils.ReadCSVInts("input.txt"); err == nil {
		part1(input)
		part2(input)
	} else {
		fmt.Println(err)
	}
}
