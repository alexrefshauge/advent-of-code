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
	fmt.Printf("Running AoC day %d solution (part %d)\n", 6, part)

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
	visited := make(map[string]bool)

	lines := strings.Split(input, "\n")
	x, y := findStart(lines)
	hDir := 0
	vDir := -1
	for x > 0 && y > 0 && x < len(lines[0])-1 && y < len(lines)-1 {
		if !visited[hash(x, y)] {
			total++
			visited[hash(x, y)] = true
		}

		if lines[y+vDir][x+hDir] == '#' {
			hDir, vDir = turn(hDir, vDir)
		}

		x += hDir
		y += vDir

	}

	return fmt.Sprintf("%d", total+1)
}

func part2(input string) string {

	lines := strings.Split(input, "\n")
	x, y := findStart(lines)
	hDir := 0
	vDir := -1

	lastTurns := make([][]int, 0)
	newObstacles := make([][]int, 0)

	for x > 0 && y > 0 && x < len(lines[0])-1 && y < len(lines)-1 {
		if lines[y+vDir][x+hDir] == '#' {
			lastTurns = append(lastTurns, []int{x, y, hDir, vDir})
			hDir, vDir = turn(hDir, vDir)
		}

		x += hDir
		y += vDir
		for _, o := range checkLoop(lines, lastTurns, x, y, hDir, vDir) {
			if len(o) != 0 && !exists(newObstacles, o) {
				newObstacles = append(newObstacles, o)
			}
		}
	}

	fmt.Println(newObstacles)
	return fmt.Sprintf("%d", len(newObstacles))
}

func exists(all [][]int, n []int) bool {
	for _, o := range all {
		if o[0] == n[0] && o[1] == n[1] {
			return true
		}
	}
	return false
}

func checkLoop(lines []string, lastTurns [][]int, x, y, hDir, vDir int) [][]int {
	options := make([][]int, 0)
	startX := x
	startY := y
	hh, vv := turn(hDir, vDir)
	obstaclePos := []int{x + hDir, y + vDir}

	for _, loopTurn := range lastTurns {

		for x > 0 && y > 0 && x < len(lines[0])-1 && y < len(lines)-1 {
			x += hh
			y += vv

			if x == loopTurn[0] && y == loopTurn[1] && hh == loopTurn[2] && vv == loopTurn[3] {
				options = append(options, obstaclePos)
			}
		}

		x = startX
		y = startY
	}

	return options
}

func turn(h, v int) (int, int) {

	if h == 1 && v == 0 {
		return 0, 1
	}

	if h == 0 && v == 1 {
		return -1, 0
	}

	if h == -1 && v == 0 {
		return 0, -1
	}

	if h == 0 && v == -1 {
		return 1, 0
	}

	return 0, 0
}

func hash(x, y int) string { return fmt.Sprintf("%d,%d", x, y) }

func findStart(lines []string) (int, int) {
	for y := range lines {
		for x := range lines[y] {
			if lines[y][x] == '^' {
				return x, y
			}
		}
	}

	return 0, 0
}
