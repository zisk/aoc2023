package main

import (
	"fmt"
	"strconv"

	"github.com/zisk/aoc2023/util"
)

func GetDigits(input string) (rune, rune) {
	result := []rune{}

	for _, s := range input {
		_, err := strconv.Atoi(string(s))
		if err != nil {
			continue
		}
		result = append(result, s)
	}
	return result[0], result[len(result)-1]
}

func main() {
	input, _ := util.InputToTxt()

	partOneSlice := []int{}

	for _, l := range input {
		first, last := GetDigits(l)
		lineresult := string([]rune{first, last})
		lineInt, _ := strconv.Atoi(lineresult)
		partOneSlice = append(partOneSlice, lineInt)
	}

	var answer int
	for _, i := range partOneSlice {
		answer = answer + i
	}
	fmt.Printf("Part 1: %d", answer)
}
