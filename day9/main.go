package main

import (
	"fmt"

	"github.com/sshilin/aoc2019/utils"
)

type intcode struct {
	ram []int

	// input
	in chan int

	// registers
	ip     int
	base   int
	halt   bool
	wait   bool
	output int
}

func createIntcode(program []int) intcode {
	ic := intcode{}
	ic.ram = make([]int, 8*1024) // 64K ought to be enough :-)
	ic.in = make(chan int, 10)
	copy([]int(ic.ram), program)
	return ic
}

func (ic *intcode) load(mode int, addr int) int {
	var value int
	switch mode {
	case 0:
		value = ic.ram[ic.ram[addr]]
	case 1:
		value = ic.ram[addr]
	case 2:
		value = ic.ram[ic.base+ic.ram[addr]]
	default:
		panic(fmt.Errorf("Unknown mode %d", mode))
	}
	return value
}

func (ic *intcode) store(mode int, addr int, val int) {
	switch mode {
	case 0:
		ic.ram[ic.ram[addr]] = val
	case 1:
		ic.ram[addr] = val
	case 2:
		ic.ram[ic.base+ic.ram[addr]] = val
	default:
		panic(fmt.Errorf("Unknown mode %d", mode))
	}
}

func parseInstruction(instruction int) (mop3, mop2, mop1, opcode int) {
	opcode = instruction % 100
	mop1 = (instruction % 1000) / 100
	mop2 = (instruction % 10000) / 1000
	mop3 = (instruction % 100000) / 10000
	return
}

func (ic *intcode) run() {
	ic.halt = false
	ic.wait = false
	for !ic.halt && !ic.wait {
		mop3, mop2, mop1, opcode := parseInstruction(ic.ram[ic.ip])
		ic.ip++
		switch opcode {
		case 1: // add
			ic.store(mop3, ic.ip+2, ic.load(mop1, ic.ip)+ic.load(mop2, ic.ip+1))
			ic.ip += 3
		case 2: // mul
			ic.store(mop3, ic.ip+2, ic.load(mop1, ic.ip)*ic.load(mop2, ic.ip+1))
			ic.ip += 3
		case 3: // input
			select {
			case input := <-ic.in:
				ic.store(mop1, ic.ip, input)
				ic.ip++
			default:
				ic.ip--
				ic.wait = true
			}
		case 4: // output
			ic.output = ic.load(mop1, ic.ip)
			ic.ip++
		case 5: // jump-if-true
			if ic.load(mop1, ic.ip) != 0 {
				ic.ip = ic.load(mop2, ic.ip+1)
			} else {
				ic.ip += 2
			}
		case 6: // jump-if-false
			if ic.load(mop1, ic.ip) == 0 {
				ic.ip = ic.load(mop2, ic.ip+1)
			} else {
				ic.ip += 2
			}
		case 7: // less
			if ic.load(mop1, ic.ip) < ic.load(mop2, ic.ip+1) {
				ic.store(mop3, ic.ip+2, 1)
			} else {
				ic.store(mop3, ic.ip+2, 0)
			}
			ic.ip += 3
		case 8: // equals
			if ic.load(mop1, ic.ip) == ic.load(mop2, ic.ip+1) {
				ic.store(mop3, ic.ip+2, 1)
			} else {
				ic.store(mop3, ic.ip+2, 0)
			}
			ic.ip += 3
		case 9: // base
			ic.base += ic.load(mop1, ic.ip)
			ic.ip++
		case 99: // halt
			ic.halt = true
		}
	}
}

func part1(input []int) {
	ic := createIntcode(input)
	ic.in <- 1
	ic.run()
	fmt.Println("Part1:", ic.output) // 3906448201
}

func part2(input []int) {
	ic := createIntcode(input)
	ic.in <- 2
	ic.run()
	fmt.Println("Part2:", ic.output) // 59785
}

func main() {
	if input, err := utils.ReadCSVInts("input.txt"); err == nil {
		part1(input)
		part2(input)
	} else {
		fmt.Println(err)
	}
}
