package main

import (
	"strconv"
	"strings"
)

func formatInput(raw string) ([]int, []int) {
	left := make([]int, 0, 1000)
	right := make([]int, 0, 1000)

	rows := strings.Split(raw, "\n")
	for _, row := range rows {
		var err error
		var lNum, rNum int
		tokens := strings.Split(row, "   ")
		if tokens[0] == "" || tokens[1] == ""{
			continue
		}
		lNum, err = strconv.Atoi(tokens[0])
		if err != nil {
			panic(err)
		}
		rNum, err = strconv.Atoi(tokens[1])
		if err != nil {
			panic(err)
		}
		left = append(left, lNum)
		right = append(right, rNum)
	}

	return left, right
}