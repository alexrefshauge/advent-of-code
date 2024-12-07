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
	for isInside(x, y, lines) {
		if !visited[hash(x, y)] {
			total++
			visited[hash(x, y)] = true
		}

		if isInside(x+hDir, y+vDir, lines) && lines[y+vDir][x+hDir] == '#' {
			hDir, vDir = turn(hDir, vDir)
		}

		x += hDir
		y += vDir
	}

	return fmt.Sprintf("%d", total)
}

func part2(input string) string {
	lines := strings.Split(strings.TrimRight(input, "\n"), "\n")
	sx, sy := findStart(lines)
	hDir := 0
	vDir := -1
	loops := 0

	for y := range lines {
		for x := range lines[y] {
			if willLoopWithObstacle(lines, sx,sy,hDir,vDir, x,y) {
				loops++
			}
		}
	}

	return fmt.Sprintf("%d", loops)
}

func willLoopWithObstacle(linesOriginal []string, x, y int, hDir, vDir int, oX, oY int) bool {
	lines := make([]string, len(linesOriginal))
	copy(lines, linesOriginal)
	if oX == x && oY == y {
		return false
	}

	if lines[oY][oX] == '#' {
		return false
	}

	l := []rune(lines[oY])
	l[oX] = '#'
	lines[oY] = string(l)

	beenHereBefore := make(map[string]bool)
	for isInside(x, y, lines) {
		if beenHereBefore[hashTurn(x,y,hDir, vDir)] {
			return true
		}
		if isInside(x+hDir, y+vDir, lines) && lines[y+vDir][x+hDir] == '#' {
			beenHereBefore[hashTurn(x,y,hDir, vDir)] = true
			hDir, vDir = turn(hDir, vDir)
			continue
		}

		x += hDir
		y += vDir
	}

	return false
}

func part2_old(input string) string {
	lines := strings.Split(strings.TrimRight(input, "\n"), "\n")
	x, y := findStart(lines)
	hDir := 0
	vDir := -1

	newObstacles := make([][]int, 0)

	for isInside(x, y, lines) {
		o := []int{x + hDir, y + vDir}
		if checkLoop(lines, x, y, hDir, vDir) && lines[y+vDir][x+hDir] != '#' {
			if len(o) != 0 && !exists(newObstacles, o) {
				newObstacles = append(newObstacles, o)
			}
		}

		if isInside(x+hDir, y+vDir, lines) && lines[y+vDir][x+hDir] == '#' {
			hDir, vDir = turn(hDir, vDir)
			continue
		}

		x += hDir
		y += vDir
	}
	return fmt.Sprintf("%d", len(newObstacles))
}

func checkLoop(lines []string, x, y int, hDir, vDir int) bool {
	hash := make(map[string]bool)
	hash[hashTurn(x, y, hDir, vDir)] = true

	if !isInside(x+hDir, y+vDir, lines) {
		return false
	}

	hDir, vDir = turn(hDir, vDir)

	for isInside(x, y, lines) {
		if !isInside(x+hDir, y+vDir, lines) {
			return false
		}

		if isInside(x+hDir, y+vDir, lines) && lines[y+vDir][x+hDir] == '#' {
			hash[hashTurn(x, y, hDir, vDir)] = true
			hDir, vDir = turn(hDir, vDir)
			continue
		}

		x += hDir
		y += vDir

		if hash[hashTurn(x, y, hDir, vDir)] {
			return true
		}

	}
	return false
}

func hashTurn(x, y, hDir, vDir int) string {
	return fmt.Sprintf("%d,%d,%d,%d", x, y, hDir, vDir)
}

func exists(all [][]int, n []int) bool {
	for _, o := range all {
		if o[0] == n[0] && o[1] == n[1] {
			fmt.Println("dupe: ", o, n)
			return true
		}
	}
	return false
}

func isInside(x, y int, lines []string) bool {
	return x >= 0 && y >= 0 && x < len(lines[0]) && y < len(lines)
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
