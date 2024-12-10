package main

import (
	mapcache "day10/mapCache"
	. "day10/point"
	_ "embed"
	"flag"
	"fmt"
	"os"
)

//go:embed input.txt
var input string

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Printf("Running AoC day %d solution (part %d)\n", 10, part)

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

func tryTravel(fullRef *mapcache.TopoMap, pathRef *mapcache.TopoMap, level int, at Point) {
	full := *fullRef
	path := *pathRef
	if (full[level][at]) {
		path[level][at] = true
	}
}

func findPaths(topoMapRef *mapcache.TopoMap, start Point) (mapcache.TopoMap, int) {
	count := 1
	topoMap := *topoMapRef
	accessible := make(mapcache.TopoMap, len(topoMap))
	for h := range topoMap {
		accessible[h] = make(map[Point]bool)
	}

	for level := range len(topoMap) {
		if level == 0 {
			accessible[level][start] = true
			continue
		}
		for lastPoint := range accessible[level-1] {
			tryTravel(&topoMap, &accessible, level, lastPoint.Shift( 0, -1))
			tryTravel(&topoMap, &accessible, level, lastPoint.Shift( 0,  1))
			tryTravel(&topoMap, &accessible, level, lastPoint.Shift(-1,  0))
			tryTravel(&topoMap, &accessible, level, lastPoint.Shift( 1,  0))
		}
	}

	return accessible, count
}

func part1(input string) string {
	topoMap := mapcache.FromString(input)
	total := 0

	for p := range topoMap[0] {
		trails, _ := findPaths(&topoMap, p)
		total += len(trails[9])
	}
	
	return fmt.Sprintf("%d", total)
}

func travel(topoMap *mapcache.TopoMap, level int, p Point) int {
	lMap := (*topoMap)[level]
	if !lMap[p] {
		return 0
	}
	if level == 9 && lMap[p] {
		return 1
	}
	count := 0

	count += travel(topoMap, level+1, p.Shift( 0, -1))
	count += travel(topoMap, level+1, p.Shift( 0,  1))
	count += travel(topoMap, level+1, p.Shift(-1,  0))
	count += travel(topoMap, level+1, p.Shift( 1,  0))
	return count
}

func part2(input string) string {
	topoMap := mapcache.FromString(input)
	total := 0

	for start := range topoMap[0] {
		trailCount := travel(&topoMap, 0, start)
		total += trailCount
	}
	
	return fmt.Sprintf("%d", total)
}