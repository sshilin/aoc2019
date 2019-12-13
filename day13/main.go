package main

import (
	"fmt"

	"github.com/sshilin/aoc2019/utils"
)

type intcode struct {
	ram []int

	// I/O
	in  chan int
	out chan int

	// registers
	ip        int
	base      int
	halt      bool
	wait      bool
	interrupt int // -1 - running; 0 - blocked on output; 1 - blocked on input
}

func createIntcode(program []int) intcode {
	ic := intcode{}
	ic.ram = make([]int, 8*1024) // 64K ought to be enough :-)
	ic.in = make(chan int, 10)
	ic.out = make(chan int, 3)
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
	ic.interrupt = -1
	for !ic.halt && ic.interrupt == -1 {
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
				ic.interrupt = 1
			}
		case 4: // output
			select {
			case ic.out <- ic.load(mop1, ic.ip):
				ic.ip++
			default:
				ic.ip--
				ic.interrupt = 0
			}
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
	field := [24][44]string{}

	blockTiles := 0
	ic := createIntcode(input)
	for !ic.halt {
		ic.run()
		x := <-ic.out
		y := <-ic.out
		id := <-ic.out

		switch id {
		case 0:
			field[y][x] = " "
		case 1:
			field[y][x] = "."
		case 2:
			blockTiles++
			field[y][x] = "="
		case 3:
			field[y][x] = "_"
		case 4:
			field[y][x] = "o"
		}
	}
	fmt.Println("Part1:", blockTiles) // 412

	for y := 0; y < 24; y++ {
		for x := 0; x < 44; x++ {
			fmt.Print(field[y][x])
		}
		fmt.Println("")
	}
}

func part2(input []int) {
	ic := createIntcode(input)
	ic.ram[0] = 2
	ball := 0
	paddle := 0
	score := 0

	field := [50][50]string{}

	for !ic.halt {
		ic.run()
		if ic.interrupt == 0 {
			x := <-ic.out
			y := <-ic.out
			id := <-ic.out

			if x == -1 && y == 0 {
				fmt.Println("SCORE")
				if id > score {
					score = id
				}
			} else {
				switch id {
				case 0:
					fmt.Println("empty", x, y)
					field[y][x] = " "
				case 1:
					fmt.Println("wall", x, y)
					field[y][x] = "."
				case 2:
					fmt.Println("brick", x, y)
					field[y][x] = "="
				case 3:
					field[y][x] = "_"
					fmt.Println("paddle", x, y)
					paddle = x
				case 4:
					fmt.Println("ball", x, y)
					field[y][x] = "o"
					ball = x
				}
			}
		} else if ic.interrupt == 1 {

			// for y := 0; y < 50; y++ {
			// 	for x := 0; x < 50; x++ {
			// 		fmt.Print(field[y][x])
			// 	}
			// 	fmt.Println("")
			// }
			fmt.Println("=========================================")
			if ball > paddle {
				fmt.Println("RIGHT")
				ic.in <- 1
			} else if ball < paddle {
				fmt.Println("LEFT")
				ic.in <- -1
			} else {
				fmt.Println("NO MOVE")
				ic.in <- 0
			}
		}
	}
	fmt.Println("Score", score)
}

func main() {
	if input, err := utils.ReadCSVInts("input.txt"); err == nil {
		// part1(input)
		part2(input)
	} else {
		fmt.Println(err)
	}
}
