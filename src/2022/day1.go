package aoc2022

import (
	"strconv"
	"strings"
)

const dayNum = 1

func init() {
	Days[dayNum] = &day{}
}

type day struct{}

func (d *day) Day() int {
	return dayNum
}

func (d *day) Parse(input string) ([][]int, error) {
	strGroups := strings.Split(input, "\n\n")

	intGroups := make([][]int, len(strGroups))

	for i, v := range strGroups {
		strLines := strings.Split(v, "\n")

		intLines := make([]int, len(strLines))
		for j, v := range strLines {
			value, err := strconv.Atoi(v)

			if err != nil {
				continue
			}

			intLines[j] = value
		}

		intGroups[i] = intLines
	}
	return intGroups, nil
}

func (d *day) Part1(elves [][]int) (string, error) {
	most := 0

	for _, elf := range elves {
		curr := 0

		for _, v := range elf {
			curr += v
		}

		if most < curr {
			most = curr
		}
	}

	return strconv.Itoa(most), nil
}

func (d *day) Part2(parsed [][]int) (string, error) {

	for _, elf := range parsed {
		for _, calories := range elf {

		}
	}

	return "TODO", nil
}

func (d *day) Exec(input string) (*DayResult, error) {
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
