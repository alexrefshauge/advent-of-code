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
	fmt.Printf("Running AoC day %d solution (part %d)\n", 8, part)

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

var (
	WIDTH = 0
	HEIGHT = 0
)

func parse(input string) map[byte][]vec {
	grid := strings.Split(strings.TrimRight(input, "\n"), "\n")
	WIDTH = len(grid[0])
	HEIGHT = len(grid)
	antennas := make(map[byte][]vec)
	for y := range HEIGHT {
		for x := range WIDTH {
			if grid[y][x] != '.' {
				antennas[grid[y][x]] = append(antennas[grid[y][x]], vec{x,y})
			}
		}
	}
	return antennas
}

func hash(x,y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func part1(input string) string {
	antennas := parse(input)
	antinodes := findAntinodes(antennas, 1)
	return fmt.Sprintf("%d", len(antinodes))
}

func part2(input string) string {
	antennas := parse(input)
	antinodes := findAntinodes(antennas, -1)
	for _, vecs := range antennas {
		if len(vecs) > 1 {
			for _, antennaPosition := range vecs {
				antinodes[hash(antennaPosition.x, antennaPosition.y)] = true
			}
		}
	}

	return fmt.Sprintf("%d", len(antinodes))
}

func findAntinodes(antennas map[byte][]vec, limit int) map[string]bool {
	antinodes := make(map[string]bool)
	for _, vecs := range antennas {
		for _, i := range vecs {
			for _, j := range vecs {
				if i == j {continue}
				diff := j.sub(i)

				for n := 1; limit < 0 || n <= limit; n++ {
					antiAdd := j.add(diff.scale(n))
					antiSub := i.sub(diff.scale(n))
					if antiAdd.inside() {
						antinodes[hash(antiAdd.x, antiAdd.y)] = true
					}
					if antiSub.inside() {
						antinodes[hash(antiSub.x, antiSub.y)] = true
					}

					if !antiAdd.inside() && !antiSub.inside() {
						break
					}
				}
			}
		}
	}

	return antinodes
}