package main

import (
	"day5/update"
	_ "embed"
	"flag"
	"fmt"
	"os"
	"sort"
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

func parse(input string) (map[int][]int, [][]int) {

	ruleMap := make(map[int][]int)
	updates := make([][]int, 0)
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
		updates = append(updates, pages)
	}

	return ruleMap, updates
}

func part1(input string) string {
	rules, updates := parse(input)
	total := 0

	for _, pages := range updates {
		u := update.New(rules, pages)
		if sort.IsSorted(u) {
			total += u.Mid()
		}
	}

	return fmt.Sprintf("%d", total)
}

func part2(input string) string {
	rules, updates := parse(input)
	total := 0

	for _, pages := range updates {
		u := update.New(rules, pages)
		if !sort.IsSorted(u) {
			sort.Sort(u)
			total += u.Mid()
		}
	}

	return fmt.Sprintf("%d", total)
}
