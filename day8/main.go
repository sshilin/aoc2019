package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

const (
	width  = 25
	height = 6
)

func linear(x int, y int) int {
	return width*y + x
}

func main() {
	var layers [][]int

	layers = append(layers, []int{})
	currentLayer := 0

	inputFile, _ := os.Open("input.txt")
	defer inputFile.Close()
	r := bufio.NewReader(inputFile)

	i := -1
	for {
		i++
		s, _, err := r.ReadRune()
		if err == io.EOF {
			break
		}
		digit, _ := strconv.Atoi(string(s))

		if i%(width*height) == 0 && i > 0 {
			layers = append(layers, []int{})
			currentLayer++
		}
		layers[currentLayer] = append(layers[currentLayer], digit)
	}

	for _, l := range layers {
		fmt.Println(len(l))
	}

	layerWithFewestZeros := 0
	zeros := 99999
	for i, layer := range layers {
		ld := 0
		for _, digit := range layer {
			if digit == 0 {
				ld++
			}
		}
		if ld < zeros {
			zeros = ld
			layerWithFewestZeros = i
		}
	}
	fmt.Println("Layer with fewest zeros is", layerWithFewestZeros)

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

	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			if message[h][w] == 1 {
				fmt.Print(".")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println("") // CJZLP
	}

}
