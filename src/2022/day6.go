package aoc2022

import (
	"fmt"
	"strings"
)

type day6 struct{}

func init() {
	Days[6] = &day6{}
}

func (d *day6) Parse(input string) (string, error) {
	return input, nil
}

// isStartOfPacket returns true if candidate has no repeated values
func isStartOfPacket(candidate []byte) bool {
	m := make(map[byte]any)

	for _, v := range candidate {
		m[v] = true
	}

	return len(m) == len(candidate)
}

func (d *day6) Part1(input string) (string, error) {
	reader := strings.NewReader(input)

	i := 0
	for {
		msg := make([]byte, 4)

		_, err := reader.ReadAt(msg, int64(i))
		if err != nil {
			return "", err
		}

		if isStartOfPacket(msg) {
			return fmt.Sprint(i + 4), nil
		}

		i++
	}
}

func (d *day6) Part2(input string) (string, error) {
	return "TODO", nil
}

func (d *day6) Exec(input string) (*DayResult, error) {
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
