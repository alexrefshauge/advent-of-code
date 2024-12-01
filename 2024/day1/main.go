package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
)

//go:embed input.txt
var input string

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Printf("Running AoC day %d solution (part %d)\n", 1, part)

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

func part1(input string) string {
	left, right := formatInput(input)
	sort.Ints(left)
	sort.Ints(right)

	total := 0 
	for i := range left {
		diff := left[i] - right[i]
		total += int(math.Abs(float64(diff)))
	}

	return fmt.Sprintf("%d", total)
}

func part2(input string) string {
	left, right := formatInput(input)
	total := 0

	for _, l := range left {
		count := 0
		for _, r := range right {
			if l == r {
				count++
			}
		}
		total += count * l
	}

	return fmt.Sprintf("%d", total)
}
