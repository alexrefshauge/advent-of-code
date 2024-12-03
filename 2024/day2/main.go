package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Printf("Running AoC day %d solution (part %d)\n", 2, part)

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
	reports := strings.Split(strings.TrimRight(input, "\n"), "\n")
	safeReports := 0
	for _, r := range reports {
		tokens := strings.Fields(r)
		levels := make([]int, len(tokens))
		for i := range tokens {
			levels[i], _ = strconv.Atoi(tokens[i])
		}
		
		if isSafe(levels) {
			safeReports++
		}
	}

	return fmt.Sprintf("%d", safeReports)
}

func part2(input string) string {
	reports := strings.Split(strings.TrimRight(input, "\n"), "\n")

	result := make(chan int)
	resultQueue := make(chan int)

	go func(){
		safeReports := 0
		reportsChecked := 0
		for reportsChecked < len(reports) {

			reportResult := <- resultQueue
			safeReports += reportResult
			reportsChecked++
		
		}

		result <- safeReports
	}()

	for _, r := range reports {
		go func() {
			tokens := strings.Fields(r)
			levels := make([]int, len(tokens))
			for i := range tokens {
				levels[i], _ = strconv.Atoi(tokens[i])
			}
	
			if isAnySafe(levels) {
				resultQueue <- 1
			} else {
				resultQueue <- 0
			}
		}()
	}

	return fmt.Sprintf("%d", <-result)
}


func isSafe(levels []int) bool {
	last := levels[0]
	velocity := 0
	for _, level := range levels[1:] {
		diff := level - last
		if diff == 0 {
			return false
		}
		if velocity == 0 {
			if diff < 0 {
				velocity = -1
			}
			if diff > 0 {
				velocity = 1
			}
		}

		if velocity == -1 && diff > velocity {
			return false
		}
		if velocity == 1 && diff < velocity {
			return false
		}

		if int(math.Abs(float64(diff))) > 3 {
			return false
		}

		last = level
	}
	return true
}

func isAnySafe(levels []int) bool {
	reports := make([][]int, len(levels))
	for i := range len(levels) {
		reports[i] = make([]int, 0, len(levels)-1)
		for j := range len(levels) {
			if j == i {
				continue
			}
			reports[i] = append(reports[i], levels[j])
		}
	}

	for _, r := range reports {
		if isSafe(r) {
			return true
		}
	}
	return false
}