package aoc2022

import (
	"fmt"
	"math"
	"strings"
)

type day25 struct{}

func init() {
	Days[25] = &day25{}
}

func (d *day25) Parse(input string) ([]string, error) {
	split := strings.Split(input, "\n")

	return split[:len(split)-1], nil
}

var snafuToIntDigit = map[rune]int{
	'2': 2,
	'1': 1,
	'0': 0,
	'-': -1,
	'=': -2,
}

var intToSnafuDigit = map[int]string{
	2:  "2",
	1:  "1",
	0:  "0",
	-1: "-",
	-2: "=",
}

func fromSNAFU(snafu string) int {
	out := 0

	for i, r := range snafu {
		d := snafuToIntDigit[r]

		out += d * int(math.Pow(5, float64(len(snafu)-i-1)))
	}

	return out
}

func toSNAFU(x int) string {
	snafu := ""
	curr := x
	for curr > 0 {
		mod := ((curr + 2) % 5) - 2
		curr = (curr + 2) / 5

		snafu = intToSnafuDigit[mod] + snafu
	}

	return snafu
}

func (d *day25) Part1(input []string) (string, error) {
	sum := 0

	for _, s := range input {
		sum += fromSNAFU(s)
	}

	snafu := toSNAFU(sum)

	return fmt.Sprint(snafu), nil
}

func (d *day25) Part2(input []string) (string, error) {
	return "TODO", nil
}

func (d *day25) Exec(input string) (*DayResult, error) {
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
