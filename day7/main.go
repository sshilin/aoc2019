package main

import (
	"errors"
	"fmt"

	"github.com/sshilin/aoc2019/utils"
)

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

func part1() {
	if input, err := utils.ReadCSVInts("input.txt"); err == nil {
		// signal := 0

		phases := []int{0, 1, 2, 3, 4}

		maxSignal := 0
		permutation(phases, len(phases), func(phases []int) {
			fmt.Println("====>", phases)
			signal := 0

			amp1 := createPu("amp1", input)
			amp1.inputReader = func(n int) (int, error) {
				if n == 0 {
					return phases[0], nil
				} else if n == 1 {
					return signal, nil
				} else {
					return -1, errors.New("Unexpected read")
				}
			}

			amp2 := createPu("amp2", input)
			amp2.inputReader = func(n int) (int, error) {
				if n == 0 {
					return phases[1], nil
				} else if n == 1 {
					return signal, nil
				} else {
					return -1, errors.New("Unexpected read")
				}
			}

			amp3 := createPu("amp3", input)
			amp3.inputReader = func(n int) (int, error) {
				if n == 0 {
					return phases[2], nil
				} else if n == 1 {
					return signal, nil
				} else {
					return -1, errors.New("Unexpected read")
				}
			}

			amp4 := createPu("amp4", input)
			amp4.inputReader = func(n int) (int, error) {
				if n == 0 {
					return phases[3], nil
				} else if n == 1 {
					return signal, nil
				} else {
					return -1, errors.New("Unexpected read")
				}
			}

			amp5 := createPu("amp5", input)
			amp5.inputReader = func(n int) (int, error) {
				if n == 0 {
					return phases[4], nil
				} else if n == 1 {
					return signal, nil
				} else {
					return -1, errors.New("Unexpected read")
				}
			}

			amp1.run()
			signal = amp1.reg.output
			amp2.run()
			signal = amp2.reg.output
			amp3.run()
			signal = amp3.reg.output
			amp4.run()
			signal = amp4.reg.output
			amp5.run()
			signal = amp5.reg.output

			if signal > maxSignal {
				maxSignal = signal
			}
		})
		fmt.Println("res:", maxSignal)
	} else {
		fmt.Println(err)
	}
}

func main() {
	part1()
}
