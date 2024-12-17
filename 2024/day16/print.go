package main

import (
	"fmt"

	"github.com/alexrefshauge/advent-of-code/common/vector"
)

var (
	BOX = "\033[43;93m  \033[0m"
	WALL = "\033[37;100m[]\033[0m"
	EMPTY = "  "
	EMPTY_SLIM = "."
	ROBOT = "\033[41;91m  \033[0m"
	ROBOT_SLIM = "\033[41;91m0\033[0m"
)

func printWorld(walls map[vector.Vec]bool, agents []agent, end vector.Vec) {
	isAgent := make(map[vector.Vec]bool)
	for _, a := range agents {
		isAgent[a.pos] = true
	}
	for y := range HEIGHT {
		for x := range WIDTH {
			v := vector.New(x,y)
			wall := walls[v]
			switch wall {
			case true: fmt.Print(WALL); break
			case false:
				if v.X == end.X && v.Y == end.Y {
					fmt.Print(BOX)
				} else if isAgent[v] {
					fmt.Print(ROBOT)
				} else {
					fmt.Print(EMPTY)
				}
				break
			}
		}
		fmt.Print("\n")
	}
}