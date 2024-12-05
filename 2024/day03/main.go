package main

import (
	_ "embed"
	"flag"
	"fmt"
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
	fmt.Printf("Running AoC day %d solution (part %d)\n", 3, part)

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
	total := 0

	for i := range len(input) {
		l, r, skip := parseMul(input[i:])
		total += l*r
		i+= skip
	}

	return fmt.Sprintf("%d", total)
}

func part2(input string) string {
	doMul := true

	total := 0

	for i := range len(input) {
		enable, ok := parseDo(input[i:])
		if ok {
			doMul = enable
		}


		if doMul {
			l, r, skip := parseMul(input[i:])
			total += l*r
			i+= skip
		}
	}

	return fmt.Sprintf("%d", total)
}

func parseDo(buf string) (bool, bool) {
	if strings.HasPrefix(buf, "do()") {
		return true, true
	}
	if strings.HasPrefix(buf, "don't()") {
		return false, true
	}
	return false, false
}

//returns left, right and character count
func parseMul(buf string) (int, int, int) {
	if !strings.HasPrefix(buf, "mul(") {
		return 0, 0, 0
	}

	buf = strings.TrimPrefix(buf, "mul(")

	iEnd := -1
	for pos, char := range buf {
		if char == ')' {
			iEnd = pos
			break
		}
	}

	inside := buf[:iEnd]
	parts := strings.Split(inside, ",")

	if len(parts) != 2 {
		return 0, 0, 0
	}

	left := parts[0]
	right := parts[1]

	lNum, err := strconv.Atoi(left)
	if err != nil {
		return 0, 0, 0
	}
	rNum, err := strconv.Atoi(right)
	if err != nil {
		return 0, 0, 0
	}

	return lNum, rNum, iEnd-1
}