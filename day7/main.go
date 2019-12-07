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
	halt   bool
	wait   bool
	output int
}

func createIntcode(program []int) intcode {
	ic := intcode{}
	ic.ram = make([]int, len(program))
	ic.in = make(chan int, 10)
	copy([]int(ic.ram), program)
	return ic
}

func (ic *intcode) load(mode int, addr int) int {
	if mode == 1 {
		return ic.ram[addr]
	}
	return ic.ram[ic.ram[addr]]
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
		_, mop2, mop1, opcode := parseInstruction(ic.ram[ic.ip])
		ic.ip++
		switch opcode {
		case 1: // add
			ic.ram[ic.ram[ic.ip+2]] = ic.load(mop1, ic.ip) + ic.load(mop2, ic.ip+1)
			ic.ip += 3
		case 2: // mul
			ic.ram[ic.ram[ic.ip+2]] = ic.load(mop1, ic.ip) * ic.load(mop2, ic.ip+1)
			ic.ip += 3
		case 3: // input
			select {
			case input := <-ic.in:
				ic.ram[ic.ram[ic.ip]] = input
				ic.ip++
			default:
				ic.ip--
				ic.wait = true
			}
		case 4: // output
			ic.output = ic.ram[ic.ram[ic.ip]]
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
				ic.ram[ic.ram[ic.ip+2]] = 1
			} else {
				ic.ram[ic.ram[ic.ip+2]] = 0
			}
			ic.ip += 3
		case 8: // equals
			if ic.load(mop1, ic.ip) == ic.load(mop2, ic.ip+1) {
				ic.ram[ic.ram[ic.ip+2]] = 1
			} else {
				ic.ram[ic.ram[ic.ip+2]] = 0
			}
			ic.ip += 3
		case 99: // halt
			ic.halt = true
		}
	}
}

func permutation(a []int, size int, exec func([]int)) {
	if size == 1 {
		exec(a)
	}
	for i := 0; i < size; i++ {
		permutation(a, size-1, exec)
		if size%2 == 1 {
			a[0], a[size-1] = a[size-1], a[0]
		} else {
			a[i], a[size-1] = a[size-1], a[i]
		}
	}
	return
}

func part1(input []int) {
	maxSignal := 0
	phases := []int{0, 1, 2, 3, 4}

	permutation(phases, len(phases), func(phases []int) {

		amp1 := createIntcode(input)
		amp2 := createIntcode(input)
		amp3 := createIntcode(input)
		amp4 := createIntcode(input)
		amp5 := createIntcode(input)

		amp1.in <- phases[0]
		amp1.in <- 0
		amp1.run()

		amp2.in <- phases[1]
		amp2.in <- amp1.output
		amp2.run()

		amp3.in <- phases[2]
		amp3.in <- amp2.output
		amp3.run()

		amp4.in <- phases[3]
		amp4.in <- amp3.output
		amp4.run()

		amp5.in <- phases[4]
		amp5.in <- amp4.output
		amp5.run()

		if amp5.output > maxSignal {
			maxSignal = amp5.output
		}
	})
	fmt.Println("Part1:", maxSignal) // 21760
}

func part2(input []int) {
	maxSignal := 0
	phases := []int{5, 6, 7, 8, 9}

	permutation(phases, len(phases), func(phases []int) {
		amp1 := createIntcode(input)
		amp2 := createIntcode(input)
		amp3 := createIntcode(input)
		amp4 := createIntcode(input)
		amp5 := createIntcode(input)

		amp1.in <- phases[0]
		amp2.in <- phases[1]
		amp3.in <- phases[2]
		amp4.in <- phases[3]
		amp5.in <- phases[4]

		for !amp1.halt && !amp2.halt && !amp3.halt && !amp4.halt && !amp5.halt {
			amp1.in <- amp5.output
			amp1.run()

			amp2.in <- amp1.output
			amp2.run()

			amp3.in <- amp2.output
			amp3.run()

			amp4.in <- amp3.output
			amp4.run()

			amp5.in <- amp4.output
			amp5.run()

			if amp5.output > maxSignal {
				maxSignal = amp5.output
			}
		}

	})
	fmt.Println("Part2:", maxSignal) // 69816958
}

func main() {
	if input, err := utils.ReadCSVInts("input.txt"); err == nil {
		part1(input)
		part2(input)
	} else {
		fmt.Println(err)
	}
}
