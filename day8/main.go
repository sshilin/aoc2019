package main

import (
	"fmt"

	"github.com/sshilin/aoc2019/utils"
)

const (
	width  = 25
	height = 6
)

func part1(layers [][]int) {
	layerWithFewestZeros := 0
	fewestZeros := width
	for i, layer := range layers {
		zeros := 0
		for _, digit := range layer {
			if digit == 0 {
				zeros++
			}
		}
		if zeros < fewestZeros {
			fewestZeros = zeros
			layerWithFewestZeros = i
		}
	}
	ones := 0
	twos := 0
	for _, digit := range layers[layerWithFewestZeros] {
		if digit == 1 {
			ones++
		}
		if digit == 2 {
			twos++
		}
	}
	fmt.Println("Part1:", ones*twos) // 2286
}

func part2(layers [][]int) {
	var message [height][width]int

	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			message[h][w] = 2
		}
	}

	for l := len(layers) - 1; l >= 0; l-- {
		layer := layers[l]
		for h := 0; h < height; h++ {
			for w := 0; w < width; w++ {
				if layer[width*h+w] == 0 {
					message[h][w] = 0
				}
				if layer[width*h+w] == 1 {
					message[h][w] = 1
				}
			}
		}
	}

	fmt.Println("Part2:")
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			if message[h][w] == 1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println("") // CJZLP
	}
}

func main() {
	if input, err := utils.ReadDigitsLine("input.txt"); err == nil {
		var layers [][]int

		layers = append(layers, []int{})
		for i, digit := range input {
			if i%(width*height) == 0 && i > 0 {
				layers = append(layers, []int{})
			}
			layers[len(layers)-1] = append(layers[len(layers)-1], digit)
		}

		part1(layers)
		part2(layers)
	} else {
		fmt.Println(err)
	}
}
