package main

import (
	"fmt"
	"math"

	"github.com/sshilin/aoc2019/utils"
)

type point struct {
	x, y, z int
}

func part1() {
	position := [][]int{
		{13, 9, 5},
		{8, 14, -2},
		{-5, 4, 11},
		{2, -6, 1},
	}
	velocity := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	for step := 0; step < 1000; step++ {
		// apply gravity
		for i := 0; i < len(position); i++ {
			for j := 0; j < len(position); j++ {
				if i == j {
					continue
				}
				for k := 0; k < 3; k++ {
					if position[i][k] < position[j][k] {
						velocity[i][k]++
					} else if position[i][k] > position[j][k] {
						velocity[i][k]--
					}
				}
			}
		}
		// apply velocity
		for i := 0; i < len(position); i++ {
			for j := 0; j < 3; j++ {
				position[i][j] += velocity[i][j]
			}
		}
	}
	// calculate energy
	totalEnergy := 0
	for i := 0; i < len(position); i++ {
		pot := 0
		kin := 0
		for j := 0; j < 3; j++ {
			pot += int(math.Abs(float64(position[i][j])))
			kin += int(math.Abs(float64(velocity[i][j])))
		}
		totalEnergy += pot * kin
	}
	fmt.Println(position)
	fmt.Println(velocity)
	fmt.Println("Total energy:", totalEnergy) // 6490
}

func part2() {
	origPosition := [][]int{
		{13, 9, 5},
		{8, 14, -2},
		{-5, 4, 11},
		{2, -6, 1},
	}
	velocity := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	step := 0
	periodX := -1
	periodY := -1
	periodZ := -1
	position := make([][]int, len(origPosition))
	utils.Copy2dInt(position, origPosition)
	for true {
		// apply gravity
		for i := 0; i < len(position); i++ {
			for j := 0; j < len(position); j++ {
				if i == j {
					continue
				}
				for k := 0; k < 3; k++ {
					if position[i][k] < position[j][k] {
						velocity[i][k]++
					} else if position[i][k] > position[j][k] {
						velocity[i][k]--
					}
				}
			}
		}

		// apply velocity
		for i := 0; i < len(position); i++ {
			for j := 0; j < 3; j++ {
				position[i][j] += velocity[i][j]
			}
		}

		step++

		if periodX == -1 &&
			position[0][0] == origPosition[0][0] &&
			position[1][0] == origPosition[1][0] &&
			position[2][0] == origPosition[2][0] &&
			position[3][0] == origPosition[3][0] &&
			velocity[0][0] == 0 &&
			velocity[1][0] == 0 &&
			velocity[2][0] == 0 &&
			velocity[3][0] == 0 {
			periodX = step
		}
		if periodY == -1 &&
			position[0][1] == origPosition[0][1] &&
			position[1][1] == origPosition[1][1] &&
			position[2][1] == origPosition[2][1] &&
			position[3][1] == origPosition[3][1] &&
			velocity[0][1] == 0 &&
			velocity[1][1] == 0 &&
			velocity[2][1] == 0 &&
			velocity[3][1] == 0 {
			periodY = step
		}
		if periodZ == -1 &&
			position[0][2] == origPosition[0][2] &&
			position[1][2] == origPosition[1][2] &&
			position[2][2] == origPosition[2][2] &&
			position[3][2] == origPosition[3][2] &&
			velocity[0][2] == 0 &&
			velocity[1][2] == 0 &&
			velocity[2][2] == 0 &&
			velocity[3][2] == 0 {
			periodZ = step
		}

		if periodX != -1 && periodY != -1 && periodZ != -1 {
			break
		}

	}
	fmt.Println(periodX, periodY, periodZ)                     // 268296 161428 102356
	fmt.Println("Part2:", lcm(lcm(periodX, periodY), periodZ)) // 277068010964808
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	part1()
	part2()
}
