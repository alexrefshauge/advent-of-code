package mapcache

import (
	. "day10/point"
	"strconv"
	"strings"
)

const height = 10

type TopoMap []map[Point]bool

func FromString(topo string) TopoMap {
	maps := make(TopoMap, height, height)
	for h := range height {
		maps[h] = make(map[Point]bool)
	}

	for y, row := range strings.Split(strings.TrimRight(topo, "\n"), "\n") {
		for x, r := range row {
			level, err := strconv.Atoi(string(r))
			if err != nil {
				panic(err)
			}

			maps[level][NewP(x,y)] = true
		}
	}

	return maps
}

