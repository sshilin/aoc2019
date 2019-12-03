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

func intersections(wire1 []point, wire2 []point) []point {
	m := make(map[point]bool, len(wire1))
	p := make([]point, 0, len(wire2))
	for _, wire1Point := range wire1 {
		m[wire1Point] = true
	}
	delete(m, point{0, 0})
	for _, wire2Point := range wire2 {
		if _, ok := m[wire2Point]; ok {
			p = append(p, wire2Point)
		}
	}
	return p
}

func distance(a point, b point) int {
	return int(math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y)))
}

func part1(wire1 []point, wire2 []point) {
	minDistance := math.MaxInt64
	for _, p := range intersections(wire1, wire2) {
		if distance(point{0, 0}, p) < minDistance {
			minDistance = distance(point{0, 0}, p)
		}
	}
	fmt.Println("Part1:", minDistance)
}

func part2(wire1 []point, wire2 []point) {
	minSteps := math.MaxInt64
	for _, p := range intersections(wire1, wire2) {
		steps := 0
		for i, wire1Point := range wire1 {
			if wire1Point == p {
				steps += i
				break
			}
		}
		for i, wire2Point := range wire2 {
			if wire2Point == p {
				steps += i
				break
			}
		}
		if steps < minSteps {
			minSteps = steps
		}
	}
	fmt.Println("Part2:", minSteps)
}

func main() {
	defer utils.Duration(utils.Track("main"))
	if records, err := utils.ReadCSVFile("input.txt"); err == nil {
		wire1 := createWire(records[0])
		wire2 := createWire(records[1])
		part1(wire1, wire2) // 870
		part2(wire1, wire2) // 13698
	} else {
		fmt.Println(err)
	}
}
