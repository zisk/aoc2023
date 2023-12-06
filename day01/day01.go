package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

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

func GetDigitsText(input string) (rune, rune) {
	var DigitMap = map[string]rune{
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}

	resultMap := make(map[int]rune)

	for i, s := range input {
		_, err := strconv.Atoi(string(s))
		if err != nil {
			continue
		}
		resultMap[i] = s
	}

	for k := range DigitMap {
		firstIndex := strings.Index(input, k)
		lastIndex := strings.LastIndex(input, k)

		if firstIndex != -1 {
			resultMap[firstIndex] = rune(DigitMap[k])
		}

		if lastIndex != -1 {
			resultMap[lastIndex] = rune(DigitMap[k])
		}
	}

	intIndexes := make([]int, len(resultMap))
	i := 0
	for k := range resultMap {
		intIndexes[i] = k
		i++
	}
	sort.Ints(intIndexes)
	return resultMap[intIndexes[0]], resultMap[intIndexes[len(intIndexes)-1]]
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

	var answerOne, answerTwo int
	for _, i := range partOneSlice {
		answerOne = answerOne + i
	}
	partTwoSlice := []int{}

	for _, l := range input {
		first, last := GetDigitsText(l)
		lineresult := string([]rune{first, last})
		lineInt, _ := strconv.Atoi(lineresult)
		partTwoSlice = append(partTwoSlice, lineInt)
	}

	for _, i := range partTwoSlice {
		answerTwo = answerTwo + i
	}
	fmt.Printf("Part 1: %d\n", answerOne)
	fmt.Printf("Part 2: %d\n", answerTwo)
}
