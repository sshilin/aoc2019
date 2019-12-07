package main

import (
	"fmt"
	"os"
)

type registers struct {
	ip      int
	halt    bool
	inCount int // count how many times PU requested inputs
	output  int
}

type processingUnit struct {
	name        string
	ram         []int
	reg         registers
	inputReader func(n int) (int, error)
}

func createPu(name string, program []int) processingUnit {
	pu := processingUnit{}
	pu.name = name
	pu.loadProgram(program)
	return pu
}

func (pu *processingUnit) mstore(val int, addr int) {
	pu.ram[addr] = val
}

func (pu *processingUnit) mload(mode int, addr int) int {
	if mode == 1 {
		return pu.ram[addr]
	}
	return pu.ram[pu.ram[addr]]
}

func (pu *processingUnit) pmload(addr int) int {
	return pu.mload(0, addr)
}

func (pu *processingUnit) imload(addr int) int {
	return pu.mload(1, addr)
}

func (pu *processingUnit) loadProgram(program []int) {
	pu.ram = make([]int, len(program))
	copy([]int(pu.ram), program)
}

func parseInstruction(instruction int) (mop3, mop2, mop1, opcode int) {
	opcode = instruction % 100
	mop1 = (instruction % 1000) / 100
	mop2 = (instruction % 10000) / 1000
	mop3 = (instruction % 100000) / 10000
	return
}

func (pu *processingUnit) run() {
	pu.reg.halt = false
	for !pu.reg.halt {
		_, mop2, mop1, opcode := parseInstruction(pu.imload(pu.reg.ip))
		pu.reg.ip++
		switch opcode {
		case 1: // add
			pu.mstore(pu.mload(mop1, pu.reg.ip)+pu.mload(mop2, pu.reg.ip+1), pu.imload(pu.reg.ip+2))
			pu.reg.ip += 3
		case 2: // mul
			pu.mstore(pu.mload(mop1, pu.reg.ip)*pu.mload(mop2, pu.reg.ip+1), pu.imload(pu.reg.ip+2))
			pu.reg.ip += 3
		case 3: // input
			if input, err := pu.inputReader(pu.reg.inCount); err == nil {
				fmt.Println(pu.name, " read ", input)
				pu.mstore(input, pu.imload(pu.reg.ip))
				pu.reg.inCount++
				pu.reg.ip++
			} else {
				fmt.Println(pu.name, "-", err)
				os.Exit(1)
			}
		case 4: // output
			pu.reg.output = pu.pmload(pu.reg.ip)
			pu.reg.ip++
		case 5: // jump-if-true
			if pu.mload(mop1, pu.reg.ip) != 0 {
				pu.reg.ip = pu.mload(mop2, pu.reg.ip+1)
			} else {
				pu.reg.ip += 2
			}
		case 6: // jump-if-false
			if pu.mload(mop1, pu.reg.ip) == 0 {
				pu.reg.ip = pu.mload(mop2, pu.reg.ip+1)
			} else {
				pu.reg.ip += 2
			}
		case 7: // less
			if pu.mload(mop1, pu.reg.ip) < pu.mload(mop2, pu.reg.ip+1) {
				pu.mstore(1, pu.imload(pu.reg.ip+2))
			} else {
				pu.mstore(0, pu.imload(pu.reg.ip+2))
			}
			pu.reg.ip += 3
		case 8: // equals
			if pu.mload(mop1, pu.reg.ip) == pu.mload(mop2, pu.reg.ip+1) {
				pu.mstore(1, pu.imload(pu.reg.ip+2))
			} else {
				pu.mstore(0, pu.imload(pu.reg.ip+2))
			}
			pu.reg.ip += 3
		case 99: // halt
			pu.reg.halt = true
		}
	}
}
