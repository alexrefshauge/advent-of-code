package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/schollz/progressbar/v3"
)

//go:embed input.txt
var input string

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Printf("Running AoC day %d solution (part %d)\n", 7, part)

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

const (
	add = iota
	mul
	con
)

var (
	operations = map[int]func(a,b uint64)uint64 {
		add: func(a,b uint64) uint64 {return a+b},
		mul: func(a,b uint64) uint64 {return a*b},
		con: func(a,b uint64) uint64 {
			res, err := strconv.ParseUint(fmt.Sprintf("%d%d", a, b), 10, 64)
			if err != nil {
				panic(err)
			}
			return res
		},
	}
)

func part1(input string) string {
	return fmt.Sprintf("%d", sumValidCalculations(input, 2))
}

func part2(input string) string {
	return fmt.Sprintf("%d", sumValidCalculations(input, 3))
}

func sumValidCalculations(input string, base int) uint64 {
	var result uint64 = 0
	calibrations := strings.Split(strings.TrimRight(input, "\n"), "\n")
	bar := progressbar.Default(int64(len(calibrations)))
	for _, l := range calibrations {
		target, nums := parseCal(l)
		opCount := len(nums)-1
		bar.Add(1)
		runOnAll(opCount, func(ops []int) bool {
			var accumulator uint64 = uint64(nums[0])
			for iOp, op := range ops {
				accumulator = operations[op](accumulator, uint64(nums[iOp+1]))
			}
			if target == accumulator {
				result += target
				return true
			}
			return false
		}, base)
	}
	return result
}

func runOnAll(opCount int, opFunc func(ops []int) bool, base int) {
	max := base-1
	cancel := opFunc(make([]int, opCount, opCount))
	ops := make([]int, opCount, opCount)
	for !isAll(ops, max) {
		if cancel { break }
		
		i := 0
		for ops[i] == max {
			ops[i] = 0
			i++
		}

		ops[i]++

		cancel = opFunc(ops)
	}
}

func isAll(arr []int, compare int) bool {
	for _, x := range arr {
		if x != compare {
			return false
		}
	}
	return true
}

func parseCal(line string) (uint64, []int) {
	parts := strings.Split(line, ": ")
	target, err := strconv.ParseUint(parts[0], 10, 64)
	if err != nil {
		panic(fmt.Errorf("failed to parse: %s", err))
	}
	numsRaw := strings.Split(parts[1], " ")

	nums := make([]int, 0)
	for _, n := range numsRaw {
		parsed, err := strconv.Atoi(n)
		if err != nil {
			panic(fmt.Errorf("failed to parse: %s", err))
		}
		nums = append(nums, parsed)
	}

	return target, nums
}