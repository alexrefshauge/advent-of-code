package main

import (
	_ "embed"
	"testing"
)

//go:embed example1.txt
var example1 string

//go:embed example2.txt
var example2 string

func Test_day1_part1(t *testing.T) {
	innerTest(t, example1, 1, part1)
}

func Test_day1_part2(t *testing.T) {
	innerTest(t, example2, 2, part2)
}

func innerTest(t *testing.T, input string, part int, solutionFunc func(string) string) {
	if input == "" {
		t.Fatalf("please provide example input (part %d)", part)
	}
	in, out := parseTestInput(input)
	answer := solutionFunc(in)
	if answer != out {
		t.Fatalf("wanted %s, got %s", out, answer)
	}
}

func parseTestInput(input string) (string, string) {
	var token rune
	i := len(input)-1
	for token != ':' {
		token = rune(input[i])
		i--
	}

	return input[:i], input[i+2:]
}
func Benchmark_day1_part1(b *testing.B) {
	part1(input)
}

func Benchmark_day1_part2(b *testing.B) {
	part2(input)
}