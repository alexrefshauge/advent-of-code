package main

import (
	"fmt"

	"github.com/alexrefshauge/advent-of-code/common/vector"
)

func printWorld(walls map[vector.Vec]bool, agents []agent) {
	isAgent := make(map[vector.Vec]bool)
	for _, a := range agents {
		isAgent[a.pos] = true
	}
	for y := range HEIGHT {
		for x := range WIDTH {
			wall := walls[vector.New(x,y)]
			switch wall {
			case true: fmt.Print("#"); break
			case false:
				if isAgent[vector.New(x,y)] {
					fmt.Print("@")
				} else {
					fmt.Print(".")
				}
				break
			}
		}
		fmt.Print("\n")
	}
}