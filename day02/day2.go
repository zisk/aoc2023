package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/zisk/aoc2023/util"
)

type Game struct {
	id      int
	results []GameResult
}

type GameResult struct {
	red, green, blue int
}

func NewGame(gamestring string) Game {
	var newGame Game
	initialSplit := strings.Split(gamestring, ":")

	header := strings.Split(initialSplit[0], " ")
	id, _ := strconv.Atoi(header[1])
	newGame.id = id

	drawStrings := strings.Split(initialSplit[1], ";")

	parsedDraws := make([]GameResult, len(drawStrings))

	i := 0
	for _, d := range drawStrings {
		parsedDraws[i] = ParseGameResult(d)
		i++
	}
	newGame.results = parsedDraws
	return newGame
}

func (g *Game) CheckGame() bool {
	for _, r := range g.results {
		if r.red > 12 || r.green > 13 || r.blue > 14 {
			return false
		}
	}
	return true
}

func ParseGameResult(resultString string) GameResult {
	var result GameResult

	resultSplit := strings.Split(resultString, ",")
	for _, draw := range resultSplit {
		s := strings.Split(strings.Trim(draw, " "), " ")
		count, _ := strconv.Atoi(s[0])
		switch s[1] {
		case "red":
			result.red = count
		case "green":
			result.green = count
		case "blue":
			result.blue = count
		}
	}

	return result
}

func (g *Game) GetMiniumCubePower() int {
	var rMax, gMax, bMax int

	for _, r := range g.results {
		if r.red > rMax {
			rMax = r.red
		}
		if r.green > gMax {
			gMax = r.green
		}
		if r.blue > bMax {
			bMax = r.blue
		}
	}
	return rMax * gMax * bMax
}

func main() {
	in, _ := util.InputToTxt()

	//part 1
	validCount := 0

	//part 2
	gameProduct := 0
	for _, g := range in {
		game := NewGame(g)
		if game.CheckGame() {
			validCount += game.id
		}
		p := game.GetMiniumCubePower()
		gameProduct += p
	}
	fmt.Printf("Part 1: %d\n", validCount)
	fmt.Printf("Part 2: %d\n", gameProduct)
}
