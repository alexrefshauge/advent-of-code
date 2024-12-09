package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Printf("Running AoC day %d solution (part %d)\n", 9, part)

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

func parse(input string) []int {
	chars := strings.Split(strings.TrimRight(input, "\n"), "")
	nums := make([]int, len(chars), len(chars))
	var err error

	for i := range chars {
		nums[i], err = strconv.Atoi(chars[i])
		if err != nil {
			panic(err)
		}
	}
	return nums
}

func expand(blocks []int) []int {
	expanded := make([]int, 0)
	id := 0
	for i, block := range blocks {
		if i%2 == 0 {
			tail := make([]int, block, block)
			for j := range tail {
				tail[j] = id
			}
			expanded = append(expanded, tail...)
			id++
		} else {
			tail := make([]int, block, block)
			for j := range tail {
				tail[j] = -1
			}
			expanded = append(expanded, tail...)
		}
	}
	return expanded
}

func collapse(parts []int) []int {
	l := 0
	r := len(parts)-1
	
	for l+1 < r-1 {
		for parts[l] != -1 {
			l++
		}
		for parts[r] == -1 {
			r--
		}
		parts[l], parts[r] = parts[r], parts[l]
	}

	return parts
}

func checksum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	total := 0
	for n := 0; nums[n] != -1; n++ {
		total += n * nums[n]
	}
	return total
}

func part1(input string) string {
	disk := parse(input)
	expanded := expand(disk)
	collapsed := collapse(expanded)
	return fmt.Sprintf("%d", checksum(collapsed))
}

func part2(input string) string {
	disk := parse(input)
	blocks := toBlocks(disk)
	collapsed := collapseBlocks(blocks)
	return fmt.Sprintf("%d", checksumBlocks(collapsed))
}
