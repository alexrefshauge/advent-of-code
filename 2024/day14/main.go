package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/alexrefshauge/advent-of-code/common/vector"
)

//go:embed input.txt
var input string

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Printf("Running AoC day %d solution (part %d)\n", 14, part)

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

type Robot struct {
	pos vector.Vec
	vel vector.Vec
}

var WIDTH = 101
var HEIGHT = 103
var STEPS = 100

func parse(input string) []*Robot {
	robots := make([]*Robot, 0)
	for _, line := range strings.Split(strings.TrimRight(input, "\n"), "\n") {
		var px,py,vx,vy int
		_, err := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)
		if err != nil { panic(err) }
		robots = append(robots, &Robot{vector.New(px,py), vector.New(vx,vy)})
	}
	return robots
}

func (r *Robot) Move() {
	r.pos = r.pos.Add(r.vel)
	if r.pos.X < 0 {r.pos.X += WIDTH}
	if r.pos.Y < 0 {r.pos.Y += HEIGHT}
	if r.pos.X >= WIDTH {r.pos.X -= WIDTH}
	if r.pos.Y >= HEIGHT {r.pos.Y -= HEIGHT}
}

func (r *Robot) Back() {
	r.pos = r.pos.Add(r.vel.Scale(-1))
	if r.pos.X < 0 {r.pos.X += WIDTH}
	if r.pos.Y < 0 {r.pos.Y += HEIGHT}
	if r.pos.X >= WIDTH {r.pos.X -= WIDTH}
	if r.pos.Y >= HEIGHT {r.pos.Y -= HEIGHT}
}

func printRobots(robots []*Robot) {
	hash := make(map[vector.Vec]int)
	for _, r := range robots {
		_, ok := hash[r.pos]
		if !ok {
			hash[r.pos] = 1
		} else {
			hash[r.pos]++
		}
	}
	for y := range HEIGHT {
		for x := range WIDTH {
			count, ok := hash[vector.New(x,y)]
			if ok && count > 0 {
				fmt.Print("██")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Print("\n")
	}
}
func safety(robots []*Robot) int {
	midx := WIDTH / 2
	midy := HEIGHT / 2
	nw, ne, sw, se := 0,0,0,0
	for _, r := range robots {
		if r.pos.X < midx {
			if r.pos.Y < midy {
				nw++
			}
			if r.pos.Y > midy {
				sw++
			}
		}
		if r.pos.X > midx {
			if r.pos.Y < midy {
				ne++
			}
			if r.pos.Y > midy {
				se++
			}
		}
	}
	return nw * ne * sw * se
}

func part1(input string) string {
	robots := parse(input)
	for range STEPS {
		for _, robot := range robots {
			robot.Move()
		}
	}

	return fmt.Sprintf("%d", safety(robots))
}

const (
	KEY_FORWARD = 'x'
	KEY_BACK = 'z'
)
func part2(input string) string {
	robots := parse(input)
	printRobots(robots)
	step := 0
	
    exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
    exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	for {
		var b []byte = make([]byte, 1)
		os.Stdin.Read(b)
		key := string(b)[0]
		if key == KEY_FORWARD {
			for _, robot := range robots {
				robot.Move();
			}
			step++
		}
		if key == KEY_BACK {
			for _, robot := range robots {
				robot.Back();
			}
			step--
		}
		fmt.Printf("\033[%dA", HEIGHT+1)
		printRobots(robots)
		fmt.Print("\033[K")
		fmt.Println(step)
	}
}
