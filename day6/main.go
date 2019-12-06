package main

import (
	"fmt"
	"strings"

	"github.com/sshilin/aoc2019/utils"
)

type node struct {
	id       string
	orbiting string
}

func part1(input []string) {
	graph := make(map[string]node)
	for _, in := range input {
		n := strings.Split(in, ")")
		graph[n[1]] = node{n[1], n[0]}
	}

	count := 0
	for _, node := range graph {
		// fmt.Println(key, node)
		tmp := node
		count++
		for tmp.orbiting != "COM" {
			// fmt.Println("==>", tmp)
			tmp = graph[tmp.orbiting]
			count++
		}
	}
	fmt.Println(count)
	// part1: 245089
}

func part2(input []string) {
	graph := make(map[string]node)
	// sanStart := ""
	// youStart := ""
	for _, in := range input {
		n := strings.Split(in, ")")
		graph[n[1]] = node{n[1], n[0]}
	}

	san2com := make(map[string]int)
	you2com := make(map[string]int)

	tmpSan := graph["SAN"]
	countSan := 0
	for tmpSan.orbiting != "COM" {
		tmpSan = graph[tmpSan.orbiting]
		san2com[tmpSan.id] = countSan
		countSan++
	}

	fmt.Println(san2com)

	tmpYou := graph["YOU"]
	countYou := 0
	for tmpYou.orbiting != "COM" {
		tmpYou = graph[tmpYou.orbiting]
		you2com[tmpYou.id] = countYou
		countYou++
	}

	fmt.Println(you2com)

	minDistance := len(graph)
	for k, yv := range you2com {
		if sv, ok := san2com[k]; ok {
			if yv+sv < minDistance {
				minDistance = yv + sv
			}
		}
	}
	fmt.Println(minDistance) // 511
}

func main() {
	input, _ := utils.ReadStrings("input.txt")

	part2(input)

}
