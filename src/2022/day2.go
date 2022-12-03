package aoc2022

import (
	"fmt"
	"strings"
)

type day2 struct{}

type rockPaperScissorsGame struct {
	scores map[rune]int

	outcome map[string]int
}

var rockPaperScissors *rockPaperScissorsGame

func init() {
	Days[2] = &day2{}

	rockPaperScissors = &rockPaperScissorsGame{
		scores: map[rune]int{
			'A': 1, // Rock
			'B': 2, // Paper
			'C': 3, // Scissors
			'X': 1, // Rock
			'Y': 2, // Paper
			'Z': 3, // Scissors
		},
		outcome: map[string]int{
			"lost": 0,
			"draw": 3,
			"win":  6,
		},
	}
}

func (d *day2) Parse(input string) ([][2]rune, error) {
	lines := strings.Split(input, "\n")

	out := make([][2]rune, len(lines)-1)

	for i := 0; i < len(lines)-1; i++ {
		runes := []rune(lines[i])

		out[i] = [2]rune{runes[0], runes[2]}
	}

	return out, nil
}

func (d *day2) Part1(rounds [][2]rune) (string, error) {
	total := 0

	for _, round := range rounds {
		opponent := rockPaperScissors.scores[round[0]]
		me := rockPaperScissors.scores[round[1]]

		if opponent == me {
			me += 3 // Draw
		} else if me == 1 && opponent == 3 || me == 3 && opponent == 2 || me == 2 && opponent == 1 {
			me += 6 // Victory
		}

		total += me
	}

	return fmt.Sprint(total), nil
}

func (d *day2) Part2(input [][2]rune) (string, error) {
	return "TODO", nil
}

func (d *day2) Exec(input string) (*DayResult, error) {
	parsed, err := d.Parse(input)

	if err != nil {
		return nil, err
	}

	part1, err := d.Part1(parsed)

	if err != nil {
		return nil, err
	}

	part2, err := d.Part2(parsed)

	if err != nil {
		return nil, err
	}

	result := &DayResult{part1, part2}

	return result, nil
}
