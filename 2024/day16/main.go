package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/alexrefshauge/advent-of-code/common/queue"
	"github.com/alexrefshauge/advent-of-code/common/vector"
)

//go:embed input.txt
var input string

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Printf("Running AoC day %d solution (part %d)\n", 16, part)

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

func parse(input string) (map[vector.Vec]bool, vector.Vec, vector.Vec) {
	var start, end vector.Vec
	walls := make(map[vector.Vec]bool)
	grid := strings.Split(strings.TrimRight(input, "\n"), "\n")

	WIDTH = len(grid[0])
	HEIGHT = len(grid)

	for y, row := range grid {
		for x, cell := range row {
			switch cell {
			case '#': walls[vector.New(x,y)] = true; break
			case 'S': start = vector.New(x,y); break
			case 'E': end = vector.New(x,y); break
			}
		}
	}

	return walls, start, end
}

var (
	WIDTH = 0
	HEIGHT = 0

	EAST = 	vector.New( 1, 0)
	WEST = 	vector.New(-1, 0)
	NORTH = vector.New( 0,-1)
	SOUTH = vector.New( 0, 1)
)

type agent struct {
	pos vector.Vec
	dir vector.Vec
	score int
	cells []vector.Vec
}

func shuffle(slice []vector.Vec) {
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

type path struct {
	score int
	cells []vector.Vec
}

var paths []path
var animate = false

func explore(agents *queue.Queue[agent], walls map[vector.Vec]bool, exploredCache map[vector.Vec]int, lowest *int, end vector.Vec) {
	a := agents.Next()

	dirs := []vector.Vec{NORTH, EAST, SOUTH, WEST}

	for _, dir := range dirs {
		added := a.dir.Add(dir)
		if walls[a.pos.Add(dir)] || added.Equals(vector.New(0,0)) { continue }
		score := a.score+1

		if !a.dir.Equals(dir) {
			score += 1000
		}

		newAgent := agent{pos: a.pos.Add(dir), dir: dir, score: score}
		if *lowest != -1 && score > *lowest { continue }

		cells := make([]vector.Vec, len(a.cells))
		copy(cells, a.cells)
		newAgent.cells = append(cells, newAgent.pos)

		if !newAgent.pos.Equals(end) {
			exploredScore, explored := exploredCache[newAgent.pos]
			if explored && score > exploredScore+1000 {
				continue
			}
			exploredCache[newAgent.pos] = score
			agents.Push(newAgent)
			continue
		}
		
		if *lowest == -1 || score <= *lowest {
			paths = append(paths, path{score: score ,cells: newAgent.cells})
			*lowest = score
		}
	}
}

func run(input string) int {
	walls, start, end := parse(input)
	agents := queue.New[agent]()
	agents.Push(agent{pos: start, dir: EAST, score: 0, cells: []vector.Vec{start}})
	var lowest *int = new(int); *lowest = -1
	explored := make(map[vector.Vec]int)
	
	if animate {printWorld(walls, agents.All(), end)}

	i := 0
	for agents.Size() > 0 {
		explore(agents, walls, explored, lowest, end)
		if !animate {continue} 
		if agents.Size() != 0 && i % agents.Size() == 0 {
			i = 1
			time.Sleep(100*time.Millisecond)
			fmt.Printf("\033[%dA", HEIGHT)
			printWorld(walls, agents.All(), end)
		} else {
			i++
		}
	}
	
	return *lowest
}

func part1(input string) string {
	lowest := run(input)
	return fmt.Sprintf("%d", lowest)
}

func part2(input string) string {
	paths = make([]path, 0)
	seats := make(map[vector.Vec]bool)
	lowest := run(input)
	for _, p := range paths {
		if p.score != lowest {continue}
		for _, c := range p.cells {
			seats[c] = true
		}
	}
	return fmt.Sprintf("%d", len(seats))
}
