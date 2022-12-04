package aoc2022

import (
	"fmt"
	"strings"
)

type day3 struct{}

func init() {
	Days[3] = &day3{}
}

func (d *day3) Parse(input string) ([]string, error) {
	lines := strings.Split(input, "\n")

	return lines[:len(lines)-1], nil
}

func (d *day3) Part1(input []string) (string, error) {
	sum := 0

	for _, line := range input {
		m := make(map[byte]string, len(line)/2)

		for i := 0; i < len(line)/2; i++ {
			m[line[i]] = ""
		}

		repeated := 0
		for i := len(line) / 2; i < len(line); i++ {
			if _, ok := m[line[i]]; ok {
				repeated = int(line[i])
				break
			}
		}

		// sum += int(repeated)
		if repeated >= 'a' {
			// Lower case
			sum += repeated - 'a' + 1
		} else {
			// Upper case
			sum += repeated - 'A' + 27
		}
	}

	return fmt.Sprint(sum), nil
}

func (d *day3) Part2(input []string) (string, error) {
	return "TODO", nil
}

func (d *day3) Exec(input string) (*DayResult, error) {
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
