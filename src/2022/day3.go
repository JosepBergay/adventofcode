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

func getPriority(char int) int {
	if char >= 'a' {
		// Lower case
		return char - 'a' + 1
	}

	// Upper case
	return char - 'A' + 27
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

		sum += getPriority(repeated)
	}

	return fmt.Sprint(sum), nil
}

func (d *day3) Part2(input []string) (string, error) {
	sum := 0

	// For each group of 3
	for i := 0; i < len(input); i = i + 3 {

		m := make(map[rune][3]bool)

		// For each line in the group
		for j := 0; j < 3; j++ {

			// For each rune in the line
			for _, r := range input[i+j] {
				arr := m[r]
				arr[j] = true
				m[r] = arr
			}
		}

		for k, v := range m {
			if v[0] && v[1] && v[2] {
				// Only the item that appears in all three rucksacks will have [true, true, true] as value.
				sum += getPriority(int(k))
				break
			}
		}
	}

	return fmt.Sprint(sum), nil
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
