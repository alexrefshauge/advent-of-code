package main

import (
	_ "embed"
	"flag"
	"fmt"
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
}

func explore(agents *queue.Queue[agent], walls map[vector.Vec]bool, lowest *int, end vector.Vec) {
	a := agents.Next()

	for _, dir := range []vector.Vec{WEST,EAST,NORTH,SOUTH} {
		added := a.dir.Add(dir)
		if walls[a.pos.Add(dir)] || (added.X == 0 && added.Y == 0) { continue }
		score := a.score+1
		if a.dir == dir {
			score += 1000
		}
		nextAgent := agent{pos: a.pos.Add(dir), dir: dir, score: score}
		if *lowest != -1 && score > *lowest { continue }
		
		if nextAgent.pos != end {
			agents.Push(nextAgent)
		}
		if *lowest == -1 {
			*lowest = score
		}
		if score < *lowest {
			*lowest = score
		}
	}
}

func part1(input string) string {
	walls, start, end := parse(input)
	agents := queue.New[agent]()
	agents.Push(agent{pos: start, score: 0})
	var lowest *int = new(int); *lowest = -1
	
	printWorld(walls, agents.All())

	for agents.Size() > 0 {
		explore(agents, walls, lowest, end)

		fmt.Printf("\033[%dA", HEIGHT)
		printWorld(walls, agents.All())
		time.Sleep(time.Second)
	}
	
	return fmt.Sprintf("%d", lowest)
}

func part2(input string) string {
	return input
}
