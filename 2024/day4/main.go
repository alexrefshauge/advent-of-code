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

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Printf("Running AoC day %d solution (part %d)\n", 4, part)

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
	block := strings.Split(strings.TrimRight(input, "\n"), "\n")

	width := len(block[0])
	height := len(block)

	for x := range width {
		for y := range height {
			total += countXmas(block, x, y)
		}
	}

	return fmt.Sprintf("%d", total)
}

func part2(input string) string {
	total := 0
	masks := []string{
		"MMSS",
		"SMSM",
		"SSMM",
		"MSMS",
	}
	block := strings.Split(strings.TrimRight(input, "\n"), "\n")
	width := len(block[0])
	height := len(block)

	for x := range width {
		for y := range height {
			if x == 0 || y == 0 || x == width-1 || y == height-1 {
				continue
			}

			if rune(block[y][x]) != 'A' {
				continue
			}

			nw := string(rune(block[y-1][x-1]))
			ne := string(rune(block[y-1][x+1]))
			sw := string(rune(block[y+1][x-1]))
			se := string(rune(block[y+1][x+1]))

			for _, m := range masks {
				if m == nw+ne+sw+se {
					total++
					break
				}
			}
		}
	}

	return fmt.Sprintf("%d", total)
}

func countXmas(block []string, x, y int) int {
	count := 0
	for h := range 3 {
		for v := range 3 {
			w := mask(block, h-1, v-1, 4, x, y)
			if w == "XMAS" {
				count++
			}
		}
	}
	return count
}
