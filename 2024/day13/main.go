package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"strings"
)

//go:embed input.txt
var input string

const machineFormat = `Button A: X+%d, Y+%d
Button B: X+%d, Y+%d
Prize: X=%d, Y=%d`

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Printf("Running AoC day %d solution (part %d)\n", 13, part)

	var answer string
	if part == 1 {
		answer = part1(input)
	} else if part == 2 {
		answer = part2(input)
	} else {
		panic(fmt.Errorf("part must be 1 or 2 %d", part))	
	}

	fmt.Println("Output:", answer)
	err := os.WriteFile(fmt.Sprintf("output%d.txt", part), []byte(answer), 0644)
	if err != nil {
		panic(err)
	}
}

type vec struct {
	x, y int
}

func (v vec) mag() int {
	return v.x +v.y
}

func (v vec) scale(s int) vec {
	return vec{v.x*s, v.y*s}
}

func (v vec) add(b vec) vec {
	return vec{v.x + b.x, v.y + b.y}
}

func (v vec) equals(b vec) bool {
	return v.x == b.x && v.y == b.y
}

const (
	maxPresses int = 100
	costRatio int = 3
	offset int = 10000000000000
)

func minCost(vA, vB, vPrize vec) int {
	cost := 0
	var big, small vec = vA, vB
	var bigCost, smallCost int = 3, 1
	if vA.mag() < vB.scale(costRatio).mag() {
		big, small = vB, vA
		bigCost, smallCost = 1, 3
	}

	for bigPress := range maxPresses {
		for smallPress := range maxPresses {
			target := big.scale(bigPress).add(small.scale(smallPress))
			if target.x > vPrize.x || target.y > vPrize.y { break }
			if !target.equals(vPrize) { continue }
			thisCost := bigPress*bigCost + smallPress*smallCost
			if cost == 0 || thisCost < cost { cost = thisCost }
			
		}
	}

	return cost
}

func minCostNoLimit(vA, vB, vPrize vec) int {
	cost := 0
	var big, small vec = vA, vB
	var bigCost, smallCost int = 3, 1
	if vA.mag() < vB.scale(costRatio).mag() {
		big, small = vB, vA
		bigCost, smallCost = 1, 3
	}

	for bigPress := 0;; bigPress++ {
		bs := big.scale(bigPress)
		if bs.x > vPrize.x || bs.y > vPrize.y { break }
		for smallPress := 0;; smallPress++ {
			target := big.scale(bigPress).add(small.scale(smallPress))
			if target.x > vPrize.x || target.y > vPrize.y { break }
			if !target.equals(vPrize) { continue }
			thisCost := bigPress*bigCost + smallPress*smallCost
			if cost == 0 || thisCost < cost { cost = thisCost }
		}
	}

	return cost
}

func part1(input string) string {
	total := 0
	for _, machine := range strings.Split(strings.TrimRight(input, "\n"), "\n\n") {
		vA, vB, vPrize := vec{0,0},vec{0,0},vec{0,0}
		_, err := fmt.Sscanf(machine, machineFormat, &vA.x,&vA.y,&vB.x,&vB.y,&vPrize.x, &vPrize.y)
		if err != nil {
			panic(err)
		}
		total += minCost(vA, vB, vPrize)
	}
	return fmt.Sprintf("%d", total)
}

func part2(input string) string {
	total := 0
	for _, machine := range strings.Split(strings.TrimRight(input, "\n"), "\n\n") {
		fmt.Printf("\033[1K%d", total)
		vA, vB, vPrize := vec{0,0},vec{0,0},vec{0,0}
		_, err := fmt.Sscanf(machine, machineFormat, &vA.x,&vA.y,&vB.x,&vB.y,&vPrize.x, &vPrize.y)
		if err != nil {
			panic(err)
		}
		total += minCostNoLimit(vA, vB, vPrize.add(vec{offset,offset}))
	}
	fmt.Println()
	return fmt.Sprintf("%d", total)
}
