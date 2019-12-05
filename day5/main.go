package main

import (
	"fmt"

	"github.com/sshilin/aoc2019/utils"
)

type memory []int

func (m memory) add(op1 int, op2 int, resPtr int) {
	m[resPtr] = op1 + op2
}

func (m memory) mul(op1 int, op2 int, resPtr int) {
	m[resPtr] = op1 * op2
}

func (m memory) load(ptr int) int {
	return m[m[ptr]]
}

func (m memory) store(val int, ptr int) {
	m[ptr] = val
}

func (m memory) run() {
	ip := 0
	halt := false
	for !halt {
		// fmt.Println("ip", ip, m[ip])
		_, mop2, mop1, opcode := parseInstruction(m[ip])
		ip++
		// fmt.Println("Instruction:", mop3, mop2, mop1, opcode)
		switch opcode {
		case 1: // opcode 1 - add
			op1, op2 := 0, 0
			if mop1 == 1 {
				op1 = m[ip]
			} else {
				op1 = m.load(ip)
			}
			if mop2 == 1 {
				op2 = m[ip+1]
			} else {
				op2 = m.load(ip + 1)
			}
			m.store(op1+op2, m[ip+2])
			ip += 3
		case 2: // opcode 2 - mul
			op1, op2 := 0, 0
			if mop1 == 1 {
				op1 = m[ip]
			} else {
				op1 = m.load(ip)
			}
			if mop2 == 1 {
				op2 = m[ip+1]
			} else {
				op2 = m.load(ip + 1)
			}
			m.store(op1*op2, m[ip+2])
			ip += 3
		case 3: // opcode 3 - input
			m.store(1, m[ip])
			ip += 1
		case 4: // opcode 3 - output
			fmt.Println("=>", m[m[ip]])
			ip += 1
		case 99: // opcode 99 - halt
			halt = true
		}
	}
}

func part1(m memory) {
	tmp := make(memory, len(m))
	copy([]int(tmp), []int(m))
	tmp[1] = 12
	tmp[2] = 2
	tmp.run()
	fmt.Println("Part 1:", tmp[0]) // 2894520
}

func part2(m memory) {
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			tmp := make(memory, len(m))
			copy([]int(tmp), []int(m))
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

func parseInstruction(instruction int) (mop3, mop2, mop1, opcode int) {
	opcode = instruction % 100
	mop1 = (instruction % 1000) / 100
	mop2 = (instruction % 10000) / 1000
	mop3 = (instruction % 100000) / 10000
	return
}

func main() {
	if input, err := utils.ReadCSVInts("input.txt"); err == nil {
		// part1(input)
		// part2(input)
		mem := make(memory, len(input))
		copy([]int(mem), input)
		// fmt.Println(mem)
		mem.run() // part1: 15508323
		// fmt.Println(mem)
	} else {
		fmt.Println(err)
	}
}
