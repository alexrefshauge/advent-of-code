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
	fmt.Printf("Running AoC day %d solution (part %d)\n", 12, part)

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

type vec struct {
	x, y int
}

func (v vec) shift(x,y int ) vec {
	return vec{v.x + x, v.y + y}
}

type plot struct {
	perimeter int
	plant byte
	points map[vec]bool
}

func NewPlot(v vec, plant byte) plot {
	return plot{
		perimeter: 0,
		plant: plant,
		points: make(map[vec]bool),
	}
}

var width, height int

func checkPlant(grid []string, v vec) (byte, bool) {
	if v.x < 0 || v.x >= width || v.y < 0 || v.y >= height {
		return 0, false
	}
	return grid[v.y][v.x], true
}


func (p *plot) expand(grid []string, v vec, explored map[vec]bool, exploreStack []vec) (map[vec]bool, []vec, int) {
	left, 	ok := checkPlant(grid, v)
	if ok && !explored[v] && p.plant == left {
		exploreStack = append(exploreStack, v)
		explored[v] = true
	} else if !explored[v] {
		return explored, exploreStack, 1
	}
	
	return explored, exploreStack, 0
}
	
var visited = make(map[vec]bool)
func fillPlot(grid []string, pos vec, discount bool) plot {
	exploreStack := []vec{pos}
	explored := make(map[vec]bool)
	explored[pos] = true

	plant := grid[pos.y][pos.x]
	newPlot := NewPlot(pos, plant)

	for len(exploreStack) > 0 {
		v := exploreStack[len(exploreStack)-1]
		exploreStack = exploreStack[:len(exploreStack)-1]
		newPlot.points[v] = true
		visited[v] = true

		var pLeft, pRight, pUp, pDown int
		explored, exploreStack, pLeft = newPlot.expand(grid, v.shift(-1, 0), explored, exploreStack)
		explored, exploreStack, pRight = newPlot.expand(grid, v.shift(1, 0), explored, exploreStack)
		explored, exploreStack, pUp = newPlot.expand(grid, v.shift(0, -1), explored, exploreStack)
		explored, exploreStack, pDown = newPlot.expand(grid, v.shift(0, 1), explored, exploreStack)

		
		if !discount {
			newPlot.perimeter += pLeft + pRight + pUp + pDown
		}
	}
	if discount {
		newPlot.perimeter = discountPerimeter(newPlot)
	}
	return newPlot
}

func flipped(last []bool, this []bool) bool {
	return last[0] != this[0] && last[1] != this[1]
}

func discountPerimeter(p plot) int {
	horizontalSides := 0
	for y := range height+1 {
		lastEdge := false
		lastComp := []bool{false,false}
		for x := range width {
			_, compare := p.points[vec{x, y-1}]
			_, cursor := p.points[vec{x,y}]
			comp := []bool{compare,cursor}
			if !(cursor || compare) { lastEdge = false; continue }//not part of plot

			edge := y == 0 && cursor
			if y != 0 {
				edge = cursor != compare
			}
			
			if edge && (!lastEdge || flipped(lastComp, comp)) {horizontalSides++}
			lastEdge = edge
			lastComp = comp
		}
		lastEdge = false
	}

	verticalSides := 0
	for x := range width+1 {
		lastEdge := false
		lastComp := []bool{false,false}
		for y := range height {
			_, compare := p.points[vec{x-1, y}]
			_, cursor := p.points[vec{x,y}]
			comp := []bool{compare,cursor}
			if !(cursor || compare) { lastEdge = false; continue }//not part of plot

			edge := cursor != compare
			
			if edge && (!lastEdge || flipped(lastComp, comp)) {verticalSides++}
			lastEdge = edge
			lastComp = comp
		}
	}

	return horizontalSides + verticalSides
}

func part1(input string) string {
	grid := strings.Split(strings.TrimRight(input, "\n"), "\n")
	plots := make([]plot, 0)
	width, height = len(grid[0]), len(grid)
	total := 0
	for y, row := range grid {
		for x := range row {
			if visited[vec{x,y}] { continue }
			newPlot := fillPlot(grid, vec{x,y}, false)
			plots = append(plots, newPlot)
			total += newPlot.perimeter * len(newPlot.points)
		}
	}
	return fmt.Sprintf("%d", total)
}

func part2(input string) string {
	visited = make(map[vec]bool)
	grid := strings.Split(strings.TrimRight(input, "\n"), "\n")
	plots := make([]plot, 0)
	width, height = len(grid[0]), len(grid)
	total := 0
	for y, row := range grid {
		for x := range row {
			if visited[vec{x,y}] { continue }
			newPlot := fillPlot(grid, vec{x,y}, true)
			plots = append(plots, newPlot)
			total += newPlot.perimeter * len(newPlot.points)
		}
	}
	return fmt.Sprintf("%d", total)
}
