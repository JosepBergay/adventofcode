package aoc2022

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type day4 struct{}

func init() {
	Days[4] = &day4{}
}

func (d *day4) Parse(input string) ([][2][2]int, error) {
	lines := strings.Split(input, "\n")

	out := make([][2][2]int, len(lines)-1)

	for i, line := range lines[:len(lines)-1] {
		assignments := strings.Split(line, ",")

		for j, assignment := range assignments {
			sections := strings.Split(assignment, "-")

			for k, section := range sections {
				parsed, err := strconv.Atoi(section)

				if err != nil {
					return nil, err
				}

				out[i][j][k] = parsed
			}

		}
	}

	return out, nil
}

func (d *day4) Part1(input [][2][2]int) (string, error) {
	out := 0

	for _, v := range input {
		lowerBound := int(math.Min(float64(v[0][0]), float64(v[1][0])))
		upperBound := int(math.Max(float64(v[0][1]), float64(v[1][1])))

		firstFullyContains := lowerBound == v[0][0] && upperBound == v[0][1]
		secondFullyContains := lowerBound == v[1][0] && upperBound == v[1][1]

		if firstFullyContains || secondFullyContains {
			out++
		}
	}

	return fmt.Sprint(out), nil
}

func (d *day4) Part2(input [][2][2]int) (string, error) {
	return "TODO", nil
}

func (d *day4) Exec(input string) (*DayResult, error) {
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
