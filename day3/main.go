package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/sshilin/aoc2019/utils"
)

type point struct {
	x, y int
}

func createWire(turns []string) []point {
	wire := make([]point, 0)
	wire = append(wire, point{})
	for _, turn := range turns {
		direction := turn[0:1]
		distance, _ := strconv.Atoi(turn[1:])
		lastPoint := wire[len(wire)-1]
		switch direction {
		case "U":
			for y := lastPoint.y + 1; y <= lastPoint.y+distance; y++ {
				wire = append(wire, point{lastPoint.x, y})
			}
		case "R":
			for x := lastPoint.x + 1; x <= lastPoint.x+distance; x++ {
				wire = append(wire, point{x, lastPoint.y})
			}
		case "D":
			for y := lastPoint.y - 1; y >= lastPoint.y-distance; y-- {
				wire = append(wire, point{lastPoint.x, y})
			}
		case "L":
			for x := lastPoint.x - 1; x >= lastPoint.x-distance; x-- {
				wire = append(wire, point{x, lastPoint.y})
			}
		}
	}
	return wire
}

func intersection(wire1 []point, wire2 []point) []point {
	zeroPoint := point{0, 0}
	p := make([]point, 0)
	for _, wire1Point := range wire1 {
		for _, wire2Point := range wire2 {
			if wire1Point == wire2Point && wire1Point != zeroPoint {
				p = append(p, wire1Point)
			}
		}
	}
	return p
}

func distance(a point, b point) int {
	return int(math.Abs(float64(a.x)-float64(b.x)) + math.Abs(float64(a.y)-float64(b.y)))
}

func part1(wire1 []point, wire2 []point) {
	minDistance := math.MaxInt64
	for _, p := range intersection(wire1, wire2) {
		if distance(point{0, 0}, p) < minDistance {
			minDistance = distance(point{0, 0}, p)
		}
	}
	fmt.Println("Part1:", minDistance)
}

func main() {
	if records, err := utils.ReadCSVFile("input.txt"); err == nil {
		wire1 := createWire(records[0])
		wire2 := createWire(records[1])
		part1(wire1, wire2) // 870
	} else {
		fmt.Println(err)
	}
}
