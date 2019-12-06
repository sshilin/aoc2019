package main

import (
	"fmt"

	"github.com/sshilin/aoc2019/utils"
)

type memory []int

func (m memory) store(val int, addr int) {
	m[addr] = val
}

func (m memory) load(mode int, addr int) int {
	if mode == 1 {
		return m[addr]
	}
	return m[m[addr]]
}

func parseInstruction(instruction int) (mop3, mop2, mop1, opcode int) {
	opcode = instruction % 100
	mop1 = (instruction % 1000) / 100
	mop2 = (instruction % 10000) / 1000
	mop3 = (instruction % 100000) / 10000
	return
}

func (m memory) run(intput int) {
	ip := 0
	halt := false
	for !halt {
		_, mop2, mop1, opcode := parseInstruction(m[ip])
		ip++
		switch opcode {
		case 1: // add
			m.store(m.load(mop1, ip)+m.load(mop2, ip+1), m[ip+2])
			ip += 3
		case 2: // mul
			m.store(m.load(mop1, ip)*m.load(mop2, ip+1), m[ip+2])
			ip += 3
		case 3: // input
			m.store(intput, m[ip])
			ip++
		case 4: // output
			fmt.Println(m[m[ip]])
			ip++
		case 5: // jump-if-true
			if m.load(mop1, ip) != 0 {
				ip = m.load(mop2, ip+1)
			} else {
				ip += 2
			}
		case 6: // jump-if-false
			if m.load(mop1, ip) == 0 {
				ip = m.load(mop2, ip+1)
			} else {
				ip += 2
			}
		case 7: // less
			if m.load(mop1, ip) < m.load(mop2, ip+1) {
				m.store(1, m[ip+2])
			} else {
				m.store(0, m[ip+2])
			}
			ip += 3
		case 8: // equals
			if m.load(mop1, ip) == m.load(mop2, ip+1) {
				m.store(1, m[ip+2])
			} else {
				m.store(0, m[ip+2])
			}
			ip += 3
		case 99: // halt
			halt = true
		}
	}
}

func part1(input []int) {
	mem := make(memory, len(input))
	copy([]int(mem), input)
	mem.run(1) // 15508323
}

func part2(input []int) {
	mem := make(memory, len(input))
	copy([]int(mem), input)
	mem.run(5) // 9006327
}

func main() {
	if input, err := utils.ReadCSVInts("input.txt"); err == nil {
		part1(input)
		part2(input)
	} else {
		fmt.Println(err)
	}
}
