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
	fmt.Printf("Running AoC day %d solution (part %d)\n", 11, part)

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

func parse(in string) []string {
	return strings.Split(strings.TrimRight(in, "\n"), " ")
}

func handleStone(stone string) []string {
	stones := make([]string, 1, 1) 
	length := len(stone)
	if stone == "0" {
		stones[0] = "1"
	} else if length % 2 == 0 {
		mid := length/2
		l := stone[:mid]
		r := stone[mid:]
		rNum, _ := strconv.Atoi(r)

		stones[0] = l
		stones = append(stones, fmt.Sprintf("%d", rNum))
	} else {
		num, _ := strconv.Atoi(stone)
		stones[0] = fmt.Sprintf("%d", num * 2024)
	}

	return stones
}

var cache []map[string]int
func mutateStone(stone string, blink, maxBlinks int) int {
	if blink >= maxBlinks {
		return 1
	}

	cached, ok := cache[blink][stone]
	if ok {
		return cached
	}

	stones := handleStone(stone)
	count := 0

	for _, s := range stones {
		count += mutateStone(s, blink+1, maxBlinks)
	}

	cache[blink][stone] = count
	return count
}

func getStoneCountAtBlink(stones []string, blinks int) int {
	cache = make([]map[string]int, blinks, blinks)
	for i := range blinks {
		cache[i] = make(map[string]int)
	}

	total := 0

	for _, stone := range stones {
		total += mutateStone(stone, 0, blinks)
	}


	return total
}

func part1(input string) string {
	count := getStoneCountAtBlink(parse(input), 25)
	return fmt.Sprintf("%d", count)
}

func part2(input string) string {
	count := getStoneCountAtBlink(parse(input), 75)
	return fmt.Sprintf("%d", count)
}
