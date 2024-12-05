package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Printf("Running AoC day %d solution (part %d)\n", 5, part)

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
	total := 0

	ruleMap := make(map[int][]int)
	parts := strings.Split(input, "\n\n")
	ruleLines := strings.Split(parts[0], "\n")
	updateLines := strings.Split(parts[1], "\n")

	for _, line := range ruleLines {
		nums := strings.Split(line, "|")
		l, _ := strconv.Atoi(nums[0])
		r, _ := strconv.Atoi(nums[1])
		ruleMap[l] = append(ruleMap[l], r)
	}

	for _, line := range updateLines {
		pages := make([]int, 0)
		for _, n := range strings.Split(line, ",") {
			num, _ := strconv.Atoi(n)
			pages = append(pages, num)

		}
		if isValid(ruleMap, pages) {
			total += mid(pages)
		}

	}

	return fmt.Sprintf("%d", total)
}

func part2(input string) string {
	total := 0

	ruleMap := make(map[int][]int)
	parts := strings.Split(input, "\n\n")
	ruleLines := strings.Split(parts[0], "\n")
	updateLines := strings.Split(parts[1], "\n")

	for _, line := range ruleLines {
		nums := strings.Split(line, "|")
		l, _ := strconv.Atoi(nums[0])
		r, _ := strconv.Atoi(nums[1])
		ruleMap[l] = append(ruleMap[l], r)
	}

	for _, line := range updateLines {
		pages := make([]int, 0)
		for _, n := range strings.Split(line, ",") {
			num, _ := strconv.Atoi(n)
			pages = append(pages, num)

		}
		if !isValid(ruleMap, pages) {
			sorted := false
			for sorted == false {
				sorted = sort(ruleMap, pages)
			}
			total += mid(pages)
		}
	}

	return fmt.Sprintf("%d", total)
}

func mid(nums []int) int {
	i := (len(nums) - 1) / 2
	return nums[i]
}

func isValid(ruleMap map[int][]int, pages []int) bool {
	for i, p := range pages {
		rules, ok := ruleMap[p]
		if !ok {
			continue
		}
		j := i
		for j >= 0 {
			if slices.Contains[[]int](rules, pages[j]) {
				return false
			}
			j--
		}
	}
	return true
}

func sort(ruleMap map[int][]int, pages []int) bool {
	for i, p := range pages {
		rules, ok := ruleMap[p]
		if !ok {
			continue
		}
		j := i - 1
		for j >= 0 {
			if slices.Contains[[]int](rules, pages[j]) {
				temp := pages[i]
				pages[i] = pages[j]
				pages[j] = temp
				return false
			}
			j--
		}
	}
	return true
}
