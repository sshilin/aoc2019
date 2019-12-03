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

func main() {
	if input, err := utils.ReadCSVInts("input.txt"); err == nil {
		input[1] = 12
		input[2] = 2
		program(input).run()
		fmt.Println("Part 1:", input[0]) // 2894520
	} else {
		fmt.Println(err)
	}
}
