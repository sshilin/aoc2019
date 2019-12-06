package main

import (
	"fmt"
	"strings"

	"github.com/sshilin/aoc2019/utils"
)

type node struct {
	object   string
	orbiting string
}

func part1(input []string) {
	graph := make(map[string]node)
	for _, line := range input {
		part := strings.Split(line, ")")
		graph[part[1]] = node{part[1], part[0]}
	}
	count := 0
	for _, obj := range graph {
		currentObj := obj
		count++
		for currentObj.orbiting != "COM" {
			currentObj = graph[currentObj.orbiting]
			count++
		}
	}
	fmt.Println("Part 1:", count) // 245089
}

func part2(input []string) {
	graph := make(map[string]node)
	for _, line := range input {
		part := strings.Split(line, ")")
		graph[part[1]] = node{part[1], part[0]}
	}

	san2com := make(map[string]int)
	you2com := make(map[string]int)

	tmpSan := graph["SAN"]
	jump := 0
	for tmpSan.orbiting != "COM" {
		tmpSan = graph[tmpSan.orbiting]
		san2com[tmpSan.object] = jump
		jump++
	}

	tmpYou := graph["YOU"]
	jump = 0
	for tmpYou.orbiting != "COM" {
		tmpYou = graph[tmpYou.orbiting]
		you2com[tmpYou.object] = jump
		jump++
	}

	minDistance := len(graph)
	for obj, youJumps := range you2com {
		if sanJumps, ok := san2com[obj]; ok {
			if distance := youJumps + sanJumps; distance < minDistance {
				minDistance = distance
			}
		}
	}

	fmt.Println("Part 2:", minDistance) // 511
}

func main() {
	if input, err := utils.ReadStrings("input.txt"); err == nil {
		part1(input)
		part2(input)
	} else {
		fmt.Println(err)
	}
}
